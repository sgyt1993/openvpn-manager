module ovpn-admin

go 1.14

require (
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/common v0.15.0 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)
