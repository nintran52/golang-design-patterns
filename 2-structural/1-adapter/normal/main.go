package main

import "fmt"

// Đây là phiên bản không sử dụng Adapter Pattern.

// Service class (Adaptee)
// `Windows` là một máy tính chỉ hỗ trợ cổng USB.
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Client class
// `Client` là đối tượng muốn sử dụng cổng Lightning.
// Client phải trực tiếp kiểm tra và gọi phương thức tương ứng trên `Windows`.
type Client struct{}

func (c *Client) InsertLightningConnectorIntoWindows(windowsMachine *Windows) {
	fmt.Println("Client inserts Lightning connector into Windows directly.")
	fmt.Println("Manual conversion of Lightning signal to USB...")
	windowsMachine.insertIntoUSBPort() // Trực tiếp gọi phương thức của Windows.
}

// Một máy tính Mac hỗ trợ Lightning sẵn, không cần chuyển đổi.
// Đây là class trực tiếp hỗ trợ giao diện Lightning.
type Mac struct{}

func (m *Mac) InsertLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

func main() {
	// Tạo đối tượng Client
	client := &Client{}

	// Tạo đối tượng Windows (Service class)
	windowsMachine := &Windows{}

	// Client phải tự mình xử lý sự không tương thích giữa Lightning và USB.
	client.InsertLightningConnectorIntoWindows(windowsMachine)

	// Nếu Client muốn sử dụng Mac, cần triển khai một cách riêng lẻ.
	// Tạo đối tượng Mac
	mac := &Mac{}
	fmt.Println("Client manually calls Mac's method for Lightning support.")
	mac.InsertLightningPort()
}

/*
Nhược điểm và khó khăn khi không sử dụng Adapter Pattern:

1. **Không hỗ trợ tính trừu tượng:**
   - Client phải biết cụ thể đối tượng mà nó đang làm việc (Windows hoặc Mac).
   - Không thể xử lý đa dạng các đối tượng thông qua một giao diện chung.

2. **Khó mở rộng:**
   - Nếu cần thêm thiết bị mới (ví dụ: một máy tính Linux chỉ hỗ trợ HDMI), client phải viết logic tương thích riêng cho thiết bị đó.

3. **Mã khó bảo trì:**
   - Logic chuyển đổi tín hiệu Lightning sang USB được xử lý trực tiếp trong client.
   - Nếu cần thay đổi cách chuyển đổi, phải sửa đổi mã client, gây ảnh hưởng lớn đến toàn bộ ứng dụng.

4. **Thiếu khả năng tái sử dụng:**
   - Logic chuyển đổi không thể tái sử dụng được cho các client khác.
   - Các phần xử lý bị phân tán, không tập trung trong một lớp.

5. **Không tuân thủ nguyên tắc Single Responsibility:**
   - Client vừa phải thực hiện nhiệm vụ giao tiếp, vừa phải xử lý logic chuyển đổi giao diện (trách nhiệm kép).

6. **Không phù hợp với quy mô lớn:**
   - Khi số lượng thiết bị tăng, mã sẽ trở nên phức tạp và khó quản lý hơn.

Việc sử dụng Adapter Pattern giúp giải quyết các vấn đề trên bằng cách tạo một lớp trung gian (Adapter) để xử lý sự không tương thích giữa các giao diện.
*/
