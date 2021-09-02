
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

function solve(N::I,M::I,A::VI,B::VI,K::I,C::VI)::I
    adj::VVI = [VI() for i in 1:N]
    for i in 1:M
        push!(adj[A[i]],B[i]); push!(adj[B[i]],A[i])
    end
    inf::I = 10^18
    dist::VI = fill(inf,N)
    q::VI= []
    csb::Array{I,2} = fill(-1,K,K)
    for i in 1:K
        fill!(dist,inf)
        s::I = C[i]
        dist[s] = 0; push!(q,s)
        while !isempty(q)
            n::I = popfirst!(q)
            d::I = dist[n]
            for n2 in adj[n]
                if dist[n2] < inf; continue; end
                dist[n2] = d + 1
                push!(q,n2)
            end
        end
        for j in 1:K
            if dist[C[j]] == inf; return -1; end
            csb[i,j] = dist[C[j]]
        end
    end
    ## Now for the bitmask dp.  Based on subset and last element
    dp::Array{I,2} = fill(inf,2^K,K)
    for mask in 0:2^K-1
        for last in 1:K
            if mask & (1<<(last-1)) == 0; continue; end
            childmask::I = mask ⊻ (1<<(last-1))
            if childmask == 0; dp[mask+1,last] = 1; continue; end
            for l2 in 1:K
                if childmask & (1<<(l2-1)) == 0; continue; end
                dp[mask+1,last] = min(dp[mask+1,last],dp[childmask+1,l2]+csb[l2,last])
            end
        end
    end
    ## Final answer prep
    ans::I = inf
    for i in 1:K; ans = min(ans,dp[2^K,i]); end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    for i in 1:M; A[i],B[i] = gis(); end
    K::I = gi()
    C::VI = gis()
    ans = solve(N,M,A,B,K,C)
    println(ans)
end

function main2(niter::I,infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    for i in 1:M; A[i],B[i] = gis(); end
    K::I = gi()
    C::VI = gis()
    for i in 1:niter
        ans = solve(N,M,A,B,K,C)
        println(ans)
    end
end


Random.seed!(8675309)
#main()

#main2(1,"junk.in")
#main2(20,"junk.in")

using Profile, StatProfilerHTML
Profile.clear()
@profile main2(1,"junk.in")
Profile.clear()
@profilehtml main2(20,"junk.in")

