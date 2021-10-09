
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

function solve(N::I,x0::I,y0::I,xh::I,yh::I)::String
    xc::F = 0.5*(x0+xh)
    yc::F = 0.5*(y0+yh)
    xr::F = x0-xc
    yr::F = y0-yc
    ang::F = 2*π / N
    x1::F = cos(ang)*xr-sin(ang)*yr + xc
    y1::F = sin(ang)*xr+cos(ang)*yr + yc
    return "$x1 $y1"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    x0::I,y0::I = gis()
    xh::I,yh::I = gis()
    ans = solve(N,x0,y0,xh,yh)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

