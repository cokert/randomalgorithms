package main

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
)

func random(len int) []int {
	r := make([]int, len)
	for i := 0; i < len; i++ {
		r[i] = rand.Intn(len)
	}
	return r
}

// minConcurrentSize is the limit for running current MergeSort operations.  Once the array is below this size, we
// switch to running all further MergeSort operations on the current gofunc.
const minConcurrentSize = 10_000

//const minConcurrentSize = 0 // this is stupid.  run this to start an absurd number of gofuncs and watch the program consume all ram.

func main() {
	const arrSize = 80_000_000
	format := new(log.TextFormatter)
	format.TimestampFormat = "2006-01-02T15:04:05.999999999Z"
	format.FullTimestamp = true
	log.SetFormatter(format)
	log.Infof("creating array of length %v", arrSize)
	arr := random(arrSize)
	log.Infof("start mergesort")
	m := MergeSort(arr)
	log.Infof("end mergesort")
	log.Infof("is sorted: %v", check(m))

	log.Infof("creating array of length %v", arrSize)
	arr = random(arrSize)
	ch := make(chan []int)
	log.Infof("start chansort")
	go MergeSortChan(arr, ch)
	n := <-ch
	log.Infof("end chansort")
	log.Infof("is sorted: %v", check(n))
}

func InsertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

func check(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		a := MergeSort(arr[:mid])
		b := MergeSort(arr[mid:])
		n := make([]int, len(arr))
		i := 0
		j := 0
		for k := 0; k < len(n); k++ {
			if j >= len(b) {
				n[k] = a[i]
				i++
			} else if i >= len(a) {
				n[k] = b[j]
				j++
			} else if a[i] < b[j] {
				n[k] = a[i]
				i++
			} else {
				n[k] = b[j]
				j++
			}
		}
		return n
	}
	return arr
}

func MergeSortChan(arr []int, ch chan []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		var a []int
		var b []int
		if len(arr) > minConcurrentSize {
			log.Debugf("channel sorting: %v", len(arr))
			chSortA := make(chan []int)
			chSortB := make(chan []int)
			go MergeSortChan(arr[:mid], chSortA)
			go MergeSortChan(arr[mid:], chSortB)
			a = <-chSortA
			b = <-chSortB
		} else {
			a = MergeSort(arr[:mid])
			b = MergeSort(arr[mid:])
		}
		n := make([]int, len(arr))
		i := 0
		j := 0
		for k := 0; k < len(n); k++ {
			if j >= len(b) {
				n[k] = a[i]
				i++
			} else if i >= len(a) {
				n[k] = b[j]
				j++
			} else if a[i] < b[j] {
				n[k] = a[i]
				i++
			} else {
				n[k] = b[j]
				j++
			}
		}
		ch <- n
	}
	ch <- arr
}
