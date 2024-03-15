package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	x := os.Args[1]
	ts, _ := strconv.ParseInt(x, 10, 64)
	out := time.Unix(ts, 0)
	fmt.Println(out)
}
