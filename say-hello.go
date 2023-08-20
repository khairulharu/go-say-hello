package gosayhello

import "fmt"

func SayHello() string {
	return "halo berhasil buat modules"
}

func SayLord(name string) string {
	return fmt.Sprintln("hallo lord", name)
}

func SayNumber(number int) any {
	return fmt.Sprintln("namor yang kamu masukkna", number)
}