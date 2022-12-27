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


CREATE TABLE `buildings` (
  `customer_id` bigint DEFAULT NULL,
  `address_id` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `FullNameOfBuildingAdmin` varchar(255) DEFAULT NULL,
  `EmailOfAdminOfBuilding` varchar(255) DEFAULT NULL,
  `PhoneNumOfBuildingAdmin` int DEFAULT NULL,
  `FullNameOfTechContactForBuilding` varchar(255) DEFAULT NULL,
  `TechContactEmailForBuilding` varchar(255) DEFAULT NULL,
  `TechContactPhoneForBuilding` int DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_buildings_on_address_id` (`address_id`),
  KEY `index_buildings_on_customer_id` (`customer_id`),
  CONSTRAINT `fk_rails_6dc7a885ab` FOREIGN KEY (`address_id`) REFERENCES `addresses` (`id`),
  CONSTRAINT `fk_rails_c29cbe7fb8` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "customer_id": 44,    "address_id": 4,    "id": 2,    "full_name_of_building_admin": "jjAUcjaLPmvdAhBJwIBgXZshd",    "email_of_admin_of_building": "XwQbXBbmtjgMGCJkkmmVXhtlx",    "phone_num_of_building_admin": 11,    "full_name_of_tech_contact_for_building": "vhLlNFlvocQcdjFSDseiAhLKj",    "tech_contact_email_for_building": "LWjXMnruwKBSDnWDXOioiuFlV",    "tech_contact_phone_for_building": 29,    "created_at": "2183-04-21T09:54:21.90221388-04:00",    "updated_at": "2116-12-14T20:39:39.841018286-05:00"}



*/

// Buildings_ struct is a row record of the buildings table in the rocket_development database
type Buildings_ struct {
	//[ 0] customer_id                                    bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CustomerID null.Int `gorm:"column:customer_id;type:bigint;" json:"customer_id"`
	//[ 1] address_id                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AddressID null.Int `gorm:"column:address_id;type:bigint;" json:"address_id"`
	//[ 2] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 3] FullNameOfBuildingAdmin                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullNameOfBuildingAdmin null.String `gorm:"column:FullNameOfBuildingAdmin;type:varchar;size:255;" json:"full_name_of_building_admin"`
	//[ 4] EmailOfAdminOfBuilding                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	EmailOfAdminOfBuilding null.String `gorm:"column:EmailOfAdminOfBuilding;type:varchar;size:255;" json:"email_of_admin_of_building"`
	//[ 5] PhoneNumOfBuildingAdmin                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PhoneNumOfBuildingAdmin null.Int `gorm:"column:PhoneNumOfBuildingAdmin;type:int;" json:"phone_num_of_building_admin"`
	//[ 6] FullNameOfTechContactForBuilding               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullNameOfTechContactForBuilding null.String `gorm:"column:FullNameOfTechContactForBuilding;type:varchar;size:255;" json:"full_name_of_tech_contact_for_building"`
	//[ 7] TechContactEmailForBuilding                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TechContactEmailForBuilding null.String `gorm:"column:TechContactEmailForBuilding;type:varchar;size:255;" json:"tech_contact_email_for_building"`
	//[ 8] TechContactPhoneForBuilding                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TechContactPhoneForBuilding null.Int `gorm:"column:TechContactPhoneForBuilding;type:int;" json:"tech_contact_phone_for_building"`
	//[ 9] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[10] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var buildingsTableInfo = &TableInfo{
	Name: "buildings",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "customer_id",
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
			GoFieldName:        "CustomerID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "customer_id",
			ProtobufFieldName:  "customer_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "address_id",
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
			GoFieldName:        "AddressID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "address_id",
			ProtobufFieldName:  "address_id",
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
			Name:               "FullNameOfBuildingAdmin",
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
			GoFieldName:        "FullNameOfBuildingAdmin",
			GoFieldType:        "null.String",
			JSONFieldName:      "full_name_of_building_admin",
			ProtobufFieldName:  "full_name_of_building_admin",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "EmailOfAdminOfBuilding",
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
			GoFieldName:        "EmailOfAdminOfBuilding",
			GoFieldType:        "null.String",
			JSONFieldName:      "email_of_admin_of_building",
			ProtobufFieldName:  "email_of_admin_of_building",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "PhoneNumOfBuildingAdmin",
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
			GoFieldName:        "PhoneNumOfBuildingAdmin",
			GoFieldType:        "null.Int",
			JSONFieldName:      "phone_num_of_building_admin",
			ProtobufFieldName:  "phone_num_of_building_admin",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "FullNameOfTechContactForBuilding",
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
			GoFieldName:        "FullNameOfTechContactForBuilding",
			GoFieldType:        "null.String",
			JSONFieldName:      "full_name_of_tech_contact_for_building",
			ProtobufFieldName:  "full_name_of_tech_contact_for_building",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "TechContactEmailForBuilding",
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
			GoFieldName:        "TechContactEmailForBuilding",
			GoFieldType:        "null.String",
			JSONFieldName:      "tech_contact_email_for_building",
			ProtobufFieldName:  "tech_contact_email_for_building",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "TechContactPhoneForBuilding",
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
			GoFieldName:        "TechContactPhoneForBuilding",
			GoFieldType:        "null.Int",
			JSONFieldName:      "tech_contact_phone_for_building",
			ProtobufFieldName:  "tech_contact_phone_for_building",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Buildings_) TableName() string {
	return "buildings"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Buildings_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Buildings_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Buildings_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Buildings_) TableInfo() *TableInfo {
	return buildingsTableInfo
}
