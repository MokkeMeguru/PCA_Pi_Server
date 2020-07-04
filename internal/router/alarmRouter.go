package router

import (
	"github.com/MokkeMeguru/PCA_Pi_Server/pkg/alarm"
	"github.com/gin-gonic/gin"
)

func AddAlarmRoute (r *gin.Engine) {
	r.POST("/alarm", alarm.SetAlarmList)
}
