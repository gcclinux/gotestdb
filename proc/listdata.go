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
	n := 1
	for i := 0; i <= len(rockslice)-1; i++ {
		fmt.Print(rockslice[i], " ")
		if n == 3 {
			fmt.Println()
			n = 0
		}
		n++
	}
	fmt.Println("-------------------------------------")

}
