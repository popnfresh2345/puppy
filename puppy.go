package puppy

import "github.com/popnfresh2345/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WHenGrownUp(Barks())
}
