package accountrole

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/db"
	"strconv"
)

type AccountRole struct {
	Id        int `json:"id"`
	RoleId    int `json:"roleId"`
	AccountId int `json:"accountId"`
}

type AccountRoleVO struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"roleId"`
	AccountId int    `json:"accountId"`
	roleName  string `json:"roleName"`
}

func createAccountRole(accountRole *AccountRole) (err error) {
	_, err = db.GetDb().Exec("INSERT INTO account_role(account_id,role_id) VALUES ($1,$2)", accountRole.AccountId, accountRole.RoleId)
	db.CheckErr(err)
	fmt.Printf("account_role  created\n")
	return err
}

func deleteAccountRole(id int) (err error) {
	var deleteQuery = "DELETE FROM account_role WHERE id = $1"
	res, err := db.GetDb().Exec(deleteQuery, id)
	db.CheckErr(err)
	if rowsAffected, rowsErr := res.RowsAffected(); rowsErr != nil {
		return fmt.Errorf("ERROR: due deleting account_role %d: %s\n", id, rowsErr)
	} else {
		if rowsAffected == 1 {
			fmt.Printf("account_role id %s deleted\n", id)
		}
	}
	return err

}

func queryAccountRoleByAccountId(accountId int) (accountRoleVOs []AccountRoleVO, err error) {
	var queryRoleAll = "select a.id,a.role_id,a.account_id,r.role_name from account_role a left join role r on a.role_id = r.id where a.account_id = $1"
	rows, err := db.GetDb().Query(queryRoleAll, accountId)
	if err != nil {
		err = fmt.Errorf("system is error")
		return
	}
	db.CheckErr(err)

	for rows.Next() {
		u := AccountRoleVO{}
		err := rows.Scan(&u.Id, &u.RoleId, &u.AccountId, &u.roleName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		accountRoleVOs = append(accountRoleVOs, u)
	}

	return accountRoleVOs, err
}

func Add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var accountRole AccountRole

	if req.Body == nil {
		commonresp.JsonRespFail(w, "Please send a request body")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&accountRole)
	err = createAccountRole(&accountRole)
	commonresp.JudgeError(w, "create ccdroute", err)
}

func Del(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	accountRoleIdStr := req.Form.Get("accountRoleId")
	if len(accountRoleIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	accountRoleId, _ := strconv.Atoi(accountRoleIdStr)
	err := deleteAccountRole(accountRoleId)
	commonresp.JudgeError(w, "del ccdRoute", err)
}

func QueryByAccountId(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	accountIdStr := req.Form.Get("accountId")
	if len(accountIdStr) == 0 {
		commonresp.JsonRespFail(w, "roleId is not empty")
		return
	}
	accountId, _ := strconv.Atoi(accountIdStr)

	ccdRoutes, err := queryAccountRoleByAccountId(accountId)
	commonresp.JudgeError(w, ccdRoutes, err)
}
