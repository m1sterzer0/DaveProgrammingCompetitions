import argparse
import os.path
from pathlib import Path
import shutil

def mkGoStarterFile(fn) :
    ttt = '''
package main
import (
	"bufio"
    "fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
        fmt.Fprintf(wrtr,"Case #%v: %v\\n",tt,0)
    }
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkGoLaunchJson(fn) :
    ttt = '''
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}"
        }
    ]
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkGoSettingsJson(fn) :
    ttt = '''
{
    "[go]": {"editor.formatOnSave": false }  
}    
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)


def mkGoGitignore(fn) :
    with open(fn,'wt') as fp : 
        print("*.in\n*.out\n*.exe\n*.prof", file=fp)

def parseCLArgs() :
    clargparse = argparse.ArgumentParser()
    clargparse.add_argument( '--dir', action='store', default='', help='Parent Directory for the preparations')
    clargs = clargparse.parse_args()
    if not clargs.dir  : raise Exception("Need to provide a --dir option.  Exiting...")
    if not os.path.exists(clargs.dir) : raise Exception(f"Directory '{clargs.dir}' does not exist.  Exiting...")
    return clargs

if __name__ == "__main__" :
    clargs = parseCLArgs()
    probList = []
    probList += [(f"2021",x) for x in ["Reversort","MoonsAndUmbrellas","ReversortEngineering","MedianSort","CheatingDetection",
                                      "AppendSort","PrimeTime","HackedExam",
                                      "BrokenClock","Subtransmutation","DigitBlocks",
                                      "ClosestPick","RoaringYears","DoubleOrNOTing",
                                      "MinimumSort","Matrygons","HiddenPancakes","Retiling",
                                      "BuildAPair","SquareFree","FenceDesign","BinarySearchGame",
                                      "CuttingCake","SlideCircuits","Ropes","DivisibleDivisions","Infinitree"]]
    probList += [(f"2022",x) for x in ["PunchedCards","3DPrinting","D1000000","ChainReactions","TwistyLittlePassages",
                                       "DoubleOrOneThing","EqualSum","Weightlifting"
                                    
    ]]


    if not os.path.exists(f"{clargs.dir}/.vscode") :
        os.mkdir(f"{clargs.dir}/.vscode")
        Path(f"{clargs.dir}/.vscode/launch.json").touch()
        mkGoLaunchJson(f"{clargs.dir}/.vscode/launch.json")
        mkGoSettingsJson(f"{clargs.dir}/.vscode/settings.json")
    
    if not os.path.exists(f"{clargs.dir}/.gitignore") :
        Path(f"{clargs.dir}/.gitignore").touch()
        mkGoGitignore(f"{clargs.dir}/.gitignore")

    for (d,prob) in probList :
        if not os.path.exists(f"{clargs.dir}/{d}") : os.mkdir(f"{clargs.dir}/{d}")
        if not os.path.exists(f"{clargs.dir}/{d}/.vscode") :
            os.mkdir(f"{clargs.dir}/{d}/.vscode")
            shutil.copyfile(f"{clargs.dir}/.vscode/launch.json",f"{clargs.dir}/{d}/.vscode/launch.json")
            shutil.copyfile(f"{clargs.dir}/.vscode/settings.json",f"{clargs.dir}/{d}/.vscode/settings.json")


        if not os.path.exists(f"{clargs.dir}/{d}/{prob}") : 
            os.mkdir(f"{clargs.dir}/{d}/{prob}")
            Path(f"{clargs.dir}/{d}/{prob}/{prob}.go").touch()
            Path(f"{clargs.dir}/{d}/{prob}/junk.in").touch()
            mkGoStarterFile(f"{clargs.dir}/{d}/{prob}/{prob}.go")


