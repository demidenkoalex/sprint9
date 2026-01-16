package main

// Пишите тесты в этом файле

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(1000)
	empty := generateRandomElements(0)

	assert.Len(t, slice, 1000)
	for _, v := range slice {
		require.GreaterOrEqual(t, v, 0)
	}

	require.Empty(t, empty)
}

var maxCases = []struct {
	name string
	nums []int
	want int
}{
	{"nil", nil, 0},
	{"empty", []int{}, 0},
	{"single", []int{7}, 7},
	{"positives", []int{1, 9, 3, 4}, 9},
	{"all negative", []int{-10, -2, -30}, -2},
	{"mixed", []int{-1, 0, 5, 2}, 5},
}

func TestMaximum(t *testing.T) {
	for _, tt := range maxCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maximum(tt.nums))
		})
	}
}

func TestMaxChunks_SameCasesAsMaximum(t *testing.T) {
	const chunks = 4

	for _, tt := range maxCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxChunks(tt.nums, chunks))
		})
	}
}
