
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

######################################################################################################
### BEGIN DIJKSTRA CODE (ADJ LIST)
######################################################################################################
struct MinHeapDijkstraNode; n::I; t::I; end ## n is nodeID, t is time (or dist)
Base.isless(a::MinHeapDijkstraNode,b::MinHeapDijkstraNode) = a.t < b.t

struct MinHeapDijkstra
    valtree::Vector{MinHeapDijkstraNode}; pos::VI
    MinHeapDijkstra(N::I) = new([],fill(0,N))
end

function _swap(h::MinHeapDijkstra,i::I,j::I)
    (n1::I,n2::I) = (h.valtree[i].n,h.valtree[j].n)
    h.pos[n2],h.pos[n1] = i,j
    h.valtree[i],h.valtree[j] = h.valtree[j],h.valtree[i]
end

function _bubbleUp(h::MinHeapDijkstra, i::I)
    if i == 1; return; end
    j::I = i >> 1; if h.valtree[i] < h.valtree[j]; _swap(h,i,j); _bubbleUp(h,j); end
end

function _bubbleDown(h::MinHeapDijkstra,i::I)
    len::I = length(h.valtree); l::I = i << 1; r::I = l + 1
    res1::Bool = l > len || !(h.valtree[i] > h.valtree[l])
    res2::Bool = r > len || !(h.valtree[i] > h.valtree[r])
    if res1 && res2; return
    elseif res2 || !res1 && !(h.valtree[l] > h.valtree[r]); _swap(h,i,l); _bubbleDown(h,l)
    else; _swap(h,i,r); _bubbleDown(h,r)
    end
end

function Base.push!(h::MinHeapDijkstra,node::MinHeapDijkstraNode)
    n::I = node.n; idx::I = h.pos[n]
    if idx == 0; push!(h.valtree,node); idx = length(h.valtree); h.pos[n] = idx
    elseif h.valtree[idx] > node; h.valtree[idx] = node
    end
    _bubbleUp(h,idx)
end

function Base.pop!(h::MinHeapDijkstra)
    ans::MinHeapDijkstraNode = h.valtree[1]; h.pos[ans.n] = 0
    node2::MinHeapDijkstraNode = pop!(h.valtree)
    if length(h.valtree) >= 1; h.pos[node2.n] = 1; h.valtree[1] = node2; _bubbleDown(h,1); end
    return ans
end

Base.isempty(h::MinHeapDijkstra) = isempty(h.valtree)

######################################################################################################
### END DIJKSTRA CODE (ADJ LIST)
######################################################################################################

function solve(N::I,M::I,X::I,Y::I,A::VI,B::VI,T::VI,K::VI)
    mh::MinHeapDijkstra = MinHeapDijkstra(N)
    push!(mh,MinHeapDijkstraNode(X,0))
    adj::Vector{VPI} = [VPI() for i in 1:N]
    for i in 1:M; push!(adj[A[i]],(B[i],i)); push!(adj[B[i]],(A[i],i)); end
    inf::I = 10^18
    dist::VI = fill(inf,N)
    while !isempty(mh)
        node::MinHeapDijkstraNode = pop!(mh)
        (n::I,t::I) = (node.n,node.t)
        dist[n] = t
        for (n2,i) in adj[n]
            if dist[n2] != inf; continue; end
            t2::I = (t + K[i] - 1) ÷ K[i] * K[i] + T[i]
            push!(mh,MinHeapDijkstraNode(n2,t2))
        end
    end
    return dist[Y] == inf ? -1 : dist[Y]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M,X,Y = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    T::VI = fill(0,M)
    K::VI = fill(0,M)
    for i in 1:M; A[i],B[i],T[i],K[i] = gis(); end
    ans = solve(N,M,X,Y,A,B,T,K)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

