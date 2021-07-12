
import sys

infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solveBrute(N,W,L,H) :
    cmax = 0
    for (l,w) in zip(L,W) : cmax = max(cmax,l+w)
    harr = [0] * (cmax+5)
    p = 0; running = 1; dbgarr = []
    for (l,w,h) in zip(L,W,H) :
        for x in range(l,l+w) :
            if harr[x] == 0 : p = (p + 2) % 1_000_000_007
            if harr[x] >= h : continue
            p = (p - abs(harr[x]-harr[x-1]) - abs(harr[x]-harr[x+1]) + abs(h-harr[x-1]) + abs(h-harr[x+1])) % 1_000_000_007
            harr[x] = h
        dbgarr.append(p)
        running = running * p % 1_000_000_007
    return running 

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        print(ntc,file=sys.stderr)
        N,K,WW = gis()
        L = gis()
        AL,BL,CL,DL = gis()
        H = gis()
        AH,BH,CH,DH = gis()
        for i in range(K,N) : L.append((AL*L[-2]+BL*L[-1]+CL) % DL + 1)
        for i in range(K,N) : H.append((AH*H[-2]+BH*H[-1]+CH) % DH + 1)
        W = [WW] * N
        ans = solveBrute(N,W,L,H)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

