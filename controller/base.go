package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

type baseController struct {
}

func (ctrl baseController) ResponseError(c *gin.Context, err *model.ErrorHandler) {
	c.JSON(err.StatusCode, gin.H{
		"errors": err,
	})
}

func (ctrl baseController) ResponseSucess(c *gin.Context, statusCode int, object interface{}) {
	c.JSON(statusCode, gin.H{
		"data": object,
	})
}
