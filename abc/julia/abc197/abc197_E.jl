
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

function solve(N::I,C::VI,X::VI)::I
    cc::VI   = fill(0,N) 
    cmin::VI = fill(0,N)
    cmax::VI = fill(0,N)
    for i in 1:N
        x,c = X[i],C[i]
        cc[c] += 1
        if cc[c] == 1; cmin[c] = cmax[c] = x; end
        if x < cmin[c]; cmin[c] = x; end
        if x > cmax[c]; cmax[c] = x; end
    end
    leftcost::I = 0; leftpos::I = 0; rightcost::I = 0; rightpos::I = 0
    for i in 1:N
        if cc[i] == 0; continue; end
        leftcost1::I = leftcost  + abs(leftpos-cmax[i]) + abs(cmax[i]-cmin[i])
        leftcost2::I = rightcost + abs(rightpos-cmax[i]) + abs(cmax[i]-cmin[i])
        rightcost1::I = leftcost + abs(leftpos-cmin[i]) + abs(cmax[i]-cmin[i])
        rightcost2::I = rightcost + abs(rightpos-cmin[i]) + abs(cmax[i]-cmin[i])
        leftcost,rightcost = min(leftcost1,leftcost2),min(rightcost1,rightcost2)
        leftpos,rightpos = cmin[i],cmax[i]
    end 
    return min(leftcost+abs(leftpos),rightcost+abs(rightpos))
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    C::VI = fill(0,N)
    X::VI = fill(0,N)
    for i in 1:N; X[i],C[i] = gis(); end
    ans = solve(N,C,X)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

