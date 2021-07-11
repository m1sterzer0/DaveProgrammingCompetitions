
import sys
#import random
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

    ## Nodes with no children are winning positions
    ## Nodes which only lead to winnning positions are losing positions
    ## Nodes which lead to a losing position are winning positions
    ## All other positions are draws

    N = gi()
    S = ["aaaaaaaa"] * N
    for i in range(N) : S[i] = gs()

    ans = solve(N,S)
    sys.stdout.write(str(ans)+'\n')

def solve(N,S) :
    ## First, uniquify the edges by representing each word as a 6 letter string
    edges = list(set([x[:3] + x[-3:] for x in S]))
    nodes = list(set([x[:3] for x in S] + [x[-3:] for x in S]))
    numnodes = len(nodes)
    n2nid = {}
    for (i,n) in enumerate(nodes) : n2nid[n] = i
    gr =    [ [] for i in range(numnodes) ]
    grrev = [ [] for i in range(numnodes) ]
    for e in edges :
        n1 = n2nid[e[:3]]
        n2 = n2nid[e[-3:]]
        gr[n1].append(n2)
        grrev[n2].append(n1)
    depcnt =      [len(gr[i]) for i in range(numnodes)]
    numchildren = [len(gr[i]) for i in range(numnodes)]
    losers  = [i for i in range(numnodes) if numchildren[i] == 0]
    winners = []
    sb = [2] * numnodes
    for l in losers : sb[l] = 0
    while (winners or losers) :
        if winners :
            w = winners.pop()
            for x in grrev[w] :
                if sb[x] != 2 : continue
                depcnt[x] -= 1
                if depcnt[x] == 0 :
                    sb[x] = 0
                    losers.append(x)
        else :
            l = losers.pop()
            for x in grrev[l] :
                if sb[x] != 2 : continue
                sb[x] = 1
                winners.append(x)
    ansarr = []
    for s in S :
        n = n2nid[s[-3:]]
        ansarr.append("Takahashi" if sb[n] == 0 else "Aoki" if sb[n] == 1 else "Draw")
    ans = "\n".join(ansarr)
    return ans

#def test(letstr,numwordsmin,numwordsmax,wordlenmin,wordlenmax) :
#    N = random.randrange(numwordsmin,numwordsmax+1)
#    print(N)
#    S = []
#    for i in range(N) :
#        numlet = random.randrange(wordlenmin,wordlenmax+1)
#        chars = [random.choice(letstr) for j in range(numlet)]
#        ss = "".join(chars)
#        print(ss)
#        S.append(ss)
#    ans = solve(N,S)
#    print("ANSWER:")
#    print(ans)
    
if __name__ == '__main__' :
    #random.seed(8675309)
    main()
    #for i in range(10) : test("abc",6,10,3,6)
    sys.stdout.flush()

