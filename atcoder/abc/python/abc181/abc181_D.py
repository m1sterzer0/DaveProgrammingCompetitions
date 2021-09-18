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
    ans = "No"
    if len(S) == 1 :
        if S[0] == "8" : ans = "Yes"
    else :
        c = [0] * 10
        for cc in S :
            x = int(cc)
            c[x] += 1
        if len(S) == 2 :
            for i in range(100) :
                if i % 8 > 0 : continue
                d1 = i % 10
                d2 = i // 10
                if d1 == d2 and c[d1] >= 2 : ans = "Yes"
                if d1 != d2 and c[d1] >= 1 and c[d2] >= 1 : ans = "Yes"
        else :
            for i in range(1000) :
                if i % 8 > 0 : continue
                d1 = i % 10; i2 = (i - d1) // 10
                d2 = i2 % 10; i3 = (i2 - d2) // 10
                d3 = i3 % 10
                if d1 == d2 and d1 == d3:
                    if c[d1] >= 3 : ans = "Yes"
                elif d1 == d2 :
                    if c[d1] >= 2 and c[d3] >= 1 : ans = "Yes"
                elif d1 == d3 :
                    if c[d1] >= 2 and c[d2] >= 1 : ans = "Yes"
                elif d2 == d3 :
                    if c[d2] >= 2 and c[d1] >= 1 : ans = "Yes"
                else :
                    if c[d1] >= 1 and c[d2] >= 1 and c[d3] >= 1 : ans = "Yes"
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

