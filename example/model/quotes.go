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


CREATE TABLE `quotes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `building_type` varchar(255) DEFAULT NULL,
  `service_quality` varchar(255) DEFAULT NULL,
  `number_of_apartments` varchar(255) DEFAULT NULL,
  `number_of_floors` varchar(255) DEFAULT NULL,
  `number_of_businesses` varchar(255) DEFAULT NULL,
  `number_of_basements` varchar(255) DEFAULT NULL,
  `number_of_parking` varchar(255) DEFAULT NULL,
  `number_of_cages` varchar(255) DEFAULT NULL,
  `number_of_occupants` varchar(255) DEFAULT NULL,
  `number_of_hours` varchar(255) DEFAULT NULL,
  `number_of_elevators_needed` varchar(255) DEFAULT NULL,
  `price_per_unit` varchar(255) DEFAULT NULL,
  `elevator_price` varchar(255) DEFAULT NULL,
  `installation_fee` varchar(255) DEFAULT NULL,
  `final_price` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `department` varchar(255) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `project_description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 79,    "building_type": "ESAooBBcJNoyvQbDlvusUAUPo",    "service_quality": "jseOEGuRmPxLPFPJiULjtTJuB",    "number_of_apartments": "PsiYDnwUSVrbXYQlohUWDJvFd",    "number_of_floors": "tNSTNoaeoKxLmrSPYpKeGUabE",    "number_of_businesses": "tTVSoEfbYUAhqEVpCFZDjsNSd",    "number_of_basements": "OYGSXyPqXXjMVIDKfhMuaOfsF",    "number_of_parking": "cOtjYsUhlHdvRFlBUcJsRhktT",    "number_of_cages": "xivkgMhaeIQiIDVCRKSkKCstE",    "number_of_occupants": "CGLnLHjSIsGAnCnQQmrsolFpv",    "number_of_hours": "SfyUQjILLYAfqiPAUdGRnunrN",    "number_of_elevators_needed": "BlTAFebldTIBrGGLncPgVvgRN",    "price_per_unit": "VFtaNlqYrmoXPDIbxuYsrxDsq",    "elevator_price": "xedKEPGUTwClAijhJpKNolRnd",    "installation_fee": "KLIRnievMyKKjFCNWCcHcYIbY",    "final_price": "RAHntpWhokjnOeLSMgRHVLWRA",    "created_at": "2163-08-02T09:01:32.631839129-04:00",    "updated_at": "2107-10-22T00:57:07.232250514-04:00",    "name": "xWkkeZwWulukOLhqktxrqIqBc",    "company_name": "ZNQWWWDQxfUHmWSOKPvsaqxCV",    "email": "evXiBlPWXCPWDhnoLpYZRtMOW",    "phone": "oPcNhXCPirwSeUxhYtHEPhWNA",    "department": "JbjWtyBeeGRlYWVnkHtZZrQja",    "project_name": "yvHZFQnYFXHAyLAKpvHGHOnvO",    "project_description": "PLmmpUCsrXIsWnUFCLRsywZxG"}



*/

// Quotes struct is a row record of the quotes table in the rocket_development database
type Quotes struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] building_type                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BuildingType null.String `gorm:"column:building_type;type:varchar;size:255;" json:"building_type"`
	//[ 2] service_quality                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ServiceQuality null.String `gorm:"column:service_quality;type:varchar;size:255;" json:"service_quality"`
	//[ 3] number_of_apartments                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfApartments null.String `gorm:"column:number_of_apartments;type:varchar;size:255;" json:"number_of_apartments"`
	//[ 4] number_of_floors                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfFloors null.String `gorm:"column:number_of_floors;type:varchar;size:255;" json:"number_of_floors"`
	//[ 5] number_of_businesses                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfBusinesses null.String `gorm:"column:number_of_businesses;type:varchar;size:255;" json:"number_of_businesses"`
	//[ 6] number_of_basements                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfBasements null.String `gorm:"column:number_of_basements;type:varchar;size:255;" json:"number_of_basements"`
	//[ 7] number_of_parking                              varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfParking null.String `gorm:"column:number_of_parking;type:varchar;size:255;" json:"number_of_parking"`
	//[ 8] number_of_cages                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfCages null.String `gorm:"column:number_of_cages;type:varchar;size:255;" json:"number_of_cages"`
	//[ 9] number_of_occupants                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfOccupants null.String `gorm:"column:number_of_occupants;type:varchar;size:255;" json:"number_of_occupants"`
	//[10] number_of_hours                                varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfHours null.String `gorm:"column:number_of_hours;type:varchar;size:255;" json:"number_of_hours"`
	//[11] number_of_elevators_needed                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	NumberOfElevatorsNeeded null.String `gorm:"column:number_of_elevators_needed;type:varchar;size:255;" json:"number_of_elevators_needed"`
	//[12] price_per_unit                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PricePerUnit null.String `gorm:"column:price_per_unit;type:varchar;size:255;" json:"price_per_unit"`
	//[13] elevator_price                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ElevatorPrice null.String `gorm:"column:elevator_price;type:varchar;size:255;" json:"elevator_price"`
	//[14] installation_fee                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	InstallationFee null.String `gorm:"column:installation_fee;type:varchar;size:255;" json:"installation_fee"`
	//[15] final_price                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FinalPrice null.String `gorm:"column:final_price;type:varchar;size:255;" json:"final_price"`
	//[16] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[17] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	//[18] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name null.String `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[19] company_name                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CompanyName null.String `gorm:"column:company_name;type:varchar;size:255;" json:"company_name"`
	//[20] email                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email null.String `gorm:"column:email;type:varchar;size:255;" json:"email"`
	//[21] phone                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone null.String `gorm:"column:phone;type:varchar;size:255;" json:"phone"`
	//[22] department                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Department null.String `gorm:"column:department;type:varchar;size:255;" json:"department"`
	//[23] project_name                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectName null.String `gorm:"column:project_name;type:varchar;size:255;" json:"project_name"`
	//[24] project_description                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProjectDescription null.String `gorm:"column:project_description;type:varchar;size:255;" json:"project_description"`
}

