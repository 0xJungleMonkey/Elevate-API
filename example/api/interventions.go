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

func configInterventions_Router(router *httprouter.Router) {
	router.GET("/interventions_", GetAllInterventions_)
	router.POST("/interventions_", AddInterventions_)
	router.GET("/interventions_/:argID", GetInterventions_)
	router.PUT("/interventions_/:argID", UpdateInterventions_)
	router.DELETE("/interventions_/:argID", DeleteInterventions_)
}

func configGinInterventions_Router(router gin.IRoutes) {
	router.GET("/interventions_", ConverHttprouterToGin(GetAllInterventions_))
	router.POST("/interventions_", ConverHttprouterToGin(AddInterventions_))
	router.GET("/interventions_/:argID", ConverHttprouterToGin(GetInterventions_))
	router.PUT("/interventions_/:argID", ConverHttprouterToGin(UpdateInterventions_))
	router.DELETE("/interventions_/:argID", ConverHttprouterToGin(DeleteInterventions_))
}

// GetAllInterventions_ is a function to get a slice of record(s) from interventions table in the rocket_development database
// @Summary Get list of Interventions_
// @Tags Interventions_
// @Description GetAllInterventions_ is a handler to get a slice of record(s) from interventions table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Interventions_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions_ [get]
// http "http://localhost:8080/interventions_?page=0&pagesize=20" X-Api-User:user123
func GetAllInterventions_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "interventions", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllInterventions_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetInterventions_ is a function to get a single record from the interventions table in the rocket_development database
// @Summary Get record from table Interventions_ by  argID
// @Tags Interventions_
// @ID argID
// @Description GetInterventions_ is a function to get a single record from the interventions table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Interventions_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /interventions_/{argID} [get]
// http "http://localhost:8080/interventions_/1" X-Api-User:user123
func GetInterventions_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetInterventions_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddInterventions_ add to add a single record to interventions table in the rocket_development database
// @Summary Add an record to interventions table
// @Description add to add a single record to interventions table in the rocket_development database
// @Tags Interventions_
// @Accept  json
// @Produce  json
// @Param Interventions_ body model.Interventions_ true "Add Interventions_"
// @Success 200 {object} model.Interventions_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions_ [post]
// echo '{"id": 65,"author": "KAnYaNnbOsMETHgRorXLarTfL","customer_id": 83,"building_id": 66,"battery_id": 87,"column_id": 66,"elevator_id": 84,"employee_id": 62,"start_datetime": "2078-04-19T19:56:42.25256109-04:00","end_datetime": "2133-01-30T05:31:22.685708736-05:00","result": "OuwZLFcJIuDNEigwnJFvIRXWv","report": "CttuQjQmffNkWpnQFTKCvZrlB","status": "KKypUFNTPhjaHbwMeDeftJCtd","created_at": "2206-06-15T13:57:41.061379409-04:00","updated_at": "2024-09-17T14:03:42.985744518-04:00"}' | http POST "http://localhost:8080/interventions_" X-Api-User:user123
func AddInterventions_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	interventions_ := &model.Interventions_{}

	if err := readJSON(r, interventions_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := interventions_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	interventions_.Prepare()

	if err := interventions_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	interventions_, _, err = dao.AddInterventions_(ctx, interventions_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, interventions_)
}

// UpdateInterventions_ Update a single record from interventions table in the rocket_development database
// @Summary Update an record in table interventions
// @Description Update a single record from interventions table in the rocket_development database
// @Tags Interventions_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Interventions_ body model.Interventions_ true "Update Interventions_ record"
// @Success 200 {object} model.Interventions_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions_/{argID} [put]
// echo '{"id": 65,"author": "KAnYaNnbOsMETHgRorXLarTfL","customer_id": 83,"building_id": 66,"battery_id": 87,"column_id": 66,"elevator_id": 84,"employee_id": 62,"start_datetime": "2078-04-19T19:56:42.25256109-04:00","end_datetime": "2133-01-30T05:31:22.685708736-05:00","result": "OuwZLFcJIuDNEigwnJFvIRXWv","report": "CttuQjQmffNkWpnQFTKCvZrlB","status": "KKypUFNTPhjaHbwMeDeftJCtd","created_at": "2206-06-15T13:57:41.061379409-04:00","updated_at": "2024-09-17T14:03:42.985744518-04:00"}' | http PUT "http://localhost:8080/interventions_/1"  X-Api-User:user123
func UpdateInterventions_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	interventions_ := &model.Interventions_{}
	if err := readJSON(r, interventions_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := interventions_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	interventions_.Prepare()

	if err := interventions_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	interventions_, _, err = dao.UpdateInterventions_(ctx,
		argID,
		interventions_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, interventions_)
}

// DeleteInterventions_ Delete a single record from interventions table in the rocket_development database
// @Summary Delete a record from interventions
// @Description Delete a single record from interventions table in the rocket_development database
// @Tags Interventions_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Interventions_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /interventions_/{argID} [delete]
// http DELETE "http://localhost:8080/interventions_/1" X-Api-User:user123
func DeleteInterventions_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteInterventions_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
