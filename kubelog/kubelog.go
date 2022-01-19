package kubelog

import "fmt"

func Kinfo(data string) {

	fmt.Printf("\033[1;34;40m%s\033[0m\n", data)
}

func Error(data string) {

	fmt.Printf("\033[1;31;40m%s\033[0m\n", data)
}

func Warn(data string) {

	fmt.Printf("\033[1;33;40m%s\033[0m\n", data)
}
