package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dbPath = kingpin.Flag("db.path", "path do openvpn-user db").Default("./openvpn.db").Envar("DB_PATH").String()
)

func InitDb() {
	_, err := GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS users(" +
			"id integer not null primary key autoincrement, " +
			"username string UNIQUE, password string, " +
			"revoked integer default 0, " +
			"deleted integer default 0" +
			")")
	CheckErr(err)

	// 创建account_role表
	_, err = GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS account_role(" +
			"role_id int not null DEFAULT 0 COMMENT '角色id'," +
			"account_id int not null DEFAULT 0 COMMENT '账号id'" +
			")")
	CheckErr(err)

	// 创建role表
	_, err = GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS role(" +
			"id integer not null primary key autoincrement, " +
			"role_name string UNIQUE" +
			")")
	CheckErr(err)

	//创建ccd表
	_, err = GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS ccd_route(" +
			"id integer not null primary key autoincrement, " +
			"address VARCHAR(100) not null DEFAULT '' COMMENT '地址'," +
			"mask VARCHAR(100) not null DEFAULT '' COMMENT '掩码'," +
			"description VARCHAR(100) not null DEFAULT '' COMMENT '描述'," +
			")")
	CheckErr(err)

	//创建ccdClientAddress表
	_, err = GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS ccd_client_address(" +
			"id integer not null primary key autoincrement, " +
			"account_id int not null DEFAULT 0 COMMENT '用户id'," +
			"client_address VARCHAR(100) not null DEFAULT '' COMMENT '区域code'," +
			"mask VARCHAR(100) not null DEFAULT '' COMMENT '掩码'" +
			")")
	CheckErr(err)

	//创建 role_ccdroute 表
	_, err = GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS role_ccdroute(" +
			"id integer not null primary key autoincrement, " +
			"role_id int not null DEFAULT 0 COMMENT '角色id'," +
			"ccd_route_id int not null DEFAULT 0 COMMENT '路由id'" +
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
