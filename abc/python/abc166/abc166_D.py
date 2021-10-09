
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
    X = gi()
    pow5 = [0] * 1001
    for i in range(1001) : pow5[i] = i**5
    a,b = 0,0
    for i in range(1,1000) :
        for j in range(i+1) :
            if X == pow5[i]-pow5[j] : a,b = i,j
            if X == pow5[i]+pow5[j] : a,b = i,-j
    print(f"{a} {b}")

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

