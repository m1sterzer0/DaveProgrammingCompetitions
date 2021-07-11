
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

function solve(H::I,W::I,bd::Array{Char,2})::I
    mm = 10^9+7
    dp::Array{I,2} = fill(0,H,W)
    sdown::Array{I,2} = fill(0,H,W)
    sright::Array{I,2} = fill(0,H,W)
    sdiag::Array{I,2} = fill(0,H,W)
    for i in H:-1:1
        for j in W:-1:1
            if (i,j) == (H,W); dp[i,j] = 1; sdown[i,j] = 1; sright[i,j] = 1; sdiag[i,j] = 1; continue; end
            if bd[i,j] == '#'; continue; end
            if i < H; dp[i,j] = (dp[i,j] + sdown[i+1,j]) % mm; end
            if j < W; dp[i,j] = (dp[i,j] + sright[i,j+1]) % mm; end
            if i < H && j < W; dp[i,j] = (dp[i,j] + sdiag[i+1,j+1]) % mm; end
            sdown[i,j]  = (dp[i,j] + (i == H ? 0 : sdown[i+1,j])) % mm
            sright[i,j] = (dp[i,j] + (j == W ? 0 : sright[i,j+1])) % mm
            sdiag[i,j]  = (dp[i,j] + (i == H || j == W ? 0 : sdiag[i+1,j+1])) % mm
        end
    end
    return dp[1,1]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    H,W = gis()
    bd::Array{Char,2} = fill('.',H,W)
    for i in 1:H; bd[i,:] = [x for x in gs()]; end
    ans = solve(H,W,bd)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

