package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel

	ID             int64     `bun:",pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	FirstName      string
	MiddleName     string
	LastName       string
	Email          string
	PasswordDigest string
	MobileNumber   string
}
