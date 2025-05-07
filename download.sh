#!/bin/bash

# https://www.ncbi.nlm.nih.gov/datasets/genome/GCF_000001405.40/
filename="human_genome.zip"
dirname="data/human_genome"

echo "Downloading FASTA from https://www.ncbi.nlm.nih.gov/datasets/genome/GCF_000001405.40/"
wget -O $filename https://api.ncbi.nlm.nih.gov/datasets/v2/genome/accession/GCF_000001405.40/download?include_annotation_type=GENOME_FASTA&include_annotation_type=GENOME_GFF&include_annotation_type=RNA_FASTA&include_annotation_type=CDS_FASTA&include_annotation_type=PROT_FASTA&include_annotation_type=SEQUENCE_REPORT&hydrated=FULLY_HYDRATED

echo "Unzipping, may take a while..."
unzip $filename -d $dirname
rm -fr $filename