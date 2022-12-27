package api

import (
	"net/http"

	"restapi-golang-gin-gen/dao"
	"restapi-golang-gin-gen/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configEmployeesRouter(router *httprouter.Router) {
	router.GET("/employees", GetAllEmployees)
	router.POST("/employees", AddEmployees)
	router.GET("/employees/:argID", GetEmployees)
	router.PUT("/employees/:argID", UpdateEmployees)
	router.DELETE("/employees/:argID", DeleteEmployees)
}

func configGinEmployeesRouter(router gin.IRoutes) {
	router.GET("/employees", ConverHttprouterToGin(GetAllEmployees))
	router.POST("/employees", ConverHttprouterToGin(AddEmployees))
	router.GET("/employees/:argID", ConverHttprouterToGin(GetEmployees))
	router.PUT("/employees/:argID", ConverHttprouterToGin(UpdateEmployees))
	router.DELETE("/employees/:argID", ConverHttprouterToGin(DeleteEmployees))
}

// GetAllEmployees is a function to get a slice of record(s) from employees table in the rocket_development database
// @Summary Get list of Employees
// @Tags Employees
// @Description GetAllEmployees is a handler to get a slice of record(s) from employees table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Employees}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /employees [get]
// http "http://localhost:8080/employees?page=0&pagesize=20" X-Api-User:user123
func GetAllEmployees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "employees", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllEmployees(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetEmployees is a function to get a single record from the employees table in the rocket_development database
// @Summary Get record from table Employees by  argID
// @Tags Employees
// @ID argID
// @Description GetEmployees is a function to get a single record from the employees table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Employees
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /employees/{argID} [get]
// http "http://localhost:8080/employees/1" X-Api-User:user123
func GetEmployees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "employees", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetEmployees(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddEmployees add to add a single record to employees table in the rocket_development database
// @Summary Add an record to employees table
// @Description add to add a single record to employees table in the rocket_development database
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param Employees body model.Employees true "Add Employees"
// @Success 200 {object} model.Employees
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /employees [post]
// echo '{"user_id": 76,"id": 44,"first_name": "irZMmJdQJeuvMEDNBGWqHgcon","last_name": "BrrftMXwMBxtvEMufOQJvdpJt","title": "KFHdwLlcrhXQUEbJNucIQSqKq","email": "YZkiZrJyffBQiRMHMcqpVEidY","created_at": "2173-01-29T04:55:20.654753241-05:00","updated_at": "2274-02-15T19:00:20.172408353-05:00"}' | http POST "http://localhost:8080/employees" X-Api-User:user123
func AddEmployees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	employees := &model.Employees{}

	if err := readJSON(r, employees); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := employees.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	employees.Prepare()

	if err := employees.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "employees", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	employees, _, err = dao.AddEmployees(ctx, employees)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, employees)
}

// UpdateEmployees Update a single record from employees table in the rocket_development database
// @Summary Update an record in table employees
// @Description Update a single record from employees table in the rocket_development database
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Employees body model.Employees true "Update Employees record"
// @Success 200 {object} model.Employees
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /employees/{argID} [put]
// echo '{"user_id": 76,"id": 44,"first_name": "irZMmJdQJeuvMEDNBGWqHgcon","last_name": "BrrftMXwMBxtvEMufOQJvdpJt","title": "KFHdwLlcrhXQUEbJNucIQSqKq","email": "YZkiZrJyffBQiRMHMcqpVEidY","created_at": "2173-01-29T04:55:20.654753241-05:00","updated_at": "2274-02-15T19:00:20.172408353-05:00"}' | http PUT "http://localhost:8080/employees/1"  X-Api-User:user123
func UpdateEmployees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	employees := &model.Employees{}
	if err := readJSON(r, employees); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := employees.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	employees.Prepare()

	if err := employees.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "employees", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	employees, _, err = dao.UpdateEmployees(ctx,
		argID,
		employees)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, employees)
}

// DeleteEmployees Delete a single record from employees table in the rocket_development database
// @Summary Delete a record from employees
// @Description Delete a single record from employees table in the rocket_development database
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Employees
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /employees/{argID} [delete]
// http DELETE "http://localhost:8080/employees/1" X-Api-User:user123
func DeleteEmployees(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "employees", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteEmployees(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
