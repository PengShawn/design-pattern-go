package main

// TomatoTopping 具体装饰
type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}
