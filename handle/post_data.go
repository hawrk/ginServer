package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostData(c *gin.Context) {
	firstData := c.Request.FormValue("first")
	secondData := c.Request.FormValue("second")
	//secondData := c.PostForm("first")  // 尽量不用

	fmt.Println("get post data |first:", firstData, ",second:", secondData)
	//msg := fmt.Sprintf("get data success, %s,%s", firstData, secondData)
	//c.JSON(http.StatusOK,gin.H{
	//	"msg": msg,
	//})
	// 或者定义一个通用的返回结构体 TODO:
	type JsonHolder struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	ret := JsonHolder{
		Code: 0,
		Msg:  "success post",
	}
	c.JSON(http.StatusOK, ret)

	return
}
