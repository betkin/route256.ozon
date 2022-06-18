package main

import (
	"fmt"
	fuzzing_example "gitlab.ozon.dev/betkin/device-api/homework/24-fuzzing/internal"
)

func main() {
	fmt.Print(fuzzing_example.MyReverse("�"))
}
