package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	t.Parallel()

	t.Run("valid-inputs", func(t *testing.T) {
		t.Parallel()

		min := 30
		max := 100
		n := Int(min, max)
		assert.GreaterOrEqual(t, n, min)
		assert.LessOrEqual(t, n, max)
	})

	t.Run("invalid-inputs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name        string
			min         int
			max         int
			expectedMin int
			expectedMax int
		}{
			{
				name:        "min<0",
				min:         -30,
				max:         90,
				expectedMin: 30,
				expectedMax: 90,
			},
			{
				name:        "max<0",
				min:         90,
				max:         -30,
				expectedMin: 30,
				expectedMax: 90,
			},
			{
				name:        "min<0-and-max<0",
				min:         90,
				max:         -30,
				expectedMin: 30,
				expectedMax: 90,
			},
			{
				name:        "max<min",
				min:         50,
				max:         20,
				expectedMin: 20,
				expectedMax: 50,
			},
			{
				name:        "min==max",
				min:         30,
				max:         30,
				expectedMin: 30,
				expectedMax: 30,
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				n := Int(tc.min, tc.max)
				assert.GreaterOrEqual(t, n, tc.expectedMin)
				assert.LessOrEqual(t, n, tc.expectedMax)
			})
		}
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("valid-inputs", func(t *testing.T) {
		t.Parallel()

		n := 10
		a := String(n)
		b := String(n)

		assert.Len(t, a, n)
		assert.NotEqual(t, a, b)
	})

	t.Run("invalid-inputs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name string
			n    int
		}{
			{
				name: "negative-val:-10",
				n:    -10,
			},
			{
				name: "zero-val:0",
				n:    0,
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				a := String(tc.n)
				assert.Empty(t, a)
			})
		}
	})
}
