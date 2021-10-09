
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


function crt2(r::VI,m::VI)::PI
    n = length(r)
    r1::I = mod(r[1],m[1]); m1::I = m[1]
    for i in 2:n
        r2::I = mod(r[i],m[i]); m2::I = m[i]
        (g::I,p::I,q::I) = gcdx(m1,m2)
        if (r2-r1) % g != 0; return (0,-1); end
        mm::I = m1 * (m2 ÷ g)
        tmp::I = (r2-r1)÷g * p % (m2 ÷ g)
        rr::I = mod(r1 + m1*tmp, mm)
        r1,m1 = rr,mm
    end
    return (r1,m1)
end

function solve(T::I,X::VI,Y::VI,P::VI,Q::VI)::VS
    ans::VS = []
    for i in 1:T
        x,y,p,q = X[i],Y[i],P[i],Q[i]
        inf::I = 8*10^18
        best::I = inf
        a::I = 2x+2y
        b::I = p+q
        m::VI = [a,b]
        r::VI = [0,0]
        for i in 0:y-1
            for j in 0:q-1
                r[1] = x + i
                r[2] = p + j
                (res,checklcm) = crt2(r,m)
                if checklcm > 0; best = min(best,res); end
            end
        end
        push!(ans,best==inf ? "infinity" : "$best")
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    T = gi()
    X::VI = fill(0,T)
    Y::VI = fill(0,T)
    P::VI = fill(0,T)
    Q::VI = fill(0,T)
    for i in 1:T; X[i],Y[i],P[i],Q[i] = gis(); end
    ans = solve(T,X,Y,P,Q)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

