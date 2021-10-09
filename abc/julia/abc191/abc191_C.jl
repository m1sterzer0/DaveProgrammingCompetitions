
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

function solve(H::I,W::I,bd::Array{Char,2})
    numv::I = 0
    for i in 1:H-1
        for j in 1:W-1
            numb::I = 0
            for (ii,jj) in ((i,j),(i,j+1),(i+1,j),(i+1,j+1))
                if bd[ii,jj] == '#'; numb += 1; end
            end
            if numb == 1 || numb == 3; numv += 1; end
        end
    end
    return numv
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

