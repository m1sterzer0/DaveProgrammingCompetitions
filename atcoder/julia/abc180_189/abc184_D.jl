
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

function solve(A::I,B::I,C::I)
    dp::Array{F,3} = fill(0.00,101,101,101)
    for a in 100:-1:0
        for b in 100:-1:0
            for c in 100:-1:0
                if a == 100 || b == 100 || c == 100; dp[a+1,b+1,c+1] = 0.00; continue; end
                s = a+b+c
                if s == 0; dp[a+1,b+1,c+1] = 10.0^99; continue; end
                dp[a+1,b+1,c+1] = 1.000 + a/s*dp[a+2,b+1,c+1] + b/s*dp[a+1,b+2,c+1] + c/s*dp[a+1,b+1,c+2]
            end
        end
    end
    return dp[A+1,B+1,C+1]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    A,B,C = gis()
    ans = solve(A,B,C)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

