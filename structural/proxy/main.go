package main

import "fmt"

// Step 1: Tạo giao diện service
// Đây là giao diện chung cho cả proxy và service thật, đảm bảo chúng có thể thay thế lẫn nhau.
// server.go: Subject
type server interface {
	handleRequest(string, string) (int, string)
}

// Step 2: Tạo service class (Real Subject)
// Đây là service thật thực hiện công việc chính, được proxy gọi đến.
// application.go: Real subject
type Application struct{}

func (a *Application) handleRequest(url, method string) (int, string) {
	// Logic xử lý các yêu cầu
	if url == "/app/status" && method == "GET" {
		return 200, "Ok" // Yêu cầu hợp lệ
	} else if url == "/create/user" && method == "POST" {
		return 201, "User Created" // Tạo người dùng thành công
	}
	return 404, "Not Found" // Yêu cầu không hợp lệ
}

// Step 2: Tạo proxy class
// Proxy quản lý service thật và thêm logic như kiểm soát truy cập và giới hạn tốc độ.
// nginx.go: Proxy
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

// Hàm tạo (constructor) cho proxy class, khởi tạo đối tượng service thật và các thông số proxy.
func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{}, // Service thật
		maxAllowedRequest: 2,              // Giới hạn số yêu cầu
		rateLimiter:       make(map[string]int),
	}
}

// Step 3: Triển khai các phương thức proxy
// Proxy kiểm tra giới hạn tốc độ trước khi ủy quyền yêu cầu cho service thật.
func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed" // Yêu cầu bị từ chối
	}
	// Ủy quyền công việc cho service thật
	return n.application.handleRequest(url, method)
}

// Kiểm tra số lượng yêu cầu từ cùng một URL có vượt quá giới hạn không.
func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	} else if n.rateLimiter[url] >= n.maxAllowedRequest {
		return false // Vượt quá giới hạn
	} else {
		n.rateLimiter[url]++
	}
	return true
}

// Step 4: Sử dụng proxy trong client code
// main.go: Client code
func main() {
	// Tạo proxy server
	nginxServer := newNginxServer()

	// Các URL cần kiểm tra
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	// Gửi yêu cầu đến proxy và in kết quả
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}