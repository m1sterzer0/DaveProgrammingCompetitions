
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

function solve(N::I,M::I,X::VI,Y::VI,Z::VI)::I
    ## sort the constraints by x
    constr::Vector{TI} = [(X[i],Y[i],Z[i]) for i in 1:M]
    sort!(constr)
    ## find the starting constraint index for each number
    sidx::VI = fill(0,N)
    for (i,(x,y,z)) in enumerate(constr)
        if sidx[x] == 0; sidx[x] = i; end
    end
    ## Now set up the DP
    dp::VI = fill(0,2^N-1)
    for i in 1:2^N-1
        szset::I = count_ones(i)
        if sidx[szset] != 0
            good = true
            for ptr in sidx[szset]:M
                (x,y,z) = constr[ptr]
                if x > szset; break; end
                rmask::I = 2^y-1
                if count_ones(i & rmask) > z; good = false; break; end
            end
            if !good; continue; end
        end
        if szset == 1; dp[i] = 1; continue; end
        ## Iterate on last value
        for j in 1:18
            mask::I = 1<<(j-1)
            if i & mask == 0; continue; end
            dp[i] += dp[i ⊻ mask]
        end
    end
    return dp[2^N-1]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    X::VI = fill(0,M)
    Y::VI = fill(0,M)
    Z::VI = fill(0,M)
    for i in 1:M; X[i],Y[i],Z[i] = gis(); end
    ans = solve(N,M,X,Y,Z)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