var quotesTableInfo = &TableInfo{
	Name: "quotes",
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
			Name:               "building_type",
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
			GoFieldName:        "BuildingType",
			GoFieldType:        "null.String",
			JSONFieldName:      "building_type",
			ProtobufFieldName:  "building_type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "service_quality",
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
			GoFieldName:        "ServiceQuality",
			GoFieldType:        "null.String",
			JSONFieldName:      "service_quality",
			ProtobufFieldName:  "service_quality",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "number_of_apartments",
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
			GoFieldName:        "NumberOfApartments",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_apartments",
			ProtobufFieldName:  "number_of_apartments",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "number_of_floors",
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
			GoFieldName:        "NumberOfFloors",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_floors",
			ProtobufFieldName:  "number_of_floors",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "number_of_businesses",
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
			GoFieldName:        "NumberOfBusinesses",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_businesses",
			ProtobufFieldName:  "number_of_businesses",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "number_of_basements",
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
			GoFieldName:        "NumberOfBasements",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_basements",
			ProtobufFieldName:  "number_of_basements",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "number_of_parking",
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
			GoFieldName:        "NumberOfParking",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_parking",
			ProtobufFieldName:  "number_of_parking",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "number_of_cages",
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
			GoFieldName:        "NumberOfCages",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_cages",
			ProtobufFieldName:  "number_of_cages",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "number_of_occupants",
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
			GoFieldName:        "NumberOfOccupants",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_occupants",
			ProtobufFieldName:  "number_of_occupants",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "number_of_hours",
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
			GoFieldName:        "NumberOfHours",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_hours",
			ProtobufFieldName:  "number_of_hours",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "number_of_elevators_needed",
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
			GoFieldName:        "NumberOfElevatorsNeeded",
			GoFieldType:        "null.String",
			JSONFieldName:      "number_of_elevators_needed",
			ProtobufFieldName:  "number_of_elevators_needed",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "price_per_unit",
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
			GoFieldName:        "PricePerUnit",
			GoFieldType:        "null.String",
			JSONFieldName:      "price_per_unit",
			ProtobufFieldName:  "price_per_unit",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "elevator_price",
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
			GoFieldName:        "ElevatorPrice",
			GoFieldType:        "null.String",
			JSONFieldName:      "elevator_price",
			ProtobufFieldName:  "elevator_price",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "installation_fee",
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
			GoFieldName:        "InstallationFee",
			GoFieldType:        "null.String",
			JSONFieldName:      "installation_fee",
			ProtobufFieldName:  "installation_fee",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "final_price",
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
			GoFieldName:        "FinalPrice",
			GoFieldType:        "null.String",
			JSONFieldName:      "final_price",
			ProtobufFieldName:  "final_price",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "company_name",
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
			GoFieldName:        "CompanyName",
			GoFieldType:        "null.String",
			JSONFieldName:      "company_name",
			ProtobufFieldName:  "company_name",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "email",
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "phone",
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "department",
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
			GoFieldName:        "Department",
			GoFieldType:        "null.String",
			JSONFieldName:      "department",
			ProtobufFieldName:  "department",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "project_name",
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "project_description",
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (q *Quotes) TableName() string {
	return "quotes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *Quotes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *Quotes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *Quotes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *Quotes) TableInfo() *TableInfo {
	return quotesTableInfo
}
