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

func configActiveStorageAttachmentsRouter(router *httprouter.Router) {
	router.GET("/activestorageattachments", GetAllActiveStorageAttachments)
	router.POST("/activestorageattachments", AddActiveStorageAttachments)
	router.GET("/activestorageattachments/:argID", GetActiveStorageAttachments)
	router.PUT("/activestorageattachments/:argID", UpdateActiveStorageAttachments)
	router.DELETE("/activestorageattachments/:argID", DeleteActiveStorageAttachments)
}

func configGinActiveStorageAttachmentsRouter(router gin.IRoutes) {
	router.GET("/activestorageattachments", ConverHttprouterToGin(GetAllActiveStorageAttachments))
	router.POST("/activestorageattachments", ConverHttprouterToGin(AddActiveStorageAttachments))
	router.GET("/activestorageattachments/:argID", ConverHttprouterToGin(GetActiveStorageAttachments))
	router.PUT("/activestorageattachments/:argID", ConverHttprouterToGin(UpdateActiveStorageAttachments))
	router.DELETE("/activestorageattachments/:argID", ConverHttprouterToGin(DeleteActiveStorageAttachments))
}

// GetAllActiveStorageAttachments is a function to get a slice of record(s) from active_storage_attachments table in the rocket_development database
// @Summary Get list of ActiveStorageAttachments
// @Tags ActiveStorageAttachments
// @Description GetAllActiveStorageAttachments is a handler to get a slice of record(s) from active_storage_attachments table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ActiveStorageAttachments}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageattachments [get]
// http "http://localhost:8080/activestorageattachments?page=0&pagesize=20" X-Api-User:user123
func GetAllActiveStorageAttachments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "active_storage_attachments", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllActiveStorageAttachments(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetActiveStorageAttachments is a function to get a single record from the active_storage_attachments table in the rocket_development database
// @Summary Get record from table ActiveStorageAttachments by  argID
// @Tags ActiveStorageAttachments
// @ID argID
// @Description GetActiveStorageAttachments is a function to get a single record from the active_storage_attachments table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ActiveStorageAttachments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /activestorageattachments/{argID} [get]
// http "http://localhost:8080/activestorageattachments/1" X-Api-User:user123
func GetActiveStorageAttachments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_attachments", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetActiveStorageAttachments(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddActiveStorageAttachments add to add a single record to active_storage_attachments table in the rocket_development database
// @Summary Add an record to active_storage_attachments table
// @Description add to add a single record to active_storage_attachments table in the rocket_development database
// @Tags ActiveStorageAttachments
// @Accept  json
// @Produce  json
// @Param ActiveStorageAttachments body model.ActiveStorageAttachments true "Add ActiveStorageAttachments"
// @Success 200 {object} model.ActiveStorageAttachments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageattachments [post]
// echo '{"id": 35,"name": "ohLYiHpphKyHZHCIsnLdnqnJC","record_type": "AkWEoKosxSvyMpuQWZkObDiSn","record_id": 3,"blob_id": 35,"created_at": "2309-03-16T11:22:12.141378952-04:00"}' | http POST "http://localhost:8080/activestorageattachments" X-Api-User:user123
func AddActiveStorageAttachments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	activestorageattachments := &model.ActiveStorageAttachments{}

	if err := readJSON(r, activestorageattachments); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activestorageattachments.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activestorageattachments.Prepare()

	if err := activestorageattachments.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_attachments", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	activestorageattachments, _, err = dao.AddActiveStorageAttachments(ctx, activestorageattachments)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activestorageattachments)
}

// UpdateActiveStorageAttachments Update a single record from active_storage_attachments table in the rocket_development database
// @Summary Update an record in table active_storage_attachments
// @Description Update a single record from active_storage_attachments table in the rocket_development database
// @Tags ActiveStorageAttachments
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ActiveStorageAttachments body model.ActiveStorageAttachments true "Update ActiveStorageAttachments record"
// @Success 200 {object} model.ActiveStorageAttachments
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageattachments/{argID} [put]
// echo '{"id": 35,"name": "ohLYiHpphKyHZHCIsnLdnqnJC","record_type": "AkWEoKosxSvyMpuQWZkObDiSn","record_id": 3,"blob_id": 35,"created_at": "2309-03-16T11:22:12.141378952-04:00"}' | http PUT "http://localhost:8080/activestorageattachments/1"  X-Api-User:user123
func UpdateActiveStorageAttachments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activestorageattachments := &model.ActiveStorageAttachments{}
	if err := readJSON(r, activestorageattachments); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activestorageattachments.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activestorageattachments.Prepare()

	if err := activestorageattachments.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_attachments", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activestorageattachments, _, err = dao.UpdateActiveStorageAttachments(ctx,
		argID,
		activestorageattachments)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activestorageattachments)
}

// DeleteActiveStorageAttachments Delete a single record from active_storage_attachments table in the rocket_development database
// @Summary Delete a record from active_storage_attachments
// @Description Delete a single record from active_storage_attachments table in the rocket_development database
// @Tags ActiveStorageAttachments
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ActiveStorageAttachments
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /activestorageattachments/{argID} [delete]
// http DELETE "http://localhost:8080/activestorageattachments/1" X-Api-User:user123
func DeleteActiveStorageAttachments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_attachments", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteActiveStorageAttachments(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
