
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
    N,K = gis()
    A = gis()
    q = [(a,1) for a in A]
    q.append((0,1))
    q.sort()
    ans = 0
    while(q) :
        (a,cnt) = q.pop()
        if a == 0 : break
        while q and q[-1][0] == a : cnt += q[-1][1]; q.pop()
        na = q[-1][0]
        maxrides = cnt * (a - na)
        if maxrides < K :
            ans += cnt * (a * (a+1) // 2 - na * (na+1) // 2)
            K -= maxrides
            q.append((na,cnt))
        else :
            fullsets = K // cnt
            ans += cnt * (a * (a+1) // 2 - (a-fullsets) * (a-fullsets+1) // 2)
            K -= fullsets * cnt
            ans += K * (a-fullsets)
            break
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

