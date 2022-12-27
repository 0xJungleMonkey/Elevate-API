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

func configActiveAdminCommentsRouter(router *httprouter.Router) {
	router.GET("/activeadmincomments", GetAllActiveAdminComments)
	router.POST("/activeadmincomments", AddActiveAdminComments)
	router.GET("/activeadmincomments/:argID", GetActiveAdminComments)
	router.PUT("/activeadmincomments/:argID", UpdateActiveAdminComments)
	router.DELETE("/activeadmincomments/:argID", DeleteActiveAdminComments)
}

func configGinActiveAdminCommentsRouter(router gin.IRoutes) {
	router.GET("/activeadmincomments", ConverHttprouterToGin(GetAllActiveAdminComments))
	router.POST("/activeadmincomments", ConverHttprouterToGin(AddActiveAdminComments))
	router.GET("/activeadmincomments/:argID", ConverHttprouterToGin(GetActiveAdminComments))
	router.PUT("/activeadmincomments/:argID", ConverHttprouterToGin(UpdateActiveAdminComments))
	router.DELETE("/activeadmincomments/:argID", ConverHttprouterToGin(DeleteActiveAdminComments))
}

// GetAllActiveAdminComments is a function to get a slice of record(s) from active_admin_comments table in the rocket_development database
// @Summary Get list of ActiveAdminComments
// @Tags ActiveAdminComments
// @Description GetAllActiveAdminComments is a handler to get a slice of record(s) from active_admin_comments table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ActiveAdminComments}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activeadmincomments [get]
// http "http://localhost:8080/activeadmincomments?page=0&pagesize=20" X-Api-User:user123
func GetAllActiveAdminComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "active_admin_comments", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllActiveAdminComments(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetActiveAdminComments is a function to get a single record from the active_admin_comments table in the rocket_development database
// @Summary Get record from table ActiveAdminComments by  argID
// @Tags ActiveAdminComments
// @ID argID
// @Description GetActiveAdminComments is a function to get a single record from the active_admin_comments table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ActiveAdminComments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /activeadmincomments/{argID} [get]
// http "http://localhost:8080/activeadmincomments/1" X-Api-User:user123
func GetActiveAdminComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_admin_comments", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetActiveAdminComments(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddActiveAdminComments add to add a single record to active_admin_comments table in the rocket_development database
// @Summary Add an record to active_admin_comments table
// @Description add to add a single record to active_admin_comments table in the rocket_development database
// @Tags ActiveAdminComments
// @Accept  json
// @Produce  json
// @Param ActiveAdminComments body model.ActiveAdminComments true "Add ActiveAdminComments"
// @Success 200 {object} model.ActiveAdminComments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activeadmincomments [post]
// echo '{"id": 83,"namespace": "KuqghbLLPOvvJlraCcdgilifx","body": "ifNhmOwAACQpNsTQlKbjJJlmi","resource_type": "CotwxSaePnICBVTNWJFAbkQXH","resource_id": 17,"author_type": "ANSkUErAcKYogtwnSZvhiBtfO","author_id": 55,"created_at": "2206-08-21T04:09:58.084431087-04:00","updated_at": "2167-12-19T11:20:44.400834746-05:00"}' | http POST "http://localhost:8080/activeadmincomments" X-Api-User:user123
func AddActiveAdminComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	activeadmincomments := &model.ActiveAdminComments{}

	if err := readJSON(r, activeadmincomments); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activeadmincomments.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activeadmincomments.Prepare()

	if err := activeadmincomments.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_admin_comments", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	activeadmincomments, _, err = dao.AddActiveAdminComments(ctx, activeadmincomments)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activeadmincomments)
}

// UpdateActiveAdminComments Update a single record from active_admin_comments table in the rocket_development database
// @Summary Update an record in table active_admin_comments
// @Description Update a single record from active_admin_comments table in the rocket_development database
// @Tags ActiveAdminComments
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ActiveAdminComments body model.ActiveAdminComments true "Update ActiveAdminComments record"
// @Success 200 {object} model.ActiveAdminComments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activeadmincomments/{argID} [put]
// echo '{"id": 83,"namespace": "KuqghbLLPOvvJlraCcdgilifx","body": "ifNhmOwAACQpNsTQlKbjJJlmi","resource_type": "CotwxSaePnICBVTNWJFAbkQXH","resource_id": 17,"author_type": "ANSkUErAcKYogtwnSZvhiBtfO","author_id": 55,"created_at": "2206-08-21T04:09:58.084431087-04:00","updated_at": "2167-12-19T11:20:44.400834746-05:00"}' | http PUT "http://localhost:8080/activeadmincomments/1"  X-Api-User:user123
func UpdateActiveAdminComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activeadmincomments := &model.ActiveAdminComments{}
	if err := readJSON(r, activeadmincomments); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activeadmincomments.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activeadmincomments.Prepare()

	if err := activeadmincomments.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_admin_comments", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activeadmincomments, _, err = dao.UpdateActiveAdminComments(ctx,
		argID,
		activeadmincomments)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activeadmincomments)
}

// DeleteActiveAdminComments Delete a single record from active_admin_comments table in the rocket_development database
// @Summary Delete a record from active_admin_comments
// @Description Delete a single record from active_admin_comments table in the rocket_development database
// @Tags ActiveAdminComments
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ActiveAdminComments
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /activeadmincomments/{argID} [delete]
// http DELETE "http://localhost:8080/activeadmincomments/1" X-Api-User:user123
func DeleteActiveAdminComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_admin_comments", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteActiveAdminComments(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
