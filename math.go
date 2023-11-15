package math

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Function that returns the value of the absolute value of a number
func Abs(x float64) float64 {
	if x < 0 {
		return (-x)
	}
	return x
}

// Function that returns the value of the arc cosine of a number
func Acos(x float64) float64 {
	return math.Acos(x)
}

// Function that returns the value of the arc hyperbolic cosine of a number
func Acosh(x float64) float64 {
	return math.Acosh(x)
}

// Function that returns the value of the arc sine of a number
func Asin(x float64) float64 {
	return math.Asin(x)
}

// Function that returns the value of the arc hyperbolic sine of a number
func Asinh(x float64) float64 {
	return math.Asinh(x)
}

// Function that returns the value of the arc tangent of a number
func Atan(x float64) float64 {
	return math.Atan(x)
}

// Function that returns the value of the arc hyperbolic tangent of a number
func Atanh(x float64) float64 {
	return math.Atanh(x)
}

// Function that returns the value of the cube root of a number
func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

// Function that returns the value of the ceiling of a number
func Ceil(x float64) float64 {
	var xi int
	xi = int(x)
	if x > float64(xi) {
		return float64((xi + 1))
	}
	return float64(xi)
}

// Function that returns the value of cosinus of an angle
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Function that returns the value of the hyperbolic cosine of a number
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Function that converts degrees to radians.
func DegreesToRadians(x float64) float64 {
	return x * (math.Pi / 180.0)
}

// Function that returns the value of the exponential of a number
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Function that returns the value of the factorial of a number
func Fact(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x == 0 {
		return 1
	}
	result := float64(1)
	for i := float64(2); i <= x; i++ {
		result *= i
	}
	return result
}

// Function that returns the value of the floor of a number
func Floor(x float64) float64 {
	var xi int

	xi = int(x)
	if x < float64(xi) {
		return float64((xi - 1))
	}
	return float64(xi)
}

// Function that returns the value of the natural logarithm of a number
func Ln(x float64) float64 {
	return math.Log(x)
}

// Function that returns the base-10 logarithm of a number.
func Log10(x float64) float64 {
	return math.Log10(x)
}

// Function that returns the value of the maximum of two numbers
func Max(x float64, y float64) float64 {
	if x >= y {
		return (x)
	}
	return y
}

// Function that returns the value of the minimum of two numbers
func Min(x float64, y float64) float64 {
	if y <= x {
		return (y)
	}
	return x
}

// Function that returns the remainder of a division operation.
func Modulo(x, y float64) float64 {
	return math.Mod(x, y)
}

// Function that returns the value of PI
func Pi() float64 {
	return math.Pi
}

// Function that returns the value of the power of a number
func Pow(x any, y any) (float64, error) {
	var xFloat, yFloat float64
	var err error

	switch typeX := x.(type) {
	case int:
		xFloat = float64(typeX)
	case float64:
		xFloat = typeX
	default:
		return 0, fmt.Errorf("invalid type")
	}

	switch typeY := y.(type) {
	case int:
		yFloat = float64(typeY)
	case float64:
		yFloat = typeY
	default:
		return 0, fmt.Errorf("invalid type")
	}

	if err != nil {
		return 0, fmt.Errorf("invalid type")
	}

	return math.Pow(xFloat, yFloat), nil
}

// Function that converts radians to degrees.
func RadiansToDegrees(x float64) float64 {
	return x * (180.0 / math.Pi)
}

// Function that generates a random number between a given range.
func Random(x, y float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return x + rand.Float64()*(y-x)
}

// Function that rounds a number to the nearest integer.
func Round(x float64) float64 {
	return math.Round(x)
}

// Function that returns the value of sinus of an angle
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Function that returns the value of the hyperbolic sine of a number
func Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Function that returns the value of the square root of a number
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Function that returns the value of the tangent of an angle
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Function that returns the value of the hyperbolic tangent of a number
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

// Function that returns the value of the trunc of a number
func Trunc(x float64) float64 {
	return float64(int(x))
}
