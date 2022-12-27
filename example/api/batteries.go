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

func configBatteries_Router(router *httprouter.Router) {
	router.GET("/batteries_", GetAllBatteries_)
	router.POST("/batteries_", AddBatteries_)
	router.GET("/batteries_/:argID", GetBatteries_)
	router.PUT("/batteries_/:argID", UpdateBatteries_)
	router.DELETE("/batteries_/:argID", DeleteBatteries_)
}

func configGinBatteries_Router(router gin.IRoutes) {
	router.GET("/batteries_", ConverHttprouterToGin(GetAllBatteries_))
	router.POST("/batteries_", ConverHttprouterToGin(AddBatteries_))
	router.GET("/batteries_/:argID", ConverHttprouterToGin(GetBatteries_))
	router.PUT("/batteries_/:argID", ConverHttprouterToGin(UpdateBatteries_))
	router.DELETE("/batteries_/:argID", ConverHttprouterToGin(DeleteBatteries_))
}

// GetAllBatteries_ is a function to get a slice of record(s) from batteries table in the rocket_development database
// @Summary Get list of Batteries_
// @Tags Batteries_
// @Description GetAllBatteries_ is a handler to get a slice of record(s) from batteries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Batteries_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries_ [get]
// http "http://localhost:8080/batteries_?page=0&pagesize=20" X-Api-User:user123
func GetAllBatteries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "batteries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBatteries_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBatteries_ is a function to get a single record from the batteries table in the rocket_development database
// @Summary Get record from table Batteries_ by  argID
// @Tags Batteries_
// @ID argID
// @Description GetBatteries_ is a function to get a single record from the batteries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Batteries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /batteries_/{argID} [get]
// http "http://localhost:8080/batteries_/1" X-Api-User:user123
func GetBatteries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBatteries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBatteries_ add to add a single record to batteries table in the rocket_development database
// @Summary Add an record to batteries table
// @Description add to add a single record to batteries table in the rocket_development database
// @Tags Batteries_
// @Accept  json
// @Produce  json
// @Param Batteries_ body model.Batteries_ true "Add Batteries_"
// @Success 200 {object} model.Batteries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries_ [post]
// echo '{"employee_id": 41,"building_id": 46,"id": 86,"type": "mZnFXfqWngUGRonwMnJVsODNb","status": "yGmagtNAXxaHcmCZvYulDqVAm","commission_date": "2035-08-21T21:56:14.966533016-04:00","last_inspection_date": "2232-06-20T21:05:07.239068364-04:00","operations_cert": "BVHSmAKsUOrgNHQgDjxVoikRf","information": "EdONBaGRmQXBIVttuaTVwIDNK","notes": "PxhnRaCaBPUlWUAHahGErVNqc","created_at": "2303-11-26T14:23:31.212269427-05:00","updated_at": "2115-03-05T11:08:36.574922594-05:00"}' | http POST "http://localhost:8080/batteries_" X-Api-User:user123
func AddBatteries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	batteries_ := &model.Batteries_{}

	if err := readJSON(r, batteries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := batteries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	batteries_.Prepare()

	if err := batteries_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	batteries_, _, err = dao.AddBatteries_(ctx, batteries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, batteries_)
}

// UpdateBatteries_ Update a single record from batteries table in the rocket_development database
// @Summary Update an record in table batteries
// @Description Update a single record from batteries table in the rocket_development database
// @Tags Batteries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Batteries_ body model.Batteries_ true "Update Batteries_ record"
// @Success 200 {object} model.Batteries_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries_/{argID} [put]
// echo '{"employee_id": 41,"building_id": 46,"id": 86,"type": "mZnFXfqWngUGRonwMnJVsODNb","status": "yGmagtNAXxaHcmCZvYulDqVAm","commission_date": "2035-08-21T21:56:14.966533016-04:00","last_inspection_date": "2232-06-20T21:05:07.239068364-04:00","operations_cert": "BVHSmAKsUOrgNHQgDjxVoikRf","information": "EdONBaGRmQXBIVttuaTVwIDNK","notes": "PxhnRaCaBPUlWUAHahGErVNqc","created_at": "2303-11-26T14:23:31.212269427-05:00","updated_at": "2115-03-05T11:08:36.574922594-05:00"}' | http PUT "http://localhost:8080/batteries_/1"  X-Api-User:user123
func UpdateBatteries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	batteries_ := &model.Batteries_{}
	if err := readJSON(r, batteries_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := batteries_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	batteries_.Prepare()

	if err := batteries_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	batteries_, _, err = dao.UpdateBatteries_(ctx,
		argID,
		batteries_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, batteries_)
}

// DeleteBatteries_ Delete a single record from batteries table in the rocket_development database
// @Summary Delete a record from batteries
// @Description Delete a single record from batteries table in the rocket_development database
// @Tags Batteries_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Batteries_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /batteries_/{argID} [delete]
// http DELETE "http://localhost:8080/batteries_/1" X-Api-User:user123
func DeleteBatteries_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBatteries_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
