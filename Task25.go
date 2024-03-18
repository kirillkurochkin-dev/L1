package main

import (
	"fmt"
	"time"
)

func sleepTime(milliseconds int) {
	<-time.After(time.Millisecond * time.Duration(milliseconds))
}

func main() {
	sleepTime(2)
	fmt.Println("Прошло 2 секунды")
}
