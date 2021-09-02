
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

function solve(N::I,M::I,H::VI,W::VI)
    inf::I = 4*10^18
    HH::VI = copy(H)
    sort!(HH)
    WW::VI = copy(W)
    sort!(WW)
    frombeg::VI = fill(inf,N)
    fromend::VI = fill(inf,N)
    for i in 2:2:N
        frombeg[i] = HH[i]-HH[i-1] + (i > 2 ? frombeg[i-2] : 0)
    end
    for i in N-1:-2:1
        fromend[i] = HH[i+1]-HH[i] + (i<N-1 ? fromend[i+2] : 0)
    end
    best = inf
    ptr::I = 1
    for i in 1:2:N
        while ptr != M && WW[ptr] < HH[i]; ptr += 1; end
        cand = min(abs(HH[i] - WW[ptr]), ptr==1 ? inf : abs(HH[i]-WW[ptr-1])) + (i == 1 ? 0 : frombeg[i-1]) + (i == N ? 0 : fromend[i+1])
        best = min(best,cand)
    end
    return best
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    H::VI = gis()
    W::VI = gis()
    ans = solve(N,M,H,W)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

