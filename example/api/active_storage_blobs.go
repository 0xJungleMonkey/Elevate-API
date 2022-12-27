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

func configActiveStorageBlobsRouter(router *httprouter.Router) {
	router.GET("/activestorageblobs", GetAllActiveStorageBlobs)
	router.POST("/activestorageblobs", AddActiveStorageBlobs)
	router.GET("/activestorageblobs/:argID", GetActiveStorageBlobs)
	router.PUT("/activestorageblobs/:argID", UpdateActiveStorageBlobs)
	router.DELETE("/activestorageblobs/:argID", DeleteActiveStorageBlobs)
}

func configGinActiveStorageBlobsRouter(router gin.IRoutes) {
	router.GET("/activestorageblobs", ConverHttprouterToGin(GetAllActiveStorageBlobs))
	router.POST("/activestorageblobs", ConverHttprouterToGin(AddActiveStorageBlobs))
	router.GET("/activestorageblobs/:argID", ConverHttprouterToGin(GetActiveStorageBlobs))
	router.PUT("/activestorageblobs/:argID", ConverHttprouterToGin(UpdateActiveStorageBlobs))
	router.DELETE("/activestorageblobs/:argID", ConverHttprouterToGin(DeleteActiveStorageBlobs))
}

// GetAllActiveStorageBlobs is a function to get a slice of record(s) from active_storage_blobs table in the rocket_development database
// @Summary Get list of ActiveStorageBlobs
// @Tags ActiveStorageBlobs
// @Description GetAllActiveStorageBlobs is a handler to get a slice of record(s) from active_storage_blobs table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ActiveStorageBlobs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageblobs [get]
// http "http://localhost:8080/activestorageblobs?page=0&pagesize=20" X-Api-User:user123
func GetAllActiveStorageBlobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "active_storage_blobs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllActiveStorageBlobs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetActiveStorageBlobs is a function to get a single record from the active_storage_blobs table in the rocket_development database
// @Summary Get record from table ActiveStorageBlobs by  argID
// @Tags ActiveStorageBlobs
// @ID argID
// @Description GetActiveStorageBlobs is a function to get a single record from the active_storage_blobs table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ActiveStorageBlobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /activestorageblobs/{argID} [get]
// http "http://localhost:8080/activestorageblobs/1" X-Api-User:user123
func GetActiveStorageBlobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_blobs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetActiveStorageBlobs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddActiveStorageBlobs add to add a single record to active_storage_blobs table in the rocket_development database
// @Summary Add an record to active_storage_blobs table
// @Description add to add a single record to active_storage_blobs table in the rocket_development database
// @Tags ActiveStorageBlobs
// @Accept  json
// @Produce  json
// @Param ActiveStorageBlobs body model.ActiveStorageBlobs true "Add ActiveStorageBlobs"
// @Success 200 {object} model.ActiveStorageBlobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageblobs [post]
// echo '{"id": 94,"key": "nYwThBYfiIdMXjdcZVFduLEoi","filename": "jybMAUIhMGhBUxrXaTwjvLnEC","content_type": "tDnbcjZXaywQOXvqgtEdpOBpY","metadata": "flWceGtWxKhTquaHMHtYJXsuo","byte_size": 41,"checksum": "PMSlMMyLXHXZliPKdWKvuiveJ","created_at": "2038-02-27T19:56:22.763420229-05:00"}' | http POST "http://localhost:8080/activestorageblobs" X-Api-User:user123
func AddActiveStorageBlobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	activestorageblobs := &model.ActiveStorageBlobs{}

	if err := readJSON(r, activestorageblobs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activestorageblobs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activestorageblobs.Prepare()

	if err := activestorageblobs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_blobs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	activestorageblobs, _, err = dao.AddActiveStorageBlobs(ctx, activestorageblobs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activestorageblobs)
}

// UpdateActiveStorageBlobs Update a single record from active_storage_blobs table in the rocket_development database
// @Summary Update an record in table active_storage_blobs
// @Description Update a single record from active_storage_blobs table in the rocket_development database
// @Tags ActiveStorageBlobs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ActiveStorageBlobs body model.ActiveStorageBlobs true "Update ActiveStorageBlobs record"
// @Success 200 {object} model.ActiveStorageBlobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /activestorageblobs/{argID} [put]
// echo '{"id": 94,"key": "nYwThBYfiIdMXjdcZVFduLEoi","filename": "jybMAUIhMGhBUxrXaTwjvLnEC","content_type": "tDnbcjZXaywQOXvqgtEdpOBpY","metadata": "flWceGtWxKhTquaHMHtYJXsuo","byte_size": 41,"checksum": "PMSlMMyLXHXZliPKdWKvuiveJ","created_at": "2038-02-27T19:56:22.763420229-05:00"}' | http PUT "http://localhost:8080/activestorageblobs/1"  X-Api-User:user123
func UpdateActiveStorageBlobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activestorageblobs := &model.ActiveStorageBlobs{}
	if err := readJSON(r, activestorageblobs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := activestorageblobs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	activestorageblobs.Prepare()

	if err := activestorageblobs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_blobs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	activestorageblobs, _, err = dao.UpdateActiveStorageBlobs(ctx,
		argID,
		activestorageblobs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, activestorageblobs)
}

// DeleteActiveStorageBlobs Delete a single record from active_storage_blobs table in the rocket_development database
// @Summary Delete a record from active_storage_blobs
// @Description Delete a single record from active_storage_blobs table in the rocket_development database
// @Tags ActiveStorageBlobs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ActiveStorageBlobs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /activestorageblobs/{argID} [delete]
// http DELETE "http://localhost:8080/activestorageblobs/1" X-Api-User:user123
func DeleteActiveStorageBlobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "active_storage_blobs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteActiveStorageBlobs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
