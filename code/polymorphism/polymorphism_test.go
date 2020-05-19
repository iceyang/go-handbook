package polymorphism

import (
	"testing"
)

func TestPolymorphismOne(t *testing.T) {
	kitty := &Cat{"kitty"}
	// kitty.Eat()

	spike := &Dog{"spike"}
	// spike.Eat()

	AnimalEat(kitty)
	AnimalEat(spike)
}
