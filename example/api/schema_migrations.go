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

func configSchemaMigrations_Router(router *httprouter.Router) {
	router.GET("/schemamigrations_", GetAllSchemaMigrations_)
	router.POST("/schemamigrations_", AddSchemaMigrations_)
	router.GET("/schemamigrations_/:argVersion", GetSchemaMigrations_)
	router.PUT("/schemamigrations_/:argVersion", UpdateSchemaMigrations_)
	router.DELETE("/schemamigrations_/:argVersion", DeleteSchemaMigrations_)
}

func configGinSchemaMigrations_Router(router gin.IRoutes) {
	router.GET("/schemamigrations_", ConverHttprouterToGin(GetAllSchemaMigrations_))
	router.POST("/schemamigrations_", ConverHttprouterToGin(AddSchemaMigrations_))
	router.GET("/schemamigrations_/:argVersion", ConverHttprouterToGin(GetSchemaMigrations_))
	router.PUT("/schemamigrations_/:argVersion", ConverHttprouterToGin(UpdateSchemaMigrations_))
	router.DELETE("/schemamigrations_/:argVersion", ConverHttprouterToGin(DeleteSchemaMigrations_))
}

// GetAllSchemaMigrations_ is a function to get a slice of record(s) from schema_migrations table in the rocket_development database
// @Summary Get list of SchemaMigrations_
// @Tags SchemaMigrations_
// @Description GetAllSchemaMigrations_ is a handler to get a slice of record(s) from schema_migrations table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.SchemaMigrations_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations_ [get]
// http "http://localhost:8080/schemamigrations_?page=0&pagesize=20" X-Api-User:user123
func GetAllSchemaMigrations_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "schema_migrations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllSchemaMigrations_(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetSchemaMigrations_ is a function to get a single record from the schema_migrations table in the rocket_development database
// @Summary Get record from table SchemaMigrations_ by  argVersion
// @Tags SchemaMigrations_
// @ID argVersion
// @Description GetSchemaMigrations_ is a function to get a single record from the schema_migrations table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Success 200 {object} model.SchemaMigrations_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /schemamigrations_/{argVersion} [get]
// http "http://localhost:8080/schemamigrations_/hello world" X-Api-User:user123
func GetSchemaMigrations_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetSchemaMigrations_(ctx, argVersion)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddSchemaMigrations_ add to add a single record to schema_migrations table in the rocket_development database
// @Summary Add an record to schema_migrations table
// @Description add to add a single record to schema_migrations table in the rocket_development database
// @Tags SchemaMigrations_
// @Accept  json
// @Produce  json
// @Param SchemaMigrations_ body model.SchemaMigrations_ true "Add SchemaMigrations_"
// @Success 200 {object} model.SchemaMigrations_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations_ [post]
// echo '{"version": "PVRAgcYvPDrleEjACHPhPCScF"}' | http POST "http://localhost:8080/schemamigrations_" X-Api-User:user123
func AddSchemaMigrations_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	schemamigrations_ := &model.SchemaMigrations_{}

	if err := readJSON(r, schemamigrations_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := schemamigrations_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	schemamigrations_.Prepare()

	if err := schemamigrations_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	schemamigrations_, _, err = dao.AddSchemaMigrations_(ctx, schemamigrations_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, schemamigrations_)
}

// UpdateSchemaMigrations_ Update a single record from schema_migrations table in the rocket_development database
// @Summary Update an record in table schema_migrations
// @Description Update a single record from schema_migrations table in the rocket_development database
// @Tags SchemaMigrations_
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Param  SchemaMigrations_ body model.SchemaMigrations_ true "Update SchemaMigrations_ record"
// @Success 200 {object} model.SchemaMigrations_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations_/{argVersion} [put]
// echo '{"version": "PVRAgcYvPDrleEjACHPhPCScF"}' | http PUT "http://localhost:8080/schemamigrations_/hello world"  X-Api-User:user123
func UpdateSchemaMigrations_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	schemamigrations_ := &model.SchemaMigrations_{}
	if err := readJSON(r, schemamigrations_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := schemamigrations_.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	schemamigrations_.Prepare()

	if err := schemamigrations_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	schemamigrations_, _, err = dao.UpdateSchemaMigrations_(ctx,
		argVersion,
		schemamigrations_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, schemamigrations_)
}

// DeleteSchemaMigrations_ Delete a single record from schema_migrations table in the rocket_development database
// @Summary Delete a record from schema_migrations
// @Description Delete a single record from schema_migrations table in the rocket_development database
// @Tags SchemaMigrations_
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Success 204 {object} model.SchemaMigrations_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schemamigrations_/{argVersion} [delete]
// http DELETE "http://localhost:8080/schemamigrations_/hello world" X-Api-User:user123
func DeleteSchemaMigrations_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteSchemaMigrations_(ctx, argVersion)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
