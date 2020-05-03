package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Welcome!")
	log.Println("This program expects input in the following format:")
	log.Println("")
	log.Println("===")
	log.Println("n k")
	log.Println("===")
	log.Println("")


	//var n, k int

	//fmt.Scanf("%d %d", &n, &k)
	for p :=0; p <10; p++{
	Tk, err := generator.RandomKtree(1000, 10)//传入n,k,新添加
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Printf("%d\n", len(Tk.Adj))
	fmt.Printf("%d\n", Tk.K)
	m := 0
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				m++
			}
		}
	}
	fmt.Printf("%d\n", m)

		fileName := strconv.Itoa(p)+".txt"
		f,err := os.Create(fileName)
		for u := 0; u < len(Tk.Adj); u++ {
			for i := 0; i < len(Tk.Adj[u]); i++ {
				v := Tk.Adj[u][i]
				if v > u {
					fmt.Printf("%d %d\n", u, v)
					l,err := fmt.Fprintln(f,u,v)
						if err != nil{
							fmt.Println(err)
							f.Close()
							return
						}
					fmt.Println(l,"bytes writen successfuly")
				}
			}
		}}

}
