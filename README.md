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

