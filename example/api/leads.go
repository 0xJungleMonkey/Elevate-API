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

func configLeadsRouter(router *httprouter.Router) {
	router.GET("/leads", GetAllLeads)
	router.POST("/leads", AddLeads)
	router.GET("/leads/:argID", GetLeads)
	router.PUT("/leads/:argID", UpdateLeads)
	router.DELETE("/leads/:argID", DeleteLeads)
}

func configGinLeadsRouter(router gin.IRoutes) {
	router.GET("/leads", ConverHttprouterToGin(GetAllLeads))
	router.POST("/leads", ConverHttprouterToGin(AddLeads))
	router.GET("/leads/:argID", ConverHttprouterToGin(GetLeads))
	router.PUT("/leads/:argID", ConverHttprouterToGin(UpdateLeads))
	router.DELETE("/leads/:argID", ConverHttprouterToGin(DeleteLeads))
}

// GetAllLeads is a function to get a slice of record(s) from leads table in the rocket_development database
// @Summary Get list of Leads
// @Tags Leads
// @Description GetAllLeads is a handler to get a slice of record(s) from leads table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Leads}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leads [get]
// http "http://localhost:8080/leads?page=0&pagesize=20" X-Api-User:user123
func GetAllLeads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "leads", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLeads(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLeads is a function to get a single record from the leads table in the rocket_development database
// @Summary Get record from table Leads by  argID
// @Tags Leads
// @ID argID
// @Description GetLeads is a function to get a single record from the leads table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Leads
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /leads/{argID} [get]
// http "http://localhost:8080/leads/1" X-Api-User:user123
func GetLeads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "leads", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLeads(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLeads add to add a single record to leads table in the rocket_development database
// @Summary Add an record to leads table
// @Description add to add a single record to leads table in the rocket_development database
// @Tags Leads
// @Accept  json
// @Produce  json
// @Param Leads body model.Leads true "Add Leads"
// @Success 200 {object} model.Leads
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leads [post]
// echo '{"id": 34,"full_name_of_the_contact": "VubQUclMrYnJdXEjigwVJYJpb","bussiness_name": "kmehEWaecfYpTnxbqsyjLgiPZ","email": "xIPxNLpBMEVJIYqwDvpYMsmAY","phone": "gJeiQZmqPfBfEvdORqmxAFZoS","project_name": "EnDHNvXkkfSELooLmqqwekxEX","project_description": "XfwcNpnoWSfZLLDIGWGFemTHx","department_incharge": "iZFyDVbwMIclhilMytscpMhyL","message": "srltjuVoYobsrQLNZmGVncWOw","attached_file": "GklaMxxFVQYvJz5QGyhgBEBaBhljMQUZOmIAJ15VH0ADPDxTGQpiXx8sCh8tXjo8KTI7Y0QrBz5hPUBjLT5jIFgMHg==","creation_date": "2314-02-22T11:24:02.085806613-05:00","created_at": "2187-09-09T23:29:36.135583678-04:00","updated_at": "2095-03-26T09:01:43.425372385-04:00"}' | http POST "http://localhost:8080/leads" X-Api-User:user123
func AddLeads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	leads := &model.Leads{}

	if err := readJSON(r, leads); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leads.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leads.Prepare()

	if err := leads.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "leads", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	leads, _, err = dao.AddLeads(ctx, leads)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leads)
}

// UpdateLeads Update a single record from leads table in the rocket_development database
// @Summary Update an record in table leads
// @Description Update a single record from leads table in the rocket_development database
// @Tags Leads
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Leads body model.Leads true "Update Leads record"
// @Success 200 {object} model.Leads
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leads/{argID} [put]
// echo '{"id": 34,"full_name_of_the_contact": "VubQUclMrYnJdXEjigwVJYJpb","bussiness_name": "kmehEWaecfYpTnxbqsyjLgiPZ","email": "xIPxNLpBMEVJIYqwDvpYMsmAY","phone": "gJeiQZmqPfBfEvdORqmxAFZoS","project_name": "EnDHNvXkkfSELooLmqqwekxEX","project_description": "XfwcNpnoWSfZLLDIGWGFemTHx","department_incharge": "iZFyDVbwMIclhilMytscpMhyL","message": "srltjuVoYobsrQLNZmGVncWOw","attached_file": "GklaMxxFVQYvJz5QGyhgBEBaBhljMQUZOmIAJ15VH0ADPDxTGQpiXx8sCh8tXjo8KTI7Y0QrBz5hPUBjLT5jIFgMHg==","creation_date": "2314-02-22T11:24:02.085806613-05:00","created_at": "2187-09-09T23:29:36.135583678-04:00","updated_at": "2095-03-26T09:01:43.425372385-04:00"}' | http PUT "http://localhost:8080/leads/1"  X-Api-User:user123
func UpdateLeads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leads := &model.Leads{}
	if err := readJSON(r, leads); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leads.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leads.Prepare()

	if err := leads.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "leads", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leads, _, err = dao.UpdateLeads(ctx,
		argID,
		leads)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leads)
}

// DeleteLeads Delete a single record from leads table in the rocket_development database
// @Summary Delete a record from leads
// @Description Delete a single record from leads table in the rocket_development database
// @Tags Leads
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Leads
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /leads/{argID} [delete]
// http DELETE "http://localhost:8080/leads/1" X-Api-User:user123
func DeleteLeads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "leads", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLeads(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
