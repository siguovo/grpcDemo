package main

import (
	"fmt"
	"time"

	"github.com/siguovo/grpcdemo/gutil"
)

func main() {
	timstr := fmt.Sprintf("time: %s", time.Now().UTC())

	fmt.Println(timstr)

	root := gutil.InitRoot()

	root.SetName("admin")
	fmt.Println("user init:", root)
}
