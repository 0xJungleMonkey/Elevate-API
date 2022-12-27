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

func configArInternalMetadata_Router(router *httprouter.Router) {
	router.GET("/arinternalmetadata_", GetAllArInternalMetadata_)
	router.POST("/arinternalmetadata_", AddArInternalMetadata_)
	router.GET("/arinternalmetadata_/:argKey", GetArInternalMetadata_)
	router.PUT("/arinternalmetadata_/:argKey", UpdateArInternalMetadata_)
	router.DELETE("/arinternalmetadata_/:argKey", DeleteArInternalMetadata_)
}

func configGinArInternalMetadata_Router(router gin.IRoutes) {
	router.GET("/arinternalmetadata_", ConverHttprouterToGin(GetAllArInternalMetadata_))
	router.POST("/arinternalmetadata_", ConverHttprouterToGin(AddArInternalMetadata_))
	router.GET("/arinternalmetadata_/:argKey", ConverHttprouterToGin(GetArInternalMetadata_))
	router.PUT("/arinternalmetadata_/:argKey", ConverHttprouterToGin(UpdateArInternalMetadata_))
	router.DELETE("/arinternalmetadata_/:argKey", ConverHttprouterToGin(DeleteArInternalMetadata_))
}

// GetAllArInternalMetadata_ is a function to get a slice of record(s) from ar_internal_metadata table in the rocket_development database
// @Summary Get list of ArInternalMetadata_
// @Tags ArInternalMetadata_
// @Description GetAllArInternalMetadata_ is a handler to get a slice of record(s) from ar_internal_metadata table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ArInternalMetadata_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata_ [get]
// http "http://localhost:8080/arinternalmetadata_?page=0&pagesize=20" X-Api-User:user123
func GetAllArInternalMetadata_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllArInternalMetadata_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetArInternalMetadata_ is a function to get a single record from the ar_internal_metadata table in the rocket_development database
// @Summary Get record from table ArInternalMetadata_ by  argKey
// @Tags ArInternalMetadata_
// @ID argKey
// @Description GetArInternalMetadata_ is a function to get a single record from the ar_internal_metadata table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Success 200 {object} model.ArInternalMetadata_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /arinternalmetadata_/{argKey} [get]
// http "http://localhost:8080/arinternalmetadata_/hello world" X-Api-User:user123
func GetArInternalMetadata_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetArInternalMetadata_(ctx, argKey)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddArInternalMetadata_ add to add a single record to ar_internal_metadata table in the rocket_development database
// @Summary Add an record to ar_internal_metadata table
// @Description add to add a single record to ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata_
// @Accept  json
// @Produce  json
// @Param ArInternalMetadata_ body model.ArInternalMetadata_ true "Add ArInternalMetadata_"
// @Success 200 {object} model.ArInternalMetadata_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata_ [post]
// echo '{"key": "XCFwLBcKuUmVkNvlmGPGxhHUs","value": "KGltpyOKDqFkXNAUBdynjmjbW","created_at": "2260-10-02T06:18:38.009009919-04:00","updated_at": "2035-04-08T00:31:51.454328451-04:00"}' | http POST "http://localhost:8080/arinternalmetadata_" X-Api-User:user123
func AddArInternalMetadata_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	arinternalmetadata_ := &model.ArInternalMetadata_{}

	if err := readJSON(r, arinternalmetadata_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := arinternalmetadata_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	arinternalmetadata_.Prepare()

	if err := arinternalmetadata_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	arinternalmetadata_, _, err = dao.AddArInternalMetadata_(ctx, arinternalmetadata_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, arinternalmetadata_)
}

// UpdateArInternalMetadata_ Update a single record from ar_internal_metadata table in the rocket_development database
// @Summary Update an record in table ar_internal_metadata
// @Description Update a single record from ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata_
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Param  ArInternalMetadata_ body model.ArInternalMetadata_ true "Update ArInternalMetadata_ record"
// @Success 200 {object} model.ArInternalMetadata_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata_/{argKey} [put]
// echo '{"key": "XCFwLBcKuUmVkNvlmGPGxhHUs","value": "KGltpyOKDqFkXNAUBdynjmjbW","created_at": "2260-10-02T06:18:38.009009919-04:00","updated_at": "2035-04-08T00:31:51.454328451-04:00"}' | http PUT "http://localhost:8080/arinternalmetadata_/hello world"  X-Api-User:user123
func UpdateArInternalMetadata_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	arinternalmetadata_ := &model.ArInternalMetadata_{}
	if err := readJSON(r, arinternalmetadata_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := arinternalmetadata_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	arinternalmetadata_.Prepare()

	if err := arinternalmetadata_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	arinternalmetadata_, _, err = dao.UpdateArInternalMetadata_(ctx,
		argKey,
		arinternalmetadata_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, arinternalmetadata_)
}

// DeleteArInternalMetadata_ Delete a single record from ar_internal_metadata table in the rocket_development database
// @Summary Delete a record from ar_internal_metadata
// @Description Delete a single record from ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata_
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Success 204 {object} model.ArInternalMetadata_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /arinternalmetadata_/{argKey} [delete]
// http DELETE "http://localhost:8080/arinternalmetadata_/hello world" X-Api-User:user123
func DeleteArInternalMetadata_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteArInternalMetadata_(ctx, argKey)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
