package mathlibtester

import (
	"math"
	"testing"

	mathPackage "github.com/Eclalang/math"
)

func TestPiFucntion(t *testing.T) {
	t.Log("TestAngleFucntions")
	want := math.Pi
	got := mathPackage.Pi()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestCosFucntion(t *testing.T) {
	t.Log("TestCosFucntion")
	want := math.Cos(1)
	got := mathPackage.Cos(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cos(30)
	got = mathPackage.Cos(30)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cos(90)
	got = mathPackage.Cos(90)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cos(-180)
	got = mathPackage.Cos(-180)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSinFucntion(t *testing.T) {
	t.Log("TestSinFucntion")
	want := math.Sin(1)
	got := mathPackage.Sin(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sin(45)
	got = mathPackage.Sin(45)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sin(180)
	got = mathPackage.Sin(180)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sin(0)
	got = mathPackage.Sin(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sin(-34)
	got = mathPackage.Sin(-34)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestTanFucntion(t *testing.T) {
	t.Log("TestTanFucntion")
	want := math.Tan(1)
	got := mathPackage.Tan(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Tan(10)
	got = mathPackage.Tan(10)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Tan(-25)
	got = mathPackage.Tan(-25)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Tan(90)
	got = mathPackage.Tan(90)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Tan(0)
	got = mathPackage.Tan(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestLnFucntion(t *testing.T) {
	t.Log("TestLnFucntion")
	want := math.Log(1)
	got := mathPackage.Ln(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Log(math.Exp(1))
	got = mathPackage.Ln(math.Exp(1))
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Log(45)
	got = mathPackage.Ln(45)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Log(0)
	got = mathPackage.Ln(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestExpFucntion(t *testing.T) {
	t.Log("TestExpFucntion")
	want := math.Exp(1)
	got := mathPackage.Exp(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Exp(45)
	got = mathPackage.Exp(45)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Exp(0)
	got = mathPackage.Exp(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Exp(-5)
	got = mathPackage.Exp(-5)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSqrtFucntion(t *testing.T) {
	t.Log("TestSqrtFucntion")
	want := math.Sqrt(1)
	got := mathPackage.Sqrt(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sqrt(16)
	got = mathPackage.Sqrt(16)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sqrt(81)
	got = mathPackage.Sqrt(81)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Sqrt(27)
	got = mathPackage.Sqrt(27)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestCbrtFucntion(t *testing.T) {
	t.Log("TestCbrtFucntion")
	want := math.Cbrt(1)
	got := mathPackage.Cbrt(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cbrt(27)
	got = mathPackage.Cbrt(27)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cbrt(16)
	got = mathPackage.Cbrt(16)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cbrt(125)
	got = mathPackage.Cbrt(125)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cbrt(-10)
	got = mathPackage.Cbrt(-10)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Cbrt(0)
	got = mathPackage.Cbrt(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestPowFucntion(t *testing.T) {
	t.Log("TestPowFucntion")
	want := math.Pow(1, 1)
	got := mathPackage.Pow(1, 1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Pow(10, 3)
	got = mathPackage.Pow(10, 3)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Pow(-8, 7)
	got = mathPackage.Pow(-8, 7)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Pow(0, 0)
	got = mathPackage.Pow(0, 0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestFactFucntion(t *testing.T) {
	t.Log("TestFactFucntion")
	want := mathPackage.Fact(1)
	got := 1
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = mathPackage.Fact(0)
	got = 1
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = mathPackage.Fact(3)
	got = 6
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestAbsFunction(t *testing.T) {
	t.Log("TestAbsFunction")
	want := math.Abs(-1)
	got := mathPackage.Abs(-1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Abs(34)
	got = mathPackage.Abs(34)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Abs(-56)
	got = mathPackage.Abs(-56)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Abs(-103)
	got = mathPackage.Abs(-103)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Abs(23)
	got = mathPackage.Abs(23)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Abs(0)
	got = mathPackage.Abs(0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestFloorFucntion(t *testing.T) {
	t.Log("TestFloorFucntion")
	want := math.Floor(1)
	got := mathPackage.Floor(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Floor(5.95)
	got = mathPackage.Floor(5.95)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Floor(5.05)
	got = mathPackage.Floor(5.05)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Floor(-5.05)
	got = mathPackage.Floor(-5.05)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestCeilFucntion(t *testing.T) {
	t.Log("TestCeilFucntion")
	want := math.Ceil(1)
	got := mathPackage.Ceil(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Ceil(0.95)
	got = mathPackage.Ceil(0.95)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Ceil(7.004)
	got = mathPackage.Ceil(7.004)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Ceil(-7.004)
	got = mathPackage.Ceil(-7.004)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestTruncFunction(t *testing.T) {
	t.Log("TestTruncFunction")
	want := math.Trunc(1)
	got := mathPackage.Trunc(1)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Trunc(13.37)
	got = mathPackage.Trunc(13.37)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Trunc(42.84)
	got = mathPackage.Trunc(42.84)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Trunc(0.123)
	got = mathPackage.Trunc(0.123)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
func TestMaxFunction(t *testing.T) {
	t.Log("TestMaxFunction")
	want := math.Max(1, 2)
	got := mathPackage.Max(1, 2)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Max(-6, -2)
	got = mathPackage.Max(-6, -2)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Max(0, 0)
	got = mathPackage.Max(0, 0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Max(6, 7)
	got = mathPackage.Max(6, 7)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Max(-63, 7)
	got = mathPackage.Max(-63, 7)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
func TestMinFunction(t *testing.T) {
	t.Log("TestMinFunction")
	want := math.Min(1, 2)
	got := mathPackage.Min(1, 2)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Min(-6, -2)
	got = mathPackage.Min(-6, -2)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Min(0, 0)
	got = mathPackage.Min(0, 0)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Min(6, 7)
	got = mathPackage.Min(6, 7)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	want = math.Min(-63, 7)
	got = mathPackage.Min(-63, 7)
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
