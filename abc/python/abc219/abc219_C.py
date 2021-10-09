
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
    N = gi()
    S = []
    for _ in range(N) : S.append(gs())
    t1 = {}; t2 = {}
    for c,d in zip(X,"abcdefghijklmnopqrstuvwxyz") :
        t1[ord(c)] = ord(d)
        t2[ord(d)] = ord(c)
    for i in range(N) : S[i] = S[i].translate(t1)
    S.sort()
    for i in range(N) : S[i] = S[i].translate(t2)
    for s in S : print(s)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

