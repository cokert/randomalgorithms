package main

import "testing"

func run(loops int, size int) {
	size = size * 1000
	for i := 0; i < loops; i++ {
		arr := random(size)
		MergeSort(arr)
	}
}

func runChan(loops int, size int) {
	size = size * 1000
	for i := 0; i < loops; i++ {
		arr := random(size)
		ch := make(chan []int)
		go MergeSortChan(arr, ch)
		_ = <-ch
	}
}

func BenchmarkMergeSort1k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 1)
}
func BenchmarkMergeSort10k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 10)
}

func BenchmarkMergeSort100k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 100)
}

func BenchmarkMergeSort500k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 500)
}

func BenchmarkMergeSort1000k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 1000)
}

func BenchmarkMergeSort80000k(b *testing.B) {
	b.ReportAllocs()
	run(b.N, 80000)
}

func BenchmarkMergeSortChan1k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 1)
}

func BenchmarkMergeSortChan10k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 10)
}

func BenchmarkMergeSortChan100k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 100)
}

func BenchmarkMergeSortChan500k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 500)
}

func BenchmarkMergeSortChan1000k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 1000)
}

func BenchmarkMergeSortChan80000k(b *testing.B) {
	b.ReportAllocs()
	runChan(b.N, 80000)
}

func BenchmarkInsertionSort1k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := random(1000)
		InsertionSort(arr)
	}
}

func BenchmarkInsertionSort10k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := random(10000)
		InsertionSort(arr)
	}
}
