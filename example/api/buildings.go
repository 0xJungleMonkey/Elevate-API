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

func configBuildings_Router(router *httprouter.Router) {
	router.GET("/buildings_", GetAllBuildings_)
	router.POST("/buildings_", AddBuildings_)
	router.GET("/buildings_/:argID", GetBuildings_)
	router.PUT("/buildings_/:argID", UpdateBuildings_)
	router.DELETE("/buildings_/:argID", DeleteBuildings_)
}

func configGinBuildings_Router(router gin.IRoutes) {
	router.GET("/buildings_", ConverHttprouterToGin(GetAllBuildings_))
	router.POST("/buildings_", ConverHttprouterToGin(AddBuildings_))
	router.GET("/buildings_/:argID", ConverHttprouterToGin(GetBuildings_))
	router.PUT("/buildings_/:argID", ConverHttprouterToGin(UpdateBuildings_))
	router.DELETE("/buildings_/:argID", ConverHttprouterToGin(DeleteBuildings_))
}

// GetAllBuildings_ is a function to get a slice of record(s) from buildings table in the rocket_development database
// @Summary Get list of Buildings_
// @Tags Buildings_
// @Description GetAllBuildings_ is a handler to get a slice of record(s) from buildings table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Buildings_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings_ [get]
// http "http://localhost:8080/buildings_?page=0&pagesize=20" X-Api-User:user123
func GetAllBuildings_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "buildings", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBuildings_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBuildings_ is a function to get a single record from the buildings table in the rocket_development database
// @Summary Get record from table Buildings_ by  argID
// @Tags Buildings_
// @ID argID
// @Description GetBuildings_ is a function to get a single record from the buildings table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Buildings_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /buildings_/{argID} [get]
// http "http://localhost:8080/buildings_/1" X-Api-User:user123
func GetBuildings_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBuildings_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBuildings_ add to add a single record to buildings table in the rocket_development database
// @Summary Add an record to buildings table
// @Description add to add a single record to buildings table in the rocket_development database
// @Tags Buildings_
// @Accept  json
// @Produce  json
// @Param Buildings_ body model.Buildings_ true "Add Buildings_"
// @Success 200 {object} model.Buildings_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings_ [post]
// echo '{"customer_id": 44,"address_id": 4,"id": 2,"full_name_of_building_admin": "jjAUcjaLPmvdAhBJwIBgXZshd","email_of_admin_of_building": "XwQbXBbmtjgMGCJkkmmVXhtlx","phone_num_of_building_admin": 11,"full_name_of_tech_contact_for_building": "vhLlNFlvocQcdjFSDseiAhLKj","tech_contact_email_for_building": "LWjXMnruwKBSDnWDXOioiuFlV","tech_contact_phone_for_building": 29,"created_at": "2183-04-21T09:54:21.90221388-04:00","updated_at": "2116-12-14T20:39:39.841018286-05:00"}' | http POST "http://localhost:8080/buildings_" X-Api-User:user123
func AddBuildings_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	buildings_ := &model.Buildings_{}

	if err := readJSON(r, buildings_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildings_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildings_.Prepare()

	if err := buildings_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	buildings_, _, err = dao.AddBuildings_(ctx, buildings_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildings_)
}

// UpdateBuildings_ Update a single record from buildings table in the rocket_development database
// @Summary Update an record in table buildings
// @Description Update a single record from buildings table in the rocket_development database
// @Tags Buildings_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Buildings_ body model.Buildings_ true "Update Buildings_ record"
// @Success 200 {object} model.Buildings_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings_/{argID} [put]
// echo '{"customer_id": 44,"address_id": 4,"id": 2,"full_name_of_building_admin": "jjAUcjaLPmvdAhBJwIBgXZshd","email_of_admin_of_building": "XwQbXBbmtjgMGCJkkmmVXhtlx","phone_num_of_building_admin": 11,"full_name_of_tech_contact_for_building": "vhLlNFlvocQcdjFSDseiAhLKj","tech_contact_email_for_building": "LWjXMnruwKBSDnWDXOioiuFlV","tech_contact_phone_for_building": 29,"created_at": "2183-04-21T09:54:21.90221388-04:00","updated_at": "2116-12-14T20:39:39.841018286-05:00"}' | http PUT "http://localhost:8080/buildings_/1"  X-Api-User:user123
func UpdateBuildings_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildings_ := &model.Buildings_{}
	if err := readJSON(r, buildings_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildings_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildings_.Prepare()

	if err := buildings_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildings_, _, err = dao.UpdateBuildings_(ctx,
		argID,
		buildings_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildings_)
}

// DeleteBuildings_ Delete a single record from buildings table in the rocket_development database
// @Summary Delete a record from buildings
// @Description Delete a single record from buildings table in the rocket_development database
// @Tags Buildings_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Buildings_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /buildings_/{argID} [delete]
// http DELETE "http://localhost:8080/buildings_/1" X-Api-User:user123
func DeleteBuildings_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuildings_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
