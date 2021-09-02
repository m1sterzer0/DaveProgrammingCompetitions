
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

# S for Takahashi
## T for Aoiki
function solve(K::I,S::String,T::String)::F
    deck::VI = fill(K,9)
    t::VI = fill(0,9)
    a::VI = fill(0,9)
    for c in S[1:4]; cc::I = c - '0'; deck[cc] -= 1; t[cc] += 1; end
    for c in T[1:4]; cc = c - '0';    deck[cc] -= 1; a[cc] += 1; end
    winningways::I = 0
    for i in 1:9
        for j in 1:9
            ways = i == j ? deck[i] * (deck[i]-1) : deck[i] * deck[j]
            t[i] += 1; a[j] += 1
            if sum(i*10^t[i] for i in 1:9) > sum(i*10^a[i] for i in 1:9); winningways += ways; end
            t[i] -= 1; a[j] -= 1
        end
    end
    return winningways / (9K-8) / (9K-9)
end
   
function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    K = gi()
    S = gs()
    T = gs()
    ans = solve(K,S,T)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

