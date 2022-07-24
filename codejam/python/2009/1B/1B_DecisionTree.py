import sys
from collections import defaultdict, deque, namedtuple

## Input crap
infile = sys.stdin
intokens = deque()
def getTokens(): 
    while not intokens:
        for c in infile.readline().rstrip().split() :
            intokens.append(c)    
def gs(): getTokens(); return intokens.popleft()
def gi(): return int(gs())
def gf(): return float(gs())
def gbs(): return [c for c in gs()]
def gis(n): return [gi() for i in range(n)]
def ia(m): return [0] * m
def iai(m,v): return [v] * m
def twodi(n,m,v): return [iai(m,v) for i in range(n)]
def fill2(m) : r = gis(2*m); return r[0::2],r[1::2]
def fill3(m) : r = gis(3*m); return r[0::3],r[1::3],r[2::3]
def fill4(m) : r = gis(4*m); return r[0::4],r[1::4],r[2::4],r[3::4]
MOD = 998244353
##MOD = 1000000007

def main() :
    if len(sys.argv) > 1 : global infile; infile = open(sys.argv[1],'rt')
    ## PROGRAM STARTS HERE
    T = gi()
    for tt in range(1,T+1) :
        L = gi()
        raw = " ".join([infile.readline().rstrip() for _ in range(L)] )
        tokens = raw.replace("("," ( ").replace(")"," ) ").rstrip().lstrip().split()
        Treenode = namedtuple("Treenode","w f t1 t0")
        tree = []
        def tparse(curs) :
            tidx = len(tree); tree.append(Treenode(float(tokens[curs+1]),"",-1,-1))
            if tokens[curs+2] == ")" : return curs+2
            w = tree[tidx].w; f = tokens[curs+2]; t1 = tidx+1
            curs = tparse(curs+3); t0 = len(tree); curs = tparse(curs+1)
            tree[tidx] = Treenode(w,f,t1,t0)
            return curs+1
        tparse(0)
        features = defaultdict(bool)
        def evaltree(tidx) :
            if tree[tidx].f == "" : return tree[tidx].w
            if features[tree[tidx].f] : return tree[tidx].w * evaltree(tree[tidx].t1)
            return tree[tidx].w * evaltree(tree[tidx].t0)
        print(f"Case #{tt}:")
        A = gi()
        for i in range(A) :
            aname = gs(); n = gi(); farr = [gs() for _ in range(n)]
            for f in farr : features[f] = True
            ans = evaltree(0)
            print(ans)
            for f in farr : features[f] = False

if __name__ == "__main__" :
    main()

