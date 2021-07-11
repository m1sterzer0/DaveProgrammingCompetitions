import argparse
import os.path
from pathlib import Path

def mkStarterFile(fn) :
    ttt = '''
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    ans = 0 
    sys.stdout.write(str(ans)+'\\n')

if __name__ == '__main__' :
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
    probList += [f"arc100_{x}" for x in ("C","D","E","F")]
    probList += [f"arc101_{x}" for x in ("C","D","E","F")]
    #probList += [f"arc102_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc103_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc104_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc105_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc106_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc107_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc108_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"arc109_{x}" for x in ("A","B","C","D","E","F")]




    #probList += [f"abc175_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc174_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc173_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc172_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc171_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc170_{x}" for x in ("A","B","C","D","E","F")]

    for prob in probList :
        if not os.path.exists(f"{clargs.dir}/{prob}.py") :
            Path(f"{clargs.dir}/{prob}.py").touch()
            mkStarterFile(f"{clargs.dir}/{prob}.py")


