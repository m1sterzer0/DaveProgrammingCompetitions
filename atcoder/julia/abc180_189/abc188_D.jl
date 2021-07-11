
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

function solve(N::I,C::I,a::VI,b::VI,c::VI)
    events::VPI = []
    for i in 1:N; push!(events,(a[i],c[i])); push!(events,(b[i]+1,-c[i])); end
    sort!(events)
    ans::I = 0
    perday::I = 0
    lastprocessed = 0
    for (t,cc) in events
        if t-1 > lastprocessed
            ans += min(perday,C) * (t-1-lastprocessed)
            lastprocessed = t-1
        end
        perday += cc
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,C = gis()
    a::VI = fill(0,N)
    b::VI = fill(0,N)
    c::VI = fill(0,N)
    for i in 1:N; a[i],b[i],c[i] = gis(); end
    ans = solve(N,C,a,b,c)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

