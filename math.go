package math

import (
	"math"
	"math/rand"
	"time"
)

// Function that returns the value of PI
func Pi() float64 {
	return math.Pi
}

// Function that returns the value of cosinus of an angle
func Cos(Angle float64) float64 {
	return math.Cos(Angle)
}

// Function that returns the value of sinus of an angle
func Sin(Angle float64) float64 {
	return math.Sin(Angle)
}

// Function that returns the value of the tangent of an angle
func Tan(Angle float64) float64 {
	return math.Tan(Angle)
}

// Function that returns the value of the natural logarithm of a number
func Ln(nb float64) float64 {
	return math.Log(nb)
}

// Function that returns the value of the exponential of a number
func Exp(nb float64) float64 {
	return math.Exp(nb)
}

// Function that returns the value of the square root of a number
func Sqrt(nb float64) float64 {
	return math.Sqrt(nb)
}

// Function that returns the value of the cube root of a number
func Cbrt(nb float64) float64 {
	return math.Cbrt(nb)
}

// Function that returns the value of the power of a number
func Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

// Function that returns the value of the factorial of a number
func Fact(nb uint64) uint64 {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := uint64(1)
	for i := uint64(2); i <= nb; i++ {
		result *= i
	}
	return result
}

// Function that returns the value of the absolute value of a number
func Abs(nb float64) float64 {
	if nb < 0 {
		return (-nb)
	}
	return nb
}

// Function that returns the base-10 logarithm of a number.
func Log10(nb float64) float64 {
	return math.Log10(nb)
}

// Function that rounds a number to the nearest integer.
func Round(nb float64) float64 {
	return math.Round(nb)
}

// Function that converts degrees to radians.
func DegreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

// Function that converts radians to degrees.
func RadiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

// Function that returns the remainder of a division operation.
func Modulo(a, b float64) float64 {
	return math.Mod(a, b)
}

// Function that generates a random number between a given range.
func Random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

// Function that returns the value of the floor of a number
func Floor(nb float64) float64 {
	var nbi int

	nbi = int(nb)
	if nb < float64(nbi) {
		return float64((nbi - 1))
	}
	return float64(nbi)
}

// Function that returns the value of the ceiling of a number
func Ceil(nb float64) float64 {
	var nbi int
	nbi = int(nb)
	if nb > float64(nbi) {
		return float64((nbi + 1))
	}
	return float64(nbi)
}

// Function that returns the value of the trunc of a number
func Trunc(nb float64) float64 {
	return float64(int(nb))
}

// Function that returns the value of the maximum of two numbers
func Max(a float64, b float64) float64 {
	if a >= b {
		return (a)
	}
	return b
}

// Function that returns the value of the minimum of two numbers
func Min(a float64, b float64) float64 {
	if b <= a {
		return (b)
	}
	return a
}

// Function that returns the value of the arc cosine of a number
func Acos(x float64) float64 {
	return math.Acos(x)
}

// Function that returns the value of the arc sine of a number
func Asin(x float64) float64 {
	return math.Asin(x)
}

// Function that returns the value of the arc tangent of a number
func Atan(x float64) float64 {
	return math.Atan(x)
}

// Function that returns the value of the hyperbolic cosine of a number
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Function that returns the value of the hyperbolic sine of a number
func Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Function that returns the value of the hyperbolic tangent of a number
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

// Function that returns the value of the arc hyperbolic cosine of a number
func Acosh(x float64) float64 {
	return math.Acosh(x)
}

// Function that returns the value of the arc hyperbolic sine of a number
func Asinh(x float64) float64 {
	return math.Asinh(x)
}

// Function that returns the value of the arc hyperbolic tangent of a number
func Atanh(x float64) float64 {
	return math.Atanh(x)
}
