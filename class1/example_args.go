package class1

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

//ControlProgramWithArgs Controls Programs flag inputs
//Note: flaga -h automatycznie wypisze wszystkie
func ControlProgramWithArgs() {
	sumMode := flag.Bool("sum", false, "Sumuje podane argumenty")
	powerMode := flag.Bool("power", false, "Podnosi arg1 do potegi arg2")
	flag.Parse()

	if *sumMode {
		numberSlice, err := sliceAtoi(flag.Args())
		handleError(err)
		fmt.Println(sumElements(numberSlice))
	}

	if *powerMode {
		numberSlice, err := sliceAtoi(flag.Args())
		handleError(err)
		args := numberSlice[0]
		pow := numberSlice[1]
		fmt.Println(powerUp(args, pow))
	}
}

func powerUp(number, pow int) int {

	return int(math.Pow(float64(number), float64(pow)))
}

func sumElements(elements []int) int {
	fmt.Println("Summing elements")
	sum := 0

	for _, el := range elements {

		sum += el
	}
	return sum
}

func sliceAtoi(list []string) ([]int, error) {
	intSlice := []int{}

	for _, el := range list {
		number, err := strconv.Atoi(el)

		if err != nil {
			return nil, err
		}

		intSlice = append(intSlice, number)
	}
	return intSlice, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Program exit due to %v", err)
	}
}
