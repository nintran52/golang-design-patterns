package main

import "fmt"

// Phiên bản không sử dụng Decorator Pattern.

// Loại pizza cụ thể với giá cố định.
type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

// Pizza với topping phô mai.
type CheeseTopping struct {
	basePizza *VeggieMania
}

func (c *CheeseTopping) getPrice() int {
	if c.basePizza == nil {
		return 10
	}
	return c.basePizza.getPrice() + 10
}

// Pizza với topping cà chua.
type TomatoTopping struct {
	basePizza *VeggieMania
}

func (t *TomatoTopping) getPrice() int {
	if t.basePizza == nil {
		return 7
	}
	return t.basePizza.getPrice() + 7
}

func main() {
	// Tạo pizza cơ bản
	basePizza := &VeggieMania{}

	// Tạo pizza với topping phô mai
	pizzaWithCheese := &CheeseTopping{
		basePizza: basePizza,
	}

	// Tạo pizza với topping cà chua và phô mai
	pizzaWithCheeseAndTomato := &TomatoTopping{
		basePizza: pizzaWithCheese.basePizza,
	}

	// Hiển thị giá
	fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

/*
Nhược điểm và khó khăn khi không sử dụng Decorator Pattern:

1. **Sự phụ thuộc chặt chẽ:**
   - Các lớp topping phụ thuộc trực tiếp vào lớp cơ bản (VeggieMania).
   - Điều này làm mã trở nên khó bảo trì khi cần thêm loại pizza mới hoặc loại topping mới.

2. **Không linh hoạt:**
   - Khó thêm các topping mới mà không thay đổi lớp cơ bản hoặc các lớp topping hiện có.
   - Nếu cần kết hợp nhiều topping, phải tạo thêm các lớp cho từng tổ hợp topping.

3. **Trùng lặp mã:**
   - Logic tính giá lặp lại trong mỗi lớp topping.
   - Dẫn đến việc khó bảo trì và dễ gây lỗi khi cần thay đổi logic.

4. **Không mở rộng được dễ dàng:**
   - Khi cần thêm một loại topping mới, phải viết lại toàn bộ lớp mới thay vì chỉ cần tạo decorator.

5. **Thiếu khả năng tái sử dụng:**
   - Mỗi lớp topping chỉ hoạt động với một loại pizza cụ thể, khó tái sử dụng cho các loại pizza khác.

Decorator Pattern giải quyết các vấn đề này bằng cách bọc động (dynamically wrapping) các đối tượng, cho phép kết hợp nhiều lớp tùy chọn mà không cần thay đổi cấu trúc của lớp cơ bản hoặc lớp topping.
*/
