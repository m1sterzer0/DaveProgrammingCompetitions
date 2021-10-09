
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
    N,K = gis()
    c = gis()
    s = {}
    for i in range(K) : 
        if c[i] not in s : s[c[i]] = 0
        s[c[i]] += 1
    ans = len(s)
    for i in range(K,N) :
        j = i-K
        if c[i] not in s : s[c[i]] = 0
        s[c[i]] += 1
        s[c[j]] -= 1
        if s[c[j]] == 0 : del s[c[j]]
        ans = max(ans,len(s))
    print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

