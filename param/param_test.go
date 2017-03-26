package param

import (
	"math"
	"testing"
)

type Val string

func (v Val) Value(key string) string {
	return string(v)
}

func TestString(t *testing.T) {
	var tests = []struct {
		val string
		out string
	}{
		{"test", "test"},
	}
	for _, tt := range tests {
		if got := String(Val(tt.val), "v"); got != tt.out {
			t.Errorf("String() = %v, want %v", got, tt.out)
		}
	}
}

func TestBool(t *testing.T) {
	var tests = []struct {
		val string
		out bool
	}{
		{"true", true},
		{"TRUE", true},
		{"True", true},
		{"t", true},
		{"T", true},
		{"1", true},

		{"false", false},
		{"FALSE", false},
		{"f", false},
		{"F", false},
		{"0", false},
		// other value is flase
		{"2", false},
		{"test", false},
		{"TRue", false},
		{"truE", false},
	}
	for _, tt := range tests {
		if got := Bool(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Bool() = %v, want %v", got, tt.out)
		}
	}
}

func TestIntBase(t *testing.T) {
	var tests = []struct {
		val string
		out int
	}{
		{"0", 0},
		{"-123", -123},
		{"111", 111},
		{"test", 0},
		{"11a", 0},
		{"0x1", 0},
		{"-1.1", 0},
	}
	t.Run("Int", func(t *testing.T) {
		for _, tt := range tests {
			if got := Int(Val(tt.val), "v"); got != tt.out {
				t.Errorf("Int(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Int8", func(t *testing.T) {
		for _, tt := range tests {
			if got := Int8(Val(tt.val), "v"); got != int8(tt.out) {
				t.Errorf("Int8(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Int16", func(t *testing.T) {
		for _, tt := range tests {
			if got := Int16(Val(tt.val), "v"); got != int16(tt.out) {
				t.Errorf("Int16(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Int32", func(t *testing.T) {
		for _, tt := range tests {
			if got := Int32(Val(tt.val), "v"); got != int32(tt.out) {
				t.Errorf("Int32(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Int64", func(t *testing.T) {
		for _, tt := range tests {
			if got := Int64(Val(tt.val), "v"); got != int64(tt.out) {
				t.Errorf("Int64(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
}

func TestInt8(t *testing.T) {
	var tests = []struct {
		val string
		out int8
	}{
		{"-12333", -1 << 7},
		{"12333", 1<<7 - 1},
	}
	for _, tt := range tests {
		if got := Int8(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Int8(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestInt16(t *testing.T) {
	var tests = []struct {
		val string
		out int16
	}{
		{"-99999999", -1 << 15},
		{"9999999", 1<<15 - 1},
	}
	for _, tt := range tests {
		if got := Int16(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Int16(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestInt32(t *testing.T) {
	var tests = []struct {
		val string
		out int32
	}{
		{"-99999999999999999", -1 << 31},
		{"999999999999999999", 1<<31 - 1},
		{"-12345678", -12345678},
		{"12345678", 12345678},
	}
	for _, tt := range tests {
		if got := Int32(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Int32(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		val string
		out int64
	}{
		{"-999999999999999999999", -1 << 63},
		{"9999999999999999999999", 1<<63 - 1},
		{"-12345678", -12345678},
		{"12345678", 12345678},
	}
	for _, tt := range tests {
		if got := Int64(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Int64(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestUIntBase(t *testing.T) {
	var tests = []struct {
		val string
		out uint
	}{
		{"0", 0},
		{"-123", 0},
		{"111", 111},
		{"test", 0},
		{"11a", 0},
		{"0x1", 0},
		{"-1.1", 0},
	}
	t.Run("Uint", func(t *testing.T) {
		for _, tt := range tests {
			if got := Uint(Val(tt.val), "v"); got != tt.out {
				t.Errorf("Uint(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Uint8", func(t *testing.T) {
		for _, tt := range tests {
			if got := Uint8(Val(tt.val), "v"); got != uint8(tt.out) {
				t.Errorf("Uint8(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Byte", func(t *testing.T) {
		for _, tt := range tests {
			if got := Byte(Val(tt.val), "v"); got != byte(tt.out) {
				t.Errorf("Byte(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Uint16", func(t *testing.T) {
		for _, tt := range tests {
			if got := Uint16(Val(tt.val), "v"); got != uint16(tt.out) {
				t.Errorf("Uint16(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Uint32", func(t *testing.T) {
		for _, tt := range tests {
			if got := Uint32(Val(tt.val), "v"); got != uint32(tt.out) {
				t.Errorf("Uint32(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Uint64", func(t *testing.T) {
		for _, tt := range tests {
			if got := Uint64(Val(tt.val), "v"); got != uint64(tt.out) {
				t.Errorf("Uint64(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
}

func TestUint8(t *testing.T) {
	var tests = []struct {
		val string
		out uint8
	}{
		{"-12333", 0},
		{"12333", 1<<8 - 1},
	}
	for _, tt := range tests {
		if got := Uint8(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Uint8(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestUint16(t *testing.T) {
	var tests = []struct {
		val string
		out uint16
	}{
		{"-99999999", 0},
		{"9999999", 1<<16 - 1},
	}
	for _, tt := range tests {
		if got := Uint16(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Uint16(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestUint32(t *testing.T) {
	var tests = []struct {
		val string
		out uint32
	}{
		{"-99999999999999999", 0},
		{"999999999999999999", 1<<32 - 1},
	}
	for _, tt := range tests {
		if got := Uint32(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Uint32(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestUint64(t *testing.T) {
	var tests = []struct {
		val string
		out uint64
	}{
		{"-999999999999999999999", 0},
		{"9999999999999999999999", 1<<64 - 1},
	}
	for _, tt := range tests {
		if got := Uint64(Val(tt.val), "v"); got != tt.out {
			t.Errorf("Uint64(%v) = %v, want %v", tt.val, got, tt.out)
		}
	}
}

func TestFloatBase(t *testing.T) {
	var tests = []struct {
		val string
		out float64
	}{
		{"0", 0},
		{"-0", 0},
		{"-0.0", 0},
		{"-0.0", 0},
		{"-123", -123},
		{"123", 123},
		{"-1.11", -1.11},
		{"1.11", 1.11},
		{".11", 0.11},
		{"0.11", 0.11},
		{"0.0.1", 0},
		{"test", 0},
		{"11a", 0},
		{"0x1", 0},
	}
	t.Run("Float32", func(t *testing.T) {
		for _, tt := range tests {
			if got := Float32(Val(tt.val), "v"); got != float32(tt.out) {
				t.Errorf("Float32(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
	t.Run("Float64", func(t *testing.T) {
		for _, tt := range tests {
			if got := Float64(Val(tt.val), "v"); got != float64(tt.out) {
				t.Errorf("Float64(%v) = %v, want %v", tt.val, got, tt.out)
			}
		}
	})
}

func TestFloatNaN(t *testing.T) {
	// NaN is ignore letter case
	NaNList := []string{"nan", "NaN", "NAN", "nAn"}
	for _, NaN := range NaNList {
		if got := Float32(Val(NaN), "v"); !math.IsNaN(float64(got)) {
			t.Errorf("Float32(%v) = %v, want %v", NaN, got, math.NaN())
		}
		if got := Float64(Val(NaN), "v"); !math.IsNaN(got) {
			t.Errorf("Float64(%v) = %v, want %v", NaN, got, math.NaN())
		}
	}
}

func TestFloatInfnite(t *testing.T) {
	// infinity is ignore letter case
	infinities := []struct {
		val  string
		sign int
	}{
		{"inf", 1},
		{"+inf", 1},
		{"INF", 1},
		{"InF", 1},
		{"infinity", 1},
		{"inFinity", 1},
		{"+infinity", 1},

		{"-inf", -1},
		{"-INF", -1},
		{"-InF", -1},
		{"-infinity", -1},
		{"-inFinity", -1},
		{"-infinity", -1},
	}
	for _, inf := range infinities {
		if got := Float32(Val(inf.val), "v"); !math.IsInf(float64(got), inf.sign) {
			t.Errorf("Float32(%v) = %v, want %v", inf.val, got, math.Inf(inf.sign))
		}
		if got := Float64(Val(inf.val), "v"); !math.IsInf(got, inf.sign) {
			t.Errorf("Float32(%v) = %v, want %v", inf.val, got, math.Inf(inf.sign))
		}
	}
}
