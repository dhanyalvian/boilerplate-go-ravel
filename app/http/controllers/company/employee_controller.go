package company

import (
	"goravel/app/http/controllers"
	"goravel/app/models/company"

	"github.com/goravel/framework/contracts/http"
)

type CompanyEmployeeController struct {
	//Dependent services
}

func NewCompanyEmployeeControler() *CompanyEmployeeController {
	return &CompanyEmployeeController{
		//Inject services
	}
}

func (r *CompanyEmployeeController) List(ctx http.Context) http.Response {
	page := controllers.GetReqDataPage(ctx)
	limit := controllers.GetReqDataLimit(ctx)

	var employee company.Employee
	pagination, records := employee.GetRecords(page, limit)

	return controllers.GenerateSuccessList(ctx, pagination, records)
}

func (r *CompanyEmployeeController) Detail(ctx http.Context) http.Response {
	id := controllers.GetReqId(ctx)
	return ctx.Response().Success().Json(http.Json{
		"url":    "/company/employees/" + id,
		"module": "Company > Employee",
	})
}
