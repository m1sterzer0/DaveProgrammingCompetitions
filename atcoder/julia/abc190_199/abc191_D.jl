
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

function solve(X::F,Y::F,R::F)
    XX::I = round(Int64,10000*X) + 2*10^9 + 20000  ## Keep it positive
    YY::I = round(Int64,10000*Y) + 2*10^9 + 20000  ## Keep it positive
    RR::I = round(Int64,10000*R)
    ans = 0
    xmin = (XX - RR - 20000) ÷ 10000 * 10000
    xmax = (XX + RR + 20000) ÷ 10000 * 10000
    r2::I = RR*RR
    for x in xmin:10000:xmax
        x2 = (x-XX)^2
        if x2 > r2; continue; end
        residual = r2-x2
        dy = isqrt(residual)
        ymin = (YY - dy + 9999) ÷ 10000 * 10000
        ymax = (YY + dy) ÷ 10000 * 10000
        if ymin > ymax; continue; end
        ans += 1 + (ymax-ymin)÷10000
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    X,Y,R = gfs()
    ans = solve(X,Y,R)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

