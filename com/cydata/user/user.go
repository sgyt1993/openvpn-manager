package user

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"os"
	"ovpn-admin/com/cydata/db"
	"text/tabwriter"
)

type User struct {
	Id       int
	Name     string
	Password string
	Revoked  bool
	Deleted  bool
}

type UserRole struct {
	Id       int
	Name     string
	Password string
	Revoked  bool
	Deleted  bool
	RoleId   *int
	RoleName *string
}

func CreateUser(username, password string) {
	if !CheckUserExistent(username) {
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		_, err := db.GetDb().Exec("INSERT INTO users(username, password) VALUES ($1, $2)", username, string(hash))
		checkErr(err)
		fmt.Printf("User %s created\n", username)
	} else {
		fmt.Printf("ERROR: User %s already registered\n", username)
	}

}

func DeleteUser(username string) {
	deleteQuery := "UPDATE users SET deleted = 1 WHERE username = $1"
	res, err := db.GetDb().Exec(deleteQuery, username)
	checkErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		if rowsAffected == 1 {
			fmt.Printf("User %s deleted\n", username)
		}
	} else {
		fmt.Printf("ERROR: due deleting user %s: %s\n", username, rowsErr)
	}
}

func RevokedUser(username string) {
	if !userDeleted(username) {
		res, err := db.GetDb().Exec("UPDATE users SET revoked = 1 WHERE username = $1", username)
		checkErr(err)
		if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
			if rowsAffected == 1 {
				fmt.Printf("User %s revoked\n", username)
			}
		} else {
			fmt.Printf("ERROR: due reoking user %s: %s\n", username, rowsErr)
		}
	}
}

func RestoreUser(username string) {
	if !userDeleted(username) {
		res, err := db.GetDb().Exec("UPDATE users SET revoked = 0 WHERE username = $1", username)
		checkErr(err)
		if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
			if rowsAffected == 1 {
				fmt.Printf("User %s restored\n", username)
			}
		} else {
			fmt.Printf("ERROR: due restoring user %s: %s\n", username, rowsErr)
		}
	}
}

func QueryByUserName(username string) (user User) {
	// we need to check if there is already such a user
	// return true if user exist
	rows, err := db.GetDb().Query("SELECT * FROM users WHERE username = $1", username)
	checkErr(err)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Revoked, &user.Deleted)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return user
}

func CheckUserExistent(username string) bool {
	// we need to check if there is already such a user
	// return true if user exist
	var c int
	_ = db.GetDb().QueryRow("SELECT count(*) FROM users WHERE username = $1", username).Scan(&c)
	if c == 1 {
		fmt.Printf("User %s exist\n", username)
		return true
	} else {
		return false
	}
}

func userDeleted(username string) bool {
	// return true if user marked as deleted
	u := User{}
	_ = db.GetDb().QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&u)
	if u.Deleted {
		fmt.Printf("User %s marked as deleted\n", username)
		return true
	} else {
		return false
	}
}

func userIsActive(username string) bool {
	// return true if user exist and not deleted and revoked
	u := User{}
	_ = db.GetDb().QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&u)
	if !u.Revoked && !u.Deleted {
		fmt.Printf("User %s is active\n", username)
		return true
	} else {
		fmt.Println("User may be deleted or revoked")
		return false
	}
}

func ListUsers() []User {
	condition := "WHERE deleted = 0 AND revoked = 0"
	var users []User
	query := "SELECT * FROM users " + condition
	rows, err := db.GetDb().Query(query)
	checkErr(err)

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Password, &u.Revoked, &u.Deleted)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

func ListUserRoles() []UserRole {
	var users []UserRole
	query := "SELECT u.id,u.username,u.password,u.revoked,u.deleted,r.id,r.role_name FROM users u " +
		"left join account_role ar on ar.account_id = u.id and u.deleted = 0 " +
		"left join role r on r.id = ar.role_id"
	rows, err := db.GetDb().Query(query)
	checkErr(err)

	for rows.Next() {
		u := UserRole{}
		err := rows.Scan(&u.Id, &u.Name, &u.Password, &u.Revoked, &u.Deleted, &u.RoleId, &u.RoleName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

func PrintUsers() {
	ul := ListUsers()
	if len(ul) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent|tabwriter.Debug)
		_, _ = fmt.Fprintln(w, "id\t username\t revoked\t deleted")
		for _, u := range ul {
			fmt.Fprintf(w, "%d\t %s\t %v\t %v\n", u.Id, u.Name, u.Revoked, u.Deleted)
		}
		_ = w.Flush()
	} else {
		fmt.Println("No users created yet")
	}
}

func ChangeUserPassword(username, password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err := db.GetDb().Exec("UPDATE users SET password = $1 WHERE username = $2", hash, username)
	checkErr(err)

	fmt.Println("Password changed")
}

func AuthUser(username, password string) {

	row := db.GetDb().QueryRow("select * from users where username = $1", username)
	u := User{}
	err := row.Scan(&u.Id, &u.Name, &u.Password, &u.Revoked, &u.Deleted)
	checkErr(err)

	if userIsActive(username) {
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err != nil {
			fmt.Println("Authorization failed")
			fmt.Println("Passwords mismatched")
			os.Exit(1)
		} else {
			fmt.Println("Authorization successful")
			os.Exit(0)
		}
	}
	fmt.Println("Authorization failed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
