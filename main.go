package main

import "testdb/proc"

func main() {

	// List available data
	proc.ListData()

	// Read data from database
	proc.ReadTest()

	// Upload data into database
	upload()

}

func upload() {
	proc.Insert("Cyndi", "Lauper", "69")
	proc.Insert("Elvis", "Presley", "87")
	proc.Insert("Bruce", "Springsteen", "73")
	proc.Insert("Billy", "Idol", "67")
	proc.Insert("Ozzy", "Osbourne", "74")
	proc.Insert("Brandon", "Flowers", "41")
}
