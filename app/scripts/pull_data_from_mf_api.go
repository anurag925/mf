package scripts

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/anurag925/mf/app/models"
	"github.com/anurag925/mf/core"
	"github.com/anurag925/mf/core/logger"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type allMFApiResponse []struct {
	SchemeCode int    `json:"schemeCode"`
	SchemeName string `json:"schemeName"`
}

type NavApiResponse struct {
	Meta struct {
		FundHouse      string `json:"fund_house"`
		SchemeType     string `json:"scheme_type"`
		SchemeCategory string `json:"scheme_category"`
		SchemeCode     int    `json:"scheme_code"`
		SchemeName     string `json:"scheme_name"`
	} `json:"meta"`
	Data []NavData `json:"data"`
}

type NavData struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

func PullDataFromMfAPI() {
	ctx := logger.AppendCtx(context.TODO(), slog.Attr{Key: "request_id", Value: slog.StringValue(uuid.NewString())})
	client := resty.New()
	url := "https://api.mfapi.in/mf"
	logger.Info(ctx, "calling mf api", "url", url)
	allMFApiResponse := allMFApiResponse{}
	res, err := client.R().SetContext(ctx).EnableTrace().SetResult(&allMFApiResponse).Get(url)
	if err != nil {
		logger.Error(ctx, "unable to complete request")
		return
	}
	if res.IsError() {
		logger.Error(ctx, "unable to complete request", "status", res.Status(), "res", res)
		return
	}
	logger.Info(ctx, "successfully completed request", "status", res.Status(), "res", res)
	logger.Info(ctx, "all mf api response", "res", allMFApiResponse)
	failedSchemeCode := []int{}
	for i, v := range allMFApiResponse {
		logger.Info(ctx, "scheme", "index", i, "schemeCode", v.SchemeCode, "schemeName", v.SchemeName)
		schemeUrl := fmt.Sprintf("%s/%d", url, v.SchemeCode)
		navResponse := NavApiResponse{}
		res, err := client.R().SetContext(ctx).EnableTrace().SetResult(&navResponse).Get(schemeUrl)
		if err != nil {
			logger.Error(ctx, "unable to complete request")
			failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
			continue
		}
		if res.IsError() {
			logger.Error(ctx, "unable to complete request", "status", res.Status(), "res", res)
			failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
			continue
		}
		logger.Info(ctx, "successfully completed request", "status", res.Status(), "res", res)
		scheme := models.Scheme{
			FundHouse:      navResponse.Meta.FundHouse,
			SchemeType:     navResponse.Meta.SchemeType,
			SchemeCategory: navResponse.Meta.SchemeCategory,
			SchemeName:     navResponse.Meta.SchemeName}

		if _, err := core.DB().NewInsert().Model(&scheme).Exec(ctx); err != nil {
			logger.Error(ctx, "unable to insert scheme", "err", err)
			failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
			continue
		}
		logger.Info(ctx, "inserted scheme for", "scheme_code", v.SchemeCode, "scheme", scheme)
		mfApiRelations := models.MfApiRelation{RelationID: int64(v.SchemeCode), SchemeID: scheme.ID}
		if _, err := core.DB().NewInsert().Model(&mfApiRelations).Exec(ctx); err != nil {
			logger.Error(ctx, "unable to insert mfApiRelations", "err", err)
			failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
			continue
		}
		logger.Info(ctx, "inserted mfApiRelations for", "scheme_code", v.SchemeCode, "mfApiRelations", mfApiRelations)
		navs := make([]models.Nav, len(navResponse.Data))

		for i, val := range navResponse.Data {
			date, err := time.Parse("02-01-2006", val.Date)
			if err != nil {
				logger.Error(ctx, "unable to parse date", "err", err)
				failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
				continue
			}
			nalVal, err := strconv.ParseFloat(val.Nav, 64)
			if err != nil {
				logger.Error(ctx, "unable to parse nav", "err", err)
				failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
				continue
			}
			navs[i] = models.Nav{Date: date, Value: nalVal, SchemeID: scheme.ID}
		}
		if _, err := core.DB().NewInsert().Model(&navs).Exec(ctx); err != nil {
			logger.Error(ctx, "unable to insert navs", "err", err)
			failedSchemeCode = append(failedSchemeCode, v.SchemeCode)
			continue
		}
		logger.Info(ctx, "inserted navs for", "scheme_code", v.SchemeCode, "navs", navs)
		time.Sleep(1 * time.Second)
	}
	logger.Info(ctx, "failed scheme code", "res", failedSchemeCode)
}
