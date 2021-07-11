
import sys
import collections
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,M,C) :
    st = collections.deque()
    st.append((0,0))
    for i in range(1,N) :
        while st and i - st[0][0] > M : st.popleft()
        if not st : return -1
        if C[i] > 0 :
            tc = st[0][1] + C[i]
            while st and st[-1][1] >= tc : st.pop()
            st.append((i,tc))
    return st[0][1]

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    T = gi()
    for ntc in range(1,T+1) :
        print(f"Case #{ntc}: ", end="")
        N,M = gis()
        C = [0] * N
        for i in range(N) : C[i] = gi()
        ans = solve(N,M,C)
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

