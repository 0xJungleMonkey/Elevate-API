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

func configBlazerDashboardQueries_Router(router *httprouter.Router) {
	router.GET("/blazerdashboardqueries_", GetAllBlazerDashboardQueries_)
	router.POST("/blazerdashboardqueries_", AddBlazerDashboardQueries_)
	router.GET("/blazerdashboardqueries_/:argID", GetBlazerDashboardQueries_)
	router.PUT("/blazerdashboardqueries_/:argID", UpdateBlazerDashboardQueries_)
	router.DELETE("/blazerdashboardqueries_/:argID", DeleteBlazerDashboardQueries_)
}

func configGinBlazerDashboardQueries_Router(router gin.IRoutes) {
	router.GET("/blazerdashboardqueries_", ConverHttprouterToGin(GetAllBlazerDashboardQueries_))
	router.POST("/blazerdashboardqueries_", ConverHttprouterToGin(AddBlazerDashboardQueries_))
	router.GET("/blazerdashboardqueries_/:argID", ConverHttprouterToGin(GetBlazerDashboardQueries_))
	router.PUT("/blazerdashboardqueries_/:argID", ConverHttprouterToGin(UpdateBlazerDashboardQueries_))
	router.DELETE("/blazerdashboardqueries_/:argID", ConverHttprouterToGin(DeleteBlazerDashboardQueries_))
}

// GetAllBlazerDashboardQueries_ is a function to get a slice of record(s) from blazer_dashboard_queries table in the rocket_development database
// @Summary Get list of BlazerDashboardQueries_
// @Tags BlazerDashboardQueries_
// @Description GetAllBlazerDashboardQueries_ is a handler to get a slice of record(s) from blazer_dashboard_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerDashboardQueries_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries_ [get]
// http "http://localhost:8080/blazerdashboardqueries_?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerDashboardQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerDashboardQueries_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerDashboardQueries_ is a function to get a single record from the blazer_dashboard_queries table in the rocket_development database
// @Summary Get record from table BlazerDashboardQueries_ by  argID
// @Tags BlazerDashboardQueries_
// @ID argID
// @Description GetBlazerDashboardQueries_ is a function to get a single record from the blazer_dashboard_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerDashboardQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerdashboardqueries_/{argID} [get]
// http "http://localhost:8080/blazerdashboardqueries_/1" X-Api-User:user123
func GetBlazerDashboardQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerDashboardQueries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerDashboardQueries_ add to add a single record to blazer_dashboard_queries table in the rocket_development database
// @Summary Add an record to blazer_dashboard_queries table
// @Description add to add a single record to blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries_
// @Accept  json
// @Produce  json
// @Param BlazerDashboardQueries_ body model.BlazerDashboardQueries_ true "Add BlazerDashboardQueries_"
// @Success 200 {object} model.BlazerDashboardQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries_ [post]
// echo '{"id": 83,"dashboard_id": 66,"query_id": 60,"position": 37,"created_at": "2102-07-07T03:40:07.52114824-04:00","updated_at": "2254-04-24T07:25:54.102440438-04:00"}' | http POST "http://localhost:8080/blazerdashboardqueries_" X-Api-User:user123
func AddBlazerDashboardQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerdashboardqueries_ := &model.BlazerDashboardQueries_{}

	if err := readJSON(r, blazerdashboardqueries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboardqueries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboardqueries_.Prepare()

	if err := blazerdashboardqueries_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerdashboardqueries_, _, err = dao.AddBlazerDashboardQueries_(ctx, blazerdashboardqueries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboardqueries_)
}

// UpdateBlazerDashboardQueries_ Update a single record from blazer_dashboard_queries table in the rocket_development database
// @Summary Update an record in table blazer_dashboard_queries
// @Description Update a single record from blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerDashboardQueries_ body model.BlazerDashboardQueries_ true "Update BlazerDashboardQueries_ record"
// @Success 200 {object} model.BlazerDashboardQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries_/{argID} [put]
// echo '{"id": 83,"dashboard_id": 66,"query_id": 60,"position": 37,"created_at": "2102-07-07T03:40:07.52114824-04:00","updated_at": "2254-04-24T07:25:54.102440438-04:00"}' | http PUT "http://localhost:8080/blazerdashboardqueries_/1"  X-Api-User:user123
func UpdateBlazerDashboardQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboardqueries_ := &model.BlazerDashboardQueries_{}
	if err := readJSON(r, blazerdashboardqueries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboardqueries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboardqueries_.Prepare()

	if err := blazerdashboardqueries_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboardqueries_, _, err = dao.UpdateBlazerDashboardQueries_(ctx,
		argID,
		blazerdashboardqueries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboardqueries_)
}

// DeleteBlazerDashboardQueries_ Delete a single record from blazer_dashboard_queries table in the rocket_development database
// @Summary Delete a record from blazer_dashboard_queries
// @Description Delete a single record from blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerDashboardQueries_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerdashboardqueries_/{argID} [delete]
// http DELETE "http://localhost:8080/blazerdashboardqueries_/1" X-Api-User:user123
func DeleteBlazerDashboardQueries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerDashboardQueries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
