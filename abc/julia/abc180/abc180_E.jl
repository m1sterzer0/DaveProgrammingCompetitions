
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

function solve(N::I,X::VI,Y::VI,Z::VI)
    inf::I = 10^18
    dp::Array{I,2} = fill(inf,2^N,N)
    dp[1+1,1] = 0
    for m in 2:2^N-1
        for i in 1:N
            if m & (1 << (i-1)) == 0; continue; end
            s = m ⊻ (1 << (i-1))
            for j in 1:N
                if j == i; continue; end
                if s & (1 << (j-1)) == 0; continue; end
                dp[m+1,i] = min(dp[m+1,i], dp[s+1,j] + abs(X[i]-X[j]) + abs(Y[i]-Y[j]) + max(0,Z[i]-Z[j]))
            end
        end
    end
    return minimum(dp[2^N,x] + abs(X[1]-X[x]) + abs(Y[1]-Y[x]) + max(0,Z[1]-Z[x]) for x in 2:N)
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    X::VI = fill(0,N)
    Y::VI = fill(0,N)
    Z::VI = fill(0,N)
    for i in 1:N; X[i],Y[i],Z[i] = gis(); end
    ans = solve(N,X,Y,Z)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

