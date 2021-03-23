module github.com/xiahongjian/pi-status

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shirou/gopsutil v3.21.2+incompatible
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/xiahongjian/pi-status/conf => ./conf
	github.com/xiahongjian/pi-status/models => ./models
	github.com/xiahongjian/pi-status/pkg/setting => ./pkg/setting
	github.com/xiahongjian/pi-status/routers => ./routers
	github.com/xiahongjian/pi-status/service => ./service
)
