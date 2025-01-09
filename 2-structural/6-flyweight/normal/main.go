package main

import (
	"fmt"
)

// Mỗi người chơi sẽ giữ tất cả các thuộc tính của mình, bao gồm cả `color`, `lat`, `long`.
// Điều này dẫn đến việc không tối ưu hóa việc sử dụng bộ nhớ khi các thuộc tính giống nhau được lặp lại nhiều lần.

// Player: Không chia Intrinsic và Extrinsic state
type Player struct {
	playerType string // Loại người chơi (Terrorist hoặc CounterTerrorist)
	color      string // Màu sắc (không thay đổi, lặp lại nhiều lần)
	lat        int    // Vị trí trên bản đồ
	long       int
}

func newPlayer(playerType, color string) *Player {
	return &Player{
		playerType: playerType,
		color:      color,
	}
}

func (p *Player) setLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// Game: Không sử dụng Flyweight
type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

func (g *Game) addTerrorist() {
	player := newPlayer("T", "red") // Màu `red` được lặp lại cho tất cả Terrorist
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist() {
	player := newPlayer("CT", "green") // Màu `green` được lặp lại cho tất cả CounterTerrorist
	g.counterTerrorists = append(g.counterTerrorists, player)
}

func main() {
	game := newGame()

	// Thêm Terrorist
	game.addTerrorist()
	game.addTerrorist()
	game.addTerrorist()

	// Thêm Counter-Terrorist
	game.addCounterTerrorist()
	game.addCounterTerrorist()

	// In thông tin các player
	for _, player := range game.terrorists {
		fmt.Printf("PlayerType: %s, Color: %s\n", player.playerType, player.color)
	}

	for _, player := range game.counterTerrorists {
		fmt.Printf("PlayerType: %s, Color: %s\n", player.playerType, player.color)
	}
}

/*
Nhược điểm và khó khăn:

1. Lãng phí bộ nhớ:
   - Màu (`color`) là một thuộc tính không thay đổi nhưng được lưu trữ riêng cho từng người chơi.
   - Khi số lượng người chơi tăng, việc lưu trữ giá trị giống nhau cho từng người chơi sẽ tiêu tốn bộ nhớ không cần thiết.

2. Hiệu suất giảm:
   - Không có cơ chế tái sử dụng các giá trị lặp lại. Khi số lượng người chơi tăng lên (hàng nghìn hoặc hàng triệu), hiệu suất của chương trình sẽ giảm đáng kể.

3. Khó bảo trì:
   - Nếu có sự thay đổi về màu sắc của một loại người chơi (ví dụ: tất cả `Terrorist` chuyển sang màu `blue`), cần phải cập nhật từng người chơi thay vì cập nhật một đối tượng chia sẻ.

4. Thiếu tính mở rộng:
   - Khi thêm nhiều thuộc tính không thay đổi khác (intrinsic state), ví dụ như "loại súng", sẽ phải lưu trữ riêng lẻ trong từng đối tượng, tiếp tục làm lãng phí bộ nhớ.

5. Khả năng tái sử dụng thấp:
   - Không có cơ chế chia sẻ các trạng thái dùng chung, dẫn đến mã không thể tái sử dụng cho các trường hợp tương tự trong các phần khác của ứng dụng.
*/
