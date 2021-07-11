
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

function solve(N::I, S::String, Q::I, T::VI, A::VI, B::VI)
    swapped = false
    SS::VC = [x for x in S]
    for i in 1:Q
        if T[i] == 2; swapped = !swapped; continue; end
        a = !swapped ? A[i] : A[i] <= N ? A[i] + N : A[i] - N
        b = !swapped ? B[i] : B[i] <= N ? B[i] + N : B[i] - N
        SS[a],SS[b] = SS[b],SS[a]
    end
    if !swapped; return join(SS,""); else; return join(SS[N+1:end],"")*join(SS[1:N],""); end
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    S = gs()
    Q = gi()
    T::VI = fill(0,Q)
    A::VI = fill(0,Q)
    B::VI = fill(0,Q)
    for i in 1:Q; T[i],A[i],B[i] = gis(); end
    ans = solve(N,S,Q,T,A,B)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

