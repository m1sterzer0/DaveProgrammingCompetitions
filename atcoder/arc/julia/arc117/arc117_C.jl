
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

function solve(N::I,C::String)
    ## Let B == 0, R == 1, W == 2
    ## a b c  a+b+c  
    ## 0 0 0    0 
    ## 1 1 1    0
    ## 2 2 2    0
    ## 0 1 2    0
    ## 0 2 1    0
    ## 1 2 0    0
    ## Observation 1, a+b+c sum to 0, so c = -a-b mod 3
    ## Observation 2, -- like pascals triangle, the number of copies of
    ## a square at the top of the pyramid is related to the number of paths
    ## to the top, which is a binomial coefficient.
    num3s::I = 0; coeff::I = 1; ans::I = 0
    for i in 1:N
        num = i == 1 ? 1 : N+1-i
        den = i == 1 ? 1 : i-1
        while num%3 == 0; num3s += 1; num ÷= 3; end
        while den%3 == 0; num3s -= 1; den ÷= 3; end
        coeff = (coeff * (num % 3)) % 3 
        coeff = (coeff * (den % 3)) % 3 ## numbers are their own inverse in mod 3
        v::I = C[i] == 'B' ? 0 : C[i] == 'R' ? 1 : 2
        if N % 2 == 0; v = mod(-v,3); end
        if num3s == 0; ans = (ans + v * coeff) % 3; end
    end
    return ans == 0 ? "B" : ans == 1 ? "R" : "W"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    C = gs()
    ans = solve(N,C)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

