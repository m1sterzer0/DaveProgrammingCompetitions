
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,K,A) :
    mm = 10**9+7
    Apos = sorted((a for a in A if a > 0), reverse=True)
    Aneg = sorted(a for a in A if a < 0)
    Azero = sum(1 for a in A if a == 0)
    if K > len(Apos) + len(Aneg) : return 0
    if len(Apos) == 0 and K%2 == 1 and Azero > 0 : return 0
    if len(Apos) + len(Aneg) == K and Azero > 0 and len(Aneg) % 2 == 1 : return 0
    ans = 1
    if len(Apos) + len(Aneg) == K and (Azero == 0 or len(Aneg) % 2 == 0) :
        ans = 1
        for a in Apos : ans = ans * a % mm
        for a in Aneg : ans = ans * a % mm
        return (mm+ans) % mm
    if len(Apos) == 0 :
        if K%2 == 1 : Aneg.reverse()
        for i in range(K) : ans = ans * Aneg[i] % mm
        return (mm+ans) % mm
    pi,ni = 0,0
    if K % 2 == 1 : pi = 1
    while pi + 1 < K and pi + 1 < len(Apos) : pi += 2
    while pi + ni < K : ni += 2
    while ni + 1 < len(Aneg) and pi - 2 >= 0 and Aneg[ni]*Aneg[ni+1] > Apos[pi-1]*Apos[pi-2] : ni += 2; pi -= 2
    for i in range(pi) : ans = ans * Apos[i] % mm
    for i in range(ni) : ans = ans * Aneg[i] % mm
    return (mm + ans) % mm

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N,K = gis()
    A = gis()
    ans = solve(N,K,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

