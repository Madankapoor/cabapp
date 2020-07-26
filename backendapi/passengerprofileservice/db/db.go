package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/models"
	"github.com/gin-gonic/gin"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

//DB ...
type DB struct {
	*sql.DB
}

// MYSQL Docker -
// docker stop cabapp
// docker rm /cabapp
// docker run --name cabapp  -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=cabapp -e MYSQL_PASSWORD=cabapp -e MYSQL_DATABASE=passenger mysql:5.7.31

var db *gorm.DB

func dbinfo() string {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	return dataSource
}

//Init ...
func Init() {
	var err error
	db, err = ConnectDB(dbinfo())
	db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Passenger{})
	Migrate()
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

// Inject BD into gin
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// Migrate is used to run db migrations. We don't use gorm but a wraper lib over it.
func Migrate() {
	db, err := sql.Open("mysql", dbinfo())
	if err != nil {
		log.Fatal(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	m.Steps(2)
}

//GetDB ...
func GetDB() *gorm.DB {
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
