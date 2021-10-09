
using Random
infile = stdin
## Type Shortcuts (to save my wrists and fingers :))
const I = Int64; const VI = Vector{I}; const SI = Set{I}; const PI = NTuple{2,I};
const TI = NTuple{3,I}; const QI = NTuple{4,I}; const VPI = Vector{PI}; const SPI = Set{PI}
const VC = Vector{Char}; const VS = Vector{String}; VB = Vector{Bool}; VVI = Vector{Vector{Int64}}
const F = Float64; const VF = Vector{F}; const PF = NTuple{2,F}

gs()::String = rstrip(readline(infile))
gi()::Int64 = parse(Int64, gs())
gf()::Float64 = parse(Float64,gs())
gss()::Vector{String} = split(gs())
gis()::Vector{Int64} = [parse(Int64,x) for x in gss()]
gfs()::Vector{Float64} = [parse(Float64,x) for x in gss()]

## Each operation reduces inversion count by 1 or 2
## Note that an odd number of inversions only happens when we start on a different character from which we end
## If we count the "wrap around inversion", then every string has an even number of inversions.
## Thus, we can just count total raw inversions (instead of doing per-string accounting) and divide by 2 at the end.
## Now the code is simple, the conversion of per-string to total is the key insight to make this easy.

function gencase(Nmin,Nmax,Kmin,Kmax)
    N = rand(Nmin:Nmax)
    K = rand(Kmin:Kmax)
    S = join(rand(['?','0','1'],N),"")
    return (S,K)
end

function test(ntc::I,Nmin::I,Nmax::I,Kmin::I,Kmax::I)
    pass = 0
    for ttt in 1:ntc
        (S,K) = gencase(Nmin,Nmax,Kmin,Kmax)
        ans1 = solve(S,K)
        ans2 = solve(S^K,1)
        if ans1 == ans2
            pass += 1
        else
            print("ERROR Case:$ttt S:$S K:$K ans1:$ans1 ans2:$ans2\n")
            ans1 = solve(S,K)
            ans2 = solve(S^K,1)
        end
    end
    print("$pass/$ntc passed\n")
end

function solve(S::String,K::I)
    if length(S)*K == 1 return 0; end
    mm::I = 10^9+7
    n::I = length(S)
    q::I = count(x->x=='?',S)
    numstrings = powermod(2,K*q,mm)
    halfstrings = numstrings * invmod(2,mm) % mm
    ans::I = 0
    for i in 1:n
        j = i == n ? 1 : i+1 
        s = S[i]*S[j]
        if      s in ("01","10");                ans = (ans + (K * numstrings  % mm)) % mm
        elseif  s in ("0?","1?","?0","?1","??"); ans = (ans + (K * halfstrings % mm)) % mm
        end
    end
    ans = ans * invmod(2,mm) % mm
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    S = gs()
    K = gi()
    ans = solve(S,K)
    println(ans)
end

Random.seed!(8675309)
main()
#test(1000,1,3,1,2)
#test(100,1,100,1,100)
#test(1000,1,100,1,100)


#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

