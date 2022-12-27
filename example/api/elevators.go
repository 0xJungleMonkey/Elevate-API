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

func configElevators_Router(router *httprouter.Router) {
	router.GET("/elevators_", GetAllElevators_)
	router.POST("/elevators_", AddElevators_)
	router.GET("/elevators_/:argID", GetElevators_)
	router.PUT("/elevators_/:argID", UpdateElevators_)
	router.DELETE("/elevators_/:argID", DeleteElevators_)
}

func configGinElevators_Router(router gin.IRoutes) {
	router.GET("/elevators_", ConverHttprouterToGin(GetAllElevators_))
	router.POST("/elevators_", ConverHttprouterToGin(AddElevators_))
	router.GET("/elevators_/:argID", ConverHttprouterToGin(GetElevators_))
	router.PUT("/elevators_/:argID", ConverHttprouterToGin(UpdateElevators_))
	router.DELETE("/elevators_/:argID", ConverHttprouterToGin(DeleteElevators_))
}

// GetAllElevators_ is a function to get a slice of record(s) from elevators table in the rocket_development database
// @Summary Get list of Elevators_
// @Tags Elevators_
// @Description GetAllElevators_ is a handler to get a slice of record(s) from elevators table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Elevators_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators_ [get]
// http "http://localhost:8080/elevators_?page=0&pagesize=20" X-Api-User:user123
func GetAllElevators_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "elevators", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllElevators_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetElevators_ is a function to get a single record from the elevators table in the rocket_development database
// @Summary Get record from table Elevators_ by  argID
// @Tags Elevators_
// @ID argID
// @Description GetElevators_ is a function to get a single record from the elevators table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Elevators_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /elevators_/{argID} [get]
// http "http://localhost:8080/elevators_/1" X-Api-User:user123
func GetElevators_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetElevators_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddElevators_ add to add a single record to elevators table in the rocket_development database
// @Summary Add an record to elevators table
// @Description add to add a single record to elevators table in the rocket_development database
// @Tags Elevators_
// @Accept  json
// @Produce  json
// @Param Elevators_ body model.Elevators_ true "Add Elevators_"
// @Success 200 {object} model.Elevators_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators_ [post]
// echo '{"column_id": 58,"id": 94,"serial_number": 29,"model": "aTWVkrgnDpBTrjAaLFjmfjQuw","type": "KBsdoQXmoiPEJumhJfONxrhQb","status": "OutKALHimskroHgLbOdOlWHZs","commision_date": "2094-10-23T00:06:11.490859579-04:00","last_inspection_date": "2164-03-02T09:16:49.879178419-05:00","inspection_cert": "LBexhLjMQbjpHqJwqjLrQkpqP","information": "HsHBZJwoOjaeFNtsWwqSCNUUQ","notes": "luZCZfOtXhbHcYUcEVElUxGwm","created_at": "2210-12-04T23:23:38.463627408-05:00","updated_at": "2087-02-06T06:40:08.275658106-05:00"}' | http POST "http://localhost:8080/elevators_" X-Api-User:user123
func AddElevators_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	elevators_ := &model.Elevators_{}

	if err := readJSON(r, elevators_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := elevators_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	elevators_.Prepare()

	if err := elevators_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	elevators_, _, err = dao.AddElevators_(ctx, elevators_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, elevators_)
}

// UpdateElevators_ Update a single record from elevators table in the rocket_development database
// @Summary Update an record in table elevators
// @Description Update a single record from elevators table in the rocket_development database
// @Tags Elevators_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Elevators_ body model.Elevators_ true "Update Elevators_ record"
// @Success 200 {object} model.Elevators_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators_/{argID} [put]
// echo '{"column_id": 58,"id": 94,"serial_number": 29,"model": "aTWVkrgnDpBTrjAaLFjmfjQuw","type": "KBsdoQXmoiPEJumhJfONxrhQb","status": "OutKALHimskroHgLbOdOlWHZs","commision_date": "2094-10-23T00:06:11.490859579-04:00","last_inspection_date": "2164-03-02T09:16:49.879178419-05:00","inspection_cert": "LBexhLjMQbjpHqJwqjLrQkpqP","information": "HsHBZJwoOjaeFNtsWwqSCNUUQ","notes": "luZCZfOtXhbHcYUcEVElUxGwm","created_at": "2210-12-04T23:23:38.463627408-05:00","updated_at": "2087-02-06T06:40:08.275658106-05:00"}' | http PUT "http://localhost:8080/elevators_/1"  X-Api-User:user123
func UpdateElevators_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	elevators_ := &model.Elevators_{}
	if err := readJSON(r, elevators_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := elevators_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	elevators_.Prepare()

	if err := elevators_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	elevators_, _, err = dao.UpdateElevators_(ctx,
		argID,
		elevators_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, elevators_)
}

// DeleteElevators_ Delete a single record from elevators table in the rocket_development database
// @Summary Delete a record from elevators
// @Description Delete a single record from elevators table in the rocket_development database
// @Tags Elevators_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Elevators_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /elevators_/{argID} [delete]
// http DELETE "http://localhost:8080/elevators_/1" X-Api-User:user123
func DeleteElevators_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteElevators_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
