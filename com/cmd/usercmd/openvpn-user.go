package main

import (
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
	"ovpn-admin/com/cydata/db"
	"ovpn-admin/com/cydata/user"
)

const (
	version = "1.0.4"
)

var (
	dbPath = kingpin.Flag("db.path", "path do openvpn-user db").Default("./openvpn-user.db").String()

	dbInitCommand    = kingpin.Command("db-init", "Init db.")
	dbMigrateCommand = kingpin.Command("db-migrate", "STUB: Migrate db.")

	createCommand             = kingpin.Command("create", "Create user.")
	createCommandUserFlag     = createCommand.Flag("user", "Username.").Required().String()
	createCommandPasswordFlag = createCommand.Flag("password", "Password.").Required().String()

	deleteCommand              = kingpin.Command("delete", "Delete user.")
	deleteCommandUserForceFlag = deleteCommand.Flag("force", "delete from db.").Default("false").Bool()
	deleteCommandUserFlag      = deleteCommand.Flag("user", "Username.").Required().String()

	revokeCommand         = kingpin.Command("revoke", "Revoke user.")
	revokeCommandUserFlag = revokeCommand.Flag("user", "Username.").Required().String()

	restoreCommand         = kingpin.Command("restore", "Restore user.")
	restoreCommandUserFlag = restoreCommand.Flag("user", "Username.").Required().String()

	listCommand = kingpin.Command("list", "List active users.")
	listAll     = listCommand.Flag("all", "Show all users include revoked and deleted.").Default("false").Bool()

	checkCommand         = kingpin.Command("check", "check user existent.")
	checkCommandUserFlag = checkCommand.Flag("user", "Username.").Required().String()

	authCommand             = kingpin.Command("auth", "Auth user.")
	authCommandUserFlag     = authCommand.Flag("user", "Username.").Required().String()
	authCommandPasswordFlag = authCommand.Flag("password", "Password.").Required().String()

	changePasswordCommand             = kingpin.Command("change-password", "Change password")
	changePasswordCommandUserFlag     = changePasswordCommand.Flag("user", "Username.").Required().String()
	changePasswordCommandPasswordFlag = changePasswordCommand.Flag("password", "Password.").Required().String()

	debug = kingpin.Flag("debug", "Enable debug mode.").Default("false").Bool()
)

func main() {
	kingpin.Version(version)
	switch kingpin.Parse() {
	case listCommand.FullCommand():
		user.PrintUsers()
	case checkCommand.FullCommand():
		user.CheckUserExistent(*authCommandUserFlag)
	case authCommand.FullCommand():
		user.AuthUser(*authCommandUserFlag, *authCommandPasswordFlag)
	case dbInitCommand.FullCommand():
		db.InitDb()
	}
}
