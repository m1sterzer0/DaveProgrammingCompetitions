
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

## Modular inverse of 10 mod 2019 is 202
## Load up an array of all of the cases that end in the rightmost digit

## Lets say we want to chop off 3 digits and shift
## Let a be one of the original numbers
## Then we want (a - a3)*10^-3 == 0 --> a*10^-3 == a3*10^-3 --> a == a3
## This works because 202 is a modular inverse of 10 in 2019.

def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    S = gs()
    digits = [int(x) for x in S[::-1]]
    pv = 1
    vals = []
    rem = [0] * 2019
    running = 0
    for d in digits :
        running = (running + d * pv) % 2019
        rem[running] += 1
        vals.append(running)
        pv = pv * 10 % 2019
    ans = rem[0]
    for v in vals :
        rem[v] -= 1
        ans += rem[v]
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

