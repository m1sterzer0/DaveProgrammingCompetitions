
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
    ## Use bitmask to encode connectivity to prevent a factor of N on the check
    G::VI = fill(0,N)
    for i in 1:M
        G[A[i]] |= 1 << (B[i]-1)
        G[B[i]] |= 1 << (A[i]-1)
    end
    dp::VI = fill(N,2^N)
    dp[1] = 1

    ## Now we do the clique looping and set the dp value to 1
    for m in 0:2^N-1
        if dp[m+1] != 1; continue; end
        for n in 1:N
            if m | n == m; continue; end
            if G[n] & m != m; continue; end
            newm = m | (1<<(n-1)); dp[newm+1] = 1
        end
    end

    ## BITMASK TRICK: Looping through subsets trick -- subtract and then and
    for m::I in 1:2^N-1
        if dp[m+1] == 1; continue; end
        m2::I = m; best::I = N
        while m2 != 0
            best = min(best,dp[m2+1] + dp[m & ~m2 + 1])
            m2 = (m2-1)&m
        end
        dp[m+1] = best
    end
    return dp[2^N]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    for i in 1:M; A[i],B[i] = gis(); end
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

