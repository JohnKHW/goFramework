package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dbHost string = "localhost"
var dbUser string = "root"
var dbPassword string = ""
var dbPort string = "3306"
var dbName string = ""

type IDriver interface {
	DriverDialector() gorm.Dialector
}

type IPGSQLDriver interface {
	IDriver
}
type PGSQLDriver struct {
}

type IMYSQLDriver interface {
	IDriver
}
type MYSQLDriver struct {
}

type ISQLIDriver interface {
	IDriver
}
type SQLIDriver struct {
}

type ISQLSDriver interface {
	IDriver
}
type SQLSDriver struct {
}

func GetDriver(driver string) IDriver {
	switch driver {
	case "postgres":
		return &PGSQLDriver{}
	case "sqlite":
		return &SQLIDriver{}
	case "sqlserver":
		return &SQLSDriver{}
	default:
		return &MYSQLDriver{}
	}
}

func (PGSQLDriver) DriverDialector() gorm.Dialector {
	dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
}

func (MYSQLDriver) DriverDialector() gorm.Dialector {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	return mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
}

func (SQLIDriver) DriverDialector() gorm.Dialector {
	dsn := "gorm.db"
	return sqlite.Open(dsn)
}

func (SQLSDriver) DriverDialector() gorm.Dialector {
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	return sqlserver.Open(dsn)
}

func DBInit() (*gorm.DB, error) {
	db, err := gorm.Open(GetDriver("").DriverDialector(), &gorm.Config{})
	return db, err
}
