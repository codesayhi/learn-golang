package main

// Khai báo biến
var (
	count = 0 //private
	Count = 0 //public

	//	Mảng - kích thước mảng không thể thay đổi sau khi khai báo
	color = [3]string{
		"red",
		"blue",
		"green",
	}
)

// Struct
type Student struct {
	Name string
	Age  uint32
}

//func (s *Student) Print() { //(Con trỏ)Thay đổi trực tiếp giá trị *Student
//
//}
//
//func (s Student) PrintSInhVien() { //giá trị - tạo ra 1 bản sao và thao tác trên bản sao đó
//
//}

func (s Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {
	//Slide - nó giống như mảng nhưng có thể thay đổi kích thước
	//colors := []string{"red", "blue", "green"}
	//colors = append(colors, "y")
	//fmt.Println("Hello World")
	s := Student{"Steve", 30}
	s.GetName()
	s.SetName("Steve")
}
