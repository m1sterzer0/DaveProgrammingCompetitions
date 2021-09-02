
import sys
import collections
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
    K = gi()
    q = collections.deque()
    for i in range(1,10) : q.append(i)
    nin = 9; nout = 0; ans = 0
    while True :
        ans = q.popleft(); nout += 1
        if nout == K : break
        ld = ans % 10
        if ld > 0 and nin < K : q.append(10*ans+ld-1); nin += 1
        if nin < K : q.append(10*ans+ld); nin += 1
        if ld < 9 and nin < K : q.append(10*ans+ld+1); nin += 1
    print(ans)
    
if __name__ == '__main__' :
    main()
    sys.stdout.flush()

