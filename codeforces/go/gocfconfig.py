import argparse
import os.path
from pathlib import Path
import shutil

def mkGoStarterFile(fn) :
    ttt = '''
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func gi64() int64     { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gis64(n int) []int64  { res := make([]int64,n); for i:=0;i<n;i++ { res[i] = gi64() }; return res }

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

func ia64(m int) []int64 { return make([]int64,m) }
func iai64(m int,v int64) []int64 { a := make([]int64,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi64(n int,m int,v int64) [][]int64 {
	r := make([][]int64,n); for i:=0;i<n;i++ { x := make([]int64,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill264(m int) ([]int64,[]int64) { a,b := ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i] = gi64(),gi64()}; return a,b }
func fill364(m int) ([]int64,[]int64,[]int64) { a,b,c := ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi64(),gi64(),gi64()}; return a,b,c }
func fill464(m int) ([]int64,[]int64,[]int64,[]int64) { a,b,c,d := ia64(m),ia64(m),ia64(m),ia64(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi64(),gi64(),gi64(),gi64()}; return a,b,c,d }
func abs64(a int64) int64 { if a < 0 { return -a }; return a }
func rev64(a []int64) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max64(a,b int64) int64 { if a > b { return a }; return b }
func min64(a,b int64) int64 { if a > b { return b }; return a }
func maxarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr64(a []int64) int64 { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr64(a []int64) int64 { ans := int64(0); for _,aa := range(a) { ans += aa }; return ans }
func zeroarr64(a []int64) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod64(a,e,mod int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint64(a,e int64) int64 { res, m := int64(1), a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd64(a,b int64) int64 { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended64(a,b int64) (int64,int64,int64) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended64(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv64(a,m int64) (int64,bool) { g,x,_ := gcdExtended64(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecint64string(a []int64) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.FormatInt(a,10) }; return strings.Join(astr," ") }
func makefact64(n int,mod int64) ([]int64,[]int64) {
	fact,factinv := make([]int64,n+1),make([]int64,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * int64(i) % mod }
	factinv[n] = powmod64(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * int64(i+1) % mod }
	return fact,factinv
}
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1_000_000_000)
	// PROGRAM STARTS HERE
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
            "program": "${file}",
            "args": [ "${fileDirname}/junk.in" ]
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
    probList += [(f"r786",  f"r786_{x}")   for x in ("A","B","C","D","E","F","G")]
    probList += [(f"r787",  f"r787_{x}")   for x in ("A","B","C","D","E","F","G")]
    probList += [(f"r788",  f"r788_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r789",  f"r789_{x}")   for x in ("A","B1","B2","C","D","E","F")]
    probList += [(f"r791",  f"r791_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r792",  f"r792_{x}")   for x in ("A","B","C","D","E")]
    probList += [(f"r793",  f"r793_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r794",  f"r794_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r795",  f"r795_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r796",  f"r796_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r798",  f"r798_{x}")   for x in ("A","B","C","D","E")]
    probList += [(f"r800",  f"r800_{x}")   for x in ("A","B","C","D","E")]
    probList += [(f"r801",  f"r801_{x}")   for x in ("A","B","C","D1","D2","E")]
    probList += [(f"r802",  f"r802_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"r803",  f"r803_{x}")   for x in ("A","B","C","D","E","F","G")]
    probList += [(f"r804",  f"r804_{x}")   for x in ("A","B","C","D","E")]

    probList += [(f"gr1",  f"gr1_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"gr2",  f"gr2_{x}")   for x in ("A","B","C","D","E")]
    probList += [(f"gr3",  f"gr3_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"gr4",  f"gr4_{x}")   for x in ("A","B","C","D","E","F1","F2")]
    probList += [(f"gr5",  f"gr5_{x}")   for x in ("A","B","C1","C2","D","E","F")]
    probList += [(f"gr6",  f"gr6_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"gr7",  f"gr7_{x}")   for x in ("A","B","C","D1","D2","E","F1")]
    probList += [(f"gr8",  f"gr8_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"gr9",  f"gr9_{x}")   for x in ("A","B","C","D","E","F")]
    probList += [(f"gr10", f"gr10_{x}")  for x in ("A","B","C","D","E","F","G")]
    probList += [(f"gr11", f"gr11_{x}")  for x in ("A","B","C","D","E")]


    #probList += [(f"abc216",f"abc216_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc217",f"abc217_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc218",f"abc218_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc222",f"abc222_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc223",f"abc223_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc224",f"abc224_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc225",f"abc225_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc226",f"abc226_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc227",f"abc227_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc228",f"abc228_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc229",f"abc229_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc230",f"abc230_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc231",f"abc231_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc232",f"abc232_{x}") for x in ("A","B","C","D","E","F","G","H")]
    #probList += [(f"abc233",f"abc233_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc234",f"abc234_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc235",f"abc235_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc236",f"abc236_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc237",f"abc237_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc238",f"abc238_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc239",f"abc239_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc240",f"abc240_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc241",f"abc241_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc242",f"abc242_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc243",f"abc243_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc244",f"abc244_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc245",f"abc245_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc246",f"abc246_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc247",f"abc247_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc248",f"abc248_{x}") for x in ("A","B","C","D","E","F","G","Ex")]
    #probList += [(f"abc249",f"abc249_{x}") for x in ("A","B","C","D","E","F","G","Ex")]


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

