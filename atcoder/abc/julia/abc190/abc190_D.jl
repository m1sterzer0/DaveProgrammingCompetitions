
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

## a + (a+1) + (a+2) + ... + (a+n-1) = S
## n*(2a+n-1)) = 2S
## Thus n needs to be a factor of 2S
## We also need 2S/n+1-n to be even
## Also, n determines a, so this is just counting factors. 
function solve(N::I)
    n2::I = 2*N
    ans::I = 0
    factors::VI = []
    for i in 1:isqrt(n2); if n2 % i == 0; push!(factors,i); push!(factors,n2 ÷ i); end; end
    unique!(factors)
    for f in factors
        if (n2 ÷ f + 1 - f) & 1 == 0; ans += 1; end
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    ans = solve(N)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

