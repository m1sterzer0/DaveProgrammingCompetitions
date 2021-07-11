
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
    H,W = gis()
    A = []
    for i in range(H) :
        x = gs(); A.append(x)
    dp = [0] * H
    for i in range(H) : dp[i] = [0] * W
    for i in range(H-1,-1,-1) :
        for j in range(W-1,-1,-1) :
            if i == H-1 and j == W-1 : continue
            if (i+j) % 2 == 0 : ## Takahashi's turn (maximizer)
                if i == H-1 : dp[i][j] = dp[i][j+1] + (1 if A[i][j+1] == "+" else -1); continue
                if j == W-1 : dp[i][j] = dp[i+1][j] + (1 if A[i+1][j] == "+" else -1); continue
                v1 = dp[i][j+1] + (1 if A[i][j+1] == "+" else -1)
                v2 = dp[i+1][j] + (1 if A[i+1][j] == "+" else -1)
                dp[i][j] = max(v1,v2)
            else : ## Aoki's turn (minimizer)
                if i == H-1 : dp[i][j] = dp[i][j+1] + (-1 if A[i][j+1] == "+" else 1); continue
                if j == W-1 : dp[i][j] = dp[i+1][j] + (-1 if A[i+1][j] == "+" else 1); continue
                v1 = dp[i][j+1] + (-1 if A[i][j+1] == "+" else 1)
                v2 = dp[i+1][j] + (-1 if A[i+1][j] == "+" else 1)
                dp[i][j] = min(v1,v2)
    ans = "Takahashi" if dp[0][0] > 0 else "Aoki" if dp[0][0] < 0 else "Draw"
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

