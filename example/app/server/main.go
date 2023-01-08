package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/droundy/goopt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"restapi-golang-gin-gen/api"
	"restapi-golang-gin-gen/dao"

	_ "restapi-golang-gin-gen/docs"
	"restapi-golang-gin-gen/model"
)

var (
	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string

	// OsSignal signal used to shutdown
	OsSignal chan os.Signal
)

// GinServer launch gin server
func GinServer() (err error) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api.ConfigGinRouter(router)
	// router.Run(":8080")
	router.Run(":80")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return
}

// @title Sample CRUD api for rocket_development db
// @version 1.0
// @description Sample CRUD api for rocket_development db
// @termsOfService

// @contact.name Me
// @contact.url http://me.com/terms.html
// @contact.email me@me.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	OsSignal = make(chan os.Signal, 1)

	// Define version information
	goopt.Version = fmt.Sprintf(
		`Application build information
  Build date      : %s
  Build number    : %s
  Git commit      : %s
  Runtime version : %s
  Built on OS     : %s
`, BuildDate, BuildNumber, LatestCommit, RuntimeVer, BuiltOnOs)
	goopt.Parse(nil)

	db, err := gorm.Open("mysql", "root@/rocket_development?parseTime=true")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(true)
	dao.DB = db

	db.AutoMigrate(
		&model.ActiveAdminComments{},
		&model.ActiveStorageAttachments{},
		&model.ActiveStorageBlobs{},
		&model.Addresses{},
		&model.AdminUsers{},
		&model.ArInternalMetadata_{},
		&model.Batteries_{},
		&model.BlazerAudits_{},
		&model.BlazerChecks_{},
		&model.BlazerDashboardQueries_{},
		&model.BlazerDashboards_{},
		&model.BlazerQueries_{},
		&model.BuildingDetails_{},
		&model.Buildings_{},
		&model.Columns_{},
		&model.Customers_{},
		&model.Elevators_{},
		&model.Employees{},
		&model.Interventions_{},
		&model.Leads{},
		&model.Maps_{},
		&model.Quotes{},
		&model.SchemaMigrations_{},
		&model.Users_{},
	)

	dao.Logger = func(ctx context.Context, sql string) {
		fmt.Printf("SQL: %s\n", sql)
	}

	go GinServer()
	LoopForever()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")

}
