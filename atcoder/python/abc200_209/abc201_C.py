
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
    S = gs()
    ans = 0
    for pin in range(10000) :
        a = pin // 1000; pin -= 1000*a
        b = pin // 100; pin -= 100*b
        c = pin // 10; pin -= 10*c
        d = pin
        good = True
        for i in range(10) :
            if S[i] == 'o' and i not in (a,b,c,d) : good = False; break
            if S[i] == 'x' and i in (a,b,c,d) : good = False; break
        if good : ans += 1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

