package main

import (
	"fmt"
	"os"
)

func normalize_frequencies(freqs map[byte]uint64) map[byte]float64 {
	normalized := map[byte]float64{} // char -> count/total_count

	// total_count is sum of all frequencies
	var total_count uint64 = 0
	for _, count := range freqs {
		total_count += count
	}

	// divide and save to normalized (also return this)
	for char, count := range freqs {
		normalized[char] = float64(count) / float64(total_count)
	}
	return normalized
}

func count_frequencies(data []byte) map[byte]uint64 {
	char_count := map[byte]uint64{} // char -> count
	for _, d := range data {
		_, key_exists := char_count[d]
		if key_exists {
			char_count[d]++
		} else {
			char_count[d] = 1
		}
	}
	return char_count
}

func main() {
	data, err := os.ReadFile("./data/GCF_000863945.3_ViralProj15505_genomic.fna")
	if err != nil {
		panic(err)
	}
	out := normalize_frequencies(count_frequencies(data))
	for k, v := range out {
		fmt.Printf("'%c'=%f\n", k, v)
	}
}
