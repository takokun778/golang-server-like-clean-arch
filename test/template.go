package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {

	tests := make([]Case, 0)

	test1 := Case{
		Name:     "",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     nil,
		Expected: nil,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result := ""

			if tt.IsErr {
				return
			}

			assert.Equal(t, tt.Expected, result)
		})
	}
}
