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


CREATE TABLE `blazer_checks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `creator_id` bigint DEFAULT NULL,
  `query_id` bigint DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `schedule` varchar(255) DEFAULT NULL,
  `emails` text,
  `slack_channels` text,
  `check_type` varchar(255) DEFAULT NULL,
  `message` text,
  `last_run_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_blazer_checks_on_creator_id` (`creator_id`),
  KEY `index_blazer_checks_on_query_id` (`query_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 9,    "creator_id": 71,    "query_id": 7,    "state": "crRNNIGcevJZjpPLSccRDOdme",    "schedule": "iIwHlFKttyxPYrhmuoPIwqFXl",    "emails": "vJQkEcIyhTbXCkfugrbsjoCpX",    "slack_channels": "GJTsMGnSSkbrIaMDFAxgAldAK",    "check_type": "xIsdZpENAnmNRgWOYMZEumeYm",    "message": "nvnwCTPFpIgeVsdktmiSyiYTi",    "last_run_at": "2255-09-05T06:19:59.846729943-04:00",    "created_at": "2268-03-31T11:53:04.773590824-04:00",    "updated_at": "2042-12-05T07:00:47.204514626-05:00"}



*/

// BlazerChecks_ struct is a row record of the blazer_checks table in the rocket_development database
type BlazerChecks_ struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] creator_id                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CreatorID null.Int `gorm:"column:creator_id;type:bigint;" json:"creator_id"`
	//[ 2] query_id                                       bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	QueryID null.Int `gorm:"column:query_id;type:bigint;" json:"query_id"`
	//[ 3] state                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	State null.String `gorm:"column:state;type:varchar;size:255;" json:"state"`
	//[ 4] schedule                                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Schedule null.String `gorm:"column:schedule;type:varchar;size:255;" json:"schedule"`
	//[ 5] emails                                         text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Emails null.String `gorm:"column:emails;type:text;size:65535;" json:"emails"`
	//[ 6] slack_channels                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	SlackChannels null.String `gorm:"column:slack_channels;type:text;size:65535;" json:"slack_channels"`
	//[ 7] check_type                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CheckType null.String `gorm:"column:check_type;type:varchar;size:255;" json:"check_type"`
	//[ 8] message                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Message null.String `gorm:"column:message;type:text;size:65535;" json:"message"`
	//[ 9] last_run_at                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	LastRunAt null.Time `gorm:"column:last_run_at;type:datetime;" json:"last_run_at"`
	//[10] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[11] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var blazer_checksTableInfo = &TableInfo{
	Name: "blazer_checks",
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
			Name:               "creator_id",
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
			GoFieldName:        "CreatorID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "creator_id",
			ProtobufFieldName:  "creator_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "query_id",
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
			GoFieldName:        "QueryID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "query_id",
			ProtobufFieldName:  "query_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "state",
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
			GoFieldName:        "State",
			GoFieldType:        "null.String",
			JSONFieldName:      "state",
			ProtobufFieldName:  "state",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "schedule",
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
			GoFieldName:        "Schedule",
			GoFieldType:        "null.String",
			JSONFieldName:      "schedule",
			ProtobufFieldName:  "schedule",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "emails",
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
			GoFieldName:        "Emails",
			GoFieldType:        "null.String",
			JSONFieldName:      "emails",
			ProtobufFieldName:  "emails",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "slack_channels",
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
			GoFieldName:        "SlackChannels",
			GoFieldType:        "null.String",
			JSONFieldName:      "slack_channels",
			ProtobufFieldName:  "slack_channels",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "check_type",
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
			GoFieldName:        "CheckType",
			GoFieldType:        "null.String",
			JSONFieldName:      "check_type",
			ProtobufFieldName:  "check_type",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "message",
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
			GoFieldName:        "Message",
			GoFieldType:        "null.String",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "last_run_at",
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
			GoFieldName:        "LastRunAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_run_at",
			ProtobufFieldName:  "last_run_at",
			ProtobufType:       "google.protobuf.Timestamp",
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
	},
}

// TableName sets the insert table name for this struct type
func (b *BlazerChecks_) TableName() string {
	return "blazer_checks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlazerChecks_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlazerChecks_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlazerChecks_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlazerChecks_) TableInfo() *TableInfo {
	return blazer_checksTableInfo
}
