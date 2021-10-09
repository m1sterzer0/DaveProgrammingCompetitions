
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

function getans(sb::VPI,i::I,j::I,k::I)::VS
    ans::VS = ["Yes"]
    b::VI = [i]
    c::VI = []
    while j != 0 && sb[j][1] != 0; push!(b,sb[j][1]); j = sb[j][2]; end
    while k != 0 && sb[k][1] != 0; push!(c,sb[k][1]); k = sb[k][2]; end
    reverse!(b); reverse!(c)
    pushfirst!(b,length(b))
    pushfirst!(c,length(c))
    push!(ans,join(b," "))
    push!(ans,join(c," "))
    return ans
end

function solve(N::I,A::VI)::VS
    sb::VPI = fill((0,0),200)
    for i in 1:N
        if sb[A[i]%200+1][1] != 0
            return getans(sb,i,0,A[i]%200+1)
        end
        cursums::VI = [j for j in 0:199 if sb[j+1][1] > 0]
        newsums::VI = [(A[i]+s) % 200 for s in cursums]
        for (k,s) in enumerate(newsums)
            if sb[s+1][1] != 0; return getans(sb,i,cursums[k]+1,s+1); end
        end
        ## Now update with my sums
        for (k,s) in enumerate(newsums); sb[s+1] = (i,cursums[k]+1); end
        sb[A[i]%200+1] = (i,0)
    end
    return ["No"]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    A::VI = gis()
    ans = solve(N,A)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

