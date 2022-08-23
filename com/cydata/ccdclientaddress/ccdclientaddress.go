package ccdclientaddress

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/db"
	"strconv"
)

type CcdClientAddress struct {
	Id            int    `json:"id"`
	AccountId     int    `json:"accountId"`
	ClientAddress string `json:"clientAddress"`
	Mask          string `json:"mask"`
}

type CcdClientAddressVO struct {
	Id            int    `json:"id"`
	AccountId     int    `json:"accountId"`
	ClientAddress string `json:"clientAddress"`
	Mask          string `json:"mask"`
	AccountName   string `json:"accountName"`
}

func checkUserExistent(accountId int) bool {
	// we need to check if there is already such a user
	// return true if user exist
	var c int
	_ = db.GetDb().QueryRow("SELECT count(*) FROM users WHERE account_id = $1", accountId).Scan(&c)
	if c == 1 {
		fmt.Printf("ccdClientAddress account_id %s exist\n", accountId)
		return true
	} else {
		return false
	}
}

func createCcdClientAddress(ccdClientAddress *CcdClientAddress) (err error) {
	_, err = db.GetDb().Exec("INSERT INTO ccd_client_address(account_id,client_address,mask) VALUES ($1,$2,$3)", ccdClientAddress.AccountId, ccdClientAddress.ClientAddress, ccdClientAddress.Mask)
	db.CheckErr(err)
	fmt.Printf("ccdClientAddress %s created\n", ccdClientAddress.AccountId)
	return err
}

func deleteCcdClientAddress(id int) (err error) {
	var deleteQuery = "DELETE FROM ccd_client_address WHERE id = $1"
	res, err := db.GetDb().Exec(deleteQuery, id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR: due deleting ccdClientAddress %d: %s\n", id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("ccdClientAddress id %s deleted\n", id)
		}
	}
	return err

}

func updateCcdClientAddress(ccdClientAddress *CcdClientAddress) (err error) {
	var updateSql = "update ccd_client_address set  account_id = $1,client_address = $2,mask = $3 where id = $4"
	res, err := db.GetDb().Exec(updateSql, ccdClientAddress.AccountId, ccdClientAddress.ClientAddress, ccdClientAddress.Mask, ccdClientAddress.Id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR:update ccdClientAddress id %d: %s\n", ccdClientAddress.Id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf(" ccdClientAddress id %d update\n", ccdClientAddress.Id)
		}
	}
	return err
}

func queryAllCcdClientAddress() (ccdClientAddress []CcdClientAddress, err error) {
	var queryRoleAll = "select id,account_id,client_address,mask from ccd_client_address"
	rows, err := db.GetDb().Query(queryRoleAll)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		u := CcdClientAddress{}
		err := rows.Scan(&u.Id, &u.AccountId, &u.ClientAddress, &u.Mask)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ccdClientAddress = append(ccdClientAddress, u)
	}

	return ccdClientAddress, err
}

func QueryCcdClientAddressByAccountId(accountId int) (ccdClientAddress CcdClientAddressVO, err error) {
	var queryRoleAll = "select ca.id,ca.account_id,ca.client_address,ca.mask,u.username from users u left join ccd_client_address ca on u.id = ca.account_id where u.account_id = $1"
	rows, err := db.GetDb().Query(queryRoleAll, accountId)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		err := rows.Scan(&ccdClientAddress.Id, &ccdClientAddress.AccountId, &ccdClientAddress.ClientAddress, &ccdClientAddress.Mask, &ccdClientAddress.AccountName)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return ccdClientAddress, err
}

func Add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var ccdClientAddress CcdClientAddress

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&ccdClientAddress)
	err = createCcdClientAddress(&ccdClientAddress)
	commonresp.JudgeError(w, "create ccdClientAddress", err)
}

func Del(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ccdRouteIdStr := req.Form.Get("ccdClientAddress")
	if len(ccdRouteIdStr) == 0 {
		commonresp.JsonRespFail(w, "ccdClientAddress is not empty")
		return
	}
	ccdRouteId, _ := strconv.Atoi(ccdRouteIdStr)
	err := deleteCcdClientAddress(ccdRouteId)
	commonresp.JudgeError(w, "del ccdClientAddress", err)
}

func Update(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var ccdClientAddress CcdClientAddress

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&ccdClientAddress)

	err = updateCcdClientAddress(&ccdClientAddress)
	commonresp.JudgeError(w, "update ccdClientAddress", err)
}

func Query(w http.ResponseWriter, req *http.Request) {
	ccdRoutes, err := queryAllCcdClientAddress()
	commonresp.JudgeError(w, ccdRoutes, err)
}

func QueryByAccountId(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	accountIdStr := req.Form.Get("accountId")
	if len(accountIdStr) == 0 {
		commonresp.JsonRespFail(w, "ccdClientAddress is not empty")
		return
	}
	accountId, _ := strconv.Atoi(accountIdStr)

	ccdRoutes, err := QueryCcdClientAddressByAccountId(accountId)
	commonresp.JudgeError(w, ccdRoutes, err)
}
