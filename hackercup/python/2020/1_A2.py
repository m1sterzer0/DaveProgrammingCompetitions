
import sys
import heapq
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## For this version of the problem, a height only gets set once for an interval.
## Seems like we should do some coordinate compression
## Also seems like offline processing from left to right could be easier than trying to do it in real time
def solve(N,W,L,H) :
    coords = [0]
    for (l,w) in zip(L,W) : coords.append(l); coords.append(l+w)
    coords = list(set(coords))
    coords.sort()
    ## First pass -- figure out who is the first person to occupy a particular interval and the height after that occupation
    events = [(L[i],i) for i in range(N)]
    events.sort(reverse=True)
    minh = []
    ht   = [0] * len(coords)
    fidx = [-1] * len(coords)
    for (i,c) in enumerate(coords) :
        while events and events[-1][0] == c :
            (_,x) = events.pop()
            heapq.heappush(minh,x)
        while minh and L[minh[0]] + W[minh[0]] <= c : 
            heapq.heappop(minh)
        if minh : fidx[i] = minh[0]; ht[i] = H[minh[0]]
    
    perim = [0] * N
    for i in range(len(coords)) :
        if ht[i] == 0 : continue
        wid = coords[i+1]-coords[i]
        j = fidx[i]
        perim[j] += 2 * wid
        for c in (i-1,i+1) :
            if fidx[c] < 0 or fidx[c] > j :
                perim[j] += ht[i]
            elif fidx[c] < j :
                perim[j] += abs(ht[i]-ht[c]) - ht[c]
        perim[j] %= 1_000_000_007

    running = 1; p = 0
    for pdel in perim :
        p = (p + pdel) % 1_000_000_007 
        running = running * p % 1_000_000_007
    return running

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        print(f"Case #{ntc}: ",file=sys.stderr)
        N,K = gis()
        L = gis()
        AL,BL,CL,DL = gis()
        W = gis()
        AW,BW,CW,DW = gis()
        H = gis()
        AH,BH,CH,DH = gis()
        for i in range(K,N) : L.append((AL*L[-2]+BL*L[-1]+CL) % DL + 1)
        for i in range(K,N) : W.append((AW*W[-2]+BW*W[-1]+CW) % DW + 1)
        for i in range(K,N) : H.append((AH*H[-2]+BH*H[-1]+CH) % DH + 1)
        ans = solve(N,W,L,H)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

