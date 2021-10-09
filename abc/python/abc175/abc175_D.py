
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,preP,C) :
    P = [x-1 for x in preP]
    ## First, count the cost of cycles, this should be O(N)
    visited = [False for i in range(N)]
    cycsum = [0 for i in range(N)]
    cyclen = [0 for i in range(N)]
    updatel = []
    for i in range(N) :
        if visited[i] : continue
        csum = 0; j = i; updatel.clear()
        while not visited[j] : csum += C[j]; updatel.append(j); visited[j] = True; j = P[j]
        for c in updatel : cyclen[c] = len(updatel); cycsum[c] = csum
    best = -10**18
    for st in range(N) :
        cur = 0; loc = st
        for i in range(min(cyclen[st],K)) :
            loc = P[loc]; cur += C[loc]; best = max(best,cur,cur+(K-i-1)//cyclen[st]*cycsum[st])
    return best

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    P = gis()
    C = gis()
    ans = solve(N,K,P,C)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

