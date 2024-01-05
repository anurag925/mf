package scripts

import (
	"context"
	"log/slog"

	"github.com/anurag925/mf/core/logger"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type allMFApiResponse []struct {
	SchemeCode int    `json:"schemeCode"`
	SchemeName string `json:"schemeName"`
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
}
