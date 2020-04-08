package main

import (
	"fmt"
	"math"
	"os"
)

type Converter struct{}

type Feet float64
type Centimeter float64
type Minutes float64
type Milliseconds float64
type Seconds float64
type Celsius float64
type Fahrenheit float64
type Radian float64
type Degree float64
type Kilogram float64
type Pounds float64

//method attaching to struct, name of func, parameters, type of parameter, func signature
func (cvr Converter) FeetToCentimeter(c Feet) Centimeter {
	return Centimeter(30.48 * c)
}

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(0.0328 * c)
}

func (cvr Converter) MinutesToSeconds(c Minutes) Seconds {
	return Seconds(60 * c)
}

func (cvr Converter) SecondsToMilliseconds(c Seconds) Milliseconds {
	return Milliseconds(1000 * c)
}

func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(33.8 * c)
}

func (cvr Converter) FahrenheitToCelsius(c Fahrenheit) Celsius {
	return Celsius(-17.222 * c)
}

func (cvr Converter) RadianToDegree(c Radian) Degree {
	return Degree(180 / math.Pi * c)
}

func (cvr Converter) DegreeToRadian(c Degree) Radian {
	return Radian(math.Pi / 180 * c)
}

func (cvr Converter) KilogramToPounds(c Kilogram) Pounds {
	return Pounds(2.2046 * c)
}

func (cvr Converter) PoundsToKilogram(c Pounds) Kilogram {
	return Kilogram(0.45359 * c)
}

func main() {

	fmt.Printf(`
		1. Feet To Centimeter
		2. Centimeter To Feet
		3. Minutes To Seconds
		4. Seconds To Milliseconds
		5. Celsius To Fahrenheit
		6. Fahrenheit To Celsius
		7. Radian To Degree
		8. Degree To Radian
		9. Kilogram To Pounds
		10. Pounds To Kilogram
		11. Quit

		Input corresponding key to convert
	`)

	var key int
	fmt.Scan(&key)

	cvr := Converter{}

	switch key {
	case 1:
		fmt.Println("Feet To Centimeter")
		F := Feet(getInput())
		fmt.Println(cvr.FeetToCentimeter(F))
	case 2:
		fmt.Println("Centimeter To Feet")
		C := Centimeter(getInput())
		fmt.Println(cvr.CentimeterToFeet(C))
	case 3:
		fmt.Println("Minutes To Seconds")
		M := Minutes(getInput())
		fmt.Println(cvr.MinutesToSeconds(M))
	case 4:
		fmt.Println("Seconds To Milliseconds")
		S := Seconds(getInput())
		fmt.Println(cvr.SecondsToMilliseconds(S))
	case 5:
		fmt.Println("Celsius To Fahrenheit")
		Ce := Celsius(getInput())
		fmt.Println(cvr.CelsiusToFahrenheit(Ce))
	case 6:
		fmt.Println("Fahrenheit To Celsius")
		Fa := Fahrenheit(getInput())
		fmt.Println(cvr.FahrenheitToCelsius(Fa))
	case 7:
		fmt.Println("Radian To Degree")
		R := Radian(getInput())
		fmt.Println(cvr.RadianToDegree(R))
	case 8:
		fmt.Println("Degree To Radian")
		D := Degree(getInput())
		fmt.Println(cvr.DegreeToRadian(D))
	case 9:
		fmt.Println("Kilogram To Pounds")
		K := Kilogram(getInput())
		fmt.Println(cvr.KilogramToPounds(K))
	case 10:
		fmt.Println("Pounds To Kilogram")
		P := Pounds(getInput())
		fmt.Println(cvr.PoundsToKilogram(P))
	case 11:
		fmt.Println("Quitting")
		os.Exit(0)
	default:
		fmt.Println("Not an Option, Try again")
	}
}

// getInput Gets user input
func getInput() float64 {
	var value float64

	fmt.Println("Please input value: ")
	fmt.Scan(&value)

	return value
}