package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zm50/common/component/logcli"
	"github.com/zm50/common/model"
)


type GinParamHandler[T any] func(c *gin.Context, arg *T) model.GeneralResult


func ArgInto[T any](fn GinParamHandler[T]) func(c *gin.Context) {
	return ArgIntoWithValid(fn)
}

func ArgIntoWithValid[T any](fn GinParamHandler[T], vailds ...func(c *gin.Context, arg *T) *any) func(c *gin.Context) {
	return func(c *gin.Context) {
		arg := new(T)
		if err := c.ShouldBindJSON(arg); err != nil {
			c.Abort()
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		// valid param
		for _, vaild := range vailds {
			res := vaild(c, arg)
			if res != nil {
				logcli.Error("failed to valid param: %v", res)
				ParamError(c, res)
				return
			}
		}

		res := fn(c, arg)
		OK(c, res)
	}
}

func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, data)
}

func Success[T any](c *gin.Context, data T) {
	OK(c, model.Success(data))
}

func ParamError[T any](c *gin.Context, data T) {
	OK(c, model.ParamError(data))
}

func InternalError[T any](c *gin.Context, data T) {
	OK(c, model.InternalError(data))
}
