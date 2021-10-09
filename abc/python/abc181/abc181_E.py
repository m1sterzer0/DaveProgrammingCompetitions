import sys
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
    N,M = gis()
    H = gis()
    W = gis()
    H.sort()
    W.sort()
    ## Calculate the sum of the pairings from the end
    ans = 8*10**18
    tail = sum(H[i+1]-H[i] for i in range(1,N,2))
    head = 0
    ptr = 0
    #print(f"DBG: N:{N} M:{M} H:{H} W:{W}")
    for s in range(0,N) :
        while ptr < (M-1) and abs(W[ptr+1]-H[s]) <= abs(W[ptr]-H[s]) :
            ptr += 1
        if s % 2 == 0 :
            if s > 0 : head += H[s-1]-H[s-2]
            cand = head + tail + abs(W[ptr]-H[s])
            #print(f"DBG: s:{s} head:{head} tail:{tail} ptr:{ptr} cand:{cand}")
            ans = min(ans,cand)
        else :
            tail -= H[s+1]-H[s]
            cand = head + tail + (H[s+1]-H[s-1]) + abs(W[ptr]-H[s])
            #print(f"DBG: s:{s} head:{head} tail:{tail} ptr:{ptr} cand:{cand}")
            ans = min(ans,cand)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

