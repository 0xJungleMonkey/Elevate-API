package dao

import (
	"context"
	"time"

	"restapi-golang-gin-gen/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllBlazerDashboards_ is a function to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBlazerDashboards_(ctx context.Context, page, pagesize int64, order string) (results []*model.BlazerDashboards_, totalRows int, err error) {

	resultOrm := DB.Model(&model.BlazerDashboards_{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetBlazerDashboards_ is a function to get a single record from the blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db Find error
func GetBlazerDashboards_(ctx context.Context, argID int64) (record *model.BlazerDashboards_, err error) {
	record = &model.BlazerDashboards_{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBlazerDashboards_ is a function to add a single record to blazer_dashboards table in the rocket_development database
// error - ErrInsertFailed, db save call failed
func AddBlazerDashboards_(ctx context.Context, record *model.BlazerDashboards_) (result *model.BlazerDashboards_, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBlazerDashboards_ is a function to update a single record from blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBlazerDashboards_(ctx context.Context, argID int64, updated *model.BlazerDashboards_) (result *model.BlazerDashboards_, RowsAffected int64, err error) {

	result = &model.BlazerDashboards_{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteBlazerDashboards_ is a function to delete a single record from blazer_dashboards table in the rocket_development database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBlazerDashboards_(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.BlazerDashboards_{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
