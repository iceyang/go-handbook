package svc

import (
	"fmt"

	"github.com/iceyang/m-go-cookbook/web/internal/model"
)

type ExampleService struct {
}

func (es *ExampleService) Create(example *model.Example) {
	fmt.Println("Create Example, ", example)
}

func (es *ExampleService) List() (examples []*model.Example) {
	examples = append(examples, &model.Example{
		Name: "Justin",
		Age:  11,
	})
	return examples
}

var Example = &ExampleService{}
