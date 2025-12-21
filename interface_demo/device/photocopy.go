package device

import "fmt"

type Photocopy struct{}

func (p Photocopy) Read(input string) {
	fmt.Println("Đã đọc được tài liệu qua máy Photocopy, tài liệu: ", input)
}

func (p Photocopy) Write(input string) {
	fmt.Println("Đã ghi được tài liệu qua máy Photocopy, tài liệu: ", input)
}
