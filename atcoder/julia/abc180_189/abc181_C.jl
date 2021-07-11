
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

struct Pt2; x::I; y::I; end
Pt2() = Pt2(0,0)
Base.:+(a::Pt2,b::Pt2)::Pt2 =   Pt2(a.x+b.x,a.y+b.y)
Base.:-(a::Pt2,b::Pt2)::Pt2 =   Pt2(a.x-b.x,a.y-b.y)
Base.:*(a::Int64,b::Pt2)::Pt2 = Pt2(a*b.x,a*b.y)
Base.:*(b::Pt2,a::Int64)::Pt2 = Pt2(a*b.x,a*b.y) 
dot(a::Pt2,b::Pt2)::Int128 = Int128(a.x)*b.x + Int128(a.y)*b.y
cross(a::Pt2,b::Pt2)::Int128 = Int128(a.x)*b.y-Int128(a.y)*b.x

function solve(N::I,x::VI,y::VI)
    p::Vector{Pt2} = [Pt2(x[i],y[i]) for i in 1:N]
    for i in 1:N
        for j in i+1:N
            for k in j+1:N
                if cross(p[i]-p[j],p[k]-p[j]) == 0; return "Yes"; end

            end
        end
    end
    return "No"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    x::VI = fill(0,N)
    y::VI = fill(0,N)
    for i in 1:N; x[i],y[i] = gis(); end
    ans = solve(N,x,y)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

