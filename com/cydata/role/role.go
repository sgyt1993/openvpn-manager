package role

import (
	"fmt"
	"net/http"
	"ovpn-admin/com/cydata/ccdroute"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/db"
	"strconv"
)

type Role struct {
	Id        int                 `json:"id"`
	RoleName  string              `json:"roleName"`
	CcdRoutes []ccdroute.CcdRoute `json:"ccdRoutes"`
}

// @title    函数名称
// @description   函数的详细描述
// @auth      作者             时间（2019/6/18   10:57 ）
// @param     roleName        string         "角色名称"
// @return    返回参数名        bool           "是否正确"
func checkRoleExistent(roleName string) bool {
	// we need to check if there is already such a user
	// return true if user exist
	var c int
	_ = db.GetDb().QueryRow("SELECT count(*) FROM role WHERE role_name = $1", roleName).Scan(&c)
	if c == 1 {
		fmt.Printf("User %s exist\n", roleName)
		return true
	} else {
		return false
	}
}

func createRole(roleName string) (err error) {
	if !checkRoleExistent(roleName) {
		_, err := db.GetDb().Exec("INSERT INTO role(role_name) VALUES ($1)", roleName)
		db.CheckErr(err)
		fmt.Printf("role %s created\n", roleName)
	} else {
		return fmt.Errorf("ERROR: role %s already registered\n", roleName)
	}
	return err
}

func deleteRole(id int) (err error) {
	var deleteQuery = "DELETE FROM role WHERE id = $1"
	res, err := db.GetDb().Exec(deleteQuery, id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR: due deleting role %d: %s\n", id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("role id %s deleted\n", id)
		}
	}
	return err

}

func updateRole(r *Role) (err error) {
	var updateSql = "update role set  role_name = $1 where id = $2"
	res, err := db.GetDb().Exec(updateSql, r.RoleName, r.Id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR:update role %s: %s\n", r.RoleName, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("role %s update\n", r.RoleName)
		}
	}
	return err
}

func queryRoleAll() (roles []Role, err error) {
	var queryRoleAllSql = "select id,role_name from role"
	rows, err := db.GetDb().Query(queryRoleAllSql)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		u := Role{}
		err := rows.Scan(&u.Id, &u.RoleName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		roles = append(roles, u)
	}

	return roles, err
}

func Add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	roleName := req.Form.Get("roleName")
	if len(roleName) == 0 {
		commonresp.JsonRespFail(w, "roleName is not empty")
		return
	}

	err := createRole(roleName)
	commonresp.JudgeError(w, "create role", err)
}

func Del(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	roleIdStr := req.Form.Get("roleId")
	if len(roleIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	roleId, _ := strconv.Atoi(roleIdStr)
	err := deleteRole(roleId)
	commonresp.JudgeError(w, "del role", err)
}

func Update(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	idStr := req.Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		commonresp.JsonRespFail(w, "参数类型错误")
		return
	}

	roleName := req.Form.Get("roleName")
	role := Role{Id: id, RoleName: roleName}

	err = updateRole(&role)
	commonresp.JudgeError(w, "update role", err)
}

func Query(w http.ResponseWriter, req *http.Request) {
	roles, err := queryRoleAll()
	commonresp.JudgeError(w, roles, err)
}
