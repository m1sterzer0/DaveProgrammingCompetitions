
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

## Rotate coordinates
function solve(r1,c1,r2,c2)
    if r1==r2 && c1==c2; return 0; end
    if r1+c1 == r2+c2; return 1; end  ## One diagonal move
    if r1-c1 == r2-c2; return 1; end  ## One diagonal move
    if abs(r1-r2)+abs(c1-c2) <= 3; return 1; end  ## One local move
    if (r1+c1)%2 == (r2+c2)%2; return 2; end ## Two diagonal moves (bishops stay on color)
    if abs(abs(r1+c1)-abs(r2+c2)) <= 3; return 2; end ## local + diagonal
    if abs(abs(r1-c1)-abs(r2-c2)) <= 3; return 2; end ## local + diagonal
    if abs(r1-r2)+abs(c1-c2) <= 6; return 2; end  ## Two local moves
    return 3
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    r1,c1 = gis()
    r2,c2 = gis()
    ans = solve(r1,c1,r2,c2)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

