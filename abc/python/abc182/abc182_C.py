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
    N = gi()
    digits = [int(x) for x in str(N)]
    cur = N % 3
    mod1 = len([x for x in digits if x % 3 == 1])
    mod2 = len([x for x in digits if x % 3 == 2])
    if cur == 0 :
        ans = 0
    elif cur == 1 :
        ans = 1 if mod1 >= 1 and len(digits) > 1 else 2 if mod2 >= 2 and len(digits) > 2 else -1
    else :
        ans = 1 if mod2 >= 1 and len(digits) > 1 else 2 if mod1 >= 2 and len(digits) > 2 else -1
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

