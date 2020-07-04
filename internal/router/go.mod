module github.com/MokkeMeguru/PCA_Pi_Server/internal/handler

go 1.14

replace (
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/alarm => ../../pkg/alarm
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/sound => ../../pkg/sound
)

require (
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/alarm v0.0.0-00010101000000-000000000000
	github.com/MokkeMeguru/PCA_Pi_Server/pkg/sound v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)
