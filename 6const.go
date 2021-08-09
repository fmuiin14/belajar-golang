package main

// hrs didalem main

import "fmt"

// untuk mencetak tulisan

func main() {
	const namaCoba = "Ping"
	fmt.Println(namaCoba)
	const (
		depan    = "Depan"
		samping  = "Samping"
		belakang = "Belakang"
	)
	fmt.Println(depan)
	fmt.Println(samping)
	fmt.Println(belakang)
}
