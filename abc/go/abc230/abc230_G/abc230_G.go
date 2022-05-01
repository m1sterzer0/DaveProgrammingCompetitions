package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
	N := gi(); P := make([]int,N+1); for i:=1;i<=N;i++ { P[i] = gi() }
	pr := make([]bool,N+1); for i:=1;i<=N;i++ { pr[i] = true }
	mu := make([]int,N+1); for i:=1;i<=N;i++ { mu[i] = -1 }
	pr[1] = false; mu[1] = 0
	m := make([]int,0)
	d := make([][]int,N+1)
	// Modifying mobius function so that 
	// mu[n] = 1 iff n is product of an odd number of distinct primes
	// mu[n] = -1 iff n is a product of an even number of distinct primes
	// mu[n] = 0 otherwise
	for i:=2;i<=N;i++ {
		if pr[i] {
			mu[i] = 1
			for j,p:=2,2*i;p<=N;j,p=j+1,p+i {
				pr[p] = false
				if j % i == 0 { mu[p] = 0} else { mu[p] *= -1 }
			} 
		}
		if mu[i] != 0 {
			m = append(m,i)
			for j:=i; j<=N; j+=i { d[j] = append(d[j],i) }
		}
	}
	used := make([]bool,N+1)
	num := make([]int,N+1)
	cand := make([]int,0)
	ans := 0
	for _,a := range m {
		for i:=a;i<=N;i+=a {
			for _,j := range d[P[i]] {
				num[j]++
				if !used[j] { used[j] = true; cand = append(cand,j) }
			}
		}
		for _,b := range cand {
			ans += mu[a] * mu[b] * ((num[b] * (num[b]+1))/2)
			num[b] = 0
			used[b] = false
		}
		cand = cand[:0]
	}
	fmt.Println(ans)
}

