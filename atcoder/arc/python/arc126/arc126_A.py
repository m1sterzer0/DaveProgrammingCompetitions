
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
    T = gi()
    for _ in range(T) :
        n2,n3,n4 = gis()
        ## First build 3+3+4
        ## Second build 3+3+2+2
        ## Thrid build 4+4+2
        ## Fourth build 4+2+2+2
        ## Finally build 2+2+2+2+2
        ans = 0
        c1 = min(n3//2,n4);    ans += c1; n3 -= 2*c1; n4 -= c1
        c2 = min(n3//2,n2//2); ans += c2; n3 -= 2*c2; n2 -= 2*c2
        c3 = min(n4//2,n2);    ans += c3; n4 -= 2*c3; n2 -= c3
        c4 = min(n4,n2//3);    ans += c4; n4 -= c4;   n2 -= 3*c4
        c5 = n2//5;            ans += c5; n2 -= 5*c5
        print(ans)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

