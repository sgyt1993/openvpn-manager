package login

import (
	"fmt"
	"ovpn-admin/login"
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
