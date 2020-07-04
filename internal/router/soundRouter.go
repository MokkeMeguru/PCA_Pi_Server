package router

import (
	"github.com/MokkeMeguru/PCA_Pi_Server/pkg/sound"
	"github.com/gin-gonic/gin"
)

func AddSoundRoute(r *gin.Engine) {
	r.GET("/sound", sound.ListSound)
	r.POST("/sound", sound.AddSound)
	r.DELETE("/clear-sound", sound.ClearSound)
}


