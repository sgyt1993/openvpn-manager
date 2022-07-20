package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/xerrors"
	"net/http"
	"ovpn-admin/commonresp"
	"strings"
	"time"
)

var sysAccountName = "sgyt"
var sysPassword = "sgyt"

const (
	jwtIssuer = "openvpn"
	jwtSecret = "sgyt"
)

type Account struct {
	userName string
	password string
}

type Claims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func AccountLogin(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	userName := req.Form.Get("userName")
	password := req.Form.Get("password")
	if strings.Compare(userName, sysAccountName) != 0 || strings.Compare(password, sysPassword) != 0 {
		commonresp.JsonRespFail(w, "password or account error")
		return
	}
	toke, _ := CreateToken(userName)
	commonresp.JsonRespOK(w, toke)
}

func JudgeLogin(rw http.ResponseWriter, req *http.Request) error {
	token := req.Header.Get("token")
	if len(token) == 0 {
		rw.Header().Set("Cache-Control", "must-revalidate, no-store")
		rw.Header().Set("Content-Type", " text/html;charset=UTF-8")
		// 模拟重定向到login接口下
		rw.Header().Set("Location", "/login.html")
		rw.WriteHeader(http.StatusFound)
		return fmt.Errorf("have no token")
	}

	// Verify the token
	_, err := ParseToken(token)
	if err != nil {
		rw.Header().Set("Cache-Control", "must-revalidate, no-store")
		rw.Header().Set("Content-Type", " text/html;charset=UTF-8")
		// 模拟重定向到login接口下
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(http.StatusFound)
		return fmt.Errorf("have no token")
	}
	return err
}

func CreateToken(username string) (string, error) {
	// 当前时间
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(24 * time.Hour)
	//   签发人
	issuer := jwtIssuer
	//	 赋值给结构体
	claims := Claims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime),
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	return toke, err
}

func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return c, nil
		}
	}

	var uErr *jwt.UnverfiableTokenError
	var expErr *jwt.TokenExpiredError
	var nbfErr *jwt.TokenNotValidYetError

	// Use xerrors.Is to see what kind of error(s) occurred
	if tokenClaims.Valid {
		fmt.Println("You look nice today")
	} else if xerrors.As(err, &uErr) {
		fmt.Println("That's not even a token")
	} else if xerrors.As(err, &expErr) {
		fmt.Println("Timing is everything")
	} else if xerrors.As(err, &nbfErr) {
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil, err
}
