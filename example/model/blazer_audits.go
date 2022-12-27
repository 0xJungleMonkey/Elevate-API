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


CREATE TABLE `blazer_audits` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `query_id` bigint DEFAULT NULL,
  `statement` text,
  `data_source` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_blazer_audits_on_user_id` (`user_id`),
  KEY `index_blazer_audits_on_query_id` (`query_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 96,    "user_id": 76,    "query_id": 47,    "statement": "FeadqVcmKFJuGrZomvHXHeVWO",    "data_source": "MvmUyFXTlDwQOtsnFEAwGGkiW",    "created_at": "2249-05-10T02:56:46.439506764-04:00"}



*/

// BlazerAudits_ struct is a row record of the blazer_audits table in the rocket_development database
type BlazerAudits_ struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] user_id                                        bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	UserID null.Int `gorm:"column:user_id;type:bigint;" json:"user_id"`
	//[ 2] query_id                                       bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	QueryID null.Int `gorm:"column:query_id;type:bigint;" json:"query_id"`
	//[ 3] statement                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Statement null.String `gorm:"column:statement;type:text;size:65535;" json:"statement"`
	//[ 4] data_source                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DataSource null.String `gorm:"column:data_source;type:varchar;size:255;" json:"data_source"`
	//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
}

var blazer_auditsTableInfo = &TableInfo{
	Name: "blazer_audits",
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
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
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
			Name:               "statement",
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
			GoFieldName:        "Statement",
			GoFieldType:        "null.String",
			JSONFieldName:      "statement",
			ProtobufFieldName:  "statement",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "data_source",
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
			GoFieldName:        "DataSource",
			GoFieldType:        "null.String",
			JSONFieldName:      "data_source",
			ProtobufFieldName:  "data_source",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "created_at",
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
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BlazerAudits_) TableName() string {
	return "blazer_audits"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlazerAudits_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlazerAudits_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlazerAudits_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlazerAudits_) TableInfo() *TableInfo {
	return blazer_auditsTableInfo
}
