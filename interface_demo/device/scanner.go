package device

import "fmt"

type Scanner struct{}

func (s Scanner) Read(input string) {
	fmt.Println("Đã đọc được tài liệu qua máy scan, tài liệu: ", input)
}
