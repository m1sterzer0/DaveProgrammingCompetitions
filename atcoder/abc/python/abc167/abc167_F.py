
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def check(up) :
    s = 0
    for (m,d) in up :
        if m > s : return False
        s += d
    return True

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi(); S = []
    for _ in range(N) : S.append(gs())
    up = []; dn = []
    for s in S :
        d,m = 0,0
        for c in s :
            d += (1 if c == '(' else -1)
            m = min(m,d)
        if d >= 0 : up.append((0-m,d))
        else      : dn.append((d-m,-d))
    up.sort()
    dn.sort()
    ok = True
    if sum(x[1] for x in up) != sum(x[1] for x in dn) : ok = False
    if ok and not check(up) : ok = False
    if ok and not check(dn) : ok = False
    ans = "Yes" if ok else "No"
    print(ans)
    
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

