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

func configBlazerAudits_Router(router *httprouter.Router) {
	router.GET("/blazeraudits_", GetAllBlazerAudits_)
	router.POST("/blazeraudits_", AddBlazerAudits_)
	router.GET("/blazeraudits_/:argID", GetBlazerAudits_)
	router.PUT("/blazeraudits_/:argID", UpdateBlazerAudits_)
	router.DELETE("/blazeraudits_/:argID", DeleteBlazerAudits_)
}

func configGinBlazerAudits_Router(router gin.IRoutes) {
	router.GET("/blazeraudits_", ConverHttprouterToGin(GetAllBlazerAudits_))
	router.POST("/blazeraudits_", ConverHttprouterToGin(AddBlazerAudits_))
	router.GET("/blazeraudits_/:argID", ConverHttprouterToGin(GetBlazerAudits_))
	router.PUT("/blazeraudits_/:argID", ConverHttprouterToGin(UpdateBlazerAudits_))
	router.DELETE("/blazeraudits_/:argID", ConverHttprouterToGin(DeleteBlazerAudits_))
}

// GetAllBlazerAudits_ is a function to get a slice of record(s) from blazer_audits table in the rocket_development database
// @Summary Get list of BlazerAudits_
// @Tags BlazerAudits_
// @Description GetAllBlazerAudits_ is a handler to get a slice of record(s) from blazer_audits table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerAudits_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits_ [get]
// http "http://localhost:8080/blazeraudits_?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerAudits_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_audits", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerAudits_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerAudits_ is a function to get a single record from the blazer_audits table in the rocket_development database
// @Summary Get record from table BlazerAudits_ by  argID
// @Tags BlazerAudits_
// @ID argID
// @Description GetBlazerAudits_ is a function to get a single record from the blazer_audits table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerAudits_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazeraudits_/{argID} [get]
// http "http://localhost:8080/blazeraudits_/1" X-Api-User:user123
func GetBlazerAudits_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerAudits_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerAudits_ add to add a single record to blazer_audits table in the rocket_development database
// @Summary Add an record to blazer_audits table
// @Description add to add a single record to blazer_audits table in the rocket_development database
// @Tags BlazerAudits_
// @Accept  json
// @Produce  json
// @Param BlazerAudits_ body model.BlazerAudits_ true "Add BlazerAudits_"
// @Success 200 {object} model.BlazerAudits_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits_ [post]
// echo '{"id": 96,"user_id": 76,"query_id": 47,"statement": "FeadqVcmKFJuGrZomvHXHeVWO","data_source": "MvmUyFXTlDwQOtsnFEAwGGkiW","created_at": "2249-05-10T02:56:46.439506764-04:00"}' | http POST "http://localhost:8080/blazeraudits_" X-Api-User:user123
func AddBlazerAudits_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazeraudits_ := &model.BlazerAudits_{}

	if err := readJSON(r, blazeraudits_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazeraudits_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazeraudits_.Prepare()

	if err := blazeraudits_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazeraudits_, _, err = dao.AddBlazerAudits_(ctx, blazeraudits_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazeraudits_)
}

// UpdateBlazerAudits_ Update a single record from blazer_audits table in the rocket_development database
// @Summary Update an record in table blazer_audits
// @Description Update a single record from blazer_audits table in the rocket_development database
// @Tags BlazerAudits_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerAudits_ body model.BlazerAudits_ true "Update BlazerAudits_ record"
// @Success 200 {object} model.BlazerAudits_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits_/{argID} [put]
// echo '{"id": 96,"user_id": 76,"query_id": 47,"statement": "FeadqVcmKFJuGrZomvHXHeVWO","data_source": "MvmUyFXTlDwQOtsnFEAwGGkiW","created_at": "2249-05-10T02:56:46.439506764-04:00"}' | http PUT "http://localhost:8080/blazeraudits_/1"  X-Api-User:user123
func UpdateBlazerAudits_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazeraudits_ := &model.BlazerAudits_{}
	if err := readJSON(r, blazeraudits_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazeraudits_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazeraudits_.Prepare()

	if err := blazeraudits_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazeraudits_, _, err = dao.UpdateBlazerAudits_(ctx,
		argID,
		blazeraudits_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazeraudits_)
}

// DeleteBlazerAudits_ Delete a single record from blazer_audits table in the rocket_development database
// @Summary Delete a record from blazer_audits
// @Description Delete a single record from blazer_audits table in the rocket_development database
// @Tags BlazerAudits_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerAudits_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazeraudits_/{argID} [delete]
// http DELETE "http://localhost:8080/blazeraudits_/1" X-Api-User:user123
func DeleteBlazerAudits_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerAudits_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
