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
	// return (((11/60)*float32(Pow(int((nb-1)), 3)) + float32(Pow(int((nb-1)), 2))) / (-0.5 + (3/2)*nb + (3/5)*float32(Pow(int(nb-1), 2)) + ((1 / 20) * float32(Pow(int(nb-1), 3)))))
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
	if nb == 0 {
		return (0)
	}
	if nb < 0 {
		return (-Cbrt(-nb))
	}

	var r float32 = float32(nb)
	var ex int = 0

	for r < 0.125 {
		r *= 8
		ex--
	}
	for r > 1.0 {
		r *= 0.125
		ex++
	}
	r = (-0.46946116*r+1.072302)*r + 0.3812513

	for ex < 0 {
		r *= 0.5
		ex++
	}
	for ex > 0 {
		r *= 2
		ex--
	}
	r = (2.0/3.0)*r + (1.0/3.0)*float32(nb)/(r*r)
	r = (2.0/3.0)*r + (1.0/3.0)*float32(nb)/(r*r)
	r = (2.0/3.0)*r + (1.0/3.0)*float32(nb)/(r*r)
	r = (2.0/3.0)*r + (1.0/3.0)*float32(nb)/(r*r)

	return (int(r))
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
