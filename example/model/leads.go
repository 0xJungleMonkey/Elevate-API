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


CREATE TABLE `leads` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `Full_name_of_the_contact` varchar(255) DEFAULT NULL,
  `Bussiness_name` varchar(255) DEFAULT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Phone` varchar(255) DEFAULT NULL,
  `Project_name` varchar(255) DEFAULT NULL,
  `Project_description` varchar(255) DEFAULT NULL,
  `Department_incharge` varchar(255) DEFAULT NULL,
  `Message` varchar(255) DEFAULT NULL,
  `Attached_file` mediumblob,
  `Creation_date` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 34,    "full_name_of_the_contact": "VubQUclMrYnJdXEjigwVJYJpb",    "bussiness_name": "kmehEWaecfYpTnxbqsyjLgiPZ",    "email": "xIPxNLpBMEVJIYqwDvpYMsmAY",    "phone": "gJeiQZmqPfBfEvdORqmxAFZoS",    "project_name": "EnDHNvXkkfSELooLmqqwekxEX",    "project_description": "XfwcNpnoWSfZLLDIGWGFemTHx",    "department_incharge": "iZFyDVbwMIclhilMytscpMhyL",    "message": "srltjuVoYobsrQLNZmGVncWOw",    "attached_file": "GklaMxxFVQYvJz5QGyhgBEBaBhljMQUZOmIAJ15VH0ADPDxTGQpiXx8sCh8tXjo8KTI7Y0QrBz5hPUBjLT5jIFgMHg==",    "creation_date": "2314-02-22T11:24:02.085806613-05:00",    "created_at": "2187-09-09T23:29:36.135583678-04:00",    "updated_at": "2095-03-26T09:01:43.425372385-04:00"}



*/

// Leads struct is a row record of the leads table in the rocket_development database
type Leads struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] Full_name_of_the_contact                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullNameOfTheContact null.String `gorm:"column:Full_name_of_the_contact;type:varchar;size:255;" json:"full_name_of_the_contact"`
	//[ 2] Bussiness_name                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BussinessName null.String `gorm:"column:Bussiness_name;type:varchar;size:255;" json:"bussiness_name"`
	//[ 3] Email                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email null.String `gorm:"column:Email;type:varchar;size:255;" json:"email"`
	//[ 4] Phone                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone null.String `gorm:"column:Phone;type:varchar;size:255;" json:"phone"`
	//[ 5] Project_name                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectName null.String `gorm:"column:Project_name;type:varchar;size:255;" json:"project_name"`
	//[ 6] Project_description                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectDescription null.String `gorm:"column:Project_description;type:varchar;size:255;" json:"project_description"`
	//[ 7] Department_incharge                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DepartmentIncharge null.String `gorm:"column:Department_incharge;type:varchar;size:255;" json:"department_incharge"`
	//[ 8] Message                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Message null.String `gorm:"column:Message;type:varchar;size:255;" json:"message"`
	//[ 9] Attached_file                                  blob                 null: true   primary: false  isArray: false  auto: false  col: blob            len: -1      default: []
	AttachedFile []byte `gorm:"column:Attached_file;type:blob;" json:"attached_file"`
	//[10] Creation_date                                  datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreationDate null.Time `gorm:"column:Creation_date;type:datetime;" json:"creation_date"`
	//[11] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[12] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var leadsTableInfo = &TableInfo{
	Name: "leads",
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
			Name:               "Full_name_of_the_contact",
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
			GoFieldName:        "FullNameOfTheContact",
			GoFieldType:        "null.String",
			JSONFieldName:      "full_name_of_the_contact",
			ProtobufFieldName:  "full_name_of_the_contact",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "Bussiness_name",
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
			GoFieldName:        "BussinessName",
			GoFieldType:        "null.String",
			JSONFieldName:      "bussiness_name",
			ProtobufFieldName:  "bussiness_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "Email",
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
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "Phone",
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
			GoFieldName:        "Phone",
			GoFieldType:        "null.String",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "Project_name",
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
			GoFieldName:        "ProjectName",
			GoFieldType:        "null.String",
			JSONFieldName:      "project_name",
			ProtobufFieldName:  "project_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "Project_description",
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
			GoFieldName:        "ProjectDescription",
			GoFieldType:        "null.String",
			JSONFieldName:      "project_description",
			ProtobufFieldName:  "project_description",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "Department_incharge",
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
			GoFieldName:        "DepartmentIncharge",
			GoFieldType:        "null.String",
			JSONFieldName:      "department_incharge",
			ProtobufFieldName:  "department_incharge",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "Message",
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
			GoFieldName:        "Message",
			GoFieldType:        "null.String",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "Attached_file",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "blob",
			DatabaseTypePretty: "blob",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "blob",
			ColumnLength:       -1,
			GoFieldName:        "AttachedFile",
			GoFieldType:        "[]byte",
			JSONFieldName:      "attached_file",
			ProtobufFieldName:  "attached_file",
			ProtobufType:       "bytes",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "Creation_date",
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
			GoFieldName:        "CreationDate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "creation_date",
			ProtobufFieldName:  "creation_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *Leads) TableName() string {
	return "leads"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *Leads) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *Leads) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *Leads) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *Leads) TableInfo() *TableInfo {
	return leadsTableInfo
}
