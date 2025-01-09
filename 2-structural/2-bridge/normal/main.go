package main

import "fmt"

// Phiên bản không sử dụng Bridge Pattern.

// Máy tính Mac, trực tiếp tích hợp logic in ấn cho từng loại máy in.
type Mac struct{}

func (m *Mac) PrintWithHp() {
	fmt.Println("Yêu cầu in từ máy Mac")
	fmt.Println("Đang in bằng máy in HP")
}

func (m *Mac) PrintWithEpson() {
	fmt.Println("Yêu cầu in từ máy Mac")
	fmt.Println("Đang in bằng máy in EPSON")
}

// Máy tính Windows, trực tiếp tích hợp logic in ấn cho từng loại máy in.
type Windows struct{}

func (w *Windows) PrintWithHp() {
	fmt.Println("Yêu cầu in từ máy Windows")
	fmt.Println("Đang in bằng máy in HP")
}

func (w *Windows) PrintWithEpson() {
	fmt.Println("Yêu cầu in từ máy Windows")
	fmt.Println("Đang in bằng máy in EPSON")
}

func main() {
	// Tạo đối tượng máy tính
	macComputer := &Mac{}
	winComputer := &Windows{}

	// Máy Mac in bằng HP
	macComputer.PrintWithHp()
	fmt.Println()

	// Máy Mac in bằng Epson
	macComputer.PrintWithEpson()
	fmt.Println()

	// Máy Windows in bằng HP
	winComputer.PrintWithHp()
	fmt.Println()

	// Máy Windows in bằng Epson
	winComputer.PrintWithEpson()
	fmt.Println()
}

/*
Nhược điểm và khó khăn khi không sử dụng Bridge Pattern:

1. **Không hỗ trợ mở rộng dễ dàng:**
   - Nếu thêm một loại máy in mới, cần phải cập nhật logic in ấn trong cả lớp Mac và Windows.
   - Nếu thêm một loại máy tính mới, cần sao chép lại toàn bộ logic liên quan đến tất cả các máy in hiện có.

2. **Sự trùng lặp mã:**
   - Logic in ấn của từng loại máy in bị lặp lại trong cả hai lớp Mac và Windows.
   - Dẫn đến mã khó bảo trì và dễ phát sinh lỗi khi cần chỉnh sửa.

3. **Phụ thuộc chặt chẽ giữa Abstraction và Implementation:**
   - Máy Mac và Windows phải biết chi tiết từng loại máy in (HP, Epson).
   - Vi phạm nguyên tắc **Dependency Inversion Principle** trong lập trình hướng đối tượng.

4. **Không linh hoạt:**
   - Client không thể thay đổi máy in được liên kết với máy tính tại thời điểm runtime mà không viết thêm mã mới.
   - Khó tích hợp các biến thể mới hoặc logic phức tạp.

5. **Khó quản lý trong dự án lớn:**
   - Số lượng lớp và phương thức sẽ tăng nhanh chóng khi có thêm nhiều loại máy in và máy tính.
   - Với `n` loại máy tính và `m` loại máy in, cần `n * m` phương thức để hỗ trợ tất cả các tổ hợp.

Bridge Pattern giải quyết các vấn đề này bằng cách tách logic cấp cao (Abstraction) khỏi chi tiết thực hiện (Implementation), cho phép mở rộng cả hai chiều một cách độc lập.
*/
