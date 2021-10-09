
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

function solve(N::I,M::I,A::VI,B::VI,K::I,C::VI,D::VI)
    ## Just do bitmasks
    sb::VB = fill(false,N)
    best::I = 0
    for m in 0:2^K-1
        fill!(sb,false)
        for i in 1:K
            if m & (1 << (i-1)) == 0; sb[C[i]] = true; else; sb[D[i]] = true; end
        end
        cand::I = 0
        for i in 1:M
            if sb[A[i]] && sb[B[i]]; cand += 1; end
        end
        best = max(cand,best)
    end
    return best 
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    for i in 1:M; A[i],B[i] = gis(); end
    K = gi()
    C::VI = fill(0,K)
    D::VI = fill(0,K)
    for i in 1:K; C[i],D[i] = gis(); end
    ans = solve(N,M,A,B,K,C,D)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

