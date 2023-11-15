package mathlibtester

import (
	"math"
	"testing"

	mathPackage "github.com/Eclalang/math"
)

func TestAbs(t *testing.T) {
	abs := -1.0
	expected := math.Abs(abs)
	got := mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
	abs = -1
	expected = math.Abs(abs)
	got = mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
	abs = -1
	expected = math.Abs(abs)
	got = mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
	abs = -1
	expected = math.Abs(abs)
	got = mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
	abs = -1
	expected = math.Abs(abs)
	got = mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
	abs = -1
	expected = math.Abs(abs)
	got = mathPackage.Abs(abs)
	if expected != got {
		t.Errorf("mathPackage.Abs(%v), got %v expected %v", abs, got, expected)
	}
}

func TestAcos(t *testing.T) {
	acos := 0.5
	expected := mathPackage.Acos(acos)
	got := math.Acos(acos)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Acos(%v) got %v, expected %v", acos, got, expected)
	}
}

func TestAcosh(t *testing.T) {
	acosh := 2.0
	expected := mathPackage.Acosh(float64(acosh))
	got := math.Acosh(acosh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Acosh(%v) got %v, expected %v", acosh, got, expected)
	}
}

func TestAsin(t *testing.T) {
	asin := 0.5
	expected := mathPackage.Asin(asin)
	got := math.Asin(asin)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Asin(%v) got %v, expected %v", asin, got, expected)
	}
}

func TestAsinh(t *testing.T) {
	asinh := 2.0
	expected := mathPackage.Asinh(asinh)
	got := math.Asinh(asinh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Asinh(%v) got %v, expected %v", asinh, got, expected)
	}
}

func TestAtan(t *testing.T) {
	atan := 0.5
	expected := mathPackage.Atan(atan)
	got := math.Atan(atan)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Atan(%v) got %v, expected %v", atan, got, expected)
	}
}

func TestAtanh(t *testing.T) {
	atanh := 0.5
	expected := mathPackage.Atanh(atanh)
	got := math.Atanh(atanh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Atanh(%v) got %v, expected %v", atanh, got, expected)
	}
}

func TestCbrt(t *testing.T) {
	cbrt := 1.0
	expected := math.Cbrt(cbrt)
	got := mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
	cbrt = 27.0
	expected = math.Cbrt(cbrt)
	got = mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
	cbrt = 16.0
	expected = math.Cbrt(cbrt)
	got = mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
	cbrt = 125.0
	expected = math.Cbrt(cbrt)
	got = mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
	cbrt = -10.0
	expected = math.Cbrt(cbrt)
	got = mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
	cbrt = 0
	expected = math.Cbrt(cbrt)
	got = mathPackage.Cbrt(cbrt)
	if expected != got {
		t.Errorf("mathPackage.Cbrt(%v), got %v expected %v", cbrt, got, expected)
	}
}

func TestCeil(t *testing.T) {
	ceil := 1.0
	expected := math.Ceil(ceil)
	got := mathPackage.Ceil(ceil)
	if expected != got {
		t.Errorf("mathPackage.ceil(%v) got %v expected %v", ceil, got, expected)
	}
	ceil = 0.95
	expected = math.Ceil(ceil)
	got = mathPackage.Ceil(ceil)
	if expected != got {
		t.Errorf("mathPackage.ceil(%v) got %v expected %v", ceil, got, expected)
	}
	ceil = 7.004
	expected = math.Ceil(ceil)
	got = mathPackage.Ceil(ceil)
	if expected != got {
		t.Errorf("mathPackage.ceil(%v) got %v expected %v", ceil, got, expected)
	}
	ceil = -7.004
	expected = math.Ceil(ceil)
	got = mathPackage.Ceil(ceil)
	if expected != got {
		t.Errorf("mathPackage.ceil(%v) got %v expected %v", ceil, got, expected)
	}
}

