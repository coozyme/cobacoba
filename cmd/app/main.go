package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_activityHttpDelivery "ddd-to-do-list/internal/delivery/handler"
	routers "ddd-to-do-list/internal/delivery/router"
	_activityRepo "ddd-to-do-list/internal/infrastructure/database/mysql/repository"
	_activityUcase "ddd-to-do-list/internal/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// if viper.GetBool(`debug`) {
	// 	log.Println("Service RUN on DEBUG mode")
	// }
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _activityHttpDelivery.InitMiddleware()
	e.Use(middL.CORS)
	ar := _activityRepo.NewMysqlActivityRepository(dbConn)
	au := _activityUcase.NewActivityUsecase(ar)
	_activityHttpDelivery.NewHandler(au)

	routers.Router(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
