package main

import "line"

func main() {
	e := line.NewServer()

	e.Run("localhost:8009")
}
