package class1

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

//ControlProgramWithArgs Controls Programs flag inputs
//Note: flaga -h automatycznie wypisze wszystkie
func ControlProgramWithArgs() {
	helperMsg := "Sums all passed args"
	sumMode := flag.Bool("sum", false, helperMsg)

	flag.Parse()

	if *sumMode {
		w := sumElements(flag.Args())
		fmt.Println(w)
	}
}

func sumElements(elements []string) int {
	fmt.Println("Summing elements")
	sum := 0

	for _, el := range elements {
		number, err := strconv.Atoi(el)

		if err != nil {
			log.Fatalf("Unable to convert element '%v' to int. Ended with Error %v", el, err)
		}

		sum += number
	}
	return sum
}
