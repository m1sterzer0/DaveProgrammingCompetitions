
import random
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(H,W,Ch,Cw,Dh,Dw,bd) :
    Ch -= 1; Cw -= 1; Dh -= 1; Dw -= 1
    if Ch == Dh and Cw == Dw : return 0
    inf = 10**18
    d = [[inf for j in range(W)] for i in range(H)]
    d[Ch][Cw] = 0
    level = 0; q=collections.deque()
    q.append((Ch,Cw))
    while q :
        q2 = []
        while q :
            (i,j) = q.popleft()
            q2.append((i,j))
            for (i2,j2) in ((i-1,j),(i+1,j),(i,j-1),(i,j+1)) :
                if i2 < 0 or i2 >= H or j2 < 0 or j2 >= W : continue
                if bd[i2][j2] == '#' or d[i2][j2] < inf : continue
                if i2 == Dh and j2 == Dw : return level
                d[i2][j2] = level
                q.append((i2,j2))
        for (i,j) in q2 :
            for di in (-2,-1,0,1,2) :
                for dj in (-2,-1,0,1,2) :
                    i2 = i+di; j2 = j+dj
                    if i2 < 0 or i2 >= H or j2 < 0 or j2 >= W : continue
                    if bd[i2][j2] == '#' or d[i2][j2] < inf : continue
                    if i2 == Dh and j2 == Dw : return level+1
                    d[i2][j2] = level+1
                    q.append((i2,j2))
        level += 1
    return -1

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W = gis()
    Ch,Cw = gis()
    Dh,Dw = gis()
    bd = []
    for i in range(H) : bd.append(gs())
    ans = solve(H,W,Ch,Cw,Dh,Dw,bd)
    sys.stdout.write(str(ans)+"\n")

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

