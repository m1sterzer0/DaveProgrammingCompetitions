
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

function solve(N::I,M::I,A::VI,B::VI,C::VC)
    inf::I = 10^18
    dist::VI     = fill(inf,(N+1)^2)
    oneaway::VB  = fill(false,(N+1)^2)
    zeroaway::VB = fill(false,(N+1)^2)
    gr::VVI = [VI() for i in 1:(N+1)^2]
    for i::I in 1:N; zeroaway[(N+1)*i+i] = true; end
    for i::I in 1:M
        oneaway[(N+1)*A[i]+B[i]] = true
        oneaway[(N+1)*B[i]+A[i]] = true
    end
    for cc::Char in "abcdefghijklmnopqrstuvwxyz"
        ee::VI = [i for i in 1:M if C[i] == cc]
        if length(ee) <= 1; continue; end
        for i::I in 1:length(ee)
            a::I,b::I = A[ee[i]],B[ee[i]]
            for j::I in 1:length(ee)
                if i == j; continue; end
                c::I,d::I = A[ee[j]],B[ee[j]]
                push!(gr[(N+1)*a+c],(N+1)*b+d)
                push!(gr[(N+1)*b+c],(N+1)*a+d)
                push!(gr[(N+1)*a+d],(N+1)*b+c)
                push!(gr[(N+1)*b+d],(N+1)*a+c)
            end
        end
    end

    q::VI = [(N+1)*1+N]; dist[(N+1)*1+N] = 0
    ans = inf
    while !isempty(q)
        n::I = popfirst!(q)
        d = dist[n]
        if d > ans; return ans; end
        if zeroaway[n]; ans = min(ans,d); end
        if oneaway[n];  ans = min(ans,d+1); end ## Still might be a better answer in the queue
        for c in gr[n]
            if dist[c] < inf; continue; end
            dist[c] = d+2
            push!(q,c)
        end
    end
    return ans == inf ? -1 : ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = fill(0,M)
    B::VI = fill(0,M)
    C::VC = fill('.',M)
    for i in 1:M
        xx::VS = gss()
        A[i] = parse(Int64,xx[1])
        B[i] = parse(Int64,xx[2])
        C[i] = xx[3][1]
    end
    ans = solve(N,M,A,B,C)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

