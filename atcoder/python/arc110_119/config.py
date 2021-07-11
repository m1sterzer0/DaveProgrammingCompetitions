import argparse
import os.path
from pathlib import Path

def mkStarterFile(fn) :
    ttt = '''
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve() :
    return 0

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ans = solve()
    sys.stdout.write(str(ans)+'\\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()
'''
    with open(fn,'wt') as fp :
        print(ttt, file=fp)

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
    probList += [f"arc119_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc118_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc117_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc116_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc115_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc114_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc113_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc112_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc111_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc110_{x}" for x in ("A","B","C","D","E","F")]

    for prob in probList :
        if not os.path.exists(f"{clargs.dir}/{prob}.py") :
            Path(f"{clargs.dir}/{prob}.py").touch()
            mkStarterFile(f"{clargs.dir}/{prob}.py")


