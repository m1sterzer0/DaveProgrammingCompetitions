
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


function solveit(n::I,a::VI)::VI
    dp::VI = fill(0,2^n)
    for m in 1:2^n-1
        if count_ones(m) == 1
            idx = 1
            while (1<<(idx-1)) & m == 0; idx += 1; end
            dp[m+1] = a[idx]
        else
            s1 = (m-1)&m
            m2 = m ⊻ s1
            dp[m+1] = dp[s1+1] + dp[m2+1]
        end
    end
    return dp
end

function solve(N::I,T::I,A::VI)::I
    if N == 1; return A[1] <= T ? A[1] : 0; end
    lhalf = N ÷ 2
    v1 = solveit(lhalf,A[1:lhalf])
    v2 = solveit(N-lhalf,A[lhalf+1:N])
    best::I = 0
    sort!(v1)
    sort!(v2)
    l2 = length(v2)
    idx2 = 1;
    while idx2 < l2 && v2[idx2+1] <= T; idx2 += 1; end
    for idx1 in 1:length(v1)
        if v1[idx1] > T; break; end
        while v1[idx1] + v2[idx2] > T; idx2 -= 1; end
        best = max(best,v1[idx1]+v2[idx2])
    end
    return best
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,T = gis()
    A::VI = gis()
    ans = solve(N,T,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

