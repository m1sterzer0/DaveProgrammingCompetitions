
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

function checkforbad(N::I,M::I,K::I,A::VI)
    run = 1
    largestrun = 1
    for i in 2:K
        if A[i] == A[i-1]+1; run += 1; largestrun = max(largestrun,run); else; run = 1; end
    end
    return largestrun >= M
end

function solve(N::I,M::I,K::I,A::VI)
    if K > 0 && checkforbad(N,M,K,A); return -1; end
    bad::VB = fill(false,N); for a in A; bad[a] = true; end
    dp1::VF = fill(0.00,N)
    dp2::VF = fill(0.00,N)
    csum::F = 0.00
    xsum::F = 0.00
    minv::F = 1.00/M
    for i in N-1:-1:0
        remsq::I = i+M+1; addsq::I = i+1
        if remsq < N; csum -= dp1[remsq]; xsum -= dp2[remsq]; end
        csum += dp1[addsq]; xsum += dp2[addsq]
        if i > 0
            dp1[i] = bad[i] ? 0.00 : 1.00 + minv * csum
            dp2[i] = bad[i] ? 1.00 : minv * xsum
        end
    end
    return (1.00 + minv * csum) / (1.00 - minv * xsum) 
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M,K = gis()
    A::VI = gis()
    ans = solve(N,M,K,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

