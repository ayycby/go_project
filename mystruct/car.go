package mystruct

import "fmt"

type Car struct {
	Description string
	Price       int
	TopSpeed    int
}

//struct传入方法时也是值传递，想要修改传入值需要传入指针
//由于是值传递，向方法中传入struct时尽量使用指针，从而避免拷贝备份在内存上的开销
func ChangeCarPrice(car *Car, newPrice int) {
	car.Price = newPrice

	//通过*访问字段
	fmt.Println((*car).Price)
	//直接访问字段
	fmt.Println(car.Price)
}
