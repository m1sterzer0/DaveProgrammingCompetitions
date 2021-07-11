
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

function solve(N::I,M::I,A::VI,X::VI,Y::VI)
    adj::VVI = [VI() for i in 1:N]
    for i in 1:M; push!(adj[X[i]],Y[i]); end
    inf::I = 10^18
    bdp::VI = fill(0,N)  ## bdp == "Best Downstream Price"
    maxprofit::I = -inf
    for n in N:-1:1
        for x in adj[n]; bdp[n] = max(bdp[n],bdp[x]); end
        if bdp[n] > 0; maxprofit = max(maxprofit,bdp[n]-A[n]); end
        bdp[n] = max(bdp[n],A[n]) ## Add our price to bdp after we process the "buy here" step above
    end
    return maxprofit
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = gis()
    X::VI = fill(0,M)
    Y::VI = fill(0,M)
    for i in 1:M; X[i],Y[i] = gis(); end
    ans = solve(N,M,A,X,Y)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

