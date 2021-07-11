
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

## We have to do DFS, but we have to do it w/o recursion becaue of julia recusion limit (fun)
function solve(N::I,C::VI,A::VI,B::VI)
    adj::VVI = [VI() for i in 1:N]
    for i in 1:N-1; (a,b) = (A[i],B[i]); push!(adj[a],b); push!(adj[b],a); end
    colorcnt::VI = fill(0,10^5)
    good::VI = []
    st::Vector{TI} = [(1,0,1)]
    while !isempty(st)
        (n::I,p::I,mode::I) = pop!(st)
        if mode == 1
            push!(st,(n,p,2))
            if colorcnt[C[n]] == 0; push!(good,n); end
            colorcnt[C[n]] += 1
            for c in adj[n]
                if c == p; continue; end
                push!(st,(c,n,1))
            end
        else
            colorcnt[C[n]] -= 1
        end
    end
    sort!(good)
    return good
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    C::VI = gis()
    A::VI = fill(0,N-1)
    B::VI = fill(0,N-1)
    for i in 1:N-1; A[i],B[i] = gis(); end
    ans = solve(N,C,A,B)
    for i in ans; println(i); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

