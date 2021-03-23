package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type ConfigList struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
	ServerPort     string
}

func NewLocalDB(user, pass, dbName string) (*ConfigList, error) {
	config := &ConfigList{
		DbDriverName:   "mysql",
		DbName:         dbName,
		DbUserName:     user,
		DbUserPassword: pass,
		DbHost:         "127.0.0.1",
		DbPort:         "3306",
		ServerPort:     "8080",
	}

	return config, nil
}
func NewDB(file string) (*ConfigList, error) {
	if err := godotenv.Load(file); err != nil {
		return nil, err
	}
	config := &ConfigList{
		DbDriverName:   os.Getenv("DRIVER_NAME"),
		DbName:         os.Getenv("DB_NAME"),
		DbUserName:     os.Getenv("DB_USER"),
		DbUserPassword: os.Getenv("DB_PASSWORD"),
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
		ServerPort:     os.Getenv("PORT"),
	}
	return config, nil
}
func (c *ConfigList) Connect() (*gorm.DB, error) {
	// _ "github.com/go-sql-driver/mysql"

	DBMS := c.DbDriverName
	USER := c.DbUserName
	PASS := c.DbUserPassword
	PROTOCOL := "tcp(" + c.DbHost + ":" + c.DbPort + ")"
	DBNAME := c.DbName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func (c *ConfigList) NewMakeDB(tables ...interface{}) error {
	db, err := c.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(tables...)
	return nil
}
