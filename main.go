package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	src := rand.NewSource(time.Now().Unix())
	for i := range data {
		data[i] = int(src.Int63())
	}
	return data
}

// // maximum returns the maximum number of elements.
func maximum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// // maxChunks returns the maximum number of elements in a chunks.
func maxChunks(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Compute minInt without importing extra packages.
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	chunkSize := len(nums) / CHUNKS
	maxima := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(nums)
		}
		sub := nums[start:end]

		go func(idx int, s []int) {
			defer wg.Done()
			if len(s) == 0 {
				maxima[idx] = minInt
				return
			}
			maxima[idx] = maximum(s)
		}(i, sub)
	}

	wg.Wait()
	return maximum(maxima)
}

func main() {

	numbers := generateRandomElements(SIZE)

	start := time.Now()
	maxValueSingle := maximum(numbers)
	elapsedSingle := time.Since(start).Microseconds()

	start = time.Now()
	maxValueChunks := maxChunks(numbers)
	elapsedChunks := time.Since(start).Microseconds()

	fmt.Printf("SIZE=%d\n", SIZE)
	fmt.Printf("maximum():   max=%d, time=%d µs\n", maxValueSingle, elapsedSingle)
	fmt.Printf("maxChunks(): max=%d, time=%d µs\n", maxValueChunks, elapsedChunks)

	if maxValueSingle != maxValueChunks {
		fmt.Printf("WARNING: results differ! maximum=%d, maxChunks=%d\n", maxValueSingle, maxValueChunks)
	}
}
