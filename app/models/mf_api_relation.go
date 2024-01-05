package models

import "github.com/uptrace/bun"

type MfApiRelation struct {
	bun.BaseModel `bun:"table:mfi_api_relations,alias:mar"`

	ID         int64 `bun:",pk,autoincrement"`
	RelationID int64
	SchemeID   int64

	Scheme Scheme `bun:"rel:belongs-to,join:scheme_id=id"`
}
