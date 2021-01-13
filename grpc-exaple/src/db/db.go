package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"

	_redis "github.com/go-redis/redis"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//InitDB ...
func InitDB() {
	dbinfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: sqlDB, Dialect: gorp.MySQLDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests

	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}

//RedisClient ...
var RedisClient *_redis.Client

//InitRedis ...
func InitRedis(params ...string) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	db, _ := strconv.Atoi(params[0])

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
