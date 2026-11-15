package main

import (
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "negative size", size: -10, wantLen: 0},
		{name: "zero size", size: 0, wantLen: 0},
		{name: "single element", size: 1, wantLen: 1},
		{name: "small slice", size: 10, wantLen: 10},
		{name: "large slice", size: 10_000, wantLen: 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d) length = %d, want %d", tt.size, len(data), tt.wantLen)
			}

			for i, value := range data {
				if value < 1 || value > maxRandomValue {
					t.Fatalf("generateRandomElements(%d)[%d] = %d, want value in range [1, %d]", tt.size, i, value, maxRandomValue)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single zero", data: []int{0}, want: 0},
		{name: "single positive", data: []int{42}, want: 42},
		{name: "single negative", data: []int{-42}, want: -42},
		{name: "positive values", data: []int{3, 1, 4, 1, 5, 9, 2, 6}, want: 9},
		{name: "negative values", data: []int{-3, -1, -4, -1, -5}, want: -1},
		{name: "mixed values", data: []int{-10, 0, 7, -2}, want: 7},
		{name: "maximum first", data: []int{9, 1, 2, 3}, want: 9},
		{name: "maximum last", data: []int{1, 2, 3, 9}, want: 9},
		{name: "duplicates", data: []int{5, 5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	allEqual := make([]int, 100)
	for i := range allEqual {
		allEqual[i] = 7
	}

	negative := make([]int, 16)
	for i := range negative {
		negative[i] = -i - 1
	}

	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "single negative", data: []int{-42}, want: -42},
		{name: "less than CHUNKS", data: []int{4, 8, 2, 7}, want: 8},
		{name: "exactly CHUNKS", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{
			name: "not divisible by CHUNKS",
			data: []int{5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 99},
			want: 99,
		},
		{
			name: "maximum in first chunk",
			data: []int{33, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			want: 33,
		},
		{
			name: "maximum in last chunk",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 33},
			want: 33,
		},
		{name: "negative values", data: negative, want: -1},
		{name: "all equal", data: allEqual, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for size := 0; size <= 200; size++ {
		data := generateRandomElements(size)
		want := maximum(data)

		if got := maxChunks(data); got != want {
			t.Fatalf("maxChunks(generateRandomElements(%d)) = %d, want %d", size, got, want)
		}
	}

	data := generateRandomElements(10_000)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Fatalf("maxChunks(generateRandomElements(10000)) = %d, want %d", got, want)
	}
}

func TestMaxChunksDoesNotModifyInput(t *testing.T) {
	data := generateRandomElements(1_000)
	original := slices.Clone(data)

	maxChunks(data)

	if !slices.Equal(data, original) {
		t.Fatal("maxChunks modified the input slice")
	}
}
