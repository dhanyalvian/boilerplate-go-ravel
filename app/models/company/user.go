package company

import (
	"goravel/config/tables"

	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model

	Uid      string `json:"uid"`
	Username string `json:"username"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (r *User) TableName() string {
	return tables.TABLE_USER
}
