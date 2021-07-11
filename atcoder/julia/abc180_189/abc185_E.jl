
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

function solve(N::I,M::I,A::VI,B::VI)
    inf = 10^18
    dp::Array{I,2} = fill(inf,N,M)
    for i in N:-1:1
        for j in M:-1:1
            if i == N && j == M
                dp[i,j] = A[i] == B[j] ? 0 : 1
            elseif i == N
                ## Choices are delete j or match and delete the suffixsum
                opt1 = ((A[i] == B[j]) ? 0 : 1) + M - j
                opt2 = 1 + dp[i,j+1]
                dp[i,j] = min(opt1,opt2)
            elseif j == M
                ## Choices are delete i or match and delete the suffixsum
                opt1 = ((A[i] == B[j]) ? 0 : 1) + N - i
                opt2 = 1 + dp[i+1,j]
                dp[i,j] = min(opt1,opt2)
            else
                ## 3 choices -- keep both, delete i, or delete j
                opt1 = ((A[i] == B[j]) ? 0 : 1) + dp[i+1,j+1]
                opt2 = 1 + dp[i,j+1]
                opt3 = 1 + dp[i+1,j]
                dp[i,j] = min(opt1,opt2,opt3)
            end
        end
    end
    return dp[1,1]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = gis()
    B::VI = gis()
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

