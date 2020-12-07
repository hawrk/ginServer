package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success %s\n", value)))

	return

}