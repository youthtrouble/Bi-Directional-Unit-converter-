package main

import "fmt"
import "math"

type Converter struct{
	Feet float64
	Minute float64
	Seconds float64
	Celsius float64
	Fahrenheit float64
	Degrees float64
	Radians float64
	Kilograms float64
	Pounds float64
  	Centimeter float64		
}


func (cvr Converter) centtofeet() float64{
	return cvr.Centimeter /30.48
}

func (cvr Converter) feettocent() float64{
	return cvr.Feet *30.48
}

func (cvr Converter) mintosec() float64{
	return cvr.Minute *60
}

func (cvr Converter) sectomin() float64{
	return cvr.Seconds /60
}

func (cvr Converter) celtofah() float64{
	return cvr.Celsius *1.8 + 32
}

func (cvr Converter) Fahtocel() float64{
	return (cvr.Fahrenheit - 32) * 0.556 
}

func (cvr Converter) degtorad() float64{
	return cvr.Degrees * math.Pi/180
}

func (cvr Converter) radtodeg() float64{
	return cvr.Radians * 180/math.Pi
}

func (cvr Converter) kgtopounds() float64{
	return cvr.Kilograms *2.20462
}

func (cvr Converter) poundstokg() float64{
	return cvr.Centimeter /2.20462
}

func main() {
	cvr := Converter{Feet: 6.0, Centimeter: 10.0, Minute: 2, Seconds:120, Celsius:36, Fahrenheit:100, Degrees:45, Radians:4.5, Kilograms:20, Pounds:5}
	fmt.Println("10 centimeters in feet is :", cvr.centtofeet(), "ft")
	fmt.Println("6 Feet in Centimeters is", cvr.feettocent(), "Centimeters")
	fmt.Println("2 Minutes is equivalent to", cvr.mintosec(), "seconds")
	fmt.Println("120 seconds is equivalent to", cvr.sectomin(), "Minutes")
	fmt.Println("36 Degrees celsius is equivalent to", cvr.celtofah(), "Fahrenheits")
	fmt.Println("100 Fahrenheits is equivalent to", cvr.Fahtocel(), "Degrees")
	fmt.Println("45 Degrees is equivalent to", cvr.degtorad(), "Rads")
	fmt.Println("4.5 Radians is equivalent to", cvr.radtodeg(), "Degrees")
	fmt.Println("20 Kilograms is equivalent to", cvr.kgtopounds(), "Pounds")
	fmt.Println("5 Pounds is equivalent to", cvr.poundstokg(), "Kilograms")
}


	
