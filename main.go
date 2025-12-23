//package main
//
//import (
//	"basic/interface_demo"
//	"basic/interface_demo/Animal"
//)
//
//func main() {
//	dog := Animal.NewDog("Dog")
//	PlayWithAnimal(dog)
//}
//
//func PlayWithAnimal(a interface_demo.Animal) {
//	a.Speak()
//	a.Run()
//}

//package main
//
//import "fmt"
//
//// --- BƯỚC 1: CHIA NHỎ INTERFACE (The Go Way) ---
//
//type Reader interface {
//	Read() string
//}
//
//type Writer interface {
//	Write(content string)
//}
//
//// Composition: Máy nào làm được cả 2 thì dùng cái này
//type ReadWriter interface {
//	Reader
//	Writer
//}
//
//// --- BƯỚC 2: CÁC THIẾT BỊ CỤ THỂ ---
//
//// 1. Máy in (Chỉ biết Ghi)
//type Printer struct{}
//
//func (p Printer) Write(content string) {
//	fmt.Printf("Printer đang in: %s\n", content)
//}
//
//// 2. Máy Scan (Chỉ biết Đọc)
//type Scanner struct{}
//
//func (s Scanner) Read() string {
//	return "Nội dung tờ giấy đã scan"
//}
//
//// 3. Máy Photocopy (Biết cả Đọc và Ghi)
//type PhotoCopier struct{}
//
//func (pc PhotoCopier) Read() string {
//	return "Dữ liệu từ khay kính"
//}
//func (pc PhotoCopier) Write(content string) {
//	fmt.Printf("PhotoCopier đang in bản sao: %s\n", content)
//}
//
//// --- BƯỚC 3: CÁC HÀM XỬ LÝ LOGIC ---
//
//// Hàm này CHỈ cần ghi dữ liệu (Máy in, Photocopy dùng được)
//func Task_InTaiLieu(w Writer, data string) {
//	w.Write(data)
//}
//
//// Hàm này CHỈ cần đọc dữ liệu (Máy Scan, Photocopy dùng được)
//func Task_QuetTaiLieu(r Reader) {
//	fmt.Println("Đã quét được:", r.Read())
//}
//
//// Hàm này CẦN CẢ HAI (Chỉ Photocopy dùng được)
//func Task_SaoChep(rw ReadWriter) {
//	data := rw.Read()
//	rw.Write(data)
//}
//
//func main() {
//	p := Printer{}
//	s := Scanner{}
//	pc := PhotoCopier{}
//
//	fmt.Println("--- TEST HÀM GHI ---")
//	Task_InTaiLieu(p, "Báo cáo tài chính")  // OK
//	Task_InTaiLieu(pc, "Báo cáo tài chính") // OK (vì Photocopy cũng biết Ghi)
//	// Task_InTaiLieu(s, "...")             // LỖI: Scanner không biết Ghi!
//
//	fmt.Println("\n--- TEST HÀM ĐỌC ---")
//	Task_QuetTaiLieu(s)  // OK
//	Task_QuetTaiLieu(pc) // OK (vì Photocopy cũng biết Đọc)
//
//	fmt.Println("\n--- TEST HÀM SAO CHÉP (CẦN CẢ 2) ---")
//	Task_SaoChep(pc) // OK
//
//	// Task_SaoChep(p) // LỖI BIÊN DỊCH: Printer không thỏa mãn ReadWriter (thiếu hàm Read)
//	// Task_SaoChep(s) // LỖI BIÊN DỊCH: Scanner không thỏa mãn ReadWriter (thiếu hàm Write)
//}
//
//package main
//
//import (
//	"basic/interface_demo"
//	"basic/interface_demo/device"
//	"fmt"
//)
//
//func main() {
//	p := device.Printer{}
//	pc := device.Photocopy{}
//	s := device.Scanner{}
//	input := "AB.docx"
//
//	//
//	fmt.Println("Đọc ++++++++++++++++++++++++++++++++++++")
//	Task_QuetTaiLieu(s, input)
//	Task_QuetTaiLieu(pc, input)
//	//
//	fmt.Println("Ghi")
//	Task_InTaiLieu(p, input)
//	Task_InTaiLieu(pc, input) //
//	Task_SaoChep(pc, input)
//}
//
//// Hàm này CHỈ cần ghi dữ liệu (Máy in, Photocopy dùng được)
//func Task_InTaiLieu(w interface_demo.Writer, data string) {
//	w.Write(data)
//}
//
//// Hàm này CHỈ cần đọc dữ liệu (Máy Scan, Photocopy dùng được)
//func Task_QuetTaiLieu(r interface_demo.Reader, data string) {
//	r.Read(data)
//}
//
//// Hàm này CẦN CẢ HAI (Chỉ Photocopy dùng được)
//func Task_SaoChep(rw interface_demo.ReaderWriter, data string) {
//	rw.Read(data)
//	rw.Write(data)
//}

