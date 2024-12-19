package main

import "fmt"

// Step 1: Đảm bảo rằng miền nghiệp vụ của bạn có thể được biểu diễn dưới dạng một thành phần chính với nhiều lớp tùy chọn bao quanh nó.
// Trong ví dụ này, thành phần chính là một loại pizza cơ bản, và các lớp tùy chọn là các topping như phô mai và cà chua.

// Step 2: Xác định các phương thức chung cho cả thành phần chính và các lớp tùy chọn.
// Tạo một component interface `IPizza` với phương thức `getPrice` để tính giá pizza.
type IPizza interface {
	getPrice() int
}

// Step 3: Tạo một concrete component class và định nghĩa hành vi cơ bản trong đó.
// `VeggieMania` là một concrete component đại diện cho loại pizza cơ bản.
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15 // Giá cơ bản của pizza là 15.
}

// Step 4: Tạo một base decorator class.
// Base decorator phải có một trường tham chiếu đến một đối tượng được bọc, được khai báo dưới dạng component interface.
type BaseTopping struct {
	pizza IPizza // Tham chiếu đến đối tượng được bọc.
}

// Base decorator sẽ không thực hiện gì thêm mà chỉ được sử dụng như một lớp cha.
func (b *BaseTopping) getPrice() int {
	return b.pizza.getPrice()
}

// Step 5: Đảm bảo tất cả các lớp đều triển khai component interface.
// Bước này được áp dụng khi tạo các concrete decorators bên dưới.

// Step 6: Tạo các concrete decorators bằng cách mở rộng từ base decorator.
// Concrete decorator cho topping cà chua.
type TomatoTopping struct {
	pizza IPizza // Tham chiếu đến pizza hoặc topping trước đó.
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() // Gọi `getPrice` của đối tượng được bọc.
	return pizzaPrice + 7            // Thêm giá của topping cà chua.
}

// Concrete decorator cho topping phô mai.
type CheeseTopping struct {
	pizza IPizza // Tham chiếu đến pizza hoặc topping trước đó.
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() // Gọi `getPrice` của đối tượng được bọc.
	return pizzaPrice + 10           // Thêm giá của topping phô mai.
}

// Step 7: Mã client phải chịu trách nhiệm tạo các decorators và kết hợp chúng theo cách mà client cần.
func main() {
	// Tạo pizza cơ bản.
	pizza := &VeggieMania{}

	// Thêm topping phô mai.
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza, // Bọc pizza cơ bản bằng topping phô mai.
	}

	// Thêm topping cà chua.
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese, // Bọc pizza đã có phô mai bằng topping cà chua.
	}

	// Hiển thị giá của pizza với các topping.
fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}