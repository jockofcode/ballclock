package main

import (
	"fmt"
	"github.com/jockofcode/ballclock"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		num_of_balls, _ := strconv.Atoi(os.Args[1])
		if num_of_balls >= 27 && num_of_balls <= 127 {

			days := ballclock.CountDaysTillReset(num_of_balls)

			fmt.Printf("%v balls cycle after %v days.\n", num_of_balls, days)
		}
	} else if len(os.Args) == 3 {
		num_of_balls, _ := strconv.Atoi(os.Args[1])
		iterations, _ := strconv.Atoi(os.Args[2])

		if num_of_balls >= 27 && num_of_balls <= 127 {
			fmt.Printf("%v\n", ballclock.GetStateAfterCycles(num_of_balls, iterations))
		}
	}
}
