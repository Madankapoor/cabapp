package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func dbinfo() string {
	return fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}

//Init ...
func Init() {
	var err error
	db, err = ConnectDB(dbinfo())
	if err != nil {
		log.Fatal(err)
	}
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
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
		"file:///migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	m.Steps(2)
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
