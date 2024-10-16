package company

import (
	"goravel/app/models"
	"goravel/config/responses"
	"goravel/config/tables"

	"github.com/goravel/framework/facades"
)

type Employee struct {
	models.BaseStruct

	Nik    string `json:"nik"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func (m *Employee) TableName() string {
	return tables.TABLE_EMPLOYEE
}

func (m *Employee) GetRecords(page int, limit int) (responses.Pagination, []Employee) {
	var records []Employee
	var pagination responses.Pagination

	qSelect := []string{
		"id",
		"nik",
		"name",
		"gender",
	}
	qWhere := "is_active = 't'"
	qOrder := "id ASC"

	err := facades.Orm().
		Query().
		Select(qSelect).
		Where(qWhere).
		Order(qOrder).
		Paginate(page, limit, &records, &pagination.TotalRecords)
	if err != nil {
		panic(err)
	}

	return pagination, records
}
