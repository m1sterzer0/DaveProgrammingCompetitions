
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

################################################################
## BEGIN Dinic's Max Flow
################################################################

mutable struct FlowEdge; to::I; cap::I; end
struct Flow; n::I; e::Vector{FlowEdge}; g::VVI; cur::VI; h::VI; end
Flow(n::I) = Flow(n,Vector{FlowEdge}(),[VI() for i in 1:n],fill(0,n),fill(0,n))
function addEdge(f::Flow,u::I,v::I,c::I,crev::I=0)
    push!(f.e,FlowEdge(v,c));    push!(f.g[u],length(f.e))
    push!(f.e,FlowEdge(u,crev)); push!(f.g[v],length(f.e))
end
## Do a good job at limiting to the level graph that ends at t
function bfs(fl::Flow,s::I,t::I)::Bool
    h::VI = fl.h; g::VVI = fl.g; e::Vector{FlowEdge} = fl.e
    fill!(h,-1)
    q::VI = [s]; h[s] = 0
    found::Bool = false
    nodelist::VI = [s]
    while !isempty(q)
        u::I = popfirst!(q)
        for i::I in g[u]
            v::I = e[i].to; c::I = e[i].cap
            if (c > 0 && h[v] == -1)
                push!(nodelist,v)
                h[v] = h[u] + 1
                #print("DBG: Found u:$u->v:$v  level:$(h[u])->$(h[v])\n")
                if v == t; found = true; break; end
                push!(q,v)
            end
        end
        if found; break; end
    end
    if !found; return false; end
    ## Now scrub the level tree to prevent the DFS from going haywire
    h2::VI = fill(-1,fl.n); h2[t] = h[t]; pop!(nodelist)
    while !isempty(nodelist)
        u::I = pop!(nodelist)
        for i::I in g[u]
            v::I = e[i].to; c::I = e[i].cap
            if (c > 0 && h2[v] != -1)
                if h2[u] == -1; h2[u] = h2[v] - 1; else; h2[u] = min(h2[u],h2[v]-1); end
            end
        end
        if h2[u] != h[u]; h[u] = -1; h2[u] = -1; end
        #print("DBG: Pass2: Level for $u is $(h[u])\n")
    end
    return true
end
function dfs(fl::Flow,u::I,t::I,f::I)::I
    if u == t; return f; end
    r::I = f
    h::VI = fl.h; g::VVI = fl.g; e::Vector{FlowEdge} = fl.e
    for i in fl.cur[u]:length(g[u])
        j::I = g[u][i]
        v::I = e[j].to; c::I = e[j].cap
        if (c > 0 && h[v] == h[u] + 1)
            a::I = dfs(fl,v,t,min(r,c))
            e[j].cap -= a
            e[j & 1 == 0 ? j-1 : j+1].cap += a
            r -= a
            if r == 0; return f; end
        end
    end
    return f-r
end
function maxFlow(fl::Flow,s::I,t::I)::I
    ans::I = 0
    inf::I = 10^18
    while bfs(fl,s,t)
        fill!(fl.cur,1)
        ans += dfs(fl,s,t,inf)
        #print("DBG: h[t]:$(fl.h[t])\n ans:$ans\n")
    end
    return ans
end

function gencase(Nmin::I,Nmax::I)
    N = rand(Nmin:Nmax)
    Bch,Wch,Qch = rand(),rand(),rand()
    sch = Bch+Wch+Qch
    Bch /= sch; Wch /= sch; Qch /= sch
    bd::Array{Char,2} = fill('.',N,N)
    for i in 1:N; for j in 1:N
        x = rand()
        bd[i,j] = x < Bch ? 'B' : x < (Bch+Wch) ? 'W' : '?'
    end; end
    return (N,bd)
end

function test(ntc::I,Nmin::I,Nmax::I)
    for ttt in 1:ntc
        (N,bd) = gencase(Nmin,Nmax)
        ans = solve(N,bd)
        print("Case #$ttt: $ans\n")
    end
end

## We can view every node as part of the checkerboard with black in the upper lefthand corner
## or part of the checkerboard with white in the upper lefthand corner
## Edges between nodes from two different checkerboard represent places of color adjacency -- these edges we wish to minimze.
## We simply create a graph, and create sources from fixed nodes from part of 1st checkerboard, and sinks for fixed nodes
## allocated to the 2nd checkerboard, and then run maxflow (== to min cut)
function solve(N::I,bd::Array{Char,2})::I
    s::I = N^2+1; t::I = N^2+2
    inf::I = 10^18
    fl::Flow = Flow(N*N+2)
    for i::I in 1:N
        for j::I in 1:N
            n1::I = N*(i-1)+j
            for (di::I,dj::I) in [(1,0),(0,1)]
                if 1 <= i+di <= N && 1 <= j+dj <= N
                    n2::I = N*(i+di-1)+(j+dj)
                    addEdge(fl,n1,n2,1,1)
                end
            end
            if bd[i,j] == 'B' && (i+j) & 1 == 0 || bd[i,j] == 'W' && (i+j) & 1 == 1
                addEdge(fl,s,n1,inf,0)
            elseif bd[i,j] == 'B' && (i+j) & 1 == 1 || bd[i,j] == 'W' && (i+j) & 1 == 0
                addEdge(fl,n1,t,inf,0)
            end
        end
    end
    #for (n1,n2,w) in mfedges; println("DBG: $n1 $n2 $w"); end
    mf::I = maxFlow(fl,s,t)
    return 2N*(N-1)-mf
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    bd::Array{Char,2} = fill('.',N,N)
    for i in 1:N; bd[i,:] = [x for x in gs()]; end
    ans = solve(N,bd)
    println(ans)
end

Random.seed!(8675309)
main()
#test(10,2,10)
#test(10,10,20)
#test(10,50,60)
#test(10,90,100)

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("junk4.in")
#Profile.clear()
#@profilehtml main("junk4.in")

