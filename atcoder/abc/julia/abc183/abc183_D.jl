
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

function solve(N::I,W::I,S::VI,T::VI,P::VI)::String
    dd::VI = fill(0,200001)
    for i in 1:N; dd[S[i]+1] += P[i]; dd[T[i]+1] -= P[i]; end
    demand::I = 0
    for i in 1:200001
        demand += dd[i];
        if demand > W; return "No"; end
    end
    return "Yes"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,W = gis()
    S::VI = fill(0,N)
    T::VI = fill(0,N)
    P::VI = fill(0,N)
    for i in 1:N; S[i],T[i],P[i] = gis(); end
    ans = solve(N,W,S,T,P)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

