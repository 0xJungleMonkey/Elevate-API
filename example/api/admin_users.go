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

func configAdminUsersRouter(router *httprouter.Router) {
	router.GET("/adminusers", GetAllAdminUsers)
	router.POST("/adminusers", AddAdminUsers)
	router.GET("/adminusers/:argID", GetAdminUsers)
	router.PUT("/adminusers/:argID", UpdateAdminUsers)
	router.DELETE("/adminusers/:argID", DeleteAdminUsers)
}

func configGinAdminUsersRouter(router gin.IRoutes) {
	router.GET("/adminusers", ConverHttprouterToGin(GetAllAdminUsers))
	router.POST("/adminusers", ConverHttprouterToGin(AddAdminUsers))
	router.GET("/adminusers/:argID", ConverHttprouterToGin(GetAdminUsers))
	router.PUT("/adminusers/:argID", ConverHttprouterToGin(UpdateAdminUsers))
	router.DELETE("/adminusers/:argID", ConverHttprouterToGin(DeleteAdminUsers))
}

// GetAllAdminUsers is a function to get a slice of record(s) from admin_users table in the rocket_development database
// @Summary Get list of AdminUsers
// @Tags AdminUsers
// @Description GetAllAdminUsers is a handler to get a slice of record(s) from admin_users table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AdminUsers}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminusers [get]
// http "http://localhost:8080/adminusers?page=0&pagesize=20" X-Api-User:user123
func GetAllAdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "admin_users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAdminUsers(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetAdminUsers is a function to get a single record from the admin_users table in the rocket_development database
// @Summary Get record from table AdminUsers by  argID
// @Tags AdminUsers
// @ID argID
// @Description GetAdminUsers is a function to get a single record from the admin_users table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.AdminUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /adminusers/{argID} [get]
// http "http://localhost:8080/adminusers/1" X-Api-User:user123
func GetAdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_users", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAdminUsers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAdminUsers add to add a single record to admin_users table in the rocket_development database
// @Summary Add an record to admin_users table
// @Description add to add a single record to admin_users table in the rocket_development database
// @Tags AdminUsers
// @Accept  json
// @Produce  json
// @Param AdminUsers body model.AdminUsers true "Add AdminUsers"
// @Success 200 {object} model.AdminUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminusers [post]
// echo '{"id": 89,"email": "FfaEYJMOQWuRgXrqxxOlDZxvS","encrypted_password": "qBCiOEZoKtjuoJIUUGbjXsqAe","reset_password_token": "OlVBmYWyCWNAgcBlLFLrELBUV","reset_password_sent_at": "2215-02-05T14:01:44.963705025-05:00","remember_created_at": "2032-06-11T22:59:21.617071347-04:00","created_at": "2309-12-19T09:35:53.937234183-05:00","updated_at": "2287-07-31T15:03:08.954379317-04:00"}' | http POST "http://localhost:8080/adminusers" X-Api-User:user123
func AddAdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	adminusers := &model.AdminUsers{}

	if err := readJSON(r, adminusers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := adminusers.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	adminusers.Prepare()

	if err := adminusers.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_users", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	adminusers, _, err = dao.AddAdminUsers(ctx, adminusers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, adminusers)
}

// UpdateAdminUsers Update a single record from admin_users table in the rocket_development database
// @Summary Update an record in table admin_users
// @Description Update a single record from admin_users table in the rocket_development database
// @Tags AdminUsers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  AdminUsers body model.AdminUsers true "Update AdminUsers record"
// @Success 200 {object} model.AdminUsers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminusers/{argID} [put]
// echo '{"id": 89,"email": "FfaEYJMOQWuRgXrqxxOlDZxvS","encrypted_password": "qBCiOEZoKtjuoJIUUGbjXsqAe","reset_password_token": "OlVBmYWyCWNAgcBlLFLrELBUV","reset_password_sent_at": "2215-02-05T14:01:44.963705025-05:00","remember_created_at": "2032-06-11T22:59:21.617071347-04:00","created_at": "2309-12-19T09:35:53.937234183-05:00","updated_at": "2287-07-31T15:03:08.954379317-04:00"}' | http PUT "http://localhost:8080/adminusers/1"  X-Api-User:user123
func UpdateAdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	adminusers := &model.AdminUsers{}
	if err := readJSON(r, adminusers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := adminusers.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	adminusers.Prepare()

	if err := adminusers.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_users", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	adminusers, _, err = dao.UpdateAdminUsers(ctx,
		argID,
		adminusers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, adminusers)
}

// DeleteAdminUsers Delete a single record from admin_users table in the rocket_development database
// @Summary Delete a record from admin_users
// @Description Delete a single record from admin_users table in the rocket_development database
// @Tags AdminUsers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.AdminUsers
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /adminusers/{argID} [delete]
// http DELETE "http://localhost:8080/adminusers/1" X-Api-User:user123
func DeleteAdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_users", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAdminUsers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
