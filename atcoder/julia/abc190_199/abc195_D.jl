
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


mycmp(a::TI,b::TI)::Bool = a[1] > b[1] || a[1] == b[1] && a[2] < b[2]
function solve(N::I,M::I,Q::I,W::VI,V::VI,X::VI,L::VI,R::VI)::VI
    ans::VI = []
    packages::Vector{TI} = [(V[i],W[i],i) for i in 1:N]
    sort!(packages,lt=mycmp)
    boxes::VPI = []
    for q in 1:Q
        l,r = L[q],R[q]
        empty!(boxes)
        for i in 1:l-1; push!(boxes,(X[i],i)); end
        for i in r+1:M; push!(boxes,(X[i],i)); end
        sort!(boxes)
        usedboxes::VB = fill(false,M)
        lans::I = 0
        for (v,w,pidx) in packages
            for (x,bidx) in boxes
                if x < w || usedboxes[bidx]; continue; end
                usedboxes[bidx] = true
                lans += v
                break
            end
        end
        push!(ans,lans)
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M,Q = gis()
    W::VI = fill(0,N)
    V::VI = fill(0,N)
    for i in 1:N; W[i],V[i] = gis(); end
    X::VI = gis()
    L::VI = fill(0,Q)
    R::VI = fill(0,Q)
    for i in 1:Q; L[i],R[i] = gis(); end
    ans = solve(N,M,Q,W,V,X,L,R)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")


