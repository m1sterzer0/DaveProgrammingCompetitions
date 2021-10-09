
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

function solve(N::I,u::VI,v::VI,w::VI)
    adj::Vector{VPI} = [VPI() for _ in 1:N ]
    for i in 1:N-1
        push!(adj[u[i]],(v[i],w[i]))
        push!(adj[v[i]],(u[i],w[i]))
    end
    dist::VI = fill(0,N)
    q::VPI = [(1,-1)]
    while !isempty(q)
        (n::I,p::I) = popfirst!(q)
        for (c::I,ww::I) in adj[n]
            if c == p; continue; end
            dist[c] = dist[n] ⊻ ww
            push!(q,(c,n))
        end
    end
    ones::VI = fill(0,60)
    zeros::VI = fill(0,60)
    for i in 1:N
        mask::I = 1
        d::I = dist[i]
        for j in 1:60
            if d & mask != 0; ones[j] += 1; else zeros[j] += 1; end
            mask <<= 1
        end
    end
    ans::I = 0
    mm::I = 10^9+7
    for i in 1:60
        adder::I = (1 << (i-1)) % mm * ones[i] % mm * zeros[i] % mm
        ans = (ans + adder) % mm
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    u::VI = fill(0,N-1)
    v::VI = fill(0,N-1)
    w::VI = fill(0,N-1)
    for i in 1:N-1; u[i],v[i],w[i] = gis(); end
    ans = solve(N,u,v,w)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

