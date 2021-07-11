
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,I,O) :
    res = [0] * N
    for i in range(N) : res[i] = ["N"] * N; res[i][i] = "Y"
    for i in range(N) :
        for j in range(i+1,N) :
            if O[j-1] == "N" or I[j] == "N" : break
            res[i][j] = "Y"

    for i in range(N-1,-1,-1) :
        for j in range(i-1,-1,-1) :
            if O[j+1] == "N" or I[j] == "N" : break
            res[i][j] = "Y"

    ansarr = []
    for i in range(N) :
        ansarr.append("".join(res[i]))
    ans = "\n".join(ansarr)
    return ans

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}:")
        N = gi()
        I = gs()
        O = gs()
        ans = solve(N,I,O)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