func TestCos(t *testing.T) {
	cos := 1.0
	expected := math.Cos(cos)
	got := mathPackage.Cos(cos)
	if expected != got {
		t.Errorf("mathPackage.Cos(%v) got %v, expected %v", cos, got, expected)
	}
	cos = 30.0
	expected = math.Cos(cos)
	got = mathPackage.Cos(cos)
	if expected != got {
		t.Errorf("mathPackage.Cos(%v) got %v, expected %v", cos, got, expected)
	}
	cos = 90.0
	expected = math.Cos(cos)
	got = mathPackage.Cos(cos)
	if expected != got {
		t.Errorf("mathPackage.Cos(%v), got %v, expected %v", cos, got, expected)
	}
	cos = -180.0
	expected = math.Cos(cos)
	got = mathPackage.Cos(cos)
	if expected != got {
		t.Errorf("mathPackage.Cos(%v) got %v, expected %v", cos, got, expected)
	}
}

func TestCosh(t *testing.T) {
	cosh := 0.5
	expected := mathPackage.Cosh(cosh)
	got := math.Cosh(cosh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Cosh(%v) returned %v, expected %v", cosh, got, expected)
	}
}

func TestDegreesToRadians(t *testing.T) {
	deg := 180.0
	expected := math.Pi
	got := mathPackage.DegreesToRadians(deg)
	if got != expected {
		t.Errorf("DegreesToRadians(%v), got %v, expected %v", deg, got, expected)
	}
}

func TestExp(t *testing.T) {
	exp := 1.0
	expected := math.Exp(exp)
	got := mathPackage.Exp(exp)
	if expected != got {
		t.Errorf("mathPackage.Exp(%v) got %v expected %v", exp, got, expected)
	}
	exp = 45.0
	expected = math.Exp(exp)
	got = mathPackage.Exp(exp)
	if expected != got {
		t.Errorf("mathPackage.Exp(%v) got %v expected %v", exp, got, expected)
	}
	exp = 0
	expected = math.Exp(exp)
	got = mathPackage.Exp(exp)
	if expected != got {
		t.Errorf("mathPackage.Exp(%v) got %v expected %v", exp, got, expected)
	}
	exp = -5.0
	expected = math.Exp(exp)
	got = mathPackage.Exp(exp)
	if expected != got {
		t.Errorf("mathPackage.Exp(%v) got %v expected %v", exp, got, expected)
	}
}

func TestFact(t *testing.T) {
	fact := 0
	expected := 1.0
	got := mathPackage.Fact(0)
	if expected != got {
		t.Errorf("mathPackage.Fact(%v), got %v, expected %v", fact, got, expected)
	}
	fact = 1.0
	expected = 1.0
	got = mathPackage.Fact(1)
	if expected != got {
		t.Errorf("mathPackage.Fact(%v), got %v, expected %v", fact, got, expected)
	}
	fact = 6.0
	expected = 6.0
	got = mathPackage.Fact(3)
	if expected != got {
		t.Errorf("mathPackage.Fact(%v), got %v, expected %v", fact, got, expected)
	}
	fact = -2.0
	expected = 0
	got = mathPackage.Fact(-2)
	if expected != got {
		t.Errorf("mathPackage.Fact(%v), got %v, expected %v", fact, got, expected)
	}
}

func TestFloor(t *testing.T) {
	floor := 1.0
	expected := math.Floor(floor)
	got := mathPackage.Floor(floor)
	if expected != got {
		t.Errorf("mathPackage.Floor(%v) got %v expected %v", floor, got, expected)
	}
	floor = 5.95
	expected = math.Floor(floor)
	got = mathPackage.Floor(floor)
	if expected != got {
		t.Errorf("mathPackage.Floor(%v) got %v expected %v", floor, got, expected)
	}
	floor = 5.05
	expected = math.Floor(floor)
	got = mathPackage.Floor(floor)
	if expected != got {
		t.Errorf("mathPackage.Floor(%v) got %v expected %v", floor, got, expected)
	}
	floor = -5.05
	expected = math.Floor(floor)
	got = mathPackage.Floor(floor)
	if expected != got {
		t.Errorf("mathPackage.Floor(%v) got %v expected %v", floor, got, expected)
	}
}

