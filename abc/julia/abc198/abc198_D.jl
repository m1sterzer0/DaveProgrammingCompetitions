
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

mutable struct UnsafeIntPerm; n::I; r::I; indices::VI; cycles::VI; end
Base.eltype(iter::UnsafeIntPerm) = Vector{Int64}
function Base.length(iter::UnsafeIntPerm)
    ans::I = 1; for i in iter.n:-1:iter.n-iter.r+1; ans *= i; end
    return ans
end
function unsafeIntPerm(a::VI,r::I=-1) 
    n = length(a)
    if r < 0; r = n; end
    return UnsafeIntPerm(n,r,copy(a),collect(n:-1:n-r+1))
end
function Base.iterate(p::UnsafeIntPerm, s::I=0)
    n = p.n; r=p.r; indices = p.indices; cycles = p.cycles
    if s == 0; return(n==r ? indices : indices[1:r],s+1); end
    for i in (r==n ? n-1 : r):-1:1
        cycles[i] -= 1
        if cycles[i] == 0
            k = indices[i]; for j in i:n-1; indices[j] = indices[j+1]; end; indices[n] = k
            cycles[i] = n-i+1
        else
            j = cycles[i]
            indices[i],indices[n-j+1] = indices[n-j+1],indices[i]
            return(n==r ? indices : indices[1:r],s+1)
        end
    end
    return nothing
end

function process(s::String,d::Dict{Char,I},numlet::I)::VI
    ans::VI = fill(0,numlet)
    for (i,c) in enumerate(reverse(s)); ans[d[c]] += 10^(i-1); end
    return ans
end

function solve(S1::String,S2::String,S3::String)
    d::Dict{Char,I} = Dict{Char,I}()
    numlet::I = 0 
    for c in S1*S2*S3
        if !haskey(d,c); numlet += 1; d[c] = numlet; end
    end
    if numlet > 10; return "UNSOLVABLE"; end
    pv1::VI = process(S1,d,numlet)
    pv2::VI = process(S2,d,numlet)
    pv3::VI = process(S3,d,numlet)
    pv123::VI = [pv1[i]+pv2[i]-pv3[i] for i in 1:numlet]

    badzero::VI = [d[S1[1]],d[S2[1]],d[S3[1]]]
    for p in unsafeIntPerm(collect(0:9),numlet)
        good = true
        for b in badzero; if p[b] == 0; good = false; break; end; end
        if !good; continue; end
        residual = sum(pv123[i]*p[i] for i in 1:numlet)
        if residual != 0; continue; end
        a::I = sum(pv1[i]*p[i] for i in 1:numlet)
        b::I = sum(pv2[i]*p[i] for i in 1:numlet)
        c::I = sum(pv3[i]*p[i] for i in 1:numlet)
        return "$a\n$b\n$c"
    end
    return "UNSOLVABLE"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    S1 = gs()
    S2 = gs()
    S3 = gs()
    ans = solve(S1,S2,S3)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

