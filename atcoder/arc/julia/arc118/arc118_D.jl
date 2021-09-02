
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


## We know group is cyclic
## Determine order of a and b -- if either are generators, then we are done.
## Now we focus on the group a^i*b^j.
## We populate P-1 terms in a grid and see if there are any repeats.
## ** If there are repeats, we are done
## ** Otherwise, we have a Order(a) x (P-1/order(a)) grid
## One of the sides of that grid must be even

function order(P::I,a::I)
    g = 1
    for i in 1:P-1
        g = g * a % P
        if g == 1; return i; end
    end
end

function doCyclic(P::I,a::I)::VS
    arr::VI = [1]
    for i in 2:P-1; push!(arr,arr[i-1] * a % P); end
    push!(arr,1)
    return ["Yes",join(arr," ")]
end

function solve(P::I,a::I,b::I)::VS
    x = order(P,a)
    y = order(P,b)
    if x == P-1; return doCyclic(P,a); end
    if y == P-1; return doCyclic(P,b); end
    gr::Array{I,2} = fill(0,x,(P-1)÷x)
    for i in 1:x; for j in 1:(P-1)÷x
        if i == 1 && j == 1; gr[i,j] = 1
        elseif j == 1; gr[i,j] = a * gr[i-1,j] % P
        else gr[i,j] = gr[i,j-1] * b % P
        end
    end; end
    n = length(unique(gr))
    if n != P-1; return ["No"]; end

    ## Construct the graph
    (m,n) = size(gr)
    if n & 1 != 0; gr = transpose(gr); (m,n) = size(gr); end
    ans::VI = []
    push!(ans,gr[1,1])
    for j in 1:n
        if j & 1 == 1
            for i in 2:m; push!(ans,gr[i,j]); end
        else
            for i in m:-1:2; push!(ans,gr[i,j]); end
        end
    end
    for j in n:-1:1; push!(ans,gr[1,j]); end
    return ["Yes",join(ans," ")]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    P,a,b = gis()
    ans = solve(P,a,b)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

