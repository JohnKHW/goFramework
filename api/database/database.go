package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dbDriver string = "pqsql"
var dbHost string = "localhost"
var dbUser string = "root"
var dbPassword string = ""
var dbPort string = "3306"
var dbName string = ""

type IDriver interface {
	DriverDialector() gorm.Dialector
}
type PGSQLDriver struct {
	IDriver
}
type MYSQLDriver struct {
	IDriver
}
type SQLIDriver struct {
	IDriver
}
type SQLSDriver struct {
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

func GetDriver(driver string) IDriver {
	switch driver {
	case "postgres", "pgsql":
		return &PGSQLDriver{}
	case "sqlite", "sqli":
		return &SQLIDriver{}
	case "sqlserver", "mssql":
		return &SQLSDriver{}
	default:
		return &MYSQLDriver{}
	}
}

func DBInit() (*gorm.DB, error) {
	db, err := gorm.Open(GetDriver("dbDriver").DriverDialector(), &gorm.Config{})
	return db, err
}
