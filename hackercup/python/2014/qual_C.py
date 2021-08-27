
import sys
import random
from multiprocessing import Pool
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

MOD = 1_000_000_007

def getInputs(tt) :
    xx = gss()
    K = int(xx[0])
    ps = int(1000*float(xx[1])+0.5)
    pr = int(1000*float(xx[2])+0.5)
    pi = int(1000*float(xx[3])+0.5)
    pu = int(1000*float(xx[4])+0.5)
    pw = int(1000*float(xx[5])+0.5)
    pd = int(1000*float(xx[6])+0.5)
    pl = int(1000*float(xx[7])+0.5)
    return (tt,K,ps,pr,pi,pu,pw,pd,pl)

def solvemulti(xx) :
    (tt,K,ps,pr,pi,pu,pw,pd,pl) = xx
    print(f"Solving case {tt} (K={K})...",file=sys.stderr)
    return solve(K,ps,pr,pi,pu,pw,pd,pl)

def solve(K,ps,pr,pi,pu,pw,pd,pl) :
    cur = {(0,pi):1.000}; last = {}
    for _ in range(2*K-1) :
        last,cur = cur,last; cur.clear()
        for ((w,psun),prob) in last.items() :
            ## Win and increased sun prob
            p1 = 0.000000001*( psun*ps*pw               + (1000-psun)*pr*pw               )* prob
            ## Win and neutral sun probability
            p2 = 0.000000001*( psun*ps*(1000-pw)        + (1000-psun)*pr*(1000-pw)        )* prob
            ## Loss and decreased sun prob
            p3 = 0.000000001*( psun*(1000-ps)*pl        + (1000-psun)*(1000-pr)*pl        )* prob
            ## Loss and neutral sun prob
            p4 = 0.000000001*( psun*(1000-ps)*(1000-pl) + (1000-psun)*(1000-pr)*(1000-pl) )* prob

            pinc = min(1000,psun+pu)
            pdec = max(0,psun-pd)

            if (w+1,pinc) not in cur : cur[(w+1,pinc)] = 0.000
            if (w+1,psun) not in cur : cur[(w+1,psun)] = 0.000
            if (w,psun)   not in cur : cur[(w,psun)]   = 0.000
            if (w,pdec)   not in cur : cur[(w,pdec)]   = 0.000

            cur[(w+1,pinc)] += p1
            cur[(w+1,psun)] += p2
            cur[(w,psun)]   += p4
            cur[(w,pdec)]   += p3
    ans = 0.000
    for ((w,psun),prob) in cur.items() :
        if w >= K : ans += prob
    return "%.6f" % ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    multi = True
    if multi :
        inputs = []
        for tt in range(1,T+1) : inputs.append(getInputs(tt))
        with Pool(processes=8) as pool : outputs = pool.map(solvemulti,inputs)
        for tt,ans in enumerate(outputs,1) : print(f"Case #{tt}: {ans}")
    else :
        for tt in range(1,T+1) : 
            inp = getInputs(tt)
            ans = solvemulti(inp)
            print(f"Case #{tt}: {ans}")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    sys.stdout.flush()

