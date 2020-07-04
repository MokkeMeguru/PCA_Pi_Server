package sound

import (
	"github.com/gin-gonic/gin"
)

func AddSound(c *gin.Context) {
	var s Sound
	_ = c.BindJSON(&s)
	if SaveSound(&s) {
		c.JSON(200, gin.H{})
	} else {
		c.JSON(304, gin.H{})
	}

}

func ListSound(c *gin.Context) {
	c.JSON(200, gin.H{"sounds": FindSounds(saveFolder)})
}

func ClearSound(c *gin.Context) {
	RemoveAllSound(saveFolder)
	c.JSON(200, gin.H{})
}

