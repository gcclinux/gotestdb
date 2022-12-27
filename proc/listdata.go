package proc

import (
	"fmt"
)

func ListData() {

	var rockslice []string
	rockslice = append(rockslice, "Cyndi", "Lauper", "69")
	rockslice = append(rockslice, "Elvis", "Presley", "87")
	rockslice = append(rockslice, "Bruce", "Springsteen", "73")
	rockslice = append(rockslice, "Billy", "Billy", "67")
	rockslice = append(rockslice, "Ozzy", "Osbourne", "74")
	rockslice = append(rockslice, "Brandon", "Flowers", "41")

	fmt.Println("List available data to upload")
	fmt.Println()

	fmt.Println("------------METHOD I------------")
	n := 1
	for i := 0; i <= len(rockslice)-1; i++ {
		fmt.Print(rockslice[i], " ")
		if n == 3 {
			fmt.Println()
			n = 0
		}
		n++
	}
	// Method 2 allows you to insert the data from line 36
	fmt.Println("------------METHOD II-----------")
	var readyslice []string
	r := 1
	for i := 0; i <= len(rockslice)-1; i++ {
		readyslice = append(readyslice, rockslice[i])
		if r == 3 {
			fmt.Println("--> ", readyslice[0], readyslice[1], readyslice[2])
			r = 0
			readyslice = nil
		}
		r++
	}

}
