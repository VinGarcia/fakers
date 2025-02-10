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

	/*
		t.Run("should return error if we try to save something that is not a map into a nested struct", func(t *testing.T) {
			var user struct {
				ID       int    `map:"id"`
				Username string `map:"username"`
				Address  struct {
					Street  string `map:"street"`
					City    string `map:"city"`
					Country string `map:"country"`
				} `map:"address"`
			}
			err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
				"id":       42,
				"username": "fakeUsername",
				"address":  "notAMap",
			}))

			tt.AssertErrContains(t, err, "string", "Address", "Street", "City", "Country")
		})

		t.Run("using the required validation", func(t *testing.T) {
			t.Run("should ignore missing fields if they are not required", func(t *testing.T) {
				var user struct {
					ID       int    `map:"id"`
					Username string `map:"username"`
					Address  struct {
						Street  string `map:"street"`
						City    string `map:"city"`
						Country string `map:"country"`
					} `map:"address"`

					OptionalStruct struct {
						ID int `map:"id"`
					} `map:"optional_struct"`
				}
				// These three should still be present after the parsing:
				user.OptionalStruct.ID = 42
				user.Username = "presetUsername"
				user.Address.Street = "presetStreet"

				// These two should be overwritten by the parser:
				user.ID = 43
				user.Address.Country = "presetCountry"

				err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
					"id": 44,
					"address": map[string]any{
						"city":    "fakeCity",
						"country": "fakeCountry",
					},
				}))
				tt.AssertNoErr(t, err)

				tt.AssertEqual(t, user.ID, 44)
				tt.AssertEqual(t, user.Username, "presetUsername")
				tt.AssertEqual(t, user.Address.Street, "presetStreet")
				tt.AssertEqual(t, user.Address.City, "fakeCity")
				tt.AssertEqual(t, user.Address.Country, "fakeCountry")
				tt.AssertEqual(t, user.OptionalStruct.ID, 42)
			})

			t.Run("should return error for missing fields if they are required", func(t *testing.T) {
				tests := []struct {
					desc               string
					input              map[string]any
					expectErrToContain []string
				}{
					{
						desc: "required field missing on root map",
						input: map[string]any{
							"id": 42,
							"address": map[string]any{
								"street":  "fakeStreet",
								"city":    "fakeCity",
								"country": "fakeCountry",
							},
						},
						expectErrToContain: []string{"missing", "required", "config", "username"},
					},
					{
						desc: "required field missing on nested map",
						input: map[string]any{
							"id":       42,
							"username": "fakeUsername",
							"address": map[string]any{
								"city":    "fakeCity",
								"country": "fakeCountry",
							},
						},
						expectErrToContain: []string{"missing", "required", "config", "street"},
					},
					{
						desc: "required field missing is a map",
						input: map[string]any{
							"id":       42,
							"username": "fakeUsername",
						},
						expectErrToContain: []string{"missing", "required", "config", "address"},
					},
				}

				for _, test := range tests {
					t.Run(test.desc, func(t *testing.T) {
						var user struct {
							ID       int    `map:"id"`
							Username string `map:"username" validate:"required"`
							Address  struct {
								Street  string `map:"street" validate:"required"`
								City    string `map:"city"`
								Country string `map:"country"`
							} `map:"address" validate:"required"`
						}
						err := structscanner.Decode(&user, newCustomValuesMapDecoder(test.input))

						tt.AssertErrContains(t, err, test.expectErrToContain...)
					})
				}
			})

			t.Run("should return error if the validation is misspelled", func(t *testing.T) {
				var user struct {
					ID       int    `map:"id"`
					Username string `map:"username" validate:"not_required"`
				}
				err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
					"id":       42,
					"username": "fakeUsername",
				}))

				tt.AssertErrContains(t, err, "validation", "not_required")
			})

			t.Run("should not return error if the required field is empty but has a default value", func(t *testing.T) {
				var user struct {
					ID       int    `map:"id"`
					Username string `map:"username" validate:"required" default:"defaultUsername"`
				}
				err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
					"id": 42,
				}))
				tt.AssertNoErr(t, err)

				tt.AssertEqual(t, user.ID, 42)
				tt.AssertEqual(t, user.Username, "defaultUsername")
			})
		})

		t.Run("using the default tag", func(t *testing.T) {
			t.Run("should work for multiple types of fields", func(t *testing.T) {
				var user struct {
					ID       int    `map:"id"`
					Username string `map:"username" default:"defaultUsername"`
					Address  struct {
						Street  string `map:"street" default:"defaultStreet"`
						City    string `map:"city"`
						Country string `map:"country"`
					} `map:"address"`

					OptionalStruct struct {
						ID int `map:"id" default:"41"`
					} `map:"optional_struct"`
				}

				// These all these should be overwritten by the parser:
				user.ID = 43
				user.Address.Country = "presetCountry"
				user.OptionalStruct.ID = 42
				user.Username = "presetUsername"
				user.Address.Street = "presetStreet"

				err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
					"id": 44,
					"address": map[string]any{
						"city":    "fakeCity",
						"country": "fakeCountry",
					},
				}))
				tt.AssertNoErr(t, err)

				tt.AssertEqual(t, user.ID, 44)
				tt.AssertEqual(t, user.Username, "defaultUsername")
				tt.AssertEqual(t, user.Address.Street, "defaultStreet")
				tt.AssertEqual(t, user.Address.City, "fakeCity")
				tt.AssertEqual(t, user.Address.Country, "fakeCountry")
				tt.AssertEqual(t, user.OptionalStruct.ID, 41)
			})
		})

		t.Run("parsing slices", func(t *testing.T) {
			tests := []struct {
				desc          string
				inputSlice    any
				expectedSlice any
			}{
				{
					desc: "should work for string slices",
					inputSlice: []string{
						"fakeUser1",
						"fakeUser2",
					},
					expectedSlice: []string{
						"fakeUser1",
						"fakeUser2",
					},
				},
			}

			for _, test := range tests {
				t.Run(test.desc, func(t *testing.T) {
					var user struct {
						ID    int      `map:"id"`
						Slice []string `map:"slice"`
					}

					err := structscanner.Decode(&user, newCustomValuesMapDecoder(map[string]any{
						"id":    42,
						"slice": test.inputSlice,
					}))
					tt.AssertNoErr(t, err)

					tt.AssertEqual(t, user.ID, 42)
					tt.AssertEqual(t, user.Slice, test.expectedSlice)
				})
			}
		})
		// */
}

func ptr[T any](t T) *T {
	return &t
}
