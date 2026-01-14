package main

// Пишите тесты в этом файле

import (
	"math/rand"
	"testing"
	"time"
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	randRange  = rand.New(randSource)
)

func randSlice(t *testing.T, n int) []int {
	t.Helper()
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	if n <= 0 {
		return []int{}
	}
	out := make([]int, n)
	for i := range out {
		out[i] = r.Intn(10_000) - 5_000 // include negative numbers too
	}
	return out
}

func assertEqualInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func assertLen(t *testing.T, s []int, want int) {
	t.Helper()
	if len(s) != want {
		t.Fatalf("len=%d, want %d", len(s), want)
	}
}

func TestGenerateRandomElements(t *testing.T) {
	rand.Seed(1)

	s := generateRandomElements(10)
	assertLen(t, s, 10)
	for _, v := range s {
		if v < 0 {
			t.Fatalf("value must be >= 0, got %d", v)
		}
		if v >= SIZE {
			t.Fatalf("value must be < %d, got %d", SIZE, v)
		}
	}

	zero := generateRandomElements(0)
	assertLen(t, zero, 0)

	neg := generateRandomElements(-5)
	assertLen(t, neg, 0)
}

func TestMaximum(t *testing.T) {
	assertEqualInt(t, maximum(nil), 0)
	assertEqualInt(t, maximum([]int{}), 0)

	assertEqualInt(t, maximum([]int{7}), 7)
	assertEqualInt(t, maximum([]int{1, 9, 3, 4}), 9)
	assertEqualInt(t, maximum([]int{-10, -2, -30}), -2)
	assertEqualInt(t, maximum([]int{-1, 0, 5, 2}), 5)
}

func TestMaxChunks(t *testing.T) {
	assertEqualInt(t, maxChunks(nil), 0)
	assertEqualInt(t, maxChunks([]int{}), 0)

	// Small slices (< CHUNKS) should still work.
	assertEqualInt(t, maxChunks([]int{1, 4, 2}), 4)
	assertEqualInt(t, maxChunks([]int{-3, -1, -7}), -1)

	// Not divisible by CHUNKS (remainder goes to last chunk).
	data := []int{1, 8, 3, 4, 5, 6, 7, 2, 9} // 9 elements
	assertEqualInt(t, maxChunks(data), 9)

	// Randomized property: result must match the single-thread maximum.
	for i := 0; i < 20; i++ {
		r := randSlice(t, 10_000)
		assertEqualInt(t, maxChunks(r), maximum(r))
	}
}
