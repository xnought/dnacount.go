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

func uppercase_byte(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a - ('a' - 'A')
	}
	return a
}

func count_frequencies(dataPtr *[]byte, all_upper bool) map[byte]uint64 {
	data := *dataPtr
	char_count := map[byte]uint64{} // char -> count
	for _, d := range data {
		key := d
		if all_upper {
			key = uppercase_byte(d)
		}

		_, key_exists := char_count[key]
		if key_exists {
			char_count[key]++
		} else {
			char_count[key] = 1
		}
	}
	return char_count
}

type region struct {
	start int
	end   int
}

func skip_fasta_header(iptr *int, dataPtr *[]byte) {
	data := *dataPtr
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

func dna_regions(dataPtr *[]byte) []region {
	idxs := []region{}
	data := *dataPtr

	i := 0
	for i < len(data) {
		skip_fasta_header(&i, dataPtr)
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

// func main() {
// 	data, err := os.ReadFile("./data/repeat_GCF_000863945.3_ViralProj15505_genomic.fna")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Loaded data")

// 	// count up frequencies for each dna region
// 	base_pair_freqs := map[byte]uint64{'A': 0, 'C': 0, 'T': 0, 'G': 0}
// 	total_len := 0
// 	regions := dna_regions(&data)
// 	for i, region_edges := range regions {
// 		region := data[region_edges.start:region_edges.end]
// 		total_len += region_edges.end - region_edges.start
// 		freqs := count_frequencies(&region, true)
// 		fmt.Printf("Computed %d/%d\n", i+1, len(regions))
// 		for k := range base_pair_freqs {
// 			value, found := freqs[k]
// 			if found {
// 				base_pair_freqs[k] += value
// 			}
// 		}
// 	}

// 	result := normalize_frequencies(base_pair_freqs)
// 	fmt.Printf("Total length %d\n", total_len)
// 	fmt.Printf("%c %.2f%%\n", 'G', result['G']*100)
// 	fmt.Printf("%c %.2f%%\n", 'C', result['C']*100)
// 	fmt.Printf("%c %.2f%%\n", 'T', result['T']*100)
// 	fmt.Printf("%c %.2f%%\n", 'A', result['A']*100)
// 	fmt.Printf("GC bias of %.2f%%\n", (result['G']+result['C'])*100)
// }

func parallel_count_frequencies(dataPtr *[]byte, ch chan map[byte]uint64) {
	ch <- count_frequencies(dataPtr, true)
}

func main() {
	data, err := os.ReadFile("./data/repeat_GCF_000863945.3_ViralProj15505_genomic.fna")
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded data")

	// count up frequencies for each dna region (store result in channel)
	total_len := 0
	regions := dna_regions(&data)
	ch := make(chan map[byte]uint64, len(regions))
	for _, region_edges := range regions {
		region := data[region_edges.start:region_edges.end]
		total_len += region_edges.end - region_edges.start
		go parallel_count_frequencies(&region, ch)
	}

	// gather computed frequencies into one map
	gather_base_pair_freqs := map[byte]uint64{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	for i := range len(regions) {
		freqs := <-ch
		fmt.Printf("Computed %d/%d\n", i+1, len(regions))
		for k := range gather_base_pair_freqs {
			value, found := freqs[k]
			if found {
				gather_base_pair_freqs[k] += value
			}
		}
	}

	result := normalize_frequencies(gather_base_pair_freqs)
	fmt.Printf("Total length %d\n", total_len)
	fmt.Printf("%c %.2f%%\n", 'G', result['G']*100)
	fmt.Printf("%c %.2f%%\n", 'C', result['C']*100)
	fmt.Printf("%c %.2f%%\n", 'T', result['T']*100)
	fmt.Printf("%c %.2f%%\n", 'A', result['A']*100)
	fmt.Printf("GC bias of %.2f%%\n", (result['G']+result['C'])*100)
}
