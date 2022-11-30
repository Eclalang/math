package math

import "math"

// Angle functions
func Pi() float32 {
	return (math.Pi)
}

func Cos(adj float32, hyp float32) float32 {
	return (adj / hyp)
}

func Sin(opp float32, hyp float32) float32 {
	return (opp / hyp)
}

func Tan(opp float32, adj float32) float32 {
	return (opp / adj)
}

func Ln(nb float64) float64 {
	return (math.Log(nb))
}

func Exp(nb float64) float64 {
	return (math.Exp(nb))
}

func Sqrt(nb int) int {
	var sqrt float32
	var tmp float32

	if nb == 0 || nb == 1 {
		return (nb)
	}
	if nb < 0 {
		return (0)
	}
	sqrt = float32(nb) / 2
	tmp = 0
	for sqrt != tmp {
		tmp = sqrt
		sqrt = (float32(nb)/tmp + tmp)
	}
	if int(sqrt)*int(sqrt) != nb {
		return (0)
	}
	return (int(sqrt))
}

func Cbrt(nb int) int {
	return (int(math.Cbrt(float64(nb))))
}

func Pow(x int, y int) int {
	var k int
	var i int

	k = x
	if y < 0 {
		return (0)
	}
	if y == 0 {
		return (1)
	}
	i = 0
	for i < y {
		k *= x
		i++
	}
	return (k)
}

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
		nb *= -i
	}
	return (nb)
}

func Abs(nb int) int {
	if nb < 0 {
		return (-nb)
	}
	return (nb)
}

func Floor(nb float32) int {
	var nbi int

	nbi = int(nb)
	if nb < float32(nbi) {
		return (nbi - 1)
	}
	return (nbi)
}

func Ceil(nb float32) int {
	var nbi int

	nbi = int(nb)
	if nb == float32(nbi) {
		return (nbi)
	}
	return (nbi + 1)
}

func Trunc(nb float32) int {
	return (int(nb))
}

func Max(a float32, b float32) float32 {
	if a >= b {
		return (a)
	}
	return (b)
}
func Min(a float32, b float32) float32 {
	if b <= a {
		return (b)
	}
	return (a)
}
