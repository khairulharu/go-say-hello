package gosayhello

import "fmt"

func SayHello() string {
	return "halo berhasil buat modules"
}

func SayLord(name string) string {
	return fmt.Sprintln("hallo lord", name)
}
