package calc

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Error("Result: %d, Expected: %d", result, expected)
	}
	result2 := Sum(7, 8, 10)
	if result2 != expected {
		t.Error("Result: %d, Expected: %d", result2, expected)
	}
}

func TestAverage(t *testing.T) {
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Error("Result: %f Expected: %f", result, expected)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(7, 8, 10)
	}
}

func TestLongRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	time.Sleep(5 * time.Second)
}

func TestSumInParallel(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Error("Result: %d, Expected:%d", result, expected)
	}
}

func TestAverageInParallel(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}
