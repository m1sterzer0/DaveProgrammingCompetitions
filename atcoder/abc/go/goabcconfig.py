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
    "io"
	"os"
	"strconv"
	"strings"
)

type scanner struct { sc *bufio.Scanner }
func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e+9))
	return &scanner{sc}	
}
func (s *scanner) s() string  { s.sc.Scan(); return s.sc.Text() }
func (s *scanner) i() int     { i,e := strconv.Atoi(s.s()); if e != nil {panic(e)}; return i }
func (s *scanner) f() float64 { f,e := strconv.ParseFloat(s.s(),64); if e != nil {panic(e)}; return f }
func (s *scanner) bs() []byte { return []byte(s.s()) }
func (s *scanner) is(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = s.i() }; return res }
func (s *scanner) fs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = s.f() }; return res }
func (s *scanner) ss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = s.s() }; return res }

var rdr = newScanner(os.Stdin)
const BUFSIZE = 10000000
var wrtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)

func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
type PI struct { x,y int }
type TI struct { x,y,z int }

func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	infn := ""
	if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
    if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = newScanner(f) }
	
    // NON-BOILERPLATE STARTS HERE
	ans := 0
    fmt.Fprintln(wrtr, ans); wrtr.Flush()
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
    probList += [(f"abc160",f"abc160_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc161",f"abc161_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc162",f"abc162_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc163",f"abc163_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc164",f"abc164_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc165",f"abc165_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc166",f"abc166_{x}") for x in ("A","B","C","D","E","F")]
    probList += [(f"abc167",f"abc167_{x}") for x in ("A","B","C","D","E","F")]

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

