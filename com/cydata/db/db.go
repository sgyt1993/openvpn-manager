package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dbPath = kingpin.Flag("db.path", "path do openvpn db").Default("./openvpn.db").Envar("DB_PATH").String()
)

func InitDb() {
	db := GetDb()
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS users(" +
			"id integer not null primary key autoincrement, " +
			"username string UNIQUE, " +
			"password string, " +
			"revoked integer default 0, " +
			"deleted integer default 0" +
			")")
	CheckErr(err)

	// 创建account_role表
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS account_role(" +
			"id integer not null primary key autoincrement, " +
			"role_id integer not null DEFAULT 0, -- 角色id\n" +
			"account_id integer not null DEFAULT 0 -- 账号id\n" +
			")")
	CheckErr(err)

	// 创建role表
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS role(" +
			"id integer not null primary key autoincrement, " +
			"role_name string UNIQUE not null default '' -- 角色名称\n " +
			")")
	CheckErr(err)

	//创建ccd表
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS ccd_route(" +
			"id integer not null primary key autoincrement, " +
			"address VARCHAR(100) not null DEFAULT '', -- 地址 \n" +
			"mask VARCHAR(100) not null DEFAULT '', -- 掩码 \n" +
			"description VARCHAR(100) not null DEFAULT '' --描述 \n" +
			")")
	CheckErr(err)

	//创建ccdClientAddress表
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS ccd_client_address(" +
			"id integer not null primary key autoincrement, " +
			"account_id integer not null DEFAULT 0, -- 用户id\n" +
			"client_address VARCHAR(100) not null DEFAULT '', -- 区域code\n" +
			"mask VARCHAR(100) not null DEFAULT '' -- 掩码\n" +
			")")
	CheckErr(err)

	//创建 role_ccdroute 表
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS role_ccdroute(" +
			"id integer not null primary key autoincrement, " +
			"role_id integer not null DEFAULT 0, -- 角色id\n" +
			"ccd_route_id integer not null DEFAULT 0 -- 路由id\n" +
			")")
	CheckErr(err)

	fmt.Printf("Database initialized at %s\n", *dbPath)

}

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", *dbPath)
	CheckErr(err)
	if db == nil {
		panic("db is nil")
	}
	return db
}

func CheckErr(err error) {
	if err != nil {

		panic(err)
	}
}
