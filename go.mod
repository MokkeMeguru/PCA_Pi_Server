module github.com/MokkeMeguru/PCA_Pi_Server

go 1.14

replace (
	github.com/MokkeMeguru/PCA_Pi_Server/internal/router => ./internal/router
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/alarm => ./pkg/alarm
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/sound => ./pkg/sound
)

require (
	github.com/MokkeMeguru/PCA_Pi_Server/internal/router v0.0.0-00010101000000-000000000000
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/swaggo/swag v1.6.7
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200702044944-0cc1aa72b347 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
