// Package codec implements coding and decoding algorithms.
package codec

import (
	"errors"
	"log"

)

type Code struct {
	Q []int
	S *dandelion.DandelionCode
}

// CodingAlgorithm receives a k-tree Tk and returns a code (Q, S).
// See Section 5 from Caminiti et al.
func CodingAlgorithm(Tk *ktree.Ktree) (*Code, error) {
	log.Printf("Coding Algorithm received input %v\n", Tk)

	// Step 1: Identify Q. Transform Tk into Rk.
	log.Println("Step 1...")
	Rk, err := ktree.RkFrom(Tk)
	if err != nil {
		return nil, err
	}
	log.Printf("Rk = %v\nQ = %v\n", Rk.Ktree, Rk.Q)

	// Step 2: Generate the characteristic tree T for Rk.
	log.Println("Step 2...")
	T := characteristic.TreeFrom(Rk)// TreeFrom returns a characteristic tree from a given Renyi k-Tree.
	log.Printf("T = %v\n", T)

	// Identify q = min(v not in Q).
	q := getMinVNotIn(Rk.Q)
	log.Printf("q = %v\n", q)

	// Make x = phi[q].
	phi := ktree.ComputePhi(len(Tk.Adj), Tk.K, Rk.Q)
	log.Printf("phi = %v\n", phi)
	x := phi[q]
	log.Printf("x = %v\n", x)

	// Step 3: Compute the Generalized Dandelion Code for T.
	log.Println("Step 3...")
	// The +1 is required because of the reindexing.
	S := dandelion.Code(T, x+1)
	log.Printf("S (with lm) = %v\n", S)

	// Remove the pair corresponding to phi[lm].
	lm, err := ktree.FindLm(Tk)
	if err != nil {
		return nil, err
	}
	// cor is the index of the pair corresponding to phi[lm].
	cor := phi[lm]
	if x < cor {
		cor--
	}
	log.Printf("lm = %v; phi[lm] = %v; cor = %v\n", lm, phi[lm], cor)
	S.P = append(S.P[:cor], S.P[cor+1:]...)
	S.L = append(S.L[:cor], S.L[cor+1:]...)

	log.Printf("Final S = %v\n", S)

	// Step 4: Return the code (Q, S).
	return &Code{Rk.Q, S}, nil
}

// DecodingAlgorithm receives a code (Q, S) and returns a k-tree Tk.
// See Section 6 in Caminiti et al.
func    DecodingAlgorithm(code *Code) (*ktree.Ktree, error) {
	log.Printf("Decoding Algorithm received input {%v, %v}\n", code.Q, code.S)
	Q, S := code.Q, code.S

	// Step 1: Compute phi, q, x, lm.
	log.Println("Step 1...")
	k := len(Q)
	n := len(S.P) + k + 2
	phi := ktree.ComputePhi(n, k, Q)
	q := getMinVNotIn(Q)
	x := phi[q]
	lm := findLm(n, phi, S.P, Q)
	log.Printf("lm = %v; q = %v; x = %v\n", phi, q, x)
	if lm == -1 {
		return nil, errors.New("Can't find lm. This should never happen.")
	}
	// cor is the index of the pair corresponding to phi[lm].
	cor := phi[lm]
	if x < cor {
		cor--
	}
	log.Printf("phi = %v; q = %v; x = %v\n", phi, q, x)
	log.Printf("cor = %v" ,cor )

	// Step 2: Insert the pair (0, e) in index cor and decode S to obtain T.
	log.Println("Step 2...")
	S.P = append(S.P[:cor], append([]int{0}, S.P[cor:]...)...)
	S.L = append(S.L[:cor], append([]int{characteristic.E}, S.L[cor:]...)...)
	log.Printf("S = %v\n", S)
	// The +1 is required because of the reindexing.
	T := dandelion.Decode(S, x+1)
	log.Printf("T = %v\n", T)

	// Step 3: Rebuild Rk by visiting T.
	log.Println("Step 3...")
	Rk := characteristic.RenyiKtreeFrom(n, k, Q, T)
	log.Printf("Rk = %v\n", Rk.Ktree)

	// Step 4: Apply phi^(-1) to Rk to obtain Tk.
	log.Println("Step 4...")
	Tk := ktree.TkFrom(Rk)
	log.Printf("Tk = %v\n", Tk)

	return Tk, nil
}

// getMinVNotIn receives a vector Q and returns q = min(v not in Q).
func getMinVNotIn(Q []int) int {
	q := Q[len(Q)-1] + 1
	if Q[0] != 0 {
		q = 0
	} else {
		for i := 0; i+1 < len(Q); i++ {
			if Q[i+1] > Q[i]+1 {
				q = Q[i] + 1
				break
			}
		}
	}
	return q
}

// findLm receives n, phi, p (parent vector) and Q. It returns lm, i.e.,
// the maximum v such that d(v) = k and v is not in Q.  lm := findLm(n, phi, S.P, Q)
func findLm(n int, phi, p, Q []int) int {
	internal := make([]bool, n)

	inv := ktree.GetInverse(phi)
	log.Printf("inv = %v; \n", inv)
	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			internal[inv[p[i]-1]] = true
		}
	}

	for i := 0; i < len(Q); i++ {
		internal[Q[i]] = true
	}

	for i := n - 1; i >= 0; i-- {
		if !internal[i] {
			return i
		}
	}

	return -1
}
