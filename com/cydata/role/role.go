package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

// @Summary 角色接口
// @title 角色删除
// @Tags 角色删除
// @param roleName query string true "角色名称"
// @Accept application/x-www-form-urlencoded
// @Produce application/x-www-form-urlencoded
// @Router /api/role/add [post]
// @Success 200 {object} commonresp.JsonResult
func Add(c *gin.Context) {
	roleName, flag := c.GetQuery("roleName")
	if flag {
		c.JSON(http.StatusOK, commonresp.Failed("roleName is not empty"))
		return
	}
	err := createRole(roleName)
	commonresp.JudgeError(c, "create role", err)
}

// @Summary 角色接口
// @title 角色删除
// @Tags 角色删除
// @param roleId query int true "角色id"
// @Accept application/x-www-form-urlencoded
// @Produce application/x-www-form-urlencoded
// @Router /api/role/del [get]
// @Success 200 {object} commonresp.JsonResult
func Del(c *gin.Context) {
	roleIdStr, flag := c.GetQuery("roleId")
	if flag {
		c.JSON(http.StatusOK, commonresp.Failed("roleId is not empty"))
		return
	}
	roleId, _ := strconv.Atoi(roleIdStr)
	err := deleteRole(roleId)
	commonresp.JudgeError(c, "del role", err)
}

// @Summary 角色接口
// @title 角色跟新
// @Tags 角色跟新
// @Accept application/json
// @Produce application/json
// @Router /api/role/update [post]
// @param id query int true "角色id"
// @param roleName query string true "角色名称"
// @Success 200 {object} commonresp.JsonResult
func Update(c *gin.Context) {
	idStr, flag := c.GetPostForm("id")
	if flag {
		c.JSON(http.StatusOK, commonresp.Failed("id is not empty"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, commonresp.Failed("参数类型错误"))
		return
	}

	roleName, flag := c.GetPostForm("roleName")
	if flag {
		c.JSON(http.StatusOK, commonresp.Failed("roleName is not empty"))
		return
	}
	role := Role{Id: id, RoleName: roleName}

	err = updateRole(&role)
	commonresp.JudgeError(c, "update role", err)
}

// @Summary 角色接口
// @title 角色查询
// @Tags 角色查询
// @Accept application/json
// @Produce application/json
// @Router /api/role/query [get]
// @Success 200 {[]Role} commonresp.JsonResult
func Query(c *gin.Context) {
	roles, err := queryRoleAll()
	commonresp.JudgeError(c, roles, err)
}
