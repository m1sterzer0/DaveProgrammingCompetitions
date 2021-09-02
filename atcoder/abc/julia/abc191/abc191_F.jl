
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

function solve(N::I,A::VI)
    fmap::Dict{I,I} = Dict{I,I}()
    ub = minimum(A)
    for a in A
        for f in 1:isqrt(a)
            if f >= ub; break; end
            if a % f != 0; continue; end
            fmap[f] = haskey(fmap,f) ? gcd(fmap[f],a) : a
            f2::I = a ÷ f
            if f2 >= ub; continue; end
            fmap[f2] = haskey(fmap,f2) ? gcd(fmap[f2],a) : a
        end
    end
    ans::I = 1
    for (k,v) in fmap; if k == v; ans += 1; end; end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N::I = gi()
    A::VI = gis()
    ans = solve(N,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

