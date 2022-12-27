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

func configColumns_Router(router *httprouter.Router) {
	router.GET("/columns_", GetAllColumns_)
	router.POST("/columns_", AddColumns_)
	router.GET("/columns_/:argID", GetColumns_)
	router.PUT("/columns_/:argID", UpdateColumns_)
	router.DELETE("/columns_/:argID", DeleteColumns_)
}

func configGinColumns_Router(router gin.IRoutes) {
	router.GET("/columns_", ConverHttprouterToGin(GetAllColumns_))
	router.POST("/columns_", ConverHttprouterToGin(AddColumns_))
	router.GET("/columns_/:argID", ConverHttprouterToGin(GetColumns_))
	router.PUT("/columns_/:argID", ConverHttprouterToGin(UpdateColumns_))
	router.DELETE("/columns_/:argID", ConverHttprouterToGin(DeleteColumns_))
}

// GetAllColumns_ is a function to get a slice of record(s) from columns table in the rocket_development database
// @Summary Get list of Columns_
// @Tags Columns_
// @Description GetAllColumns_ is a handler to get a slice of record(s) from columns table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Columns_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns_ [get]
// http "http://localhost:8080/columns_?page=0&pagesize=20" X-Api-User:user123
func GetAllColumns_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "columns", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllColumns_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetColumns_ is a function to get a single record from the columns table in the rocket_development database
// @Summary Get record from table Columns_ by  argID
// @Tags Columns_
// @ID argID
// @Description GetColumns_ is a function to get a single record from the columns table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Columns_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /columns_/{argID} [get]
// http "http://localhost:8080/columns_/1" X-Api-User:user123
func GetColumns_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetColumns_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddColumns_ add to add a single record to columns table in the rocket_development database
// @Summary Add an record to columns table
// @Description add to add a single record to columns table in the rocket_development database
// @Tags Columns_
// @Accept  json
// @Produce  json
// @Param Columns_ body model.Columns_ true "Add Columns_"
// @Success 200 {object} model.Columns_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns_ [post]
// echo '{"battery_id": 40,"id": 43,"type": "MRcsyTHJDkIxTBMdAESRNbZvJ","num_of_floors_served": 59,"status": "gaJxcRAnhcJwTmnrVLMAfGtwk","information": "xEijvYGMinapPhajtKeaumxcn","notes": "vBDTVUGsGLkZweRWuqpHoDBqX","created_at": "2183-11-01T07:11:42.860060882-04:00","updated_at": "2314-02-24T23:11:20.823796502-05:00"}' | http POST "http://localhost:8080/columns_" X-Api-User:user123
func AddColumns_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	columns_ := &model.Columns_{}

	if err := readJSON(r, columns_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := columns_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	columns_.Prepare()

	if err := columns_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	columns_, _, err = dao.AddColumns_(ctx, columns_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, columns_)
}

// UpdateColumns_ Update a single record from columns table in the rocket_development database
// @Summary Update an record in table columns
// @Description Update a single record from columns table in the rocket_development database
// @Tags Columns_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Columns_ body model.Columns_ true "Update Columns_ record"
// @Success 200 {object} model.Columns_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns_/{argID} [put]
// echo '{"battery_id": 40,"id": 43,"type": "MRcsyTHJDkIxTBMdAESRNbZvJ","num_of_floors_served": 59,"status": "gaJxcRAnhcJwTmnrVLMAfGtwk","information": "xEijvYGMinapPhajtKeaumxcn","notes": "vBDTVUGsGLkZweRWuqpHoDBqX","created_at": "2183-11-01T07:11:42.860060882-04:00","updated_at": "2314-02-24T23:11:20.823796502-05:00"}' | http PUT "http://localhost:8080/columns_/1"  X-Api-User:user123
func UpdateColumns_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	columns_ := &model.Columns_{}
	if err := readJSON(r, columns_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := columns_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	columns_.Prepare()

	if err := columns_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	columns_, _, err = dao.UpdateColumns_(ctx,
		argID,
		columns_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, columns_)
}

// DeleteColumns_ Delete a single record from columns table in the rocket_development database
// @Summary Delete a record from columns
// @Description Delete a single record from columns table in the rocket_development database
// @Tags Columns_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Columns_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /columns_/{argID} [delete]
// http DELETE "http://localhost:8080/columns_/1" X-Api-User:user123
func DeleteColumns_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteColumns_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
