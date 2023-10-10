package response

import (
	"net/http"

	"gitee.com/go-course/go12/vblog/exception"
	"github.com/gin-gonic/gin"
)

// c.JSON(http.StatusBadRequest, err.Error())
// 不面向对象, 工具

// 正常请求数据返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// 异常情况的数据返回, 返回我们的业务Exception
func Failed(c *gin.Context, err error) {
	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		// 非可以预期, 没有定义业务的情况
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	c.JSON(e.HttpCode, e)
}
