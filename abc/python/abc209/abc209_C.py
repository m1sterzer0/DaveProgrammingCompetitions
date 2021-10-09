
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def fastpow(a,b,mod) :
    res = 1
    while b :
        if b & 1 :
            res = res * a % mod
        a = a * a % mod
        b >>= 1
    return res

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    mm = 10**9+7
    N = gi()
    C = gis()
    C.sort()
    ans = 1
    for (i,c) in enumerate(C) :
        numchoices = c - i
        if numchoices <= 0 : ans = 0; break
        ans = ans * numchoices % mm
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

