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

func configBlazerChecks_Router(router *httprouter.Router) {
	router.GET("/blazerchecks_", GetAllBlazerChecks_)
	router.POST("/blazerchecks_", AddBlazerChecks_)
	router.GET("/blazerchecks_/:argID", GetBlazerChecks_)
	router.PUT("/blazerchecks_/:argID", UpdateBlazerChecks_)
	router.DELETE("/blazerchecks_/:argID", DeleteBlazerChecks_)
}

func configGinBlazerChecks_Router(router gin.IRoutes) {
	router.GET("/blazerchecks_", ConverHttprouterToGin(GetAllBlazerChecks_))
	router.POST("/blazerchecks_", ConverHttprouterToGin(AddBlazerChecks_))
	router.GET("/blazerchecks_/:argID", ConverHttprouterToGin(GetBlazerChecks_))
	router.PUT("/blazerchecks_/:argID", ConverHttprouterToGin(UpdateBlazerChecks_))
	router.DELETE("/blazerchecks_/:argID", ConverHttprouterToGin(DeleteBlazerChecks_))
}

// GetAllBlazerChecks_ is a function to get a slice of record(s) from blazer_checks table in the rocket_development database
// @Summary Get list of BlazerChecks_
// @Tags BlazerChecks_
// @Description GetAllBlazerChecks_ is a handler to get a slice of record(s) from blazer_checks table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerChecks_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks_ [get]
// http "http://localhost:8080/blazerchecks_?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerChecks_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_checks", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerChecks_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerChecks_ is a function to get a single record from the blazer_checks table in the rocket_development database
// @Summary Get record from table BlazerChecks_ by  argID
// @Tags BlazerChecks_
// @ID argID
// @Description GetBlazerChecks_ is a function to get a single record from the blazer_checks table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerChecks_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerchecks_/{argID} [get]
// http "http://localhost:8080/blazerchecks_/1" X-Api-User:user123
func GetBlazerChecks_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerChecks_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerChecks_ add to add a single record to blazer_checks table in the rocket_development database
// @Summary Add an record to blazer_checks table
// @Description add to add a single record to blazer_checks table in the rocket_development database
// @Tags BlazerChecks_
// @Accept  json
// @Produce  json
// @Param BlazerChecks_ body model.BlazerChecks_ true "Add BlazerChecks_"
// @Success 200 {object} model.BlazerChecks_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks_ [post]
// echo '{"id": 9,"creator_id": 71,"query_id": 7,"state": "crRNNIGcevJZjpPLSccRDOdme","schedule": "iIwHlFKttyxPYrhmuoPIwqFXl","emails": "vJQkEcIyhTbXCkfugrbsjoCpX","slack_channels": "GJTsMGnSSkbrIaMDFAxgAldAK","check_type": "xIsdZpENAnmNRgWOYMZEumeYm","message": "nvnwCTPFpIgeVsdktmiSyiYTi","last_run_at": "2255-09-05T06:19:59.846729943-04:00","created_at": "2268-03-31T11:53:04.773590824-04:00","updated_at": "2042-12-05T07:00:47.204514626-05:00"}' | http POST "http://localhost:8080/blazerchecks_" X-Api-User:user123
func AddBlazerChecks_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerchecks_ := &model.BlazerChecks_{}

	if err := readJSON(r, blazerchecks_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerchecks_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerchecks_.Prepare()

	if err := blazerchecks_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerchecks_, _, err = dao.AddBlazerChecks_(ctx, blazerchecks_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerchecks_)
}

// UpdateBlazerChecks_ Update a single record from blazer_checks table in the rocket_development database
// @Summary Update an record in table blazer_checks
// @Description Update a single record from blazer_checks table in the rocket_development database
// @Tags BlazerChecks_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerChecks_ body model.BlazerChecks_ true "Update BlazerChecks_ record"
// @Success 200 {object} model.BlazerChecks_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks_/{argID} [put]
// echo '{"id": 9,"creator_id": 71,"query_id": 7,"state": "crRNNIGcevJZjpPLSccRDOdme","schedule": "iIwHlFKttyxPYrhmuoPIwqFXl","emails": "vJQkEcIyhTbXCkfugrbsjoCpX","slack_channels": "GJTsMGnSSkbrIaMDFAxgAldAK","check_type": "xIsdZpENAnmNRgWOYMZEumeYm","message": "nvnwCTPFpIgeVsdktmiSyiYTi","last_run_at": "2255-09-05T06:19:59.846729943-04:00","created_at": "2268-03-31T11:53:04.773590824-04:00","updated_at": "2042-12-05T07:00:47.204514626-05:00"}' | http PUT "http://localhost:8080/blazerchecks_/1"  X-Api-User:user123
func UpdateBlazerChecks_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerchecks_ := &model.BlazerChecks_{}
	if err := readJSON(r, blazerchecks_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerchecks_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerchecks_.Prepare()

	if err := blazerchecks_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerchecks_, _, err = dao.UpdateBlazerChecks_(ctx,
		argID,
		blazerchecks_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerchecks_)
}

// DeleteBlazerChecks_ Delete a single record from blazer_checks table in the rocket_development database
// @Summary Delete a record from blazer_checks
// @Description Delete a single record from blazer_checks table in the rocket_development database
// @Tags BlazerChecks_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerChecks_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerchecks_/{argID} [delete]
// http DELETE "http://localhost:8080/blazerchecks_/1" X-Api-User:user123
func DeleteBlazerChecks_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerChecks_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
