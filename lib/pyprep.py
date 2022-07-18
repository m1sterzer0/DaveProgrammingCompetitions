import argparse
import os.path
import shutil
from pathlib import Path


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
        print(f"Case #{tt}: {ans}");

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
    clargparse.add_argument( '--dir', required=True, action='store', default='', help='Parent Directory for the preparations')
    clargparse.add_argument( '--type', action='store', choices=["cj","ks","cf","abc","arc","agc"], default='def', help='Contest type for template customization')
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
            mkPyLaunchJson(f"{clargs.dir}/{d}/.vscode/launch.json")
            Path(f"{clargs.dir}/{d}/junk.in").touch()
        if not os.path.exists(f"{clargs.dir}/{d}/{prob}.py") : 
            mkPyStarterFile(f"{clargs.dir}/{d}/{prob}.py",clargs.type)

