
import random
import sys
import numpy as np
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(R,C,K,r,c,v) :
    ## Build a grid with the numbers
    pregifts = [(rr-1,cc-1,vv) for (rr,cc,vv) in zip(r,c,v)]
    pregifts.sort(reverse=True)
    gifts = np.zeros((K,3),dtype=np.int64)
    for (i,(x,y,z)) in enumerate(pregifts) :
        gifts[i,:] = x,y,z
    idx = 0
    dp1 = np.zeros(C,dtype=np.int64)
    dp2 = np.zeros(C,dtype=np.int64)
    gi = gifts[0,0]; gj = gifts[0,1]
    for i in range(R-1,-1,-1) :
        best0 = 0; best1 = 0; best2 = 0
        for j in range(C-1,-1,-1) :
            dp1j = dp1[j]
            if i == gi and j == gj :
                val = gifts[idx,2]
                if idx < K-1 : idx += 1; gi = gifts[idx,0]; gj = gifts[idx,1]
                vpdp1j = val + dp1j
                best0 = max(best0,vpdp1j,best1+val)
                best1 = max(best1,vpdp1j,best2+val)
                best2 = best2 if best2 > vpdp1j else vpdp1j
            else :
                best0 = best0 if best0 > dp1j else dp1j
                best1 = best1 if best1 > dp1j else dp1j
                best2 = best2 if best2 > dp1j else dp1j
            dp2[j] = best0
        dp1,dp2 = dp2,dp1
    return dp1[0]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    R,C,K = gis()
    r = np.zeros(K,dtype=np.int64)
    c = np.zeros(K,dtype=np.int64)
    v = np.zeros(K,dtype=np.int64)
    for i in range(K) : r[i],c[i],v[i] = gis()
    ans = solve(R,C,K,r,c,v)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    NumbaFlag = True
    if not NumbaFlag :
        random.seed(19006492568)
        main()
        sys.stdout.flush()
    elif sys.argv[-1] == 'ONLINE_JUDGE':
        from numba.pycc import CC
        cc = CC('my_module')
        cc.export('solve', 'i8(i8,i8,i8,i8[:],i8[:],i8[:])')(solve)
        cc.compile()
        exit()
    else :
        # noinspection PyUnresolvedReferences
        from my_module import solve
        random.seed(19006492568)
        main()
        sys.stdout.flush()

