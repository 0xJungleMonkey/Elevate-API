package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
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


CREATE TABLE `batteries` (
  `employee_id` bigint DEFAULT NULL,
  `building_id` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) DEFAULT NULL,
  `Status` varchar(255) DEFAULT NULL,
  `CommissionDate` date DEFAULT NULL,
  `LastInspectionDate` date DEFAULT NULL,
  `OperationsCert` varchar(255) DEFAULT NULL,
  `Information` text,
  `Notes` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_batteries_on_building_id` (`building_id`),
  KEY `index_batteries_on_employee_id` (`employee_id`),
  CONSTRAINT `fk_rails_ceeeaf55f7` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_rails_fc40470545` FOREIGN KEY (`building_id`) REFERENCES `buildings` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "employee_id": 41,    "building_id": 46,    "id": 86,    "type": "mZnFXfqWngUGRonwMnJVsODNb",    "status": "yGmagtNAXxaHcmCZvYulDqVAm",    "commission_date": "2035-08-21T21:56:14.966533016-04:00",    "last_inspection_date": "2232-06-20T21:05:07.239068364-04:00",    "operations_cert": "BVHSmAKsUOrgNHQgDjxVoikRf",    "information": "EdONBaGRmQXBIVttuaTVwIDNK",    "notes": "PxhnRaCaBPUlWUAHahGErVNqc",    "created_at": "2303-11-26T14:23:31.212269427-05:00",    "updated_at": "2115-03-05T11:08:36.574922594-05:00"}



*/

// Batteries_ struct is a row record of the batteries table in the rocket_development database
type Batteries_ struct {
	//[ 0] employee_id                                    bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	EmployeeID null.Int `gorm:"column:employee_id;type:bigint;" json:"employee_id"`
	//[ 1] building_id                                    bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	BuildingID null.Int `gorm:"column:building_id;type:bigint;" json:"building_id"`
	//[ 2] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 3] Type                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Type null.String `gorm:"column:Type;type:varchar;size:255;" json:"type"`
	//[ 4] Status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status null.String `gorm:"column:Status;type:varchar;size:255;" json:"status"`
	//[ 5] CommissionDate                                 date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	CommissionDate null.Time `gorm:"column:CommissionDate;type:date;" json:"commission_date"`
	//[ 6] LastInspectionDate                             date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	LastInspectionDate null.Time `gorm:"column:LastInspectionDate;type:date;" json:"last_inspection_date"`
	//[ 7] OperationsCert                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	OperationsCert null.String `gorm:"column:OperationsCert;type:varchar;size:255;" json:"operations_cert"`
	//[ 8] Information                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Information null.String `gorm:"column:Information;type:text;size:65535;" json:"information"`
	//[ 9] Notes                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Notes null.String `gorm:"column:Notes;type:text;size:65535;" json:"notes"`
	//[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	columns   Columns_  `gorm:"column:columns;type:Columns_;" json:"columns"`
}

var batteriesTableInfo = &TableInfo{
	Name: "batteries",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "employee_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "EmployeeID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "employee_id",
			ProtobufFieldName:  "employee_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "building_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "BuildingID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "building_id",
			ProtobufFieldName:  "building_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "Type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Type",
			GoFieldType:        "null.String",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "Status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Status",
			GoFieldType:        "null.String",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CommissionDate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "CommissionDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "commission_date",
			ProtobufFieldName:  "commission_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "LastInspectionDate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "LastInspectionDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_inspection_date",
			ProtobufFieldName:  "last_inspection_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "OperationsCert",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "OperationsCert",
			GoFieldType:        "null.String",
			JSONFieldName:      "operations_cert",
			ProtobufFieldName:  "operations_cert",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "Information",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Information",
			GoFieldType:        "null.String",
			JSONFieldName:      "information",
			ProtobufFieldName:  "information",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "Notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Notes",
			GoFieldType:        "null.String",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:    12,
			Name:     "columns",
			Comment:  ``,
			Notes:    ``,
			Nullable: false,

			IsPrimaryKey:    false,
			IsAutoIncrement: false,
			IsArray:         false,
			ColumnType:      "Columns_",
			ColumnLength:    -1,
			GoFieldName:     "columns",
			GoFieldType:     "Columns_",
			JSONFieldName:   "columns",

			ProtobufPos: 13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Batteries_) TableName() string {
	return "batteries"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Batteries_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Batteries_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Batteries_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Batteries_) TableInfo() *TableInfo {
	return batteriesTableInfo
}
