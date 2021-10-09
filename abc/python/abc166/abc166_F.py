
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
    N,A,B,C = gis()
    S = []
    for _ in range(N) : S.append(gs())
    good = True; moves = []
    def play(s1,s2) :
        nonlocal A,B,C
        if s1 == 'A'   : A += 1; moves.append('A')
        elif s1 == 'B' : B += 1; moves.append('B')
        elif s1 == 'C' : C += 1; moves.append('C')
        if s2 == 'A'   : A -= 1
        elif s2 == 'B' : B -= 1
        elif s2 == 'C' : C -= 1
    for i,s in enumerate(S) :
        c1 = A if s[0] == 'A' else B if s[0] == 'B' else C
        c2 = A if s[1] == 'A' else B if s[1] == 'B' else C
        if c1 == c2 == 0 : good = False; break
        elif c1 == 0 : play(s[0],s[1])
        elif c2 == 0 : play(s[1],s[0])
        elif i+1 == N : play(s[0],s[1])
        elif s[0] in S[i+1] : play(s[0],s[1])
        else : play(s[1],s[0])
    if not good :
        print ("No")
    else :
        print("Yes")
        ansstr = "\n".join(moves)
        print(ansstr)

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

