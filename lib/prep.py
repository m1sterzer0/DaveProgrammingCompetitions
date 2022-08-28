import argparse
import os.path
import shutil
from pathlib import Path


def mkGoStarterFile(fn,type) :
    header = '''package main
import (
    "bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10_000_000)
var rdr = bufio.NewScanner(os.Stdin)
'''
    googleheader = '''package main
import (
    "bufio"
    "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
'''
    stdfunc = '''func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
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
func sortUniq(a []int) []int {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}
'''

    int64func = '''func gi64() int64     { i,e := strconv.ParseInt(gs(),10,64); if e != nil {panic(e)}; return i }
func gis64(n int) []int64  { res := make([]int64,n); for i:=0;i<n;i++ { res[i] = gi64() }; return res }
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
func sortUniq64(a []int64) []int64 {
    sort.Slice(a,func(i,j int) bool { return a[i] < a[j] } )
    n,j := len(a),0; if n == 0 { return a }
    for i:=0;i<n;i++ { if a[i] != a[j] { j++; a[j] = a[i] } }; return a[:j+1]
}
'''
    maindefault = '''func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
}
'''
    googledefault = '''func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush(); infn := ""; if len(os.Args) > 1 { infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
	// PROGRAM STARTS HERE
    T := gi()
    for tt:=1;tt<=T;tt++ {
        fmt.Fprintf(wrtr,"Case #%v: %v\\n",tt,0)
    }
}
'''
    vars = [header,stdfunc,maindefault]
    if type == "cj" or type == "ks" : vars = [googleheader,stdfunc,googledefault]
    if type == "cf" : vars = [header,stdfunc,int64func,maindefault]

    with open(fn,'wt') as fp :
        for v in vars :
            print(v,file=fp) 
 
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

def mkCppStarterFile(fn,type) :
    header = '''#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef vector<ll> vi;
typedef pair<ll,ll> pi;
#define FOR(i,a) for (ll i = 0; i < (a); i++)
#define len(x) (ll) x.size()
const ll INF = 1LL << 62;
const ll MOD = 1000000007;
//const ll MOD = 998244353;
const double PI = 4*atan(double(1.0));
vector<vector<ll>> twodi(ll N, ll M, ll v) { vector<vector<ll>> res(N); for (auto &vv : res) vv.resize(M,v); return res; }
'''
    body = '''int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
}
'''
    googlebody = '''int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        printf("Case #%lld: %lld\\n",tt,0);
    }
}
'''
    vars = [header,body]
    if type == "cj" or type == "ks" : vars = [header,googlebody]
    with open(fn,'wt') as fp :
        for v in vars :
            print(v,file=fp) 
 
def mkCppLaunchJson(fn) :
    ttt = '''
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch C++",
            "type": "cppdbg",
            "request": "launch",
            "cwd": "${fileDirname}",
            "program": "${fileDirname}/${fileBasenameNoExtension}",
            "args": [ "${fileDirname}/junk.in" ]
        }
    ]
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkPyStarterFile(fn,type) :
    header = '''import sys
from collections import deque
## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007
'''
    body = '''def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE

if __name__ == "__main__" :
    main()
'''
    googlebody = '''def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        ans = 0
        print(f"Case #{tt}: {ans}")

if __name__ == "__main__" :
    main()
'''
    vars = [header,body]
    if type == "cj" or type == "ks" : vars = [header,googlebody]
    with open(fn,'wt') as fp :
        for v in vars :
            print(v,file=fp) 
 
def mkPyLaunchJson(fn) :
    ttt = '''
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Python: Current File",
            "type": "python",
            "request": "launch",
            "program": "${file}",
            "args": [ "${fileDirname}/junk.in" ],
            "console": "integratedTerminal",
            "justMyCode": true
        }
    ]
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkGitignore(fn) :
    with open(fn,'wt') as fp : 
        print("*.in\n*.out\n*.exe\n*.prof", file=fp)

def parseCLArgs() :
    clargparse = argparse.ArgumentParser()
    clargparse.add_argument( '--lang', required=True, action='store', choices=["go","py","cpp"], help='Language for the prep')
    clargparse.add_argument( '--dir', required=True, action='store', default='', help='Parent Directory for the preparations')
    clargparse.add_argument( '--type', action='store', choices=["cj","ks","hc","cf","abc","arc","agc"], default='def', help='Contest type for template customization')
    clargparse.add_argument( '--cname', required=True, action='store', help='Name of the contest')
    clargparse.add_argument( '--plist', required=True, action='store', help='String with comma separated list of problem names')
    clargs = clargparse.parse_args()
    if not os.path.exists(clargs.dir) : raise Exception(f"Directory '{clargs.dir}' does not exist.  Exiting...")
    return clargs

if __name__ == "__main__" :
    clargs = parseCLArgs()
    xx = clargs.plist.split(',')
    ## Prepend contest name to cf, abc, arc, and agc problems
    if clargs.type in ("cf","abc","arc","agc","cj"): 
        probList = [(clargs.cname,clargs.cname+"_"+x) for x in xx]
    else :
        probList = [(clargs.cname,x) for x in xx]

    if not os.path.exists(f"{clargs.dir}/.vscode") :
        os.mkdir(f"{clargs.dir}/.vscode")
    
    if not os.path.exists(f"{clargs.dir}/.gitignore") :
        Path(f"{clargs.dir}/.gitignore").touch()
        mkGitignore(f"{clargs.dir}/.gitignore")

    for (d,prob) in probList :
        if not os.path.exists(f"{clargs.dir}/{d}") : os.mkdir(f"{clargs.dir}/{d}")
        if not os.path.exists(f"{clargs.dir}/{d}/.vscode") :
            os.mkdir(f"{clargs.dir}/{d}/.vscode")
            if clargs.lang == "go" :  mkGoLaunchJson(f"{clargs.dir}/{d}/.vscode/launch.json")
            if clargs.lang == "py" :  mkPyLaunchJson(f"{clargs.dir}/{d}/.vscode/launch.json")
            if clargs.lang == "cpp" : mkCppLaunchJson(f"{clargs.dir}/{d}/.vscode/launch.json")

        if clargs.lang == "go" :
            if not os.path.exists(f"{clargs.dir}/{d}/{prob}") : 
                os.mkdir(f"{clargs.dir}/{d}/{prob}")
                mkGoStarterFile(f"{clargs.dir}/{d}/{prob}/{prob}.go",clargs.type)
                Path(f"{clargs.dir}/{d}/{prob}/junk.in").touch()
        elif clargs.lang == "py" :
            Path(f"{clargs.dir}/{d}/junk.in").touch()
            if not os.path.exists(f"{clargs.dir}/{d}/{prob}.py") : 
                mkPyStarterFile(f"{clargs.dir}/{d}/{prob}.py",clargs.type)
        elif clargs.lang == "cpp" :
            Path(f"{clargs.dir}/{d}/junk.in").touch()
            if not os.path.exists(f"{clargs.dir}/{d}/{prob}.cpp") : 
                mkCppStarterFile(f"{clargs.dir}/{d}/{prob}.cpp",clargs.type)



