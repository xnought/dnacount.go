# test-go

A quick script in Go to count the base pair frequencies in the Human papillomavirus genome.

https://www.ncbi.nlm.nih.gov/datasets/taxonomy/333760/

```bash
go run main.go
```

which counts the frequencies of the base pairs to result in

```txt
Human papillomavirus genome
Total length 8006
G 19.12%
C 17.42%
T 30.60%
A 32.86%
GC bias of 36.54%
```