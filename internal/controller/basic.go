package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MaximillianoNico/Go-Rest-API/pkg/e"
	app "github.com/MaximillianoNico/Go-Rest-API/pkg/formatter"
)

type candidate struct {
	name       string
	interests  []string
	language   string
	experience bool
}

// @Summary Basic Route
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /ping [get]
func BasicRoute(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "test", map[string]string{
		"message": "pong",
	})
}

func RouteWithParams(c *gin.Context) {
	appG := app.Gin{C: c}

	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action

	appG.Response(http.StatusOK, e.SUCCESS, "test params", map[string]string{
		"name":    name,
		"action":  action,
		"message": message,
	})
}
