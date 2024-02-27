package utils

import (
	"fmt"
	"time"

	"github.com/wbrijesh/utils/helpers"
)

func FunctionTimer(function func(), functionName string) {
	startTime := time.Now()
	function()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("\n" + functionName + " took " + helpers.DurationToString(duration) + " to run")
}
