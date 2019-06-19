package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/m-go-cookbook/web/internal/model"
	"github.com/iceyang/m-go-cookbook/web/internal/svc"
)

type ExampleController struct{}

var Example = &ExampleController{}

func (ec *ExampleController) Create(c *gin.Context) {
	example := &model.Example{
		Name: "Justin",
		Age:  18,
	}
	svc.Example.Create(example)
	c.JSON(200, example)
}

func (ec *ExampleController) List(c *gin.Context) {
	c.JSON(200, svc.Example.List())
}

func (ec *ExampleController) EmptyBody(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (ec *ExampleController) With404(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
