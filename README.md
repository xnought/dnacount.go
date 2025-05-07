# test-go

A quick script in Go to count the base pair frequencies in the Human papillomavirus genome. ([click here to see the file contents](data/GCF_000863945.3_ViralProj15505_genomic.fna))

Downloaded from https://www.ncbi.nlm.nih.gov/datasets/taxonomy/333760/ all credit to them

```bash
go run main.go FILENAME_HERE
```

For example

```bash
go run main.go data/repeat_GCF_000863945.3_ViralProj15505_genomic.fna
```

returns

```txt
Loaded 'data/GCF_000863945.3_ViralProj15505_genomic.fna' into RAM
 100% |████████████████████████████████████████| [0s:0s]            
Total length 8006
G 19.12%
C 17.42%
T 30.60%
A 32.86%
GC bias of 36.54%
```

**Entire Human Genome**

```bash
bash download.sh  # downloads files to data/human_genome/
```

```bash
go run main.go data/human_genome/ncbi_dataset/data/GCF_000001405.40/GCF_000001405.40_GRCh38.p14_genomic.fna
```

returns 

```txt
Loaded 'data/human_genome/ncbi_dataset/data/GCF_000001405.40/GCF_000001405.40_GRCh38.p14_genomic.fna' into RAM
 100% |████████████████████████████████████████| [26s:0s]            
Total length 3339662079
G 20.57%
C 20.48%
T 29.52%
A 29.43%
GC bias of 41.05%
```