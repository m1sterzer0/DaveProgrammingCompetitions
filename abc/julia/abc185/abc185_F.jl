
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

struct segTree1; xor::VI; lazy::VI; l::I; r::I; end

function segTree1(l::I,r::I)
    n = 1; while n < r-l+1; n *= 2; end; n *= 2
    xor::VI = fill(0,n)
    lazy::VI = fill(0,n)
    return segTree1(xor,lazy,l,r)
end

function _processLazy(st::segTree1,idx::I,l::I,r::I)
    if r-l % 2 == 0; st.xor[idx] ⊻= st.lazy[idx]; end
    if l != r
        st.lazy[2idx] ⊻= st.lazy[idx]
        st.lazy[2idx+1] ⊻= st.lazy[idx]
    end
    st.lazy[idx] = 0
end

function _xor(st::segTree1,idx::I,l::I,r::I,ll::I,rr::I,v::I)
    #print("DBG: enter _xor idx:$idx l:$l r:$r ll:$ll rr:$rr v:$v\n")
    if st.lazy[idx] != 0; _processLazy(st,idx,l,r); end
    if rr < l || r < ll; return; end
    if ll <= l && r <= rr
        if (r-l) % 2 == 0; st.xor[idx] ⊻= v; end
        if l != r; st.lazy[2idx] ⊻= v; st.lazy[2idx+1] ⊻= v; end
        return
    end
    m::I = (l+r)>>1
    _xor(st,2idx,l,m,ll,rr,v)
    _xor(st,2idx+1,m+1,r,ll,rr,v)
    st.xor[idx] = st.xor[2idx] ⊻ st.xor[2idx+1]
end

function _xorq(st::segTree1,idx::I,l::I,r::I,ll::I,rr::I)
    if st.lazy[idx] != 0; _processLazy(st,idx,l,r); end
    if rr < l || r < ll; return 0; end
    if ll <= l && r <= rr; return st.xor[idx]; end
    m::I = (l+r)>>1
    v1::I = _xorq(st,2idx,l,m,ll,rr)
    v2::I = _xorq(st,2idx+1,m+1,r,ll,rr)
    return v1 ⊻ v2
end

xor(st::segTree1,l::I,r::I,v::I) = _xor(st,1,st.l,st.r,l,r,v)
xorq(st::segTree1,l::I,r::I)     = _xorq(st,1,st.l,st.r,l,r)

function solve(N::I,Q::I,A::VI,T::VI,X::VI,Y::VI)
    ans::VI = []
    st = segTree1(1,N)
    for i in 1:N; 
        xor(st,i,i,A[i])
    end
    for i in 1:Q
        if T[i] == 1
            xor(st,X[i],X[i],Y[i])
        else
            push!(ans,xorq(st,X[i],Y[i]))
        end
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,Q = gis()
    A::VI = gis()
    T::VI = fill(0,Q)
    X::VI = fill(0,Q)
    Y::VI = fill(0,Q)
    for i in 1:Q; T[i],X[i],Y[i] = gis(); end
    ans = solve(N,Q,A,T,X,Y)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

