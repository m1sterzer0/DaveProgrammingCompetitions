
import os
import sys
import time
import collections
from fractions import Fraction

############################################################################################################################################
## FILE IO STUFF
defaultFile = ""
infh = sys.stdin if defaultFile == "" and len(sys.argv) < 2 else open(sys.argv[1],'rt') if defaultFile == "" else open(defaultFile,'rt')
tokens = collections.deque()
def gs() :
    while len(tokens) == 0 :
        for x in infh.readline().rstrip().split() : tokens.append(x)
    return tokens.popleft()
def gi() : return int(gs())
def gf() : return float(gs())
def gis(n) : return [ gi() for i in range(n) ]
############################################################################################################################################

if __name__ == "__main__" :
    start = time.time()
    T = gi()
    for tt in range(1,T+1) :
        N,W = gi(),gi(); gi(); P = gi(); gi(); R = gi(); gi()
        if R < P : P,R = R,P
        X = [0]*N; A = [0]*N; B = [0]*N
        for i in range(N) : X[i] = gi(); gi(); A[i] = gi(); B[i] = gi()
        setvals = set([0,W])
        for i in range(N) : setvals.add(X[i]); setvals.add(X[i]+P); setvals.add(X[i]+R)
        xvals = sorted(setvals)
        running,a,b = Fraction(0,1),Fraction(0,1),Fraction(0,1)
        bsum = sum(B); running += -bsum; best = abs(running); h = Fraction(2,R)
        for i,x1 in enumerate(xvals) :
            if i == 0 or i == len(xvals)-1 : continue 
            a *= 0; b *= 0; x2 = xvals[i+1]
            for i in range(N) :
                if x1 >= X[i] and x1 < X[i] + P and P != 0 :
                    b += Fraction(x1-X[i],P) * h * (A[i]+B[i])
                    a += h * Fraction(1,P) * (A[i]+B[i])
                elif x1 >= X[i] + P and x1 < X[i]+R and P != R :
                    b += Fraction(X[i]+R-x1,R-P) * h * (A[i]+B[i])
                    a += -h * Fraction(1,R-P) * (A[i]+B[i])
            if a != 0 :
                xcrit = -b/a 
                if 0 < xcrit and xcrit < (x2-x1) :
                    cand = running - b * b / 2 / a
                    if cand * running <= 0 :
                        best *= 0; break
                    if abs(cand) < best : best = abs(cand)
            newrunning = running + b * (x2-x1) + a * (x2-x1) * (x2-x1) / 2
            if running * newrunning <= 0 :
                best *= 0; break
            if abs(newrunning) < best : best = abs(newrunning)
            running = newrunning
        print(f"Case #{tt}: {best.numerator}/{best.denominator}")
    duration = time.time()-start
    print(f"Execution Time: {duration}",file=sys.stderr)




