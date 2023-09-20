package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	ShortLength string `json:"short_length"`
}

type TestRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Test(c *gin.Context) {
	var req TestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req)

	var per = TestResponse{
		Id:          req.Id,
		Name:        req.Name,
		Sex:         "男",
		ShortLength: "长",
	}
	c.JSON(http.StatusOK, per)
}
