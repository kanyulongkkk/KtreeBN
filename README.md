# KtreeBN
An Efficient Epistasis Mining Approach Using K-tree Optimizing Bayesian Network(The detailed instruction published in http://122.205.95.139/KtreeBN/)
## KtreeBN includes three parts,each part can be used separately.
### 1.Generate k-tree
#### use:
#### go install github.com/kanyulongkkk/KtreeBN/examples/generate-tree
#### $GOPATH/bin/generate-ktree
#### Parameter:
#### "n":used to adjuest the number of SNP.
#### "k":used to adjust the treewidth of ktree.
#### "p":used to adjust the sample times.
### 2. Get the k-clique
#### use:
#### $python install github.com/kanyulongkkk/KtreeBN/callr
#### $python callr.py
#### Parameter:
#### inputfile_data: the outputfile of the fist part.
#### inputfile_simulate_data: the simulate data download from  http://122.205.95.139/KtreeBN/.
#### outputfile: the output file name.
### 3.Get the Bayesian network
#### Parameter:
#### "alpha": used to adjust the threshold of fast-iamb algorithm.
#### "epi": used to adjust the dimension of SNP-SNP interation.
#### "inputfile": the outputfile of the second part.
#### "outputfile": bayesian network.
