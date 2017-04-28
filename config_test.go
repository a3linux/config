package config

import (
	"testing"

	"os"
)

func BenchmarkGet(b *testing.B) {
	c := New("test.json")

	var v string
	for i := 0; i < b.N; i++ {
		c.Get("test", &v)
	}
}

func TestGet(t *testing.T) {
	c := New("test.json")

	t.Run("Get", func(t *testing.T) {
		t.Run("should successfully retrieve string", func(t *testing.T) {
			var s string

			err := c.Get("string", &s)
			if err != nil {
				t.Error(err)
			}

			if s != "asdf" {
				t.Errorf("Expected 'asdf', got '%v'", s)
			}
		})

		t.Run("should successfully cast bool to string", func(t *testing.T) {
			t.Skip()
		})

		t.Run("should successfully cast float to string", func(t *testing.T) {
			t.Skip()
		})

		t.Run("should successfully retrieve bool", func(t *testing.T) {
			var b bool

			err := c.Get("bool", &b)
			if err != nil {
				t.Error(err)
			}

			if b != true {
				t.Errorf("Expected 'true', got '%v'", b)
			}
		})

		t.Run("should successfully cast string to bool", func(t *testing.T) {
			var b bool

			err := c.Get("string", &b)
			if err != nil {
				t.Error(err)
			}

			if b != true {
				t.Errorf("Expected 'true', got '%v'", b)
			}
		})

		t.Run("should successfully cast float to bool", func(t *testing.T) {
			var b bool

			err := c.Get("float64", &b)
			if err != nil {
				t.Error(err)
			}

			if b != true {
				t.Errorf("Expected 'true', got '%v'", b)
			}
		})

		t.Run("should successfully retrieve float64", func(t *testing.T) {
			var f float64

			err := c.Get("float64", &f)
			if err != nil {
				t.Error(err)
			}

			if f != 64.4 {
				t.Errorf("Expected '64.4', got '%v'", f)
			}
		})

		t.Run("should successfully cast string to float", func(t *testing.T) {
			t.Skip()
		})

		t.Run("should successfully cast bool to float", func(t *testing.T) {
			t.Skip()
		})

		t.Run("should successfully override with Env", func(t *testing.T) {
			var s string

			err := os.Setenv("string", "abcd")
			if err != nil {
				t.Error(err)
			}

			err = c.Get("string", &s)
			if err != nil {
				t.Error(err)
			}

			if s != "abcd" {
				t.Errorf("Expected 'abcd', got '%v'", s)
			}
		})

		t.Run("should successfully cast float to int", func(t *testing.T) {
			var i int

			err := c.Get("float64", &i)
			if err != nil {
				t.Error(err)
			}

			if i != 64 {
				t.Errorf("Expected '64', got '%v'", i)
			}
		})
	})
}
