package main

import "fmt"
import "math"

type Converter struct{}		

type Feet float64

type Centimeter float64


func (cvr Converter) centtofeet(x Centimeter) Feet{
	return Feet(float64(x)/30.48)
}

func (cvr Converter) feettocent(x Feet) Centimeter{
	return Centimeter (float64(x)*30.48)
}

type Minute float64

type Seconds float64

func (cvr Converter) mintosec(y Minute) Seconds{
	return Seconds (float64(y)*60.0)
}

func (cvr Converter) sectomin(y Seconds) Minute{
	return Minute (float64(y)/60.0)
}

type Celsius float64

type Fahrenheit float64

func (cvr Converter) celtofah(a Celsius) Fahrenheit{
	return Fahrenheit (float64(a)*1.8 + 32)
}

func (cvr Converter) Fahtocel(a Fahrenheit) Celsius{
	return Celsius ((float64(a) - 32) * 0.556) 
}

type Degrees float64

type Radians float64


func (cvr Converter) degtorad(z Degrees) Radians{
	return Radians (float64(z)* math.Pi/180)
}

func (cvr Converter) radtodeg(z Radians) Degrees{
	return Degrees (float64(z)* 180/math.Pi)
}

type Kilograms float64

type Pounds float64

func (cvr Converter) kgtopounds(c Kilograms) Pounds{
	return Pounds (float64(c)*2.20462)
}

func (cvr Converter) poundstokg(c Pounds) Kilograms{
	return Kilograms (float64(c)/2.20462)
}

func main() {
	cvr := Converter{}
	fmt.Println("10 centimeters in feet is :", cvr.centtofeet(10.0), "ft")
	fmt.Println("6 Feet in Centimeters is", cvr.feettocent(6.0), "Centimeters")
	fmt.Println("2 Minutes is equivalent to", cvr.mintosec(2), "seconds")
	fmt.Println("120 seconds is equivalent to", cvr.sectomin(120), "Minutes")
	fmt.Println("36 Degrees celsius is equivalent to", cvr.celtofah(36), "Fahrenheits")
	fmt.Println("100 Fahrenheits is equivalent to", cvr.Fahtocel(100), "Degrees")
	fmt.Println("45 Degrees is equivalent to", cvr.degtorad(45), "Rads")
	fmt.Println("4.5 Radians is equivalent to", cvr.radtodeg(4.5), "Degrees")
	fmt.Println("20 Kilograms is equivalent to", cvr.kgtopounds(20), "Pounds")
	fmt.Println("5 Pounds is equivalent to", cvr.poundstokg(5), "Kilograms")
}

	
