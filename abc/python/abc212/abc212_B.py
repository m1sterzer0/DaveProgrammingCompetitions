
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
    X = gs()
    if X[0] == X[1] == X[2] == X[3] :
        ans = "Weak"
    else :
        weak = True
        XX = [int(c) for c in X]
        for i in range(3) :
            if (XX[i] + 1) % 10 != XX[i+1] : weak = False
        ans = "Weak" if weak else "Strong"
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

