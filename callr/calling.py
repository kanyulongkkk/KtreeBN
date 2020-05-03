import networkx as nx

from networkx.algorithms.approximation import  treewidth_min_degree
G = nx.Graph()
f = open("(100,3).txt")
f_data = f.readlines()
ddaa=[]
for line in f_data:
    line = line.rstrip("\n").split(" ")
    new_line =[]
    for n in line:
            new_line.append(int(n))
    ddaa.append(tuple(new_line))

print("ddaa",ddaa)
G.add_edges_from(ddaa)
a,b = treewidth_min_degree(G)
print(a)
print(b.nodes)
import pandas as pd
frozenset_l = [list(i) for i in list(b.nodes)]
print(frozenset_l)
frozenset_pd = pd.DataFrame(frozenset_l)
print(frozenset_pd)
frozenset_pd.to_csv("columns.csv", index=None, header=None)
frozenset_pd = pd.read_csv("columns.csv", header=None)
data = pd.read_csv("H0.1MAF0.2-2locus_EDM-1_001.txt", sep='\t', header=None)
for index, row in frozenset_pd.iterrows():
     frozenset_l = list(row)
     frozenset_l.append(100)
     print(frozenset_l)
     output_data = data.loc[:, frozenset_l]
     #output_data.to_csv( str(index) + ".txt", sep='\t', index=None, header=None)
     output_data.to_csv("_".join([str(i) for i in frozenset_l])+".txt", sep='\t', index=None, header=None)
import  rpy2
print(rpy2.__version__)
from rpy2.robjects.packages import importr
BnBeeEpi = importr('BnBeeEpi')
# r_script ='''
#
# mydata <- read.table("D:\\\\数据\\\\99_65_34_91_100.txt", header = T)
#
# mydatasnp <- mydata[, 1:(length(mydata) - 1)]
# Class <- mydata[, length(mydata)]
# mydatasnp[mydatasnp == 0] <- "a"
# mydatasnp[mydatasnp == 1] <- "b"
# mydatasnp[mydatasnp == 2] <- "c"
# Class[Class == 0] <- "f"
# Class[Class == 1] <- "t"
# for(i in 1:(length(mydata) - 1)) {
#   mydatasnp[, i] <- factor(mydatasnp[, i])
# }
# Class <- factor(Class)
# mydata <- cbind(mydatasnp, Class)
# epi =2
# temp <- fast.iamb(mydata, alpha = 0.001)
# print(arcs(temp))
# print('---------')
# '''
# robjects.r(r_script)
