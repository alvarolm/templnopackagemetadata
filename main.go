package main

import (
	"context"
	"os"
)

func main() {
	Greet("test").Render(context.Background(), os.Stdout)
}
