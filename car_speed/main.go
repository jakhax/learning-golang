package main

import "fmt"

type Car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	top_speed_kmh  float64
	steering_wheel int16
}

const usixteenbitmax float64 = 65535
const kmh_to_miles float64 = 1.60934

//speed = gas_pedal * top_speed_kmh / usixteenbitmax

//value receiver to calculate speed in kph
//value receivers make a copy of the item
func (c Car) kph() float64 {
	return float64(c.gas_pedal) * c.top_speed_kmh / usixteenbitmax
}

//pointer receiver to change value of top speed to 500 the convert it to mph
//pointer receiver alters the item directly
func (c *Car) mph() float64 {
	c.top_speed_kmh = 500
	return c.kph() / kmh_to_miles
}

func main() {
	c_car := Car{gas_pedal: 22314, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	fmt.Printf("gas_pedal: %v, brake_pedal: %v, steering_wheel: %v, top_speed: %v\n", c_car.gas_pedal, c_car.brake_pedal, c_car.steering_wheel, c_car.top_speed_kmh)
	fmt.Println("Speed in kph:", c_car.kph())
	fmt.Printf("gas_pedal: %v, brake_pedal: %v, steering_wheel: %v, top_speed: %v\n", c_car.gas_pedal, c_car.brake_pedal, c_car.steering_wheel, c_car.top_speed_kmh)
	fmt.Println("New speed in mph:", c_car.mph())
	fmt.Printf("gas_pedal: %v, brake_pedal: %v, steering_wheel: %v, top_speed: %v\n", c_car.gas_pedal, c_car.brake_pedal, c_car.steering_wheel, c_car.top_speed_kmh)

}
