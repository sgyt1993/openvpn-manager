package login

import (
	"encoding/json"
	"fmt"
	"ovpn-admin/com/cydata/commonresp"
	"ovpn-admin/com/cydata/login"
	"ovpn-admin/com/cydata/role"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	token, _ := login.CreateToken("sgyt")
	fmt.Println(token)

}

func TestParseToken(t *testing.T) {
	token, _ := login.CreateToken("sgyt")
	chaim, _ := login.ParseToken(token)
	fmt.Println(chaim.UserName)
}

func TestRoleJson(t *testing.T) {
	role1 := role.Role{1, "sgyt", nil}
	resp := commonresp.OK(role1)
	jsonResp, _ := json.Marshal(resp)
	fmt.Println(jsonResp)
}
