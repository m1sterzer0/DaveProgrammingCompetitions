
import sys
infile = sys.stdin.buffer

def gs()  : return infile.readline().rstrip()
def gi()  : return int(gs())
def gf()  : return float(gs())
def gss() : return gs().split()
def gis() : return [int(x) for x in gss()]
def gfs() : return [float(x) for x in gss()]

def unwind(A,sb,rem) :
    ans = [sb[rem]+1]
    rem = (rem + 200 - A[sb[rem]]) % 200
    while rem > 0 :
        ans.append(sb[rem]+1)
        rem = (rem + 200 - A[sb[rem]]) % 200
    ans.sort()
    return ans

def solve(N,A) :

    cand = []
    sb = [-1] * 200
    rs = A[0] % 200; sb[rs] = 0
    for i in range(1,N) :
        rem = A[i] % 200
        if sb[rem] > -1 :
            l2 = [i+1]
            l1 = unwind(A,sb,rem)
            return ("Yes",l1,l2)
        elif sb[0] > -1 :
            l2 = [i+1]
            l1 = unwind(A,sb,0) + [i+1]
            return ("Yes",l1,l2)
        else :
            cand.clear()
            for j in range(200) :
                if sb[j] == -1 : continue
                cand.append(j)
            for j in cand :
                targ = (j + rem) % 200
                if sb[targ] > -1 :
                    l1 = unwind(A,sb,targ)
                    l2 = unwind(A,sb,j) + [i+1]
                    return ("Yes",l1,l2)
            sb[rem] = i
            for j in cand :
                targ = (j + rem) % 200
                sb[targ] = i
    return ("No",[],[])
        
def main(infn="") :
    global infile
    infile = open(infn,"r") if infn else open(sys.argv[1],"r") if len(sys.argv) > 1 else sys.stdin
    N = gi()
    A = gis()
    (res,l1,l2) = solve(N,A)
    if res == "No" : 
        print("No")
    else :
        print("Yes")
        print(f"{len(l1)} ",end="")
        print(" ".join(str(x) for x in l1))
        print(f"{len(l2)} ",end="")
        print(" ".join(str(x) for x in l2))

if __name__ == '__main__' :
    main()
    sys.stdout.flush()

