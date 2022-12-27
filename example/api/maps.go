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

func configMaps_Router(router *httprouter.Router) {
	router.GET("/maps_", GetAllMaps_)
	router.POST("/maps_", AddMaps_)
	router.GET("/maps_/:argID", GetMaps_)
	router.PUT("/maps_/:argID", UpdateMaps_)
	router.DELETE("/maps_/:argID", DeleteMaps_)
}

func configGinMaps_Router(router gin.IRoutes) {
	router.GET("/maps_", ConverHttprouterToGin(GetAllMaps_))
	router.POST("/maps_", ConverHttprouterToGin(AddMaps_))
	router.GET("/maps_/:argID", ConverHttprouterToGin(GetMaps_))
	router.PUT("/maps_/:argID", ConverHttprouterToGin(UpdateMaps_))
	router.DELETE("/maps_/:argID", ConverHttprouterToGin(DeleteMaps_))
}

// GetAllMaps_ is a function to get a slice of record(s) from maps table in the rocket_development database
// @Summary Get list of Maps_
// @Tags Maps_
// @Description GetAllMaps_ is a handler to get a slice of record(s) from maps table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Maps_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps_ [get]
// http "http://localhost:8080/maps_?page=0&pagesize=20" X-Api-User:user123
func GetAllMaps_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "maps", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMaps_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMaps_ is a function to get a single record from the maps table in the rocket_development database
// @Summary Get record from table Maps_ by  argID
// @Tags Maps_
// @ID argID
// @Description GetMaps_ is a function to get a single record from the maps table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Maps_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /maps_/{argID} [get]
// http "http://localhost:8080/maps_/1" X-Api-User:user123
func GetMaps_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMaps_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMaps_ add to add a single record to maps table in the rocket_development database
// @Summary Add an record to maps table
// @Description add to add a single record to maps table in the rocket_development database
// @Tags Maps_
// @Accept  json
// @Produce  json
// @Param Maps_ body model.Maps_ true "Add Maps_"
// @Success 200 {object} model.Maps_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps_ [post]
// echo '{"id": 28,"created_at": "2092-07-09T18:33:23.327500336-04:00","updated_at": "2092-07-20T02:13:53.355945152-04:00"}' | http POST "http://localhost:8080/maps_" X-Api-User:user123
func AddMaps_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	maps_ := &model.Maps_{}

	if err := readJSON(r, maps_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := maps_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	maps_.Prepare()

	if err := maps_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	maps_, _, err = dao.AddMaps_(ctx, maps_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, maps_)
}

// UpdateMaps_ Update a single record from maps table in the rocket_development database
// @Summary Update an record in table maps
// @Description Update a single record from maps table in the rocket_development database
// @Tags Maps_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Maps_ body model.Maps_ true "Update Maps_ record"
// @Success 200 {object} model.Maps_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps_/{argID} [put]
// echo '{"id": 28,"created_at": "2092-07-09T18:33:23.327500336-04:00","updated_at": "2092-07-20T02:13:53.355945152-04:00"}' | http PUT "http://localhost:8080/maps_/1"  X-Api-User:user123
func UpdateMaps_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	maps_ := &model.Maps_{}
	if err := readJSON(r, maps_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := maps_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	maps_.Prepare()

	if err := maps_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	maps_, _, err = dao.UpdateMaps_(ctx,
		argID,
		maps_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, maps_)
}

// DeleteMaps_ Delete a single record from maps table in the rocket_development database
// @Summary Delete a record from maps
// @Description Delete a single record from maps table in the rocket_development database
// @Tags Maps_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Maps_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /maps_/{argID} [delete]
// http DELETE "http://localhost:8080/maps_/1" X-Api-User:user123
func DeleteMaps_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteMaps_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
