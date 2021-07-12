
import sys
import collections
import random
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def bruteAnalyze(N,S,setx) :
    gr = [[] for i in range(N)]
    for (a,b) in setx :
        gr[a].append(b)
        gr[b].append(a)
    
    def dfs(gr,S,visited,n,p) :
        visited[n] = True
        sz = 1
        groups = []
        bsz = 0 if S[n]=="#" else 1
        for c in gr[n] :
            if visited[c] : continue
            (csz,cgr,cbsz) = dfs(gr,S,visited,c,n)
            sz += csz; groups += cgr
            if bsz > 0 : bsz += cbsz
        if S[n] == "*" and (p < 0 or S[p] == '#') : groups.append(bsz)
        return (sz,groups,bsz)

    (ntot,gr,_bsz) = dfs(gr,S,[False]*N,0,-1)
    if ntot != N : return (False,0)
    pairs = 0
    for x in gr : pairs += x * (x-1) // 2
    return (True,pairs)

def solveBrute(N,S,E) :
    set1 = set()
    set2 = set()
    for i in range(N) :
        for j in range(i+1,N) :
            set1.add((i,j))
    edgelist = []
    for (i,x) in enumerate(E,start=1) :
        x-=1
        edge = (x,i) if x < i else (i,x)
        set1.remove(edge)
        set2.add(edge)

    (conflag,pairs) = bruteAnalyze(N,S,set2)
    bestpairs = pairs
    bestways = N-1  ## Add and remove edge within graph
    for e1add in set1 :
        for e1sub in set2 :
            setx = set2.copy()
            setx.remove(e1sub)
            setx.add(e1add)
            (conflag,pairs) = bruteAnalyze(N,S,setx)
            if not conflag : continue
            if pairs > bestpairs : (bestpairs,bestways) = (pairs,1)
            elif pairs == bestpairs : bestways += 1
    return f"{bestpairs} {bestways}"


def solve(N,S,E) :
    gr = [[] for i in range(N)]
    for (i,x) in enumerate(E,start=1) :
        x -= 1
        gr[i].append(x)
        gr[x].append(i)
    
    ## Get parents and BU Order 
    q = collections.deque()
    q.append((0,-1))
    tdorder = []
    par = [-1] * N
    while q :
        (n,p) = q.popleft()
        tdorder.append(n)
        par[n] = p
        for c in gr[n] :
            if c == p : continue
            q.append((c,n))
    buorder = tdorder[::-1]

    sz = [0] * N
    ccs = []
    for n in buorder :
        if S[n] == '#' : continue
        sz[n] = 1
        for c in gr[n] : sz[n] += sz[c] ## Parents will be zero, so no need to check for them here
        if par[n] == -1 or S[par[n]] == '#' : ccs.append(sz[n])
    ccs.sort(reverse=True)
    
    ## OK, now we have one of 4 cases
    ## Case 1)  No good patients.
    ## Case 2)  One cluster of good patients
    ## Case 3a) Top two clusters of good patients are the same size
    ## Case 3b) Top two clusters of patients are different sizes

    ## In all 4 cases, we look to see if we can remove the edge above us and then count the number of connections we can make
    ans1,ans2,groupsz = 0,0,0
    if len(ccs) == 0 or len(ccs) == 1 and ccs[0] == 1:  ## Added special case for a singleton good patient
        ans1 = 0
        sz2 = [0] * N
        for n in buorder :
            if par[n] == -1 : continue
            sz2[n] = 1
            for c in gr[n] : sz2[n] += sz2[c]
            ans2 += sz2[n] * (N-sz2[n])
    elif len(ccs) == 1 :
        ans1 = ccs[0] * (ccs[0]-1) // 2
        sz2 =  [0] * N
        bsz = [0] * N
        for n in buorder :
            if par[n] == -1 : continue
            sz2[n] = 1
            bsz[n] = 1 if S[n] == '*' else 0
            for c in gr[n] : sz2[n] += sz2[c]; bsz[n] += bsz[c]
            if S[n] == '#' or S[par[n]] == '#' :
                ans2 += sz2[n] * (N-sz2[n])
            else :
                ans2 += bsz[n] * (ccs[0]-bsz[n])
    else :
        ## Color all of the groups 
        ans1 = (ccs[0]+ccs[1]) * (ccs[0]+ccs[1]-1) // 2
        for x in ccs[2:] : ans1 += x * (x-1) // 2
        bigsz,smallsz = ccs[0],ccs[1]
        case3a = bigsz == smallsz
        numbig = sz.count(bigsz)
        numsmall = 0 if case3a else sz.count(smallsz)
        numbigbelow = [0] * N
        numsmallbelow = [] if case3a else [0] * N
        for n in buorder :
            for c in gr[n] : numbigbelow[n] += numbigbelow[c]
            if sz[n] == bigsz : numbigbelow[n] += 1
            if not case3a :
                for c in gr[n] : numsmallbelow[n] += numsmallbelow[c]
                if sz[n] == smallsz : numsmallbelow[n] += 1
            if par[n] < 0 or (S[n] == '*' and S[par[n]] == '*') : continue  ## Can't disconnect a group for max
            if case3a :
                ans2 += (bigsz*bigsz) * (numbigbelow[n] * (numbig-numbigbelow[n]))
            else :
                ans2 += (bigsz*smallsz) * numbigbelow[n] * (numsmall-numsmallbelow[n])
                ans2 += (bigsz*smallsz) * numsmallbelow[n] * (numbig-numbigbelow[n])

    return f"{ans1} {ans2}"

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ",end="")
        N,K = gis()
        S = gs()
        E = gis()
        A,B,C = gis()
        for i in range(K+2,N+1) : E.append((A * E[-2] + B * E[-1] + C) % (i-1) + 1)
        print(f"    DBG: ntc:{ntc} N:{N}", file=sys.stderr)
        ans = solve(N,S,E)
        print(ans)

def test(ntc,grmin,grmax) :
    numpassed = 0
    for tc in range(1,ntc+1) :
        ## Testgen code goes here.
        N = random.randrange(grmin,grmax+1)
        sprob = random.random()
        schars = ['*' if random.random() < sprob else '#' for i in range(N)]
        S = "".join(schars)
        E = [random.randrange(1,i+2) for i in range(N-1)]

        ans1 = solveBrute(N,S,E)
        ans2 = solve(N,S,E)
        if ans1 == ans2 :
            numpassed += 1
        else :
            print(f"ERROR: tc:{tc} N:{N} sprob:{sprob} ans1:'{ans1}' ans2:'{ans2}'")
            solveBrute(N,S,E)
            solve(N,S,E)
    print(f"{numpassed}/{ntc} passed")

if __name__ == '__main__' :
    random.seed(8675309)
    main()
    #for ntc in (10,100,1000,10000,50000) : test(ntc,3,10)
    #for ntc in (10,100,1000,10000)       : test(ntc,11,16)
    #for ntc in (10,100,1000)             : test(ntc,17,25)
    sys.stdout.flush()

