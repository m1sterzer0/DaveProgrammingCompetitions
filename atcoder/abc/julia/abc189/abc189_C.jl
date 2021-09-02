
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

function solve(N::I,A::VI)
    st::VPI = [(-100,0)]
    left::VI = fill(0,N)
    right::VI = fill(0,N)
    for i in 1:N
        while A[i] < st[end][1]; pop!(st); end
        left[i] = st[end][2]+1
        push!(st,(A[i],i))
    end
    empty!(st); push!(st,(-100,N+1))
    for i in N:-1:1
        while A[i] <= st[end][1]; pop!(st); end
        right[i] = st[end][2]-1
        push!(st,(A[i],i))
    end
    best::I = 0
    for i in 1:N; best = max(best,(right[i]-left[i]+1)*A[i]); end
    return best
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    A::VI = gis()
    ans = solve(N,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

