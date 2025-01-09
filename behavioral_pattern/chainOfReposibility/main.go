package main

import "fmt"

// Bước 1: Định nghĩa một giao diện chung cho tất cả các handler.
// department.go: Handler interface
type Department interface {
	execute(*Patient)   // Xử lý yêu cầu.
	setNext(Department) // Thiết lập handler tiếp theo trong chuỗi.
}

// Bước 2: Tạo các handler cụ thể để xử lý từng phần của yêu cầu.
// reception.go: Concrete handler
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Bệnh nhân đã đăng ký xong")
		r.next.execute(p) // Chuyển tiếp yêu cầu đến handler tiếp theo.
		return
	}
	fmt.Println("Lễ tân đăng ký cho bệnh nhân")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// doctor.go: Concrete handler
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Bác sĩ đã khám cho bệnh nhân")
		d.next.execute(p) // Chuyển tiếp yêu cầu đến handler tiếp theo.
		return
	}
	fmt.Println("Bác sĩ đang khám cho bệnh nhân")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// medical.go: Concrete handler
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Bệnh nhân đã nhận thuốc")
		m.next.execute(p) // Chuyển tiếp yêu cầu đến handler tiếp theo.
		return
	}
	fmt.Println("Nhà thuốc đang cấp thuốc cho bệnh nhân")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

// cashier.go: Concrete handler
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Bệnh nhân đã thanh toán xong")
		return
	}
	fmt.Println("Thu ngân đang thu tiền từ bệnh nhân")
	p.paymentDone = true
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

// Bước 3: Định nghĩa đối tượng dữ liệu (yêu cầu).
// patient.go: Đối tượng đại diện cho yêu cầu.
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Bước 4: Xây dựng chuỗi và xử lý các yêu cầu.
// main.go: Mã client
func main() {
	// Tạo các handler riêng lẻ.
	cashier := &Cashier{}
	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	// Tạo đối tượng bệnh nhân.
	patient := &Patient{name: "abc"}

	// Bắt đầu xử lý bệnh nhân qua chuỗi các handler.
	reception.execute(patient)
}
