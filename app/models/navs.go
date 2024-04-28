package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Nav struct {
	bun.BaseModel

	ID        int64     `bun:",pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Date      time.Time
	Value     float64
	SchemeID  int64

	Scheme Scheme `bun:"rel:belongs-to,join:scheme_id=id"`
}
