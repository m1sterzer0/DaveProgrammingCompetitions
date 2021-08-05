
import sys
import collections
import re
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def prework() :
    ans = {}
    for y in (4,6,8,10,12,20) :
        maxsize = 20*y+2
        d = [0] * maxsize; nd = [0] * maxsize; d[0] = 1
        for x in range(1,20+1) :
            for i in range(maxsize) : nd[i] = 0 
            oldmax = (x-1)*y
            for j in range(oldmax+1) :
                for k in range(1,y+1) :
                    nd[j+k] += d[j]
            denom = y**x
            cum = 0
            p = []
            for j in range(maxsize) :
                p.append(1.0 - cum/denom)
                cum += nd[j]
            ans[(y,x)] = p
            nd,d = d,nd
    return ans

def solve(H,S,Sarr,pw) :
    best = 0.0
    for s in Sarr :
        x1 = re.match(r'(\d+)d(\d+)(([+-])(\d+))?',s)
        g1,g2,g3 = x1.group(1),x1.group(2),x1.group(3)
        X,Y = int(g1),int(g2)
        if g3 :
            g4,g5 = x1.group(4),x1.group(5)
            #print(f"DBG: g1:{g1} g2:{g2} g3:{g3} g4:{g4} g5:{g5}")
            Z = int(g5) * (1 if g4 == '+' else -1)
        else :
            Z = 0
        h = H - Z
        ans = pw[(Y,X)]
        if h < 0 : return 1.00
        if h >= len(ans) : continue
        best = max(best,ans[h])
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    pw = prework()
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        H,S = gis()
        Sarr = gss()
        print(f"Case {ntc} H:{H} S:{S}",file=sys.stderr)
        ans = solve(H,S,Sarr,pw)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

