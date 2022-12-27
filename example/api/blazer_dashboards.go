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

func configBlazerDashboards_Router(router *httprouter.Router) {
	router.GET("/blazerdashboards_", GetAllBlazerDashboards_)
	router.POST("/blazerdashboards_", AddBlazerDashboards_)
	router.GET("/blazerdashboards_/:argID", GetBlazerDashboards_)
	router.PUT("/blazerdashboards_/:argID", UpdateBlazerDashboards_)
	router.DELETE("/blazerdashboards_/:argID", DeleteBlazerDashboards_)
}

func configGinBlazerDashboards_Router(router gin.IRoutes) {
	router.GET("/blazerdashboards_", ConverHttprouterToGin(GetAllBlazerDashboards_))
	router.POST("/blazerdashboards_", ConverHttprouterToGin(AddBlazerDashboards_))
	router.GET("/blazerdashboards_/:argID", ConverHttprouterToGin(GetBlazerDashboards_))
	router.PUT("/blazerdashboards_/:argID", ConverHttprouterToGin(UpdateBlazerDashboards_))
	router.DELETE("/blazerdashboards_/:argID", ConverHttprouterToGin(DeleteBlazerDashboards_))
}

// GetAllBlazerDashboards_ is a function to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// @Summary Get list of BlazerDashboards_
// @Tags BlazerDashboards_
// @Description GetAllBlazerDashboards_ is a handler to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerDashboards_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards_ [get]
// http "http://localhost:8080/blazerdashboards_?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerDashboards_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerDashboards_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerDashboards_ is a function to get a single record from the blazer_dashboards table in the rocket_development database
// @Summary Get record from table BlazerDashboards_ by  argID
// @Tags BlazerDashboards_
// @ID argID
// @Description GetBlazerDashboards_ is a function to get a single record from the blazer_dashboards table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerDashboards_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerdashboards_/{argID} [get]
// http "http://localhost:8080/blazerdashboards_/1" X-Api-User:user123
func GetBlazerDashboards_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerDashboards_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerDashboards_ add to add a single record to blazer_dashboards table in the rocket_development database
// @Summary Add an record to blazer_dashboards table
// @Description add to add a single record to blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards_
// @Accept  json
// @Produce  json
// @Param BlazerDashboards_ body model.BlazerDashboards_ true "Add BlazerDashboards_"
// @Success 200 {object} model.BlazerDashboards_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards_ [post]
// echo '{"id": 76,"creator_id": 2,"name": "OJhLIJTHBAFwIWZcxwrnLosUn","created_at": "2033-08-03T17:59:33.147289267-04:00","updated_at": "2163-08-18T06:31:58.704720108-04:00"}' | http POST "http://localhost:8080/blazerdashboards_" X-Api-User:user123
func AddBlazerDashboards_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerdashboards_ := &model.BlazerDashboards_{}

	if err := readJSON(r, blazerdashboards_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboards_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboards_.Prepare()

	if err := blazerdashboards_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerdashboards_, _, err = dao.AddBlazerDashboards_(ctx, blazerdashboards_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboards_)
}

// UpdateBlazerDashboards_ Update a single record from blazer_dashboards table in the rocket_development database
// @Summary Update an record in table blazer_dashboards
// @Description Update a single record from blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerDashboards_ body model.BlazerDashboards_ true "Update BlazerDashboards_ record"
// @Success 200 {object} model.BlazerDashboards_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards_/{argID} [put]
// echo '{"id": 76,"creator_id": 2,"name": "OJhLIJTHBAFwIWZcxwrnLosUn","created_at": "2033-08-03T17:59:33.147289267-04:00","updated_at": "2163-08-18T06:31:58.704720108-04:00"}' | http PUT "http://localhost:8080/blazerdashboards_/1"  X-Api-User:user123
func UpdateBlazerDashboards_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboards_ := &model.BlazerDashboards_{}
	if err := readJSON(r, blazerdashboards_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboards_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboards_.Prepare()

	if err := blazerdashboards_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboards_, _, err = dao.UpdateBlazerDashboards_(ctx,
		argID,
		blazerdashboards_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboards_)
}

// DeleteBlazerDashboards_ Delete a single record from blazer_dashboards table in the rocket_development database
// @Summary Delete a record from blazer_dashboards
// @Description Delete a single record from blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerDashboards_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerdashboards_/{argID} [delete]
// http DELETE "http://localhost:8080/blazerdashboards_/1" X-Api-User:user123
func DeleteBlazerDashboards_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerDashboards_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
