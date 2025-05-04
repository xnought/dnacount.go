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

type region struct {
	start int
	end   int
}

func skip_fasta_header(iptr *int, data []byte) {
	if data[*iptr] == '>' {
		// keep going with i until I hit a new line
		for {
			(*iptr)++
			if (*iptr) >= len(data) || data[*iptr] == '\n' {
				break
			}
		}
	}
}

func dna_regions(data []byte) []region {
	idxs := []region{}

	i := 0
	for i < len(data) {
		skip_fasta_header(&i, data)
		// read the dna base pairs, stop once I hit > again
		r := region{start: i - 1, end: -1}
		for {
			if i >= len(data) || data[i] == '>' {
				break
			}
			i++
		}
		r.end = i - 1

		idxs = append(idxs, r)
	}

	return idxs
}

func main() {
	data, err := os.ReadFile("./data/GCF_000863945.3_ViralProj15505_genomic.fna")
	if err != nil {
		panic(err)
	}

	regions := dna_regions(data)
	for _, r := range regions {
		freqs := count_frequencies(data[r.start:r.end])
		fmt.Println(freqs)
	}
}
