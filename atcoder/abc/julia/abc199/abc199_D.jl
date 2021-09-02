
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

function colornodes(edgelist::VPI,colors::VI,idx::I)::I
    for cnt::I in idx:length(edgelist)
        (i::I,j::I) = edgelist[cnt]
        if colors[i] > 0 && colors[j] > 0
            if colors[i] == colors[j]; return 0; end
        elseif cnt == length(edgelist)
            return 2
        else
            if colors[j] == 0; (i,j) = (j,i); end
            cchoice::PI = colors[j] == 1 ? (2,3) : colors[j] == 2 ? (1,3) : (1,2)
            ans::I = 0
            colors[i] = cchoice[1]
            ans += colornodes(edgelist,colors,idx+1)
            colors[i] = cchoice[2]
            ans += colornodes(edgelist,colors,idx+1)
            colors[i] = 0
            return ans
        end
    end
    return 1
end

## This graph is connected
function solve2(N::I,adj::VVI,start::I)
    visited::VB = fill(false,N)
    q::VI = [start]; visited[start] = true
    edgeset::SPI = SPI()
    edgelist::VPI = []
    while !isempty(q)
        n = popfirst!(q)
        for c in adj[n]
            if (n,c) ∉ edgeset && (c,n) ∉ edgeset
                push!(edgeset,(n,c))
                push!(edgelist,(n,c))
            end
            if visited[c]; continue; end
            visited[c] = true
            push!(q,c)
        end
    end
    if length(edgelist) == 0; return (3,visited); end
    colors = fill(0,N)
    colors[start] = 1
    ans::I = 3 * colornodes(edgelist,colors,1)
    return (ans,visited)
end

function solve(N::I,M::I,A::VI,B::VI)
    adj::VVI = [VI() for i in 1:N]
    for i in 1:M
        push!(adj[A[i]],B[i])
        push!(adj[B[i]],A[i])
    end
    ans::I = 1
    visited::VB = fill(false,N)
    for i in 1:N
        if visited[i]; continue; end
        (lans,lvisited) = solve2(N,adj,i)
        if lans == 0; return 0; end
        ans *= lans
        for j in 1:N; visited[j] |= lvisited[j]; end
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    for i in 1:M
        A[i],B[i] = gis()
    end
    ans = solve(N,M,A,B)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

