package ccdroute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/db"
	"strconv"
)

type CcdRoute struct {
	Id          int    `json:"id"`
	Address     string `json:"address"`
	Mask        string `json:"mask"`
	Description string `json:"description"`
}

func createCcdRoute(ccd *CcdRoute) (err error) {
	_, err = db.GetDb().Exec("INSERT INTO ccd_route(address,mask,description) VALUES ($1,$2,$3)", ccd.Address, ccd.Mask, ccd.Description)
	db.CheckErr(err)
	fmt.Printf("ccd_route %s created\n", ccd.Address)
	return err
}

func deleteCcdRoute(id int) (err error) {
	var deleteQuery = "DELETE FROM ccd_route WHERE id = $1"
	res, err := db.GetDb().Exec(deleteQuery, id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR: due deleting ccd_route %d: %s\n", id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("ccd_route id %s deleted\n", id)
		}
	}
	return err

}

func updateCcdRoute(ccd *CcdRoute) (err error) {
	var updateSql = "update ccd_route set  address = $1,mask = $2,description = $3 where id = $4"
	res, err := db.GetDb().Exec(updateSql, ccd.Address, ccd.Mask, ccd.Description, ccd.Id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR:update ccd_route id %d: %s\n", ccd.Id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf(" ccd_route id %d update\n", ccd.Id)
		}
	}
	return err
}

func queryAllCcdRoute() (ccds []CcdRoute, err error) {
	var queryRoleAll = "select id,address,mask,description from ccd_route"
	rows, err := db.GetDb().Query(queryRoleAll)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		u := CcdRoute{}
		err := rows.Scan(&u.Id, &u.Address, &u.Mask, &u.Description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ccds = append(ccds, u)
	}

	return ccds, err
}

func Add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var ccd CcdRoute

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&ccd)
	err = createCcdRoute(&ccd)
	commonresp.JudgeError(w, "create ccdroute", err)
}

func Del(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ccdRouteIdStr := req.Form.Get("ccdRouteId")
	if len(ccdRouteIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	ccdRouteId, _ := strconv.Atoi(ccdRouteIdStr)
	err := deleteCcdRoute(ccdRouteId)
	commonresp.JudgeError(w, "del ccdRoute", err)
}

func Update(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var ccd CcdRoute

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&ccd)

	err = updateCcdRoute(&ccd)
	commonresp.JudgeError(w, "update role", err)
}

func Query(w http.ResponseWriter, req *http.Request) {
	ccdRoutes, err := queryAllCcdRoute()
	commonresp.JudgeError(w, ccdRoutes, err)
}
