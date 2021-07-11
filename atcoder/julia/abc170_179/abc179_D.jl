
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

struct segtree{T}; _n::I; sz::I; mylog::I; d::Vector{T}; op; e::T; end
function segtree{T}(n::I,op,e::T) where {T}
    sz::I = 1; mylog::I = 0; while sz < n; sz <<= 1; mylog += 1; end
    return segtree{T}(n,sz,mylog,fill(e,2*sz),op,e)
end
function segtree{T}(v::Vector{T},op,e::T) where {T}
    sz::I = 1; mylog::I = 0; while sz < length(v); sz <<= 1; mylog += 1; end
    d::Vector{T} = fill(e,2*sz)
    d[sz:sz+length(v)-1] .= v
    s::segtree{T} = segtree{T}(length(v),sz,mylog,d,op,e)
    for i in sz-1:-1:1; _segtreeUpdate(s,i); end
    return s
end
function segtreeSet(s::segtree{T},p::I,x::T) where {T}
    p += s.sz-1
    s.d[p] = x
    for i in 1:s.mylog; _segtreeUpdate(s,p >> i); end
end
function segtreeGet(s::segtree{T},p::I) where {T}
    p -= 1
    return s.d[s.sz+p]
end
## This version, unlike atcoder verison, returns op(a[l],a[l+1],...,a[r])
function segtreeProd(s::segtree{T},l::I,r::I) where {T}
    d::VI = s.d
    l += s.sz-1; r += s.sz-1+1 ## Add 1 for our version
    sml::T = s.e; smr::T = s.e
    while l < r
        @inbounds if (l & 1 == 1); sml = s.op(sml,d[l]); l += 1; end
        @inbounds if (r & 1 == 1); r -= 1; smr = s.op(d[r],smr); end
        l >>= 1; r >>= 1
    end
    return s.op(sml,smr)
end
function segtreeAllprod(s::segtree{T}) where {T}
    return s.d[1]
end
function _segtreeUpdate(s::segtree{T},k::I) where {T}
    if k > 0; s.d[k] = s.op(s.d[2*k],s.d[2*k+1]); end
end

@inline function segtreeop(a::I,b::I)::I
    return (a+b) % 998244353
end

################################################################################
## Begin Fenwick Tree (Adapted from datastructures.jl)
################################################################################

mutable struct FenwickTree
    bi_tree::Vector{Int64}
    n::Int64
    tot::Int64
end

FenwickTree(n::Int64) = FenwickTree(fill(0,n),n,0)
function FenwickTree(a::AbstractVector{Int64})
    n = length(a)
    tree = FenwickTree(n)
    for i in 1:n; inc!(tree,i,a[i]); end
    return tree
end
Base.length(ft::FenwickTree) = ft.n
Base.eltype(::Type{FenwickTree}) = Int64
function Base.empty!(ft::FenwickTree)
    fill!(ft.bi_tree,0)
    ft.tot = 0
end

function inc!(ft::FenwickTree, ind::Int64, val = 1)
    i = ind
    n = ft.n
    #@boundscheck 1 <= i <= n || throw(ArgumentError("$i should be in between 1 and $n"))
    #@inbounds while i <= n
    while i <= n
        ft.bi_tree[i] = (ft.bi_tree[i] + val) % 998244353
        i += i&(-i)
    end
    ft.tot = (ft.tot + val) % 998244353
end

function prefixsum(ft::FenwickTree, ind::Int64)
    if ind < 1; return 0; end
    sum = 0
    i = ind
    n = ft.n
    #@boundscheck 1 <= i <= n || throw(ArgumentError("$i should be in between 1 and $n"))
    #@inbounds while i > 0
    while i > 0
        sum = (sum + ft.bi_tree[i]) % 998244353
        i -= i&(-i)
    end
    return sum
end

rangesum(ft::FenwickTree, left::Int64, right::Int64) = (prefixsum(ft,right) - prefixsum(ft,left-1) + 998244353) % 998244353

suffixsum(ft::FenwickTree, ind::Int64) = ind > ft.n ? 0 : ft.tot - prefixsum(ft,ind-1)
Base.getindex(ft::FenwickTree, ind::Int64) = prefixsum(ft, ind)

################################################################################
## End Fenwick Tree (Adapted from datastructures.jl)
################################################################################

function solve(N::I,K::I,L::VI,R::VI)
    ft::FenwickTree = FenwickTree(N)
    inc!(ft,N,1)
    for i in N-1:-1:1
        v::I = 0
        for k::I in 1:K
            if i+L[k] > N; continue; end
            v = (v + rangesum(ft,i+L[k],min(N,i+R[k]))) % 998244353
        end
        if i == 1; return v; end
        inc!(ft,i,v)
    end
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin; _ = infile
    N,K = gis()
    L::VI = fill(0,K)
    R::VI = fill(0,K)
    for i in 1:K; L[i],R[i] = gis(); end
    ans = solve(N,K,L,R)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("junk.in")
#Profile.clear()
#@profilehtml main("junk.in")

