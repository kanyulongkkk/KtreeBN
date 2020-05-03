ombFast User Guide

Version: 1.0

1) Installation
	
remove ombFast R package to your R work directory and then type

	install.packages("ombFast_1.0.tar.gz", repos = NULL, type = "source")
	
in your R console.

2) Input Format

Users should upload right format file which contains the case-control data as the input for the programme. 
The first line of the input file contains the SNP IDs and the name for the reponse variable(last column), 
and the following lines are the genotype data which should be coded by 0, 1, 2 with each line corresponding to one individual.
The last column should contain the disease status of each individual coded by 0 and 1.
The following is a sample data file for 5 individuals(3 cases and 2 controls), each genotyped for 10 SNPs.

X0,X1,X2,X3,X4,X5,X6,X7,X8,X9,Class
1,0,2,1,0,0,1,1,0,1,0
2,2,1,2,1,0,0,0,0,0,0
1,0,0,1,1,1,0,2,1,1,1
0,1,1,0,1,0,0,0,0,1,1
1,2,0,0,1,0,0,1,0,0,1

3) Parameters

There are 2 parameters need to be specified:

	"alpha"   : used to adjust the threshold of fast-iamb algorithm.(recommend 0.05--0.001)
	"epi"       : used to adjust the dimension of SNP-SNP interaction(recommend 2 or 3)
	
We have provided a parameters Rscript for you, please run it .

4£©Output

You can get a Bayesian network of SNP through the Code in the "Parameter for ombFast.R" 

