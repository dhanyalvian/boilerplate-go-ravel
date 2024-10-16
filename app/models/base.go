package models

import "time"

type BaseStruct struct {
	// ID of as primary key
	// in: int64
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
