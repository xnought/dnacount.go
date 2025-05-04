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

	// count up frequencies for each dna region
	base_pair_freqs := map[byte]uint64{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	total_len := 0
	for _, region_edges := range dna_regions(data) {
		region := data[region_edges.start:region_edges.end]
		total_len += region_edges.end - region_edges.start
		freqs := count_frequencies(region)
		for k := range base_pair_freqs {
			value, found := freqs[k]
			if found {
				base_pair_freqs[k] += value
			}
		}
	}

	result := normalize_frequencies(base_pair_freqs)
	fmt.Println("Human papillomavirus genome")
	fmt.Printf("Total length %d\n", total_len)
	fmt.Printf("%c %.2f%%\n", 'G', result['G']*100)
	fmt.Printf("%c %.2f%%\n", 'C', result['C']*100)
	fmt.Printf("%c %.2f%%\n", 'T', result['T']*100)
	fmt.Printf("%c %.2f%%\n", 'A', result['A']*100)
	fmt.Printf("GC bias of %.2f%%\n", (result['G']+result['C'])*100)
}