//package main
//
//import (
//	"fmt"
//	"log"
//)
//
//type Notify interface {
//	Send(msg string) error
//}
//
//// ===============
//type Email struct {
//	Email string
//}
//
//type Phone struct {
//	Phone string
//}
//
//type Slack struct {
//	Channel string
//}
//
/////
//
//func (e *Email) Send(msg string) error {
//	fmt.Printf("Gửi Notify đến %s: %s\n", e.Email, msg)
//	return nil
//}
//
//func (p *Phone) Send(msg string) error {
//	fmt.Printf("Gửi Notify đến %s: %s\n", p.Phone, msg)
//	return nil
//}
//
//func (s *Slack) Send(msg string) error {
//	fmt.Printf("Gửi Notify đến %s: %s\n", s.Channel, msg)
//	return nil
//}
//
//func SendAllNotify(ntf []Notify, msg string) {
//	for _, n := range ntf {
//		if err := n.Send(msg); err != nil {
//			log.Println("Lỗi")
//		}
//	}
//}
//func main() {
//	e := &Email{"email"}
//	p := &Phone{"phone"}
//	s := &Slack{"slack"}
//
//	list := []Notify{e, p, s}
//
//	SendAllNotify(list, "Hello World")
//}
//
//package main
//
//import (
//	"fmt"
//	"sync" // 1. Import thư viện đồng bộ
//	"time"
//)
//
//type Notify interface {
//	Send(msg string) error
//}
//
//type Email struct{ Address string }
//type Phone struct{ Number string }
//type Slack struct{ Webhook string }
//
//func (e *Email) Send(msg string) error {
//	time.Sleep(time.Second) // Giả lập mạng chậm 1s
//	fmt.Printf("[EMAIL] Đã gửi tới %s\n", e.Address)
//	return nil
//}
//
//func (p *Phone) Send(msg string) error {
//	time.Sleep(time.Second)
//	fmt.Printf("[SMS] Đã gửi tới %s\n", p.Number)
//	return nil
//}
//
//func (s *Slack) Send(msg string) error {
//	time.Sleep(time.Second)
//	fmt.Printf("[SLACK] Đã gửi tới %s\n", s.Webhook)
//	return nil
//}
//
//func SendAllNotify(ntf []Notify, msg string) {
//	var wg sync.WaitGroup // 2. Khai báo bộ đếm
//
//	for _, n := range ntf {
//		wg.Add(1) // 3. Mỗi lần tạo nhánh mới, tăng bộ đếm lên 1
//
//		// 4. Chạy hàm ẩn danh với từ khóa 'go'
//		go func(target Notify) {
//			defer wg.Done() // 5. Khi hàm này chạy xong, giảm bộ đếm xuống 1
//			target.Send(msg)
//		}(n) // Truyền n vào tham số của hàm để tránh lỗi closure
//	}
//
//	wg.Wait() // 6. Chặn ở đây cho đến khi bộ đếm về 0
//}
//
//func main() {
//	start := time.Now()
//
//	list := []Notify{
//		&Email{"admin@mail.com"},
//		&Phone{"09123"},
//		&Slack{"https://slack.com/123"},
//	}
//
//	fmt.Println("Bắt đầu gửi thông báo đồng thời...")
//	SendAllNotify(list, "Hello Concurrent World!")
//
//	fmt.Printf("Hoàn thành! Tổng thời gian: %v\n", time.Since(start))
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// 1. Định nghĩa Interface và các Struct (như cũ)
//type Notify interface {
//	Send(msg string) string // Thay vì trả về error, trả về string để dễ dùng channel
//}
//
//type Email struct{ Name string }
//type Phone struct{ Number string }
//type Slack struct{ Webhook string }
//
//func (e Email) Send(msg string) string {
//	// Giả lập thời gian gửi ngẫu nhiên từ 1-5 giây
//	delay := rand.Intn(5) + 1
//	time.Sleep(time.Duration(delay) * time.Second)
//	return fmt.Sprintf("Email gửi tới %s sau %ds", e.Name, delay)
//}
//func (p Phone) Send(msg string) string {
//	delay := rand.Intn(5) + 1
//	time.Sleep(time.Duration(delay) * time.Second)
//	return fmt.Sprintf("Email gửi tới %s sau %ds", p.Number, delay)
//}
//
//func (s Slack) Send(msg string) string {
//	delay := rand.Intn(5) + 1
//	time.Sleep(time.Duration(delay) * time.Second)
//	return fmt.Sprintf("Email gửi tới %s sau %ds", s.Webhook, delay)
//}
//
//func main() {
//	results := make(chan string, 3)
//	list := []Notify{
//		Email{Name: "Sinh Viên Go"},
//		Phone{Number: "0911098342"},
//		Slack{Webhook: "http.com"},
//	}
//
//	for _, n := range list {
//		go func(target Notify) {
//			results <- target.Send("Chào bạn!")
//		}(n)
//	}
//
//	fmt.Println("Đang đợi kết quả...")
//
//	// Tối ưu: Tạo 1 đồng hồ tổng cho cả quá trình
//	timeout := time.After(3 * time.Second)
//
//	for i := 0; i < len(list); i++ {
//		select {
//		case res := <-results:
//			fmt.Println("THÀNH CÔNG:", res)
//		case <-timeout: // Dùng chung 1 channel timeout
//			fmt.Println("THẤT BẠI: Quá thời gian tổng 3s, dừng các báo cáo còn lại!")
//			return // Thoát luôn vì đã hết giờ tổng
//		}
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func HocGo() {
//	fmt.Println("HocGo")
//}
//
//func UongCafe() {
//	fmt.Println("UongCafe")
//}
//
//func main() {
//	go HocGo()
//	go UongCafe()
//	time.Sleep(10 * time.Microsecond)
//}

// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
//	func main() {
//		// 1. Tạo một channel kiểu string (Unbuffered - Không bộ đệm)
//		ch := make(chan string)
//
//		// 2. Chạy 1 Goroutine (Nhân viên)
//		go func() {
//			fmt.Println("Nhân viên: Đang chuẩn bị gửi mail...")
//
//			// Giả lập việc gửi mail tốn 2 giây
//			time.Sleep(2 * time.Second)
//
//			// Gửi báo cáo vào channel
//			ch <- "Đã gửi mail thành công!"
//
//			fmt.Println("Nhân viên: Đã gửi xong và báo cáo cho sếp.")
//		}()
//
//		// 3. Ở hàm main (Ông chủ): Đứng đợi ở "đường ống"
//		fmt.Println("Ông chủ: Đang ngồi đợi báo cáo từ nhân viên...")
//
//		// Dòng này sẽ CHẶN (Block) hàm main lại cho đến khi có dữ liệu từ ch
//		tin := <-ch
//
//		// 4. In kết quả nhận được
//		fmt.Println("Ông chủ nhận được tin:", tin)
//		fmt.Println("Ông chủ: Xong việc, đi về!")
//	}
//package main
//
//import "fmt"
//
//func main() {
//	list := []string{"Email 1", "Email 2", "Email 3"}
//	ch := make(chan string, len(list)) // Dùng Buffered để gửi không bị chặn ngay
//
//	go func() {
//		for _, item := range list {
//			ch <- item // Đẩy từng cái vào ống
//		}
//		close(ch) // Đóng ống sau khi gửi xong để bên nhận biết đường mà nghỉ
//	}()
//
//	// Bên nhận: Dùng vòng lặp để lấy từng cái cho đến khi ống đóng
//	for tin := range ch {
//		fmt.Println("Đang xử lý:", tin)
//	}
//}

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	var wg sync.WaitGroup // 1. Khai báo bộ đếm

	for _, number := range numbers {
		wg.Add(1) // 2. Thêm 1 người vào danh sách đợi
		go func(n int) { // 3. Truyền number vào tham số n để tránh lỗi closure
			defer wg.Done() // 4. Báo cáo khi làm xong
			ch <- n * n
		}(number)
	}

	// 5. Chạy một Goroutine riêng để đợi tất cả làm xong rồi mới ĐÓNG ỐNG
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 6. Nhận kết quả (Sạch sẽ, không cần time.Sleep)
	for num := range ch {
		fmt.Println("Kết quả sau khi bình phương:", num)
	}
}
