package puppy

import (
	"fmt"
	"https://github.com/jarupong555/dog"

	"github.com/GoesToEleven/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("v.1.0.0")
}
func From12() {
	fmt.Println("v.1.2.0")
}
