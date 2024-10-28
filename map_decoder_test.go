package fakers

import (
	"testing"

	tt "github.com/vingarcia/fakers/internal/testtools"
	"github.com/vingarcia/structscanner"
)

func TestMapTagDecoder(t *testing.T) {
	t.Run("should generate values for valid structs", func(t *testing.T) {
		var obj struct {
			String string
			Bool   bool

			Int    int
			Int8   int8
			Int16  int16
			Int32  int32
			Int64  int64
			Uint   uint
			Uint8  uint8
			Uint16 uint16
			Uint32 uint32
			Uint64 uint64

			F32 float32
			F64 float64

			Map       map[string]any
			Slice     []string
			Interface any
			Pointer   *int
			Func      func() string
			Chan      chan string

			IntArray [3]int
			StrArray [3]string

			Struct struct {
				String string
			}
		}
		err := structscanner.Decode(&obj, newCustomValuesMapDecoder(map[string]any{}))
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, obj.String, "FakeString")
		tt.AssertEqual(t, obj.Bool, false)

		tt.AssertEqual(t, obj.Int, 42)
		tt.AssertEqual(t, obj.Int8, int8(42))
		tt.AssertEqual(t, obj.Int16, int16(42))
		tt.AssertEqual(t, obj.Int32, int32(42))
		tt.AssertEqual(t, obj.Int64, int64(42))
		tt.AssertEqual(t, obj.Uint, uint(42))
		tt.AssertEqual(t, obj.Uint8, uint8(42))
		tt.AssertEqual(t, obj.Uint16, uint16(42))
		tt.AssertEqual(t, obj.Uint32, uint32(42))
		tt.AssertEqual(t, obj.Uint64, uint64(42))

		tt.AssertEqual(t, obj.F32, float32(42.0))
		tt.AssertEqual(t, obj.F64, 42.0)

		tt.AssertEqual(t, obj.Map, map[string]any(nil))
		tt.AssertEqual(t, obj.Slice, []string(nil))
		tt.AssertEqual(t, obj.Interface, nil)
		tt.AssertEqual(t, obj.Pointer, (*int)(nil))
		tt.AssertEqual(t, obj.Func == nil, true)
		tt.AssertEqual(t, obj.Chan == nil, true)

		tt.AssertEqual(t, obj.IntArray, [3]int{0, 0, 0})
		tt.AssertEqual(t, obj.StrArray, [3]string{"", "", ""})

		tt.AssertEqual(t, obj.Struct, struct {
			String string
		}{
			String: "FakeString",
		})
	})

	t.Run("should set custom values for valid structs", func(t *testing.T) {
		var obj struct {
			String string
			Bool   bool

			Int    int
			Int8   int8
			Int16  int16
			Int32  int32
			Int64  int64
			Uint   uint
			Uint8  uint8
			Uint16 uint16
			Uint32 uint32
			Uint64 uint64

			F32 float32
			F64 float64

			Map       map[string]any
			Slice     []string
			Func      func() string
			Interface any
			Chan      chan string
			Pointer   *int

			IntArray [3]int
			StrArray [3]string

			Struct struct {
				String string
			}
		}
		err := structscanner.Decode(&obj, newCustomValuesMapDecoder(map[string]any{
			"String": "SomeString",
			"Bool":   true,

			"Int":    20,
			"Int8":   21,
			"Int16":  22,
			"Int32":  23,
			"Int64":  24,
			"Uint":   25,
			"Uint8":  26,
			"Uint16": 27,
			"Uint32": 28,
			"Uint64": 29,

			"F32": 30.0,
			"F64": 30.1,

			"Map":       map[string]any{"key": "value"},
			"Slice":     []string{"foo", "bar"},
			"Interface": any(1),
			"Pointer":   ptr[int](24),
			"Func":      func() string { return "myFunc" },
			"Chan": func() chan string {
				ch := make(chan string, 1)
				ch <- "myChanMsg"
				return ch
			}(),

			"IntArray": [3]int{1, 2, 3},
			"StrArray": [3]string{"1", "2", "3"},

			"Struct": struct {
				String string
			}{
				String: "FakeString",
			},
		}))
		tt.AssertNoErr(t, err)

		tt.AssertEqual(t, obj.String, "SomeString")
		tt.AssertEqual(t, obj.Bool, true)

		tt.AssertEqual(t, obj.Int, 20)
		tt.AssertEqual(t, obj.Int8, int8(21))
		tt.AssertEqual(t, obj.Int16, int16(22))
		tt.AssertEqual(t, obj.Int32, int32(23))
		tt.AssertEqual(t, obj.Int64, int64(24))
		tt.AssertEqual(t, obj.Uint, uint(25))
		tt.AssertEqual(t, obj.Uint8, uint8(26))
		tt.AssertEqual(t, obj.Uint16, uint16(27))
		tt.AssertEqual(t, obj.Uint32, uint32(28))
		tt.AssertEqual(t, obj.Uint64, uint64(29))

		tt.AssertEqual(t, obj.F32, float32(30.0))
		tt.AssertEqual(t, obj.F64, 30.1)

		tt.AssertEqual(t, obj.Map, map[string]any{"key": "value"})
		tt.AssertEqual(t, obj.Slice, []string{"foo", "bar"})
		tt.AssertEqual(t, obj.Interface, any(1))
		tt.AssertEqual(t, obj.Pointer, ptr[int](24))
		tt.AssertEqual(t, obj.Func(), "myFunc")
		tt.AssertEqual(t, len(obj.Chan), 1)
		tt.AssertEqual(t, <-obj.Chan, "myChanMsg")

		tt.AssertEqual(t, obj.IntArray, [3]int{1, 2, 3})
		tt.AssertEqual(t, obj.StrArray, [3]string{"1", "2", "3"})

		tt.AssertEqual(t, obj.Struct, struct {
			String string
		}{
			String: "FakeString",
		})
	})
}

func ptr[T any](t T) *T {
	return &t
}
