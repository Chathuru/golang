package main

import "fmt"

// Define constant variables
const usixteenbitmax float64 = 65535
const kmh_multiplier float64 = 1.60934

// Struct definition
type car struct{
	gas_pedal uint16
	break_pedal uint16
	streering_wheel int16
	top_speed_kmh float64
}

// Value reciver function. These functions directly bind to the struct car. func (variable struct-name) function-name()
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiplier)
}

// Pointer reciver function. Use to update the values in the struct
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{
				gas_pedal: 65000,
				break_pedal: 0,
				streering_wheel: 0,
				top_speed_kmh: 150,
			}

	b_car := car{5600,0,0,80}

	fmt.Println("a_car gas_pedal value =",a_car.gas_pedal)
	fmt.Println("b_car top_speed_kmh =",b_car.top_speed_kmh)

	fmt.Println("a_car speed",a_car.kmh())
	fmt.Println("a_car mph",a_car.mph())

	fmt.Println("b_car speed",b_car.kmh())

	a_car.new_top_speed(250)
	fmt.Println("a_car new speed",a_car.kmh())
	fmt.Println("a_car new mph",a_car.mph())
}