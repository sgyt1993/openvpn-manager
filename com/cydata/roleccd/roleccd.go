package roleccd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/db"
	"strconv"
)

type RoleRoute struct {
	Id         int `json:"id"`
	RoleId     int `json:"roleId"`
	CcdRouteId int `json:"ccdRouteId"`
}

type RoleRouteVO struct {
	Id          int    `json:"id"`
	RoleId      int    `json:"roleId"`
	CcdRouteId  int    `json:"ccdRouteId"`
	Address     string `json:"address"`
	Mask        string `json:"mask"`
	Description string `json:"description"`
}

func createRoleRoute(roleRoutes []RoleRoute) (err error) {
	dbClient := db.GetDb()
	tx, err := dbClient.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, roleRoute := range roleRoutes {
		_, err = dbClient.Exec("INSERT INTO role_ccdroute(role_id,ccd_route_id) VALUES ($1,$2)", roleRoute.RoleId, roleRoute.CcdRouteId)
		db.CheckErr(err)
		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Printf("role_route created\n")
	}

	return err
}

func deleteRoleRoute(id int) (err error) {
	var deleteQuery = "DELETE FROM role_ccdroute WHERE id = $1"
	res, err := db.GetDb().Exec(deleteQuery, id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR: due deleting role_route %d: %s\n", id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("role_route id %s deleted\n", id)
		}
	}
	return err
}

func queryRoleRouteByRoleId(roleId int) (roleRoutes []RoleRouteVO, err error) {
	var queryRoleAll = "select r.id,r.role_id,r.ccd_route_id,c.address,c.mask,c.description from role_ccdroute r left join ccd_route c on c.id = r.ccd_route_id where r.role_id = $1"
	rows, err := db.GetDb().Query(queryRoleAll, roleId)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		u := RoleRouteVO{}
		err := rows.Scan(&u.Id, &u.RoleId, &u.CcdRouteId, &u.Address, &u.Mask, &u.Description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		roleRoutes = append(roleRoutes, u)
	}

	return roleRoutes, err
}

func Add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var roleRoute []RoleRoute

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&roleRoute)
	err = createRoleRoute(roleRoute)
	commonresp.JudgeError(w, "create ccdroute", err)
}

func Del(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	roleRouteIdStr := req.Form.Get("roleCcdRouteId")
	if len(roleRouteIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	roleRouteId, _ := strconv.Atoi(roleRouteIdStr)
	err := deleteRoleRoute(roleRouteId)
	commonresp.JudgeError(w, "del ccdRoute", err)
}

func QueryByRoleId(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	roleRouteIdStr := req.Form.Get("roleId")
	if len(roleRouteIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	roleRouteId, _ := strconv.Atoi(roleRouteIdStr)
	roleRoutes, err := queryRoleRouteByRoleId(roleRouteId)
	commonresp.JudgeError(w, roleRoutes, err)
}
