package main

import "fmt"

func main() {
	veggiePizza := &VeggieMania{}
	veggiePizzaWithCheese := &CheeseTopping{
		pizza: veggiePizza,
	}
	veggiePizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with cheese topping is %d\n", veggiePizzaWithCheese.getPrice())
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPizza := &PeppyPaneer{}
	peppyPizzaWithTomato := &TomatoTopping{
		pizza: peppyPizza,
	}
	peppyPizzaWithTomatoAndCheese := &CheeseTopping{
		pizza: peppyPizzaWithTomato,
	}

	fmt.Printf("Price of PeppyPaneer with tomato topping is %d\n", peppyPizzaWithTomato.getPrice())
	fmt.Printf("Price of PeppyPaneer with cheese and tomato topping is %d\n", peppyPizzaWithTomatoAndCheese.getPrice())
}
