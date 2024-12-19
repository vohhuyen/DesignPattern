package main

import "fmt"

// Bước 1: Xác định các chiều độc lập trong các class của bạn.
// - Ở đây chúng ta có hai chiều độc lập:
//   + Abstraction: Máy tính (Mac, Windows).
//   + Implementation: Máy in (Epson, HP).

// Bước 2: Xác định các thao tác mà client cần và định nghĩa chúng trong lớp abstraction cơ bản.
// Giao diện Computer là lớp Abstraction, chứa các phương thức mà client cần sử dụng.
type Computer interface {
	Print()             // Phương thức in ấn, client gọi tới.
	SetPrinter(Printer) // Phương thức để liên kết với một máy in (implementation).
}

// Bước 3: Xác định các thao tác khả dụng trên tất cả các platforms.
// Giao diện Printer là lớp Implementation, chứa các phương thức mà mọi máy in phải triển khai.
type Printer interface {
	PrintFile() // Thực hiện in file, được gọi từ lớp Abstraction.
}

// Bước 4: Đối với tất cả các platforms trong miền của bạn, hãy tạo các concrete implementation classes.
// Máy in Epson, triển khai giao diện Printer.
type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Đang in bằng máy in EPSON")
}

// Máy in HP, triển khai giao diện Printer.
type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Đang in bằng máy in HP")
}

// Bước 5: Trong lớp abstraction, thêm một reference field trỏ đến loại implementation.
// Máy tính Mac (Refined Abstraction) triển khai giao diện Computer.
type Mac struct {
	printer Printer // Tham chiếu đến lớp Implementation (máy in).
}

func (m *Mac) Print() {
	fmt.Println("Yêu cầu in từ máy Mac")
	m.printer.PrintFile() // Ủy thác công việc in ấn cho máy in.
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p // Gán máy in (implementation) cho máy Mac.
}

// Máy tính Windows (Refined Abstraction) triển khai giao diện Computer.
type Windows struct {
	printer Printer // Tham chiếu đến lớp Implementation (máy in).
}

func (w *Windows) Print() {
	fmt.Println("Yêu cầu in từ máy Windows")
	w.printer.PrintFile() // Ủy thác công việc in ấn cho máy in.
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p // Gán máy in (implementation) cho máy Windows.
}

// Bước 6: Nếu bạn có nhiều biến thể của logic cấp cao, hãy tạo các refined abstractions.
// Ở ví dụ này, Mac và Windows là các lớp Refined Abstraction, mở rộng từ lớp Abstraction cơ bản.
// Chúng chứa logic riêng cho từng loại máy tính, nhưng vẫn sử dụng máy in thông qua giao diện Printer.

// Bước 7: Mã client nên truyền một đối tượng implementation cho constructor của abstraction.
// Mã client kết nối abstraction (Mac, Windows) với implementation (Epson, HP) và thực hiện in ấn.
func main() {
	// Tạo các implementation cụ thể cho máy in
hpPrinter := &Hp{}       // Máy in HP
	epsonPrinter := &Epson{} // Máy in Epson

	// Tạo abstraction cụ thể cho máy Mac
	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter) // Liên kết máy in HP với máy Mac
	macComputer.Print()               // Yêu cầu in từ máy Mac
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter) // Liên kết máy in Epson với máy Mac
	macComputer.Print()                  // Yêu cầu in từ máy Mac
	fmt.Println()

	// Tạo abstraction cụ thể cho máy Windows
	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter) // Liên kết máy in HP với máy Windows
	winComputer.Print()               // Yêu cầu in từ máy Windows
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter) // Liên kết máy in Epson với máy Windows
	winComputer.Print()                  // Yêu cầu in từ máy Windows
	fmt.Println()
}