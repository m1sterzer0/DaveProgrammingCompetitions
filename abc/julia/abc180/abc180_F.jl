
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

function solve(N,M,L)
    ## Number of edges must be between N and N minus number of partitions
    ## If a connected part has as many edges as nodes, then it is a ring, and there are only (k-1)!/2 ways to make that ring
    ## If a connected part has one fewer edge than nodes, then there are k!/2 ways to make the graph
    ## I think we can do this with a dp
    mm = 10^9+7
    chain::VI = fill(0,310)
    chain[3] = 3
    for k in 4:310; chain[k] = k * chain[k-1] % mm; end
    ring::VI = fill(0,310)
    ring[3] = 1
    for k in 4:310; ring[k] = (k-1) * ring[k-1] % mm; end
    imod::VI = fill(0,310)
    imod[1] = 1
    for i in 2:310; imod[i] = invmod(i,mm); end
    comb::Array{I,2} = fill(0,310,310)
    for i in 1:310
        comb[i,1] = i
        for j in 2:i; comb[i,j] = comb[i,j-1] * (i-j+1) % mm * imod[j] % mm; end
    end
    if L == 1; return solveit(N,M,L,chain,ring,comb) end
    a1::I = solveit(N,M,L-1,chain,ring,comb)
    a2::I = solveit(N,M,L,chain,ring,comb)
    return (a2 + mm - a1) % mm
end

function solveit(N::I,M::I,L::I,chain::VI,ring::VI,comb::Array{I,2})::I
    mm::I = 10^9+7
    dp::Array{I,2} = fill(0,N+1,M+1)
    dp[0+1,0+1] = 1
    dp[1+1,0+1] = 1
    for n::I in 2:N
        for psize::I in 1:min(n,L)
            n2::I = n - psize
            if psize == 1
                carr = ((0,1),)
            elseif psize == 2
                carr = ((1,n-1),(2,n-1))
            else
                x::I = n == psize ? 1 : comb[n-1,psize-1]  ## Picking a set that contains the largest number
                carr = ((psize-1,x*chain[psize]%mm),(psize,x*ring[psize]%mm))
            end
            for (m::I,ncomb::I) in carr
                if m > M; continue; end
                if n2 == 0; 
                    dp[n+1,m+1] = (dp[n+1,m+1] + ncomb) % mm
                    continue
                end
                for m2 in 0:min(n2,M-m)
                    dp[n+1,m+m2+1] = (dp[n+1,m+m2+1] + (ncomb * dp[n2+1,m2+1] % mm)) % mm
                end
            end
        end
    end
    return dp[N+1,M+1]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M,L = gis()
    ans = solve(N,M,L)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

