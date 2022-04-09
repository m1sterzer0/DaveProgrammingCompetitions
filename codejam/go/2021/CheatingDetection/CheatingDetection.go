package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)

	T := gi()
	gi() // Throw away P value -- don't need it
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
		// Premise is that players of similar skill should have similar "inversion count" when looking at results of problems from easiest to hardest
		bd := make([]string,100); for i:=0;i<100;i++ { bd[i] = gs() }

		// 1) given a skill level s, calculate the probability p of getting a random problem right 
		//    From Algebra: p = \frac16\cdot\int_{-3}^{3} \frac{1}{1 + e^{x-s}}dx
		//    From wolfram alpha: p = (1/6) * (ln(exp(s+3)+1) - ln(exp(s-3)+1)
		// 2) Given a probability p of getting a random number right, calculate the skill level
		//    From wolfram alpha: s = 3 + ln(exp(6p)-1) - ln(exp(6)-exp(6p))
		calcSkill := func(p float64) float64 { 
			s := 0.00
			if p <= 0.0 { s = -3.00 } else if p >= 1.0 { s = 3.000 } else { s = 3.0 + math.Log(math.Exp(6*p)-1) - math.Log(math.Exp(6)-math.Exp(6*p)) }
			if s > 3.0 { s = 3.0 }; if s < -3.0 { s = -3.0 }; return s
		}

		// First, calculate the expected difficulty of each problem
		probSkill := make([]float64,10000)
		for j:=0;j<10000;j++ {
			numWrong := 0
			for i:=0;i<100;i++ { if bd[i][j] == '0' { numWrong++ } }
			p := float64(numWrong)/float64(100)
			probSkill[j] = calcSkill(p)
		}

		// Use Bayes on each player, assuming each has 50% chance of being a cheater
		// After Bayes, report the one with a highest probability of being a cheater
		best,bestlp := 0,-1e99
		for i:=0;i<100;i++ {
			numRight := 0
			for j:=0;j<10000;j++ { if bd[i][j] == '1' { numRight++} }
			sfair :=  calcSkill(float64(numRight)/float64(10000))
			scheat := calcSkill(float64(numRight-5000)/float64(5000))
			pcheat := 0.5; lpcheat := math.Log(pcheat)
			for j:=0;j<10000;j++ {
				probFair := 1.0 / (1 + math.Exp(probSkill[j]-sfair))
				probCheat := 0.5 + 0.5 / (1 + math.Exp(probSkill[j]-scheat))
				if bd[i][j] == '0' { probFair = 1.0 - probFair; probCheat = 1.0 - probCheat }
				pb := pcheat * probCheat + (1.0-pcheat) * probFair
				lpcheat += math.Log(probCheat) - math.Log(pb)
				pcheat = math.Exp(lpcheat)
			}
			if lpcheat > bestlp { best,bestlp = i,lpcheat }
		}
        fmt.Fprintf(wrtr,"Case #%v: %v\n",tt,best+1)
    }
}

