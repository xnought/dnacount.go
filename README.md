# `dnacount.go`

A quick script in Go to count the base pair frequencies in genomes w/ parallelization across named FASTA labels/regions. With `dnacount.go`, I computed that the human genome has a GC bias of 41%. See results below (bottom of page).


**Build**

```bash
go build dnacount
```

**Execute**

```bash
./dnacount data/repeat_GCF_000863945.3_ViralProj15505_genomic.fna
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

Download reference human genome https://www.ncbi.nlm.nih.gov/datasets/genome/GCF_000001405.40/  

```bash
wget -O human_genome.zip https://api.ncbi.nlm.nih.gov/datasets/v2/genome/accession/GCF_000001405.40/download?include_annotation_type=GENOME_FASTA && unzip human_genome.zip -d data/human_genome && rm -fr human_genome.zip
```

execute 

```bash
./dnacount data/human_genome/ncbi_dataset/data/GCF_000001405.40/GCF_000001405.40_GRCh38.p14_genomic.fna
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
