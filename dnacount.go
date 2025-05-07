package main

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar"
)

func normFreqs(freqs map[byte]uint64) map[byte]float64 {
	normalized := map[byte]float64{} // char -> count/total_count

	// totalCount is sum of all frequencies
	var totalCount uint64 = 0
	for _, count := range freqs {
		totalCount += count
	}

	// divide and save to normalized (also return this)
	for char, count := range freqs {
		normalized[char] = float64(count) / float64(totalCount)
	}
	return normalized
}

func byteUpper(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a - ('a' - 'A')
	}
	return a
}

func countFreqs(data []byte, all_upper bool) map[byte]uint64 {
	charCount := map[byte]uint64{} // char -> count
	for _, d := range data {
		key := d
		if all_upper {
			key = byteUpper(d)
		}

		_, exists := charCount[key]
		if exists {
			charCount[key]++
		} else {
			charCount[key] = 1
		}
	}
	return charCount
}

type region struct {
	start int
	end   int
}

func skipFastaHeader(iptr *int, data []byte) {
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

func findDNARegions(data []byte) []region {
	idxs := []region{}

	i := 0
	for i < len(data) {
		skipFastaHeader(&i, data)
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

func countFreqsSaveToChannel(data []byte, ch chan map[byte]uint64) {
	ch <- countFreqs(data, true)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("Must provide one argument as the filename")
	}

	filename := args[0]
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded '%s' into RAM\n", filename)

	// count up frequencies for each dna region (store result in channel)
	totalLen := 0
	regions := findDNARegions(data)
	ch := make(chan map[byte]uint64, len(regions))
	for _, r := range regions {
		region := data[r.start:r.end]
		totalLen += r.end - r.start
		go countFreqsSaveToChannel(region, ch) // execute countFreqs in parallel
	}

	// gather computed frequencies into one map
	gatherFreqs := map[byte]uint64{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	bar := progressbar.New(len(regions))
	for range len(regions) {
		freqs := <-ch
		// fmt.Printf("Computed %d/%d\n", i+1, len(regions))
		bar.Add(1)
		for k := range gatherFreqs {
			value, found := freqs[k]
			if found {
				gatherFreqs[k] += value
			}
		}
	}

	result := normFreqs(gatherFreqs)
	fmt.Printf("\nTotal length %d\n", totalLen)
	fmt.Printf("%c %.2f%%\n", 'G', result['G']*100)
	fmt.Printf("%c %.2f%%\n", 'C', result['C']*100)
	fmt.Printf("%c %.2f%%\n", 'T', result['T']*100)
	fmt.Printf("%c %.2f%%\n", 'A', result['A']*100)
	fmt.Printf("GC bias of %.2f%%\n", (result['G']+result['C'])*100)
}
