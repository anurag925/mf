package models

import (
	"time"

	"github.com/uptrace/bun"
)

type MfApiRelation struct {
	bun.BaseModel `bun:"table:mf_api_relations,alias:mar"`

	ID         int64     `bun:",pk,autoincrement"`
	CreatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	RelationID int64
	SchemeID   int64

	Scheme Scheme `bun:"rel:belongs-to,join:scheme_id=id"`
}
