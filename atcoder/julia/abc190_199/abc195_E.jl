
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

function solve(N::I,S::String,X::String)::String
    dp::Array{Char,2} = fill('.',N,7)
    o1::VI = [(x - '0') for x in S]
    pvs::VI = []
    pv::I = 1
    for i in 1:N; push!(pvs,pv); pv = pv * 10 % 7; end
    ## Do the final column
    for n in N:-1:1
        for i in 0:6
            opt1::I = (10*i + o1[n]) % 7
            opt2::I = (10*i + 0)     % 7
            if n == N
                if X[N] == 'T' && (opt1 == 0 || opt2 == 0); dp[N,i+1] = 'T'
                elseif X[N] == 'A' && (opt1 == 0 && opt2 == 0); dp[N,i+1] = 'T'
                else; dp[N,i+1] = 'A'
                end
            else
                if dp[n+1,opt1+1] == X[n] || dp[n+1,opt2+1] == X[n]; dp[n,i+1] = X[n]
                else; dp[n,i+1] = X[n] == 'A' ? 'T' : 'A'
                end
            end
        end
    end
    return dp[1,1] == 'T' ? "Takahashi" : "Aoki"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    S = gs()
    X = gs()
    ans = solve(N,S,X)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

