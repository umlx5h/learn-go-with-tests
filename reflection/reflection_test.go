package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Struct with with nested string fields",
			Person{"Chris", Profile{
				32,
				"Tokyo",
			}},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Pointer Struct with with nested string fields",
			&Person{"Chris", Profile{
				32,
				"Tokyo",
			}},
			[]string{"Chris", "Tokyo"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Kyoto"},
			},
			[]string{"London", "Kyoto"},
		},
		{
			"Array",
			[2]Profile{
				{33, "London"},
				{34, "Kyoto"},
			},
			[]string{"London", "Kyoto"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	walk(x, func(input string) {
		got = append(got, input)
	})

	assert.Equal(t, expected, got[0], "wrong number of function calls")

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assert.Contains(t, got, "Bar")
		assert.Contains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Tokyo"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Tokyo"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assert.Equal(t, want, got)
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Tokyo"}
		}

		var got []string
		want := []string{"Berlin", "Tokyo"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		assert.Equal(t, want, got)
	})
}
