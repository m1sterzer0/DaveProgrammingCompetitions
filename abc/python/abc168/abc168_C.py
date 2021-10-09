
import random
import sys
import math
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
    A,B,H,M = gis()
    ## 12*60 = 720
    ## Hour moves 1 tick per minute
    ## Minute moves 12 ticks per minute
    hpos = (60*H+M)
    mpos = 12 * (60*H+M) % 720
    pi = 4 * math.atan(1)
    angle = 2*pi / 720 * (mpos-hpos)
    d = math.sqrt(A*A+B*B-2*A*B*math.cos(angle))
    print(d)

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

