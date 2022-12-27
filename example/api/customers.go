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

func configCustomers_Router(router *httprouter.Router) {
	router.GET("/customers_", GetAllCustomers_)
	router.POST("/customers_", AddCustomers_)
	router.GET("/customers_/:argID", GetCustomers_)
	router.PUT("/customers_/:argID", UpdateCustomers_)
	router.DELETE("/customers_/:argID", DeleteCustomers_)
}

func configGinCustomers_Router(router gin.IRoutes) {
	router.GET("/customers_", ConverHttprouterToGin(GetAllCustomers_))
	router.POST("/customers_", ConverHttprouterToGin(AddCustomers_))
	router.GET("/customers_/:argID", ConverHttprouterToGin(GetCustomers_))
	router.PUT("/customers_/:argID", ConverHttprouterToGin(UpdateCustomers_))
	router.DELETE("/customers_/:argID", ConverHttprouterToGin(DeleteCustomers_))
}

// GetAllCustomers_ is a function to get a slice of record(s) from customers table in the rocket_development database
// @Summary Get list of Customers_
// @Tags Customers_
// @Description GetAllCustomers_ is a handler to get a slice of record(s) from customers table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Customers_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers_ [get]
// http "http://localhost:8080/customers_?page=0&pagesize=20" X-Api-User:user123
func GetAllCustomers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "customers", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCustomers_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetCustomers_ is a function to get a single record from the customers table in the rocket_development database
// @Summary Get record from table Customers_ by  argID
// @Tags Customers_
// @ID argID
// @Description GetCustomers_ is a function to get a single record from the customers table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Customers_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /customers_/{argID} [get]
// http "http://localhost:8080/customers_/1" X-Api-User:user123
func GetCustomers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCustomers_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCustomers_ add to add a single record to customers table in the rocket_development database
// @Summary Add an record to customers table
// @Description add to add a single record to customers table in the rocket_development database
// @Tags Customers_
// @Accept  json
// @Produce  json
// @Param Customers_ body model.Customers_ true "Add Customers_"
// @Success 200 {object} model.Customers_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers_ [post]
// echo '{"address_id": 44,"user_id": 2,"id": 93,"customer_creation_date": "VZNdjPMdgJbnBOXWwEcoZBddF","date": "WCdPvfuxhNmkfOVxtFOsuIgCU","company_name": "OOsKhOXcPRCxjEMEeXSeieCxS","company_hq_adress": "ySkdGuXDAqdePaJJjyavqykJm","full_name_of_company_contact": "wYCjgqRNIoKjujYqaVYxdMWiJ","company_contact_phone": "eueltgJZGMWISqwNIILkwLYRs","company_contact_e_mail": "wdCePrGEthWDpcyLxDSCZJURg","company_desc": "MVOTvwQKQVaxLkeOPMHVLucSX","full_name_service_tech_auth": "ZpWBrKhcFDTXnmbmfDnfUBWJk","tech_auth_phone_service": "vZoEcrYHnEihInwxArxRaAFdU","tech_manager_email_service": "JTxpobNqnXjHXywurCZxLoyUp","created_at": "2245-11-08T05:54:50.870909776-05:00","updated_at": "2168-06-09T12:13:23.24668422-04:00"}' | http POST "http://localhost:8080/customers_" X-Api-User:user123
func AddCustomers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	customers_ := &model.Customers_{}

	if err := readJSON(r, customers_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := customers_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	customers_.Prepare()

	if err := customers_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	customers_, _, err = dao.AddCustomers_(ctx, customers_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, customers_)
}

// UpdateCustomers_ Update a single record from customers table in the rocket_development database
// @Summary Update an record in table customers
// @Description Update a single record from customers table in the rocket_development database
// @Tags Customers_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Customers_ body model.Customers_ true "Update Customers_ record"
// @Success 200 {object} model.Customers_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers_/{argID} [put]
// echo '{"address_id": 44,"user_id": 2,"id": 93,"customer_creation_date": "VZNdjPMdgJbnBOXWwEcoZBddF","date": "WCdPvfuxhNmkfOVxtFOsuIgCU","company_name": "OOsKhOXcPRCxjEMEeXSeieCxS","company_hq_adress": "ySkdGuXDAqdePaJJjyavqykJm","full_name_of_company_contact": "wYCjgqRNIoKjujYqaVYxdMWiJ","company_contact_phone": "eueltgJZGMWISqwNIILkwLYRs","company_contact_e_mail": "wdCePrGEthWDpcyLxDSCZJURg","company_desc": "MVOTvwQKQVaxLkeOPMHVLucSX","full_name_service_tech_auth": "ZpWBrKhcFDTXnmbmfDnfUBWJk","tech_auth_phone_service": "vZoEcrYHnEihInwxArxRaAFdU","tech_manager_email_service": "JTxpobNqnXjHXywurCZxLoyUp","created_at": "2245-11-08T05:54:50.870909776-05:00","updated_at": "2168-06-09T12:13:23.24668422-04:00"}' | http PUT "http://localhost:8080/customers_/1"  X-Api-User:user123
func UpdateCustomers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	customers_ := &model.Customers_{}
	if err := readJSON(r, customers_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := customers_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	customers_.Prepare()

	if err := customers_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	customers_, _, err = dao.UpdateCustomers_(ctx,
		argID,
		customers_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, customers_)
}

// DeleteCustomers_ Delete a single record from customers table in the rocket_development database
// @Summary Delete a record from customers
// @Description Delete a single record from customers table in the rocket_development database
// @Tags Customers_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Customers_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /customers_/{argID} [delete]
// http DELETE "http://localhost:8080/customers_/1" X-Api-User:user123
func DeleteCustomers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCustomers_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
