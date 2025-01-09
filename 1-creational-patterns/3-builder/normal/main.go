package main

import "fmt"

// Nhà thông thường
func buildNormalHouse() House {
	return House{
		windowType: "Wooden Window",
		doorType:   "Wooden Door",
		floor:      2,
	}
}

// Nhà băng
func buildIglooHouse() House {
	return House{
		windowType: "Snow Window",
		doorType:   "Snow Door",
		floor:      1,
	}
}

// Cấu trúc nhà
type House struct {
	windowType string
	doorType   string
	floor      int
}

// Mã client
func main() {
	// Xây dựng nhà thông thường
	normalHouse := buildNormalHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	// Xây dựng nhà băng
	iglooHouse := buildIglooHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}

/*
Khó khăn:
1. Mã trở nên phức tạp và khó quản lý:
   - Khi số lượng loại nhà tăng lên, số lượng hàm cũng tăng, làm mã trở nên dài dòng và khó theo dõi.

2. Khó kiểm tra và bảo trì:
   - Nếu một bước xây dựng thay đổi (ví dụ, thay đổi cách thiết lập cửa sổ), bạn cần sửa đổi tất cả các hàm liên quan.

3. Thiếu sự tổ chức:
   - Không có cấu trúc rõ ràng để quản lý các bước xây dựng, làm mã dễ bị lỗi khi sửa đổi.

4. Khả năng tái sử dụng thấp:
   - Các bước xây dựng không thể được tái sử dụng trong các ngữ cảnh khác, dẫn đến lãng phí tài nguyên và thời gian.

5. Thiếu khả năng tùy chỉnh:
   - Không thể tùy chỉnh các bước xây dựng cho từng yêu cầu cụ thể (ví dụ: nhà thông thường với số tầng khác nhau).
*/
