
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

function numeval(X::String,b::I)
    pv::BigInt = 1
    ans::BigInt = 0
    for c in reverse(X)
        d::BigInt = c - '0'
        ans += d * pv; pv *= b
    end
    return ans
end

function solve(X::String,M::I)::I
    minbase::I = maximum([x for x in X]) - '0' + 1
    minval::BigInt = numeval(X,minbase)
    if minval > M; return 0; end
    if length(X) == 1; return 1; end
    ## Arithmetic sequence
    if length(X) == 2
        v1::I = Int64(minval)
        v2::BigInt = numeval(X,minbase+1)
        d::I = Int64(v2-minval)
        ## Need n such that v1 + (n-1)*d <= M --> (n-1)*d <= M - v1 --> (n-1) <= (M-v1) ÷ d --> n <= (M-v1) ÷ d + 1
        return (M-v1)÷d + 1
    end
    ub = 2*minbase; while numeval(X,ub) <= M; ub = 2ub; end
    lb = minbase
    while ub-lb > 1
        m = (ub+lb) >> 1
        if numeval(X,m) <= M; lb = m; else; ub = m; end
    end
    return lb - minbase + 1
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    X = gs()
    M = gi()
    ans = solve(X,M)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

