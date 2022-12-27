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


CREATE TABLE `interventions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `author` varchar(255) DEFAULT NULL,
  `customer_id` int DEFAULT NULL,
  `building_id` int DEFAULT NULL,
  `battery_id` int DEFAULT NULL,
  `column_id` int DEFAULT NULL,
  `elevator_id` int DEFAULT NULL,
  `employee_id` int DEFAULT NULL,
  `start_datetime` datetime DEFAULT NULL,
  `end_datetime` datetime DEFAULT NULL,
  `result` varchar(255) DEFAULT NULL,
  `report` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 65,    "author": "KAnYaNnbOsMETHgRorXLarTfL",    "customer_id": 83,    "building_id": 66,    "battery_id": 87,    "column_id": 66,    "elevator_id": 84,    "employee_id": 62,    "start_datetime": "2078-04-19T19:56:42.25256109-04:00",    "end_datetime": "2133-01-30T05:31:22.685708736-05:00",    "result": "OuwZLFcJIuDNEigwnJFvIRXWv",    "report": "CttuQjQmffNkWpnQFTKCvZrlB",    "status": "KKypUFNTPhjaHbwMeDeftJCtd",    "created_at": "2206-06-15T13:57:41.061379409-04:00",    "updated_at": "2024-09-17T14:03:42.985744518-04:00"}



*/

// Interventions_ struct is a row record of the interventions table in the rocket_development database
type Interventions_ struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] author                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Author null.String `gorm:"column:author;type:varchar;size:255;" json:"author"`
	//[ 2] customer_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CustomerID null.Int `gorm:"column:customer_id;type:int;" json:"customer_id"`
	//[ 3] building_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BuildingID null.Int `gorm:"column:building_id;type:int;" json:"building_id"`
	//[ 4] battery_id                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BatteryID null.Int `gorm:"column:battery_id;type:int;" json:"battery_id"`
	//[ 5] column_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ColumnID null.Int `gorm:"column:column_id;type:int;" json:"column_id"`
	//[ 6] elevator_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ElevatorID null.Int `gorm:"column:elevator_id;type:int;" json:"elevator_id"`
	//[ 7] employee_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	EmployeeID null.Int `gorm:"column:employee_id;type:int;" json:"employee_id"`
	//[ 8] start_datetime                                 datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	StartDatetime null.Time `gorm:"column:start_datetime;type:datetime;" json:"start_datetime"`
	//[ 9] end_datetime                                   datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	EndDatetime null.Time `gorm:"column:end_datetime;type:datetime;" json:"end_datetime"`
	//[10] result                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Result null.String `gorm:"column:result;type:varchar;size:255;" json:"result"`
	//[11] report                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Report null.String `gorm:"column:report;type:varchar;size:255;" json:"report"`
	//[12] status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status null.String `gorm:"column:status;type:varchar;size:255;" json:"status"`
	//[13] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[14] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var interventionsTableInfo = &TableInfo{
	Name: "interventions",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "author",
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
			GoFieldName:        "Author",
			GoFieldType:        "null.String",
			JSONFieldName:      "author",
			ProtobufFieldName:  "author",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "customer_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CustomerID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "customer_id",
			ProtobufFieldName:  "customer_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "building_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BuildingID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "building_id",
			ProtobufFieldName:  "building_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "battery_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BatteryID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "battery_id",
			ProtobufFieldName:  "battery_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "column_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColumnID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "column_id",
			ProtobufFieldName:  "column_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "elevator_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ElevatorID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "elevator_id",
			ProtobufFieldName:  "elevator_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "employee_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EmployeeID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "employee_id",
			ProtobufFieldName:  "employee_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "start_datetime",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "StartDatetime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "start_datetime",
			ProtobufFieldName:  "start_datetime",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "end_datetime",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "EndDatetime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "end_datetime",
			ProtobufFieldName:  "end_datetime",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "result",
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
			GoFieldName:        "Result",
			GoFieldType:        "null.String",
			JSONFieldName:      "result",
			ProtobufFieldName:  "result",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "report",
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
			GoFieldName:        "Report",
			GoFieldType:        "null.String",
			JSONFieldName:      "report",
			ProtobufFieldName:  "report",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *Interventions_) TableName() string {
	return "interventions"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *Interventions_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *Interventions_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *Interventions_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *Interventions_) TableInfo() *TableInfo {
	return interventionsTableInfo
}
