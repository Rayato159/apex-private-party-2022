package src

import (
	"fmt"
	"os"
)

// 5h operating
// 15 layer / 3 photographer = 5 groups (3 layer per 1 group)

// 1 photographer can handle -> 20 min/layer (means 1 layer must be 20 min)
// 1 group must be 20 min
// 1 round must be -> 5 * 20 = 100 min

// Conclusion
// 300/100 -> 3 round to swap per layer

// Algorithm -> Combination nCr -> 15C3

func Solution() error {
	arr := [15]string{
		"Lifeline",
		"Horizon",
		"Mirage",
		"Revenant",
		"Pathfider",
		"Valkyrie",
		"Octane",
		"Caustic",
		"Vantage",
		"Wraith",
		"Wattson",
		"Rampart",
		"Loba",
		"Bloodhound",
		"Crypto",
	}
	n := 15
	r := 3
	CombinationDisplay(arr, n, r)
	return nil
}

func CombinationDisplay(arr [15]string, n int, r int) {
	data := [15]string{""}
	count := 0
	var record string

	Combination(arr, data, 0, n-1, 0, r, &count, &record)
	fmt.Printf("\ncombination possible -> %d\n", count)

	recordByte := []byte(record)
	if err := os.WriteFile("./solution.txt", recordByte, os.FileMode(0777)); err != nil {
		panic(err.Error())
	}
}

func Combination(arr [15]string, data [15]string, start, end, index, r int, count *int, record *string) {
	if index == r {
		for i := 0; i < r; i++ {
			fmt.Printf("%s ", data[i])
			if i != r-1 {
				*record += fmt.Sprintf("%s,", data[i])
			} else {
				*record += data[i]
			}
		}
		*count++
		*record += "\n"
		fmt.Println()
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = arr[i]
		Combination(arr, data, i+1, end, index+1, r, count, record)
	}
}
