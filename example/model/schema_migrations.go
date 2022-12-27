package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `schema_migrations` (
  `version` varchar(255) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "version": "PVRAgcYvPDrleEjACHPhPCScF"}



*/

// SchemaMigrations_ struct is a row record of the schema_migrations table in the rocket_development database
type SchemaMigrations_ struct {
	//[ 0] version                                        varchar(255)         null: false  primary: true   isArray: false  auto: false  col: varchar         len: 255     default: []
	Version string `gorm:"primary_key;column:version;type:varchar;size:255;" json:"version"`
}

var schema_migrationsTableInfo = &TableInfo{
	Name: "schema_migrations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Version",
			GoFieldType:        "string",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SchemaMigrations_) TableName() string {
	return "schema_migrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SchemaMigrations_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SchemaMigrations_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SchemaMigrations_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SchemaMigrations_) TableInfo() *TableInfo {
	return schema_migrationsTableInfo
}
