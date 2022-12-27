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

func configBuildingDetails_Router(router *httprouter.Router) {
	router.GET("/buildingdetails_", GetAllBuildingDetails_)
	router.POST("/buildingdetails_", AddBuildingDetails_)
	router.GET("/buildingdetails_/:argID", GetBuildingDetails_)
	router.PUT("/buildingdetails_/:argID", UpdateBuildingDetails_)
	router.DELETE("/buildingdetails_/:argID", DeleteBuildingDetails_)
}

func configGinBuildingDetails_Router(router gin.IRoutes) {
	router.GET("/buildingdetails_", ConverHttprouterToGin(GetAllBuildingDetails_))
	router.POST("/buildingdetails_", ConverHttprouterToGin(AddBuildingDetails_))
	router.GET("/buildingdetails_/:argID", ConverHttprouterToGin(GetBuildingDetails_))
	router.PUT("/buildingdetails_/:argID", ConverHttprouterToGin(UpdateBuildingDetails_))
	router.DELETE("/buildingdetails_/:argID", ConverHttprouterToGin(DeleteBuildingDetails_))
}

// GetAllBuildingDetails_ is a function to get a slice of record(s) from building_details table in the rocket_development database
// @Summary Get list of BuildingDetails_
// @Tags BuildingDetails_
// @Description GetAllBuildingDetails_ is a handler to get a slice of record(s) from building_details table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BuildingDetails_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails_ [get]
// http "http://localhost:8080/buildingdetails_?page=0&pagesize=20" X-Api-User:user123
func GetAllBuildingDetails_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "building_details", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBuildingDetails_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBuildingDetails_ is a function to get a single record from the building_details table in the rocket_development database
// @Summary Get record from table BuildingDetails_ by  argID
// @Tags BuildingDetails_
// @ID argID
// @Description GetBuildingDetails_ is a function to get a single record from the building_details table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BuildingDetails_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /buildingdetails_/{argID} [get]
// http "http://localhost:8080/buildingdetails_/1" X-Api-User:user123
func GetBuildingDetails_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBuildingDetails_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBuildingDetails_ add to add a single record to building_details table in the rocket_development database
// @Summary Add an record to building_details table
// @Description add to add a single record to building_details table in the rocket_development database
// @Tags BuildingDetails_
// @Accept  json
// @Produce  json
// @Param BuildingDetails_ body model.BuildingDetails_ true "Add BuildingDetails_"
// @Success 200 {object} model.BuildingDetails_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails_ [post]
// echo '{"building_id": 32,"id": 43,"information_key": "mcqIsWqmIeHXTBFVPvWZtCPXK","value": "nWUeKMQoHkUAJsjeBuRnUXLTG","created_at": "2244-09-29T07:08:16.702936483-04:00","updated_at": "2038-01-19T05:44:49.149006966-05:00"}' | http POST "http://localhost:8080/buildingdetails_" X-Api-User:user123
func AddBuildingDetails_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	buildingdetails_ := &model.BuildingDetails_{}

	if err := readJSON(r, buildingdetails_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildingdetails_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildingdetails_.Prepare()

	if err := buildingdetails_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	buildingdetails_, _, err = dao.AddBuildingDetails_(ctx, buildingdetails_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildingdetails_)
}

// UpdateBuildingDetails_ Update a single record from building_details table in the rocket_development database
// @Summary Update an record in table building_details
// @Description Update a single record from building_details table in the rocket_development database
// @Tags BuildingDetails_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BuildingDetails_ body model.BuildingDetails_ true "Update BuildingDetails_ record"
// @Success 200 {object} model.BuildingDetails_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails_/{argID} [put]
// echo '{"building_id": 32,"id": 43,"information_key": "mcqIsWqmIeHXTBFVPvWZtCPXK","value": "nWUeKMQoHkUAJsjeBuRnUXLTG","created_at": "2244-09-29T07:08:16.702936483-04:00","updated_at": "2038-01-19T05:44:49.149006966-05:00"}' | http PUT "http://localhost:8080/buildingdetails_/1"  X-Api-User:user123
func UpdateBuildingDetails_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildingdetails_ := &model.BuildingDetails_{}
	if err := readJSON(r, buildingdetails_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildingdetails_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildingdetails_.Prepare()

	if err := buildingdetails_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildingdetails_, _, err = dao.UpdateBuildingDetails_(ctx,
		argID,
		buildingdetails_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildingdetails_)
}

// DeleteBuildingDetails_ Delete a single record from building_details table in the rocket_development database
// @Summary Delete a record from building_details
// @Description Delete a single record from building_details table in the rocket_development database
// @Tags BuildingDetails_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BuildingDetails_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /buildingdetails_/{argID} [delete]
// http DELETE "http://localhost:8080/buildingdetails_/1" X-Api-User:user123
func DeleteBuildingDetails_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuildingDetails_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
