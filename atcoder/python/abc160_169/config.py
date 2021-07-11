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

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin

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
    probList += [f"abc169_{x}" for x in ("A","B","C","D","E","F")]
    probList += [f"abc168_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc167_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc166_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc165_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc164_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc163_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc162_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc161_{x}" for x in ("A","B","C","D","E","F")]
    #probList += [f"abc160_{x}" for x in ("A","B","C","D","E","F")]

    for prob in probList :
        if not os.path.exists(f"{clargs.dir}/{prob}.py") :
            Path(f"{clargs.dir}/{prob}.py").touch()
            mkStarterFile(f"{clargs.dir}/{prob}.py")


