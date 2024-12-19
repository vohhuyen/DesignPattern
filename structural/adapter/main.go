package main

import "fmt"

// Bước 1: Đảm bảo rằng bạn có ít nhất hai classes với các interfaces không tương thích:
// - `Windows` là service class mà bạn không thể thay đổi, chỉ hỗ trợ cổng USB.
// - `Client` là class mong muốn sử dụng giao diện Lightning, nhưng không thể giao tiếp trực tiếp với `Windows`.

// Service class (Adaptee)
// `Windows` là một máy tính chỉ hỗ trợ cổng USB.
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Client class
// `Client` là đối tượng muốn sử dụng cổng Lightning.
// Nó chỉ làm việc với các đối tượng triển khai giao diện `Computer`.
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort() // Gọi phương thức từ interface `Computer`.
}

// Bước 2: Khai báo client interface và mô tả cách các clients giao tiếp với service.
// Giao diện `Computer` định nghĩa các phương thức mà `Client` mong đợi.
type Computer interface {
	InsertIntoLightningPort() // Phương thức để cắm cổng Lightning.
}

// Bước 3: Tạo một adapter class và làm cho nó tuân theo client interface.
// `WindowsAdapter` sẽ đóng vai trò làm cầu nối giữa `Client` và `Windows`.
// Ban đầu, lớp adapter chỉ được định nghĩa và chưa thực thi các phương thức.

type WindowsAdapter struct {
	// Bước 4: Thêm một field vào adapter class để lưu tham chiếu đến service object.
	// `WindowsAdapter` chứa một tham chiếu đến đối tượng `Windows`.
	windowMachine *Windows
}

// Bước 5: Lần lượt triển khai tất cả các phương thức của client interface trong adapter class.
// Triển khai phương thức `InsertIntoLightningPort` cho adapter.
// Adapter chuyển đổi giao diện Lightning thành USB và ủy thác công việc thực tế cho đối tượng `Windows`.
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort() // Gọi phương thức của `Windows`.
}

// Bước 6: Clients nên sử dụng adapter thông qua client interface.
// Điều này đảm bảo rằng client không cần biết về sự tồn tại của adapter hoặc service class gốc.

func main() {
	// Tạo đối tượng Client
	client := &Client{}

	// Tạo đối tượng Windows (Service class)
	windowsMachine := &Windows{}

	// Tạo adapter để kết nối giữa Client và Windows
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine, // Liên kết đối tượng Windows với adapter.
	}

	// Client sử dụng adapter để giao tiếp với Windows thông qua giao diện `Computer`.
client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	// Nếu có một máy tính Mac hỗ trợ Lightning, Client có thể sử dụng mà không cần adapter.
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
}

// Một máy tính Mac hỗ trợ Lightning sẵn, không cần adapter.
// Đây là class trực tiếp triển khai giao diện `Computer`.
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}