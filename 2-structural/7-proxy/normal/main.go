package main

import "fmt"

// Không sử dụng Proxy mà trực tiếp gọi Real Subject
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

func main() {
	// Trực tiếp sử dụng Application (Real Subject)
	application := &Application{}

	// Các URL cần kiểm tra
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	// Gửi yêu cầu đến Real Subject và in kết quả
	httpCode, body := application.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = application.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = application.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = application.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = application.handleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}

/*
Nhược điểm và khó khăn khi không sử dụng Proxy:

1. **Thiếu kiểm soát truy cập**:
   - Không có cơ chế để giới hạn số lượng yêu cầu (rate limiting) từ các client.
   - Hệ thống dễ bị quá tải nếu có nhiều yêu cầu đồng thời.

2. **Không có khả năng quản lý thống nhất**:
   - Nếu cần thêm các tính năng như log, xác thực, hoặc cache, phải thay đổi trực tiếp trong class Application, vi phạm nguyên tắc trách nhiệm duy nhất (Single Responsibility Principle).

3. **Không có lớp trung gian**:
   - Không có khả năng kiểm soát hoặc phân tích các yêu cầu trước khi chúng đến service thật.

4. **Khó mở rộng**:
   - Khi cần thêm logic xử lý bổ sung (ví dụ: tính năng bảo mật), phải sửa đổi Application, làm tăng nguy cơ ảnh hưởng đến các tính năng hiện có.

5. **Thiếu bảo mật**:
   - Không có cơ chế kiểm tra và chặn các yêu cầu không hợp lệ trước khi xử lý.

6. **Khó bảo trì**:
   - Toàn bộ logic nằm trong Application, làm tăng độ phức tạp và khó khăn khi bảo trì hoặc thay đổi hệ thống.

Tóm lại, việc không sử dụng Proxy dẫn đến mã không có cấu trúc rõ ràng, thiếu khả năng mở rộng và bảo mật.
*/
