package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
}

func TestFunctionExport() {
	fmt.Println("test function")
}

func durationToString(duration time.Duration) string {
	if duration.Seconds() < 1 {
		return fmt.Sprintf("%d", duration.Milliseconds()) + "ms"
	}
	return fmt.Sprintf("%f", duration.Seconds()) + "s"
}

func FunctionTimer(function func(), functionName string) {
	startTime := time.Now()
	function()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("\n" + functionName + " took " + durationToString(duration) + " to run")
}
