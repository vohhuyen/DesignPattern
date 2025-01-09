package main

import "fmt"

// command.go: Interface Command
// Bước 1: Tuyên bố giao diện Command với một phương thức thực thi duy nhất.
type Command interface {
	execute()
}

// onCommand.go: Lệnh Cụ thể Bật TV
// Bước 2: Tạo các lớp lệnh cụ thể thực hiện giao diện Command.
// Lệnh cụ thể để bật TV
type OnCommand struct {
	device Device // Tham chiếu đến đối tượng receiver
}

func (c *OnCommand) execute() {
	c.device.on() // Gọi phương thức on của receiver
}

// offCommand.go: Lệnh Cụ thể Tắt TV
// Lệnh cụ thể để tắt TV
type OffCommand struct {
	device Device // Tham chiếu đến đối tượng receiver
}

func (c *OffCommand) execute() {
	c.device.off() // Gọi phương thức off của receiver
}

// button.go: Invoker (Người Gửi)
// Bước 3: Xác định các lớp sẽ đóng vai trò là sender. Thêm các trường lưu trữ lệnh vào các lớp này.
// Button hoạt động như Invoker
type Button struct {
	command Command // Lưu trữ tham chiếu đến đối tượng Command
}

func (b *Button) press() {
	// Bước 4: Thay đổi sender để nó thực thi lệnh thay vì gửi yêu cầu trực tiếp đến receiver.
	b.command.execute() // Gọi phương thức execute của Command
}

// device.go: Giao diện Receiver
// Giao diện Receiver
type Device interface {
	on()
	off()
}

// tv.go: Receiver Cụ thể
// Receiver cụ thể: TV
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Bật TV")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Tắt TV")
}

// main.go: Mã Client
// Bước 5: Client khởi tạo các đối tượng theo thứ tự:
// Tạo receiver, tạo lệnh và liên kết chúng với receiver.
// Tạo sender và liên kết nó với các lệnh cụ thể.
func main() {
	// Tạo receiver
	tv := &Tv{}

	// Tạo các lệnh và liên kết chúng với receiver
	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	// Tạo sender (các nút) và liên kết chúng với các lệnh cụ thể
	onButton := &Button{
		command: onCommand,
	}
	offButton := &Button{
		command: offCommand,
	}

	// Sử dụng các nút để thực thi lệnh
	onButton.press()  // Output: Bật TV
	offButton.press() // Output: Tắt TV
}
