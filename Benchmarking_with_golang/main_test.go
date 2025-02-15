package main

import (
	"os"
	"runtime/pprof"
	"strconv"
	"testing"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap
			}
		}
	}
	return arr
}

// MergeSort sorts an array using the merge sort algorithm.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge combines two sorted arrays into one sorted array.
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func generateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

/*
goos: windows
goarch: amd64    //the architecture of the machine
pkg: example.com/benchmarking
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
//the name of benchmark function /no of iteration       /average time it took
BenchmarkBubbleSort-8           100000000               13.31 ns/op

	memory per operation   /number of allocation per operation

/op            0 B/op                 0 allocs/op
PASS                                    /time it took to run the benchmark
ok      example.com/benchmarking        1.588s
*/
// func BenchmarkBubbleSort(b *testing.B) {
// 	arr := []int{5, 4, 3, 2, 1}
// 	// b.N is the number of iterations the function will run
// 	b.ReportAllocs() //this wi;l track the memory allocation
// 	for i := 0; i < b.N; i++ {
// 		bubbleSort(arr)
// 	}
// }
// result format
//BenchmarkName-Parallelism   Iterations   TimePerOperation  MemoryUsage  Allocations

func BenchmarkSortAlgoithms(b *testing.B) {

	cpuProfile, err := os.Create("cpu_profile.prof")
	if err != nil {
		b.Fatal(err)
	}

	defer cpuProfile.Close()
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	//start memory profiling
	memProfile, err := os.Create("mem_profile.prof")
	if err != nil {
		b.Fatal(err)
	}
	defer memProfile.Close()
	defer pprof.WriteHeapProfile(memProfile)

	sizes := []int{100, 1000, 10000}
	for _, size := range sizes {
		b.Run("BubbleSort_"+strconv.Itoa(size), func(b *testing.B) {
			arr := generateArray(size)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				bubbleSort(arr)
			}
		})

		b.Run("MergeSort_"+strconv.Itoa(size), func(b *testing.B) {
			arr := generateArray(size)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				mergeSort(arr)

			}

		})
	}
}
