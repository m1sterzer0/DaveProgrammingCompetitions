
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
    l += s.sz-1; r += s.sz-1+1 ## Add 1 for our version
    sml::T = s.e; smr::T = s.e
    while l < r
        if (l & 1 == 1); sml = s.op(sml,s.d[l]); l += 1; end
        if (r & 1 == 1); r -= 1; smr = s.op(s.d[r],smr); end
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

function segtreeop(a::I,b::I)::I
    return a < b ? a : b
end

function solve(N::I,P::VI,A::VI,B::VI,C::VI)
    inf::I = 10^18
    ans::I = inf
    st = segtree{I}(N+1,segtreeop,inf)
    dp::VI = fill(inf,N)
    pos::VI = fill(0,N)
    for i in 1:N; pos[P[i]] = i; end
    Asum::VI = fill(0,N+1)
    Bsum::VI = fill(0,N+1)
    Csum::VI = fill(0,N+1)
    for i in 1:N; Asum[i+1] = Asum[i] + A[i]; end
    for i in 1:N; Bsum[i+1] = Bsum[i] + min(A[i],B[i]); end
    for i in 1:N; Csum[i+1] = Csum[i] + min(A[i],C[i]); end
    for i in 1:N
        dp[i] = min(Bsum[i],segtreeProd(st,1,pos[i])+Asum[i])
        ans = min(ans,dp[i]+Csum[N+1]-Csum[i+1])
        segtreeSet(st,pos[i],dp[i]-Asum[i+1])
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    P::VI = gis()
    A::VI = fill(0,N)
    B::VI = fill(0,N)
    C::VI = fill(0,N)
    for i in 1:N; A[i],B[i],C[i] = gis(); end
    ans = solve(N,P,A,B,C)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

