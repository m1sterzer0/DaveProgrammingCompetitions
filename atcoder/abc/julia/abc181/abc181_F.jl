
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

###########################################################################
## BEGIN UnionFindFast -- Only works with integers, but avoids dictionaries
###########################################################################
struct UnionFindFast; p::VI; sz::VI; n::I; UnionFindFast(n::I) = new(collect(1:n),fill(1,n),n); end
function findset(h::UnionFindFast,x::I)::I
    if h.p[x] == x; return x; end
    return h.p[x] = findset(h,h.p[x])
end
function getsize(h::UnionFindFast,x::I)::I
    a::I = findset(h,x); return h.sz[a]
end
function joinset(h::UnionFindFast,x::I,y::I)
    a::I = findset(h,x); b::I = findset(h,y)
    if a != b
        if h.sz[a] < h.sz[b]; (a,b) = (b,a); end
        h.p[b] = a
        h.sz[a] += h.sz[b]
    end
end
###############################################################
## END UnionFindFast
################################################################

function solve(N::I,x::VI,y::VI)
    uf::UnionFindFast = UnionFindFast(N+2)
    edges::Vector{TI} = []
    push!(edges,(200*200,N+1,N+2))
    for i in 1:N
        push!(edges,((y[i]-100)^2,i,N+1))
        push!(edges,((y[i]+100)^2,i,N+2))
    end
    for i in 1:N
        for j in i+1:N
            d2 = (x[i]-x[j])^2+(y[i]-y[j])^2
            push!(edges,(d2,i,j))
        end
    end
    sort!(edges)
    for (d2,i,j) in edges
        joinset(uf,i,j)
        if findset(uf,N+1) == findset(uf,N+2); return 0.5*sqrt(1.0*d2); end
    end
    return -1.0 ## Shouldn't get here
end


function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    x::VI = fill(0,N)
    y::VI = fill(0,N)
    for i in 1:N; x[i],y[i] = gis(); end
    ans = solve(N,x,y)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

