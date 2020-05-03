# This is parameters setting for BnBeeEpi.
library(ombFast)
# format transformation
# mydata <- read.table(file_name)
mydatasnp <- mydata[, 1:(length(mydata) - 1)]
Class <- mydata[, length(mydata)]
mydatasnp[mydatasnp == 0] <- "a"
mydatasnp[mydatasnp == 1] <- "b"
mydatasnp[mydatasnp == 2] <- "c"
Class[Class == 0] <- "f"
Class[Class == 1] <- "t"
for(i in 1:(ncol(mydata) - 1)) {
  mydatasnp[, i] <- factor(mydatasnp[, i])
}
Class <- factor(Class)
mydata <- cbind(mydatasnp, Class)


# The parameters need to be specified.
alpha = 0.001
epi = 2
# epi = 3


result <- fast.iamb(mydata, alpha = 0.001)

# check the result
arcs(result)
#write.table(arcs(temp),file ="D:\\generatedata",sep="")

