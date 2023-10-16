package response

import (
	"net/http"

	"github.com/edison626/vblog/exception"
	"github.com/gin-gonic/gin"
)

//分别创建 成功和 失败后的返回值

// 正常情况数据返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// 异常情况的数据返回， 返回我们的业务Exception
func Failed(c *gin.Context, err error) {

	//如果出现多个Handler,需要通过手动abord
	defer c.Abort()

	//声明
	var e *exception.ApiException

	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		// 非可以预期。没有定义业务的情况
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	//从exception 获取状态码
	c.JSON(e.HttpCode, e)

}
