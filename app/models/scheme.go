package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Scheme struct {
	bun.BaseModel

	ID             int64     `bun:",pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	FundHouse      string
	SchemeName     string
	SchemeType     string
	SchemeCategory string
}
