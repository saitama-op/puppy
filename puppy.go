package puppy

import (
	"github.com/saitama-op/dog"
)

var VersionStr string = "v1.0.0"

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

func Version() string {
	return VersionStr
}
