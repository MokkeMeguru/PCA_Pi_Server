package alarm

import (
	"github.com/gin-gonic/gin"
)

type List struct {
	AlarmList []string
}

var confPath = "./configs/alarm.conf"

func SetAlarmList(c *gin.Context) {
	var list List
	c.BindJSON(&list)
	ConfigWrite(list.AlarmList, confPath)
}