func TestLn(t *testing.T) {
	ln := 1.0
	expected := math.Log(ln)
	got := mathPackage.Ln(ln)
	if expected != got {
		t.Errorf("mathPackage.Ln(%v) got %v expected %v", ln, got, expected)
	}
	ln = math.Exp(1)
	expected = math.Log(math.Exp(1))
	got = mathPackage.Ln(math.Exp(1))
	if expected != got {
		t.Errorf("mathPackage.Ln(%v) got %v expected %v", ln, got, expected)
	}
	ln = 45.0
	expected = math.Log(45)
	got = mathPackage.Ln(45)
	if expected != got {
		t.Errorf("mathPackage.Ln(%v) got %v expected %v", ln, got, expected)
	}
	ln = 0
	expected = math.Log(0)
	got = mathPackage.Ln(0)
	if expected != got {
		t.Errorf("mathPackage.Ln(%v) got %v expected %v", ln, got, expected)
	}
}

func TestLog10(t *testing.T) {
	log10 := 100.0
	expected := 2.0
	got := mathPackage.Log10(log10)
	if got != expected {
		t.Errorf("mathPackage.Log10(%v) got %v, expected %v", log10, got, expected)
	}
}

func TestMax(t *testing.T) {
	x := 1.0
	y := 2.0
	expected := math.Max(x, y)
	got := mathPackage.Max(x, y)
	if expected != got {
		t.Errorf("mathPackage.Max(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = -6.0
	y = -2.0
	expected = math.Max(x, y)
	got = mathPackage.Max(x, y)
	if expected != got {
		t.Errorf("mathPackage.Max(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 0
	y = 0
	expected = math.Max(x, y)
	got = mathPackage.Max(x, y)
	if expected != got {
		t.Errorf("mathPackage.Max(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 6.0
	y = 7.0
	expected = math.Max(x, y)
	got = mathPackage.Max(x, y)
	if expected != got {
		t.Errorf("mathPackage.Max(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = -63.0
	y = 7.0
	expected = math.Max(x, y)
	got = mathPackage.Max(x, y)
	if expected != got {
		t.Errorf("mathPackage.Max(%v, %v) got %v expected %v", x, y, got, expected)
	}
}

func TestMin(t *testing.T) {
	x := 1.0
	y := 2.0
	expected := math.Min(x, y)
	got := mathPackage.Min(x, y)
	if expected != got {
		t.Errorf("mathPackage.Min(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = -6.0
	y = -2.0
	expected = math.Min(x, y)
	got = mathPackage.Min(x, y)
	if expected != got {
		t.Errorf("mathPackage.Min(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 0
	y = 0
	expected = math.Min(x, y)
	got = mathPackage.Min(x, y)
	if expected != got {
		t.Errorf("mathPackage.Min(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 6.0
	y = 7.0
	expected = math.Min(x, y)
	got = mathPackage.Min(x, y)
	if expected != got {
		t.Errorf("mathPackage.Min(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = -63.0
	y = 7.0
	expected = math.Min(x, y)
	got = mathPackage.Min(x, y)
	if expected != got {
		t.Errorf("mathPackage.Min(%v, %v) got %v expected %v", x, y, got, expected)
	}
}

func TestModulo(t *testing.T) {
	x := 5.0
	y := 3.0
	expected := math.Mod(x, y)
	got := mathPackage.Modulo(x, y)
	if got != expected {
		t.Errorf("MathPackage.Modulo(%v, %v), got %v, expected %v", x, y, got, expected)
	}
	x = -5.0
	y = -3.0
	expected = math.Mod(x, y)
	got = mathPackage.Modulo(x, y)
	if got != expected {
		t.Errorf("MathPackage.Modulo(%v, %v), got %v, expected %v", x, y, got, expected)
	}
	x = 0
	y = 2.0
	expected = math.Mod(x, y)
	got = mathPackage.Modulo(x, y)
	if got != expected {
		t.Errorf("MathPackage.Modulo(%v, %v), got %v, expected %v", x, y, got, expected)
	}
}

func TestPi(t *testing.T) {
	expected := math.Pi
	got := mathPackage.Pi()
	if expected != got {
		t.Errorf("mathPackage.Pi, got %v expected %v", got, expected)
	}
}

func TestPow(t *testing.T) {
	x := 1.0
	y := 1.0
	expected := math.Pow(x, y)
	got, _ := mathPackage.Pow(x, y)
	if expected != got {
		t.Errorf("mathPackage.Pow(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 10.0
	y = 3.0
	expected = math.Pow(x, y)
	got, _ = mathPackage.Pow(x, y)
	if expected != got {
		t.Errorf("mathPackage.Pow(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = -8.0
	y = 7.0
	expected = math.Pow(x, y)
	got, _ = mathPackage.Pow(x, y)
	if expected != got {
		t.Errorf("mathPackage.Pow(%v, %v) got %v expected %v", x, y, got, expected)
	}
	x = 0
	y = 0
	expected = math.Pow(x, y)
	got, _ = mathPackage.Pow(x, y)
	if expected != got {
		t.Errorf("mathPackage.Pow(%v, %v) got %v expected %v", x, y, got, expected)
	}
}

func TestRadiansToDegrees(t *testing.T) {
	rad := math.Pi
	expected := 180.0
	got := mathPackage.RadiansToDegrees(rad)
	if got != expected {
		t.Errorf("mathPackage.RadiansToDegrees(%v), got %v, expected %v", rad, got, expected)
	}
}

func TestRandom(t *testing.T) {
	min := 1.0
	max := 10.0
	got := mathPackage.Random(min, max)
	if got < min || got > max {
		t.Errorf("math.package.Random(%v, %v), got %v, expected a number between %v and %v", min, max, got, min, max)
	}
}

func TestRound(t *testing.T) {
	round := 2.49
	expected := math.Round(round)
	got := mathPackage.Round(round)
	if got != expected {
		t.Errorf("mathPackage.Round(%v), got %v, expected %v", round, got, expected)
	}
}

func TestSin(t *testing.T) {
	sin := 1.0
	expected := math.Sin(sin)
	got := mathPackage.Sin(sin)
	if expected != got {
		t.Errorf("mathPackage.Sin(%v), got %v expected %v", sin, got, expected)
	}
	sin = 45.0
	expected = math.Sin(sin)
	got = mathPackage.Sin(sin)
	if expected != got {
		t.Errorf("mathPackage.Sin(%v), got %v expected %v", sin, got, expected)
	}
	sin = 180.0
	expected = math.Sin(sin)
	got = mathPackage.Sin(sin)
	if expected != got {
		t.Errorf("mathPackage.Sin(%v), got %v expected %v", sin, got, expected)
	}
	sin = 0
	expected = math.Sin(sin)
	got = mathPackage.Sin(sin)
	if expected != got {
		t.Errorf("mathPackage.Sin(%v), got %v expected %v", sin, got, expected)
	}
	sin = -34.0
	expected = math.Sin(sin)
	got = mathPackage.Sin(sin)
	if expected != got {
		t.Errorf("mathPackage.Sin(%v), got %v expected %v", sin, got, expected)
	}
}

func TestSinh(t *testing.T) {
	sinh := 0.5
	expected := mathPackage.Sinh(sinh)
	got := math.Sinh(sinh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Sinh(%v) got %v, expected %v", sinh, got, expected)
	}
}

func TestSqrt(t *testing.T) {
	sqrt := 1.0
	expected := math.Sqrt(sqrt)
	got := mathPackage.Sqrt(sqrt)
	if expected != got {
		t.Errorf("mathPackage.Sqrt(%v), got %v expected %v", sqrt, got, expected)
	}
	sqrt = 27.0
	expected = math.Sqrt(sqrt)
	got = mathPackage.Sqrt(sqrt)
	if expected != got {
		t.Errorf("mathPackage.Sqrt(%v), got %v expected %v", sqrt, got, expected)
	}
	sqrt = 16.0
	expected = math.Sqrt(sqrt)
	got = mathPackage.Sqrt(sqrt)
	if expected != got {
		t.Errorf("mathPackage.Sqrt(%v), got %v expected %v", sqrt, got, expected)
	}
	sqrt = 125.0
	expected = math.Sqrt(sqrt)
	got = mathPackage.Sqrt(sqrt)
	if expected != got {
		t.Errorf("mathPackage.Sqrt(%v), got %v expected %v", sqrt, got, expected)
	}
	sqrt = 0
	expected = math.Sqrt(sqrt)
	got = mathPackage.Sqrt(sqrt)
	if expected != got {
		t.Errorf("mathPackage.Sqrt(%v), got %v expected %v", sqrt, got, expected)
	}
}

func TestTan(t *testing.T) {
	tan := 1.0
	expected := math.Tan(tan)
	got := mathPackage.Tan(tan)
	if expected != got {
		t.Errorf("mathPackage.Tan(%v), got %v expected %v", tan, got, expected)
	}
	tan = 10.0
	expected = math.Tan(tan)
	got = mathPackage.Tan(tan)
	if expected != got {
		t.Errorf("mathPackage.Tan(%v), got %v expected %v", tan, got, expected)
	}
	tan = -25.0
	expected = math.Tan(tan)
	got = mathPackage.Tan(tan)
	if expected != got {
		t.Errorf("mathPackage.Tan(%v), got %v expected %v", tan, got, expected)
	}
	tan = 90.0
	expected = math.Tan(tan)
	got = mathPackage.Tan(tan)
	if expected != got {
		t.Errorf("mathPackage.Tan(%v), got %v expected %v", tan, got, expected)
	}
	tan = 0
	expected = math.Tan(tan)
	got = mathPackage.Tan(tan)
	if expected != got {
		t.Errorf("mathPackage.Tan(%v), got %v expected %v", tan, got, expected)
	}
}

func TestTanh(t *testing.T) {
	tanh := 0.5
	expected := mathPackage.Tanh(tanh)
	got := math.Tanh(tanh)
	if math.Abs(expected-got) > 1e-9 {
		t.Errorf("mathPackage.Tanh(%v) got %v, expected %v", tanh, got, expected)
	}
}

func TestTrunc(t *testing.T) {
	trunc := 1.0
	expected := math.Trunc(trunc)
	got := mathPackage.Trunc(trunc)
	if expected != got {
		t.Errorf("mathPackage.Trunc(%v), got %v, expected %v", trunc, got, expected)
	}
	trunc = 13.37
	expected = math.Trunc(trunc)
	got = mathPackage.Trunc(trunc)
	if expected != got {
		t.Errorf("mathPackage.Trunc(%v), got %v, expected %v", trunc, got, expected)
	}
	trunc = 42.84
	expected = math.Trunc(trunc)
	got = mathPackage.Trunc(trunc)
	if expected != got {
		t.Errorf("mathPackage.Trunc(%v), got %v, expected %v", trunc, got, expected)
	}
	trunc = 0.123
	expected = math.Trunc(trunc)
	got = mathPackage.Trunc(trunc)
	if expected != got {
		t.Errorf("mathPackage.Trunc(%v), got %v, expected %v", trunc, got, expected)
	}
}
