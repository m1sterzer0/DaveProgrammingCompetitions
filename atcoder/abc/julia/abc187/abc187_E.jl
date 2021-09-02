
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

function solve(N::I,a::VI,b::VI,Q::I,t::VI,e::VI,x::VI)::VI
    adj::VVI = [VI() for i in 1:N]
    for i in 1:N-1; push!(adj[a[i]],b[i]); push!(adj[b[i]],a[i]); end

    par::VI = fill(0,N)
    q::VPI = [(1,0)]
    while !isempty(q)
        (n,p) = popfirst!(q)
        par[n] = p
        for c in adj[n]
            if c == p; continue; end
            push!(q,(c,n))
        end
    end

    glb::I = 0
    nadd::VI = fill(0,N)
    for i in 1:Q
        (aa,bb) = t[i] == 1 ? (a[e[i]],b[e[i]]) : (b[e[i]],a[e[i]])
        if bb == par[aa]
            nadd[aa] += x[i]
        else
            glb += x[i]
            nadd[bb] -= x[i]
        end
    end

    ans::VI = fill(0,N)
    q2::Vector{TI} = [(1,0,glb+nadd[1])]
    while !isempty(q2)
        (n,p,v) = popfirst!(q2)
        ans[n] = v
        for c in adj[n]
            if c == p; continue; end
            push!(q2,(c,n,v+nadd[c]))
        end
    end

    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    a::VI = fill(0,N-1)
    b::VI = fill(0,N-1)
    for i in 1:N-1; a[i],b[i] = gis(); end
    Q = gi()
    t::VI = fill(0,Q)
    e::VI = fill(0,Q)
    x::VI = fill(0,Q)
    for i in 1:Q; t[i],e[i],x[i] = gis(); end
    ans = solve(N,a,b,Q,t,e,x)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

