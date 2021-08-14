
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def main(infn="") :
    myinf = 10**18
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    H,W = gis()
    bd = []
    for _ in range(H) : bd.append(gs())
    darr = [[myinf]*W for _ in range(H)]
    q = [(0,0)]; nextq = []; dist = 0; darr[0][0] = 0
    neighbors = [(-1,0),(1,0),(0,-1),(0,1)]
    punchsquares = [(-2,-1),(-2,0),(-2,1),
                    (-1,-2),(-1,-1),(-1,0),(-1,1),(-1,2),
                    (0,-2),(0,-1),(0,1),(0,2),
                    (1,-2),(1,-1),(1,0),(1,1),(1,2),
                    (2,-1),(2,0),(2,1)]
    while(True) :
        while q :
            (x,y) = q.pop()
            for (dx,dy) in neighbors :
                nx,ny = x+dx,y+dy
                if nx >= 0 and nx < H and ny >= 0 and ny < W and bd[nx][ny] == "." and darr[nx][ny] == myinf : darr[nx][ny] = dist; q.append((nx,ny))
            for (dx,dy) in punchsquares :
                nx,ny = x+dx,y+dy
                if nx >= 0 and nx < H and ny >= 0 and ny < W and bd[nx][ny] == '#' and darr[nx][ny] == myinf : darr[nx][ny] = dist+1; nextq.append((nx,ny))
        if darr[H-1][W-1] < myinf : ans = darr[H-1][W-1]; break
        dist += 1
        q,nextq = nextq,q
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

