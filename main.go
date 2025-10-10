package main

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

//const pi float32 = 3.1415926
//
//func tinhChuVi(bankinh float32) float32 {
//	return 2 * pi * bankinh
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
}
