package main

import "fmt"

// Không sử dụng lớp cơ sở hoặc giao diện
// Mỗi loại sản phẩm là một cấu trúc độc lập

// Cấu trúc Ak47
type Ak47 struct {
	name  string
	power int
}

func newAk47() *Ak47 {
	return &Ak47{
		name:  "AK47 gun",
		power: 4,
	}
}

func (a *Ak47) getName() string {
	return a.name
}

func (a *Ak47) getPower() int {
	return a.power
}

// Cấu trúc Musket
type Musket struct {
	name  string
	power int
}

func newMusket() *Musket {
	return &Musket{
		name:  "Musket gun",
		power: 1,
	}
}

func (m *Musket) getName() string {
	return m.name
}

func (m *Musket) getPower() int {
	return m.power
}

// Cấu trúc Sniper
type Sniper struct {
	name  string
	power int
}

func newSniper() *Sniper {
	return &Sniper{
		name:  "Sniper rifle",
		power: 10,
	}
}

func (s *Sniper) getName() string {
	return s.name
}

func (s *Sniper) getPower() int {
	return s.power
}

// Mã client
func main() {
	// Tạo sản phẩm trực tiếp bằng cách gọi constructor
	// Khó khăn xuất hiện khi thêm sản phẩm mới (Sniper)
	ak47 := newAk47()
	musket := newMusket()
	sniper := newSniper() // Mỗi lần thêm sản phẩm mới phải sửa client

	// In chi tiết của các sản phẩm
	printAk47Details(ak47)
	printMusketDetails(musket)
	printSniperDetails(sniper)
}

func printAk47Details(a *Ak47) {
	fmt.Printf("Gun: %s\n", a.getName())
	fmt.Printf("Power: %d\n", a.getPower())
}

func printMusketDetails(m *Musket) {
	fmt.Printf("Gun: %s\n", m.getName())
	fmt.Printf("Power: %d\n", m.getPower())
}

func printSniperDetails(s *Sniper) {
	fmt.Printf("Gun: %s\n", s.getName())
	fmt.Printf("Power: %d\n", s.getPower())
}

// Khó Khăn Khi Không Sử Dụng Factory Pattern:
// 1. Phụ thuộc vào chi tiết cụ thể của sản phẩm:
// 	- Client phải biết chính xác cách khởi tạo từng loại sản phẩm (newAk47(), newMusket()). Điều này làm tăng sự phụ thuộc giữa client và các lớp cụ thể, vi phạm nguyên tắc Dependency Inversion.

// 2. Khó mở rộng sản phẩm:
// 	- Nếu cần thêm một loại sản phẩm mới (ví dụ: Sniper), client code phải sửa đổi để thêm constructor mới. Điều này làm giảm tính linh hoạt và dễ gây lỗi.

// 3. Không có sự trừu tượng hóa:
//	- Mọi logic tạo sản phẩm đều nằm trong client code. Nếu cần thay đổi cách tạo sản phẩm (ví dụ: thêm tham số), tất cả các nơi sử dụng phải được cập nhật.

// 4. Không tái sử dụng logic tạo:
// - Logic tạo sản phẩm không được tổ chức lại ở một nơi duy nhất (như trong factory), dẫn đến việc nhân bản mã nếu cùng một loại sản phẩm được tạo ở nhiều nơi.

// 5. Quản lý nhiều sản phẩm phức tạp:
// - Khi có nhiều loại sản phẩm, client code dễ trở nên khó quản lý và lỗi xảy ra khi logic tạo sản phẩm bị lặp lại hoặc không đồng nhất.
