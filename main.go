package main

import "fmt"

//// Khai báo biến
//var (
//	count = 0 //private
//	Count = 0 //public
//
//	//	Mảng - kích thước mảng không thể thay đổi sau khi khai báo
//	color = [3]string{
//		"red",
//		"blue",
//		"green",
//	}
//)
//
//// map
//var (
//	colorMap = map[string]int{
//		"red":   0,
//		"blue":  1,
//		"green": 2,
//	}
//)
//
////
//
//// Struct
//type Student struct {
//	StudentName `json:"name"`
//	Age         uint32 ``
//}
//type StudentName string
//
//func (n StudentName) String() string {
//	return len(n) <= 32 && n != ""
//}
//
////func (s *Student) Print() { //(Con trỏ)Thay đổi trực tiếp giá trị *Student
////
////}
////
////func (s Student) PrintSInhVien() { //giá trị - tạo ra 1 bản sao và thao tác trên bản sao đó
////
////}
//
//func (s Student) GetName() string {
//	return s.Name
//}
//
//func (s *Student) SetName(name string) {
//	s.Name = name
//}

// Map
//
//	colors = []string {
//		"green",
//		"blue",
//		"red"
//	}
//
// colorMap = map
//func kiemTraDuLieu(dai, rong int) bool {
//	return dai >= rong
//}
//func tinhChuVi(dai, rong int) int {
//	return (dai * rong) * 2
//}
//
//func tinhDienTich(dai, rong int) int {
//	return dai * rong
//}

// const pi float32 = 3.1415926
//
//	func tinhChuVi(bankinh float32) float32 {
//		return 2 * pi * bankinh
//	}
//
//	func namNhuan(nam int) bool {
//		if nam%400 == 0 {
//			return true
//		}
//		if nam%4 == 0 && nam%100 != 0 {
//			return true
//		}
//		return false
//	}
func fibonacif(n int) []int {
	if n == 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	} else {
		arr := make([]int, n)
		arr[0] = 0
		arr[1] = 1
		for i := 2; i < n; i++ {
			arr[i] = arr[i-1] + arr[i-2]
		}
		return arr
	}
}

func sumFibonaci(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	a := 0
	b := 1
	sum := a + b
	for i := 2; i < n; i++ {
		next := a + b
		sum += next
		a = b
		b = next
	}
	return sum
}
func checkFibonacci(n int) bool {
	if n < 0 {
		return false
	}
	a, b := 0, 1
	for a <= n {
		if a == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

//func checkFibonacci(n int) bool {
//	if n < 0 {
//		return false
//	}
//	a, b := 0, 1
//	for {
//		if a == n {            // trúng
//			return true
//		}
//		if a > n {             // vượt rồi thì không thuộc
//			return false
//		}
//		a, b = b, a+b          // tiến bước
//	}
//}

func main() {
	////Slide - nó giống như mảng nhưng có thể thay đổi kích thước
	////colors := []string{"red", "blue", "green"}
	////colors = append(colors, "y")
	//fmt.Println("Hello World")
	//
	////s := Student{"Steve", 30}
	////s.GetName()
	////s.SetName("Steve")
	//
	////
	//coloMap := make(map[string]int)
	//coloMap["red"] = 1
	//coloMap["blue"] = 2

	// các bài tập nhỏ
	//In
	//fmt.Println("Hello World")
	//Biến & hằng
	//Tính chu vi/diện tích hình chữ nhật từ input. Thử const, :=, fmt.
	//var (
	//	chieudai, chieurong int
	//)
	//fmt.Print("Nhập chiều dài: ")
	//fmt.Scan(&chieudai)
	//fmt.Print("Nhập chiều rông: ")
	//fmt.Scan(&chieurong)
	//if !kiemTraDuLieu(chieudai, chieurong) {
	//	fmt.Println("Chiều dài không được nhỏ hơn chiều rộng")
	//	return
	//}
	//a := tinhChuVi(chieudai, chieurong)
	//fmt.Println("Chu vi là : ", a)
	//fmt.Println("Dien tich là : ", tinhDienTich(chieudai, chieurong))
	//var bankinh float32 = 10
	//
	//fmt.Println("Chu vị hình tròn là : ", tinhChuVi(bankinh))
	//
	//var nam int
	//fmt.Print("Nhập năm: ")
	//fmt.Scanln(&nam)
	//
	//a := namNhuan(nam)
	//if a == true {
	//	fmt.Printf("Năm %v là năm nhuận \n ", nam)
	//} else {
	//	fmt.Printf("Năm %v không phải năm nhuận \n ", nam)
	//}

	//
	fmt.Println(fibonacif(7))
	fmt.Println(sumFibonaci(12))

}
