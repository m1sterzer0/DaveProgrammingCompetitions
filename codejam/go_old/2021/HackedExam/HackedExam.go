package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	combarr := make([][]big.Int,121)
	for i:=0;i<=120;i++ { combarr[i] = make([]big.Int,121) }
	combarr[0][0].SetInt64(1)
	for i:=1;i<=120;i++ {
		for j:=0;j<=i;j++ {
			if j == 0 || j == i { 
				combarr[i][j].SetInt64(1)
			} else {
				combarr[i][j].Add(&combarr[i-1][j-1],&combarr[i-1][j])
			}
		}
	}

	T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		N,Q := gi2(); A := make([]string,N); S := make([]int,N)
		for i:=0;i<N;i++ { A[i] = gs(); S[i] = gi() }
		ans,z,w := "",0,1
		if N == 1 {
			if 2*S[0] >= Q {
				ans,z,w = A[0],S[0],1
			} else {
				ansarr := make([]byte,Q)
				for i:=0;i<Q;i++ { if A[0][i] == 'T' { ansarr[i] = 'F' } else { ansarr[i] = 'T'} }
				ans,z,w = string(ansarr),Q-S[0],1
			}
			fmt.Fprintf(wrtr,"Case #%v: %v %v/%v\n",tt,ans,z,w)
		} else if N == 2 {
			// ta,ca = num where they both agree,num where they both agree that they got right
			// tb,cb = num where they disagree, num where they disagree that person 0 got right
			// ca + cb = score[0]
			// ca + (tb-cb) = score[1]
			ta,tb := 0,0
			for i:=0;i<Q;i++ { if A[0][i] == A[1][i] { ta++ } else { tb++ } }
			ca := (S[0] + S[1] - tb) / 2; cb := S[0] - ca
			z = 0
			ma := true; if 2*ca < ta { ma = false; z += ta-ca } else { z += ca }
			mb := true; if 2*cb < tb { mb = false; z += tb-cb } else { z += cb }
			ansarr := make([]byte,Q)
			for i:=0;i<Q;i++ {
				if A[0][i] == A[1][i] {
					if ma { ansarr[i] = A[0][i] } else if A[0][i] == 'T' { ansarr[i] = 'F' } else { ansarr[i] = 'T' }
				} else {
					if mb { ansarr[i] = A[0][i] } else { ansarr[i] = A[1][i] }
				}
			}
			ans = string(ansarr)
			fmt.Fprintf(wrtr,"Case #%v: %v %v/%v\n",tt,ans,z,w)
		} else {
			// tall, ta, tb, tc == total of all agree, a is different, b is different, c is different
			tall,ta,tb,tc := 0,0,0,0
			for i:=0;i<Q;i++ { 
				if A[0][i] == A[1][i] && A[0][i] == A[2][i] { tall++ } else if A[1][i] == A[2][i] { ta++ } else if A[0][i] == A[2][i] { tb++ } else { tc++ }
			}
			// 4 buckets, so 2^4 possibilities
			// call, ca, cb, cc = number the agreeing players got right
			// call + (ta-ca) + cb + cc = S[0]
			// call + ca + (tb-cb) + cc = S[1]
			// call + ca + cb + (tc-cc) = S[2]
			// loop through call values
			ways := big.NewInt(0)
			options := make([]big.Int,16)
			for i:=0;i<16;i++ { options[i].SetInt64(0) }
			temp   := big.NewInt(0)
			temp2a := big.NewInt(0); temp2b := big.NewInt(0)
			temp3a := big.NewInt(0); temp3b := big.NewInt(0)
			temp4a := big.NewInt(0); temp4b := big.NewInt(0)
			temp5a := big.NewInt(0); temp5b := big.NewInt(0)

			for call:=0; call<=tall; call++ {
				ca := (S[1]+S[2]-2*call-tb-tc) / 2
				cb := (S[0]+S[2]-2*call-ta-tc) / 2
				cc := (S[0]+S[1]-2*call-ta-tb) / 2
				if ca < 0 || ca > ta || cb < 0 || cb > tb || cc < 0 || cc > tc { continue }
				if call + (ta-ca) + cb + cc != S[0] { continue }
				if call + ca + (tb-cb) + cc != S[1] { continue }
				if call + ca + cb + (tc-cc) != S[2] { continue }
				// Calculate the number of ways
				lways := big.NewInt(1)
				if tall > 0 { lways.Mul(lways,&combarr[tall][call]) }
				if ta   > 0 { lways.Mul(lways,&combarr[ta][ca])     }
				if tb   > 0 { lways.Mul(lways,&combarr[tb][cb])     }
				if tc   > 0 { lways.Mul(lways,&combarr[tc][cc])     }
				ways.Add(ways,lways)
				temp2a.SetInt64(int64(call)); temp2b.SetInt64(int64(tall-call))
				temp3a.SetInt64(int64(ca));   temp3b.SetInt64(int64(ta-ca))
				temp4a.SetInt64(int64(cb));   temp4b.SetInt64(int64(tb-cb))
				temp5a.SetInt64(int64(cc));   temp5b.SetInt64(int64(tc-cc))
				for i:=0;i<16;i++ {
					if i % 16 < 8 { temp.Mul(temp2a,lways) } else { temp.Mul(temp2b,lways) }
					options[i].Add(&options[i],temp)
					if i % 8 < 4 { temp.Mul(temp3a,lways) } else { temp.Mul(temp3b,lways) }
					options[i].Add(&options[i],temp)
					if i % 4 < 2 { temp.Mul(temp4a,lways) } else { temp.Mul(temp4b,lways) }
					options[i].Add(&options[i],temp)
					if i % 2 < 1 { temp.Mul(temp5a,lways) } else { temp.Mul(temp5b,lways) }
					options[i].Add(&options[i],temp)
				}
			}

			best,bestval := -1,big.NewInt(0)
			for i:=0;i<16;i++ {
				c := bestval.Cmp(&options[i])
				if c < 0 { best = i; bestval.Set(&options[i]) }
			}
			ansarr := make([]byte,Q)
			for i:=0;i<Q;i++ {
				if A[0][i] == A[1][i] && A[0][i] == A[2][i] {
					if best % 16 < 8 { ansarr[i] = A[0][i] } else if A[0][i] == 'T' { ansarr[i] = 'F' } else { ansarr[i] = 'T' }
				} else if A[1][i] == A[2][i] {
					if best % 8 < 4 { ansarr[i] = A[1][i] } else { ansarr[i] = A[0][i] }
				} else if A[0][i] == A[2][i] {
					if best % 4 < 2 { ansarr[i] = A[2][i] } else { ansarr[i] = A[1][i] }
				} else {
					if best % 2 < 1 { ansarr[i] = A[0][i] } else { ansarr[i] = A[2][i] }
				}
			}
			frac := big.NewRat(1,1)
			frac.SetFrac(bestval,ways)
			num := frac.Num().String()
			denom := frac.Denom().String()
			ans := string(ansarr)
			fmt.Fprintf(wrtr,"Case #%v: %v %v/%v\n",tt,ans,num,denom)
			wrtr.Flush()
		}
	}
}

