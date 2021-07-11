
import random
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def solve(N,A) :
    nimber = 0
    for x in A[2:] : nimber ^= x
    twosum = A[0] + A[1]
    ## Want to find largest a s.t. a <= A[0] with  a xor (twosum-a) = nimber --> a xor nimber = twosum - a --> a + (a xor nimber) = twosum
    ## Now note a + b = a xor b + 2 * (a AND b) -->
    ## a xor (a xor nimber) + 2 * a & (a xor nimber) = twosum --> 2 * a & (a xor nimber) = twosum - nimber
    ## Thus parity of twosum and nimber have to match.  Assuming so
    ## define B = (twosum - nimber) // 2
    ## a & (a xor nimber) = B
    ## Now this is a bitwise operation.  Truth table
    ## a bit  nimber bit   a xor nimber  result
    ## -----  ----------   ------------  ------
    ##   0        0              0         0
    ##   1        0              1         1
    ##   0        1              1         0
    ##   1        1              0         0
    ## For bits where nimber is 0, we must set a to give the desired result
    ## For bits where nimber is 1, if B is zero we are impossible.  Otherwise we have free choice.
    if nimber > twosum : return -1
    if nimber % 2 != twosum % 2 : return -1
    b = (twosum - nimber) >> 1
    for i in range(62) : 
        if (nimber >> i) & 1 and (b >> i) & 1 :
            return -1
    baseline = 0
    for i in range(62) : 
        if (nimber >> i) & 1 == 0 :
            baseline += (b & (1 << i))
    if baseline > A[0] : return -1
    for i in range(62,-1,-1) : 
        if (nimber >> i) & 1 == 1 :
            if baseline + (1 << i) <= A[0] : 
                baseline += (1 << i)
    if baseline == 0 : return -1
    return A[0] - baseline
            
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    ans = solve(N,A)
    sys.stdout.write(str(ans)+'\n')

if __name__ == '__main__' :
    random.seed(19006492568)
    main()
    sys.stdout.flush()

