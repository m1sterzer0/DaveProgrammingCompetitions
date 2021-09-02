
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

function solve(N::I,Q::I,C::VI,QQ::Vector{TI})
    ans::VI = []
    uf = UnionFindFast(N)
    cl::Vector{Dict{I,I}} = [Dict{I,I}() for i in 1:N]
    for i in 1:N; cl[i][C[i]] = 1; end
    for (t,a,b) in QQ
        if t == 1
            pa = findset(uf,a); pb = findset(uf,b)
            if pa == pb; continue; end
            joinset(uf,a,b)
            npa = findset(uf,a)
            npb = findset(uf,b)
            (m,s) = pa == npa ? (pa,pb) : (pb,pa)
            for (c,v) in cl[s]
                if haskey(cl[m],c); cl[m][c] += v; else; cl[m][c] = v; end
            end
        else
            pa = findset(uf,a)
            push!(ans,haskey(cl[pa],b) ? cl[pa][b] : 0)
        end
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,Q = gis()
    C::VI = gis()
    QQ::Vector{TI} = []
    for i in 1:Q
        xx::VI = gis()
        push!(QQ,(xx[1],xx[2],xx[3]))
    end
    ans = solve(N,Q,C,QQ)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

