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

func configBlazerQueries_Router(router *httprouter.Router) {
	router.GET("/blazerqueries_", GetAllBlazerQueries_)
	router.POST("/blazerqueries_", AddBlazerQueries_)
	router.GET("/blazerqueries_/:argID", GetBlazerQueries_)
	router.PUT("/blazerqueries_/:argID", UpdateBlazerQueries_)
	router.DELETE("/blazerqueries_/:argID", DeleteBlazerQueries_)
}

func configGinBlazerQueries_Router(router gin.IRoutes) {
	router.GET("/blazerqueries_", ConverHttprouterToGin(GetAllBlazerQueries_))
	router.POST("/blazerqueries_", ConverHttprouterToGin(AddBlazerQueries_))
	router.GET("/blazerqueries_/:argID", ConverHttprouterToGin(GetBlazerQueries_))
	router.PUT("/blazerqueries_/:argID", ConverHttprouterToGin(UpdateBlazerQueries_))
	router.DELETE("/blazerqueries_/:argID", ConverHttprouterToGin(DeleteBlazerQueries_))
}

// GetAllBlazerQueries_ is a function to get a slice of record(s) from blazer_queries table in the rocket_development database
// @Summary Get list of BlazerQueries_
// @Tags BlazerQueries_
// @Description GetAllBlazerQueries_ is a handler to get a slice of record(s) from blazer_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerQueries_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries_ [get]
// http "http://localhost:8080/blazerqueries_?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_queries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerQueries_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerQueries_ is a function to get a single record from the blazer_queries table in the rocket_development database
// @Summary Get record from table BlazerQueries_ by  argID
// @Tags BlazerQueries_
// @ID argID
// @Description GetBlazerQueries_ is a function to get a single record from the blazer_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerqueries_/{argID} [get]
// http "http://localhost:8080/blazerqueries_/1" X-Api-User:user123
func GetBlazerQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerQueries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerQueries_ add to add a single record to blazer_queries table in the rocket_development database
// @Summary Add an record to blazer_queries table
// @Description add to add a single record to blazer_queries table in the rocket_development database
// @Tags BlazerQueries_
// @Accept  json
// @Produce  json
// @Param BlazerQueries_ body model.BlazerQueries_ true "Add BlazerQueries_"
// @Success 200 {object} model.BlazerQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries_ [post]
// echo '{"id": 8,"creator_id": 36,"name": "AxfmrEbJNxpWmooBLsmqUsglF","description": "FbJLVgJnJaSEXNArXUSGcTreu","statement": "IpIAuHuYxXToqxfHjKPidNmLy","data_source": "qIALNagNbboBuqcBTayiKpvUG","status": "OSgBPAGKijBcCDSXTbjsJFEdS","created_at": "2306-05-15T00:20:03.44145846-04:00","updated_at": "2302-08-14T07:53:36.278382772-04:00"}' | http POST "http://localhost:8080/blazerqueries_" X-Api-User:user123
func AddBlazerQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerqueries_ := &model.BlazerQueries_{}

	if err := readJSON(r, blazerqueries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerqueries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerqueries_.Prepare()

	if err := blazerqueries_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerqueries_, _, err = dao.AddBlazerQueries_(ctx, blazerqueries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerqueries_)
}

// UpdateBlazerQueries_ Update a single record from blazer_queries table in the rocket_development database
// @Summary Update an record in table blazer_queries
// @Description Update a single record from blazer_queries table in the rocket_development database
// @Tags BlazerQueries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerQueries_ body model.BlazerQueries_ true "Update BlazerQueries_ record"
// @Success 200 {object} model.BlazerQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries_/{argID} [put]
// echo '{"id": 8,"creator_id": 36,"name": "AxfmrEbJNxpWmooBLsmqUsglF","description": "FbJLVgJnJaSEXNArXUSGcTreu","statement": "IpIAuHuYxXToqxfHjKPidNmLy","data_source": "qIALNagNbboBuqcBTayiKpvUG","status": "OSgBPAGKijBcCDSXTbjsJFEdS","created_at": "2306-05-15T00:20:03.44145846-04:00","updated_at": "2302-08-14T07:53:36.278382772-04:00"}' | http PUT "http://localhost:8080/blazerqueries_/1"  X-Api-User:user123
func UpdateBlazerQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerqueries_ := &model.BlazerQueries_{}
	if err := readJSON(r, blazerqueries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerqueries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerqueries_.Prepare()

	if err := blazerqueries_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerqueries_, _, err = dao.UpdateBlazerQueries_(ctx,
		argID,
		blazerqueries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerqueries_)
}

// DeleteBlazerQueries_ Delete a single record from blazer_queries table in the rocket_development database
// @Summary Delete a record from blazer_queries
// @Description Delete a single record from blazer_queries table in the rocket_development database
// @Tags BlazerQueries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerqueries_/{argID} [delete]
// http DELETE "http://localhost:8080/blazerqueries_/1" X-Api-User:user123
func DeleteBlazerQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerQueries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
