package routes

import (
	"fmt"

	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/controllers/company"
	"goravel/config/routes"
)

func Api() {
	routeUser()
	routeCompany()
}

func routeUser() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}

func routeCompany() {
	employeeUrl := routes.ROUTE_COMPANY_EMPLOYEE
	employeeCtrl := company.NewCompanyEmployeeControler()

	facades.Route().Get(employeeUrl, employeeCtrl.List)
	facades.Route().Get(fmt.Sprintf("%s/{id}", employeeUrl), employeeCtrl.Detail)
}
