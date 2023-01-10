package main

import (
	"fmt"
	"time"
)

func main() {
	timstr := fmt.Sprintf("time: %s", time.Now().UTC())

	fmt.Println(timstr)
}
