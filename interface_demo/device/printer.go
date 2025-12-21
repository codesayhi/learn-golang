package device

import "fmt"

type Printer struct{}

func (p Printer) Write(input string) {
	fmt.Println("Đã ghi dữ liệu bằng máy in", input)
}
