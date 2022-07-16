package lib

import (
	in "go-in-action/pkg/d3/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stderr, name)
}
