package mathPackage

import "math"

// Function that returns the value of PI
func Pi() float64 {
	return (math.Pi)
}

// Function that returns the value of cosinus of an angle
func Cos(Angle float64) float64 {
	return (math.Cos(Angle))
}

// Function that returns the value of sinus of an angle
func Sin(Angle float64) float64 {
	return (math.Sin(Angle))
}

// Function that returns the value of the tangent of an angle
func Tan(Angle float64) float64 {
	return (math.Tan(Angle))
}

// Function that returns the value of the natural logarithm of a number
func Ln(nb float64) float64 {
	return (math.Log(nb))
}

// Function that returns the value of the exponential of a number
func Exp(nb float64) float64 {
	return (math.Exp(nb))
}

// Function that returns the value of the square root of a number
func Sqrt(nb float64) float64 {
	return (math.Sqrt(float64(nb)))
}

// Function that returns the value of the cube root of a number
func Cbrt(nb float64) float64 {
	return (math.Cbrt(nb))
}

// Function that returns the value of the power of a number
func Pow(x int, y int) float64 {
	var k int
	var i int

	k = x
	if y < 0 {
		return (0)
	}
	if y == 0 {
		return (1)
	}
	i = 1
	for i < y {
		k *= x
		i++
	}
	return (float64(k))
}

// Function that returns the value of the factorial of a number
func Fact(nb int) int {
	var i int
	if nb < 0 {
		return (0)
	}
	if nb == 0 {
		return (1)
	}
	i = nb - 1
	for i > 0 {
		nb *= i
		i--
	}
	return (nb)
}

// Function that returns the value of the absolute value of a number
func Abs(nb float64) float64 {
	if nb < 0 {
		return (-nb)
	}
	return (nb)
}

// Function that returns the value of the floor of a number
func Floor(nb float64) float64 {
	var nbi int

	nbi = int(nb)
	if nb < float64(nbi) {
		return float64((nbi - 1))
	}
	return (float64(nbi))
}

// Function that returns the value of the ceiling of a number
func Ceil(nb float64) float64 {
	var nbi int
	nbi = int(nb)
	if nb > float64(nbi) {
		return float64((nbi + 1))
	}
	return (float64(nbi))
}

// Function that returns the value of the trunc of a number
func Trunc(nb float64) float64 {
	return (float64(int(nb)))
}

// Function that returns the value of the maximum of two numbers
func Max(a float64, b float64) float64 {
	if a >= b {
		return (a)
	}
	return (b)
}

// Function that returns the value of the minimum of two numbers
func Min(a float64, b float64) float64 {
	if b <= a {
		return (b)
	}
	return (a)
}
