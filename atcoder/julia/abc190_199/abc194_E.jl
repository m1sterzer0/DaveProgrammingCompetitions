
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

function solve(N::I,M::I,A::VI)::I
    ## Step 1: find Mex of A[1:M]
    ## Step 2: find the occurance times of the numbers less than mex in A[1:M]
    ## Step 3: cycle through rest of points, looking for >M gaps in those numbers
    ## Step 4: deal with the end state 
    ## O(N)
    sb::VI = fill(0,M+1)
    for i in 1:M
        if A[i] > M; continue; end
        sb[A[i]+1] = i
    end
    mymex::I = 10^18
    for i in 0:M; if sb[i+1] == 0; mymex = i; break; end; end
    for i in M+1:N
        if A[i] < mymex
            if i - sb[A[i]+1] > M; mymex = A[i]; else; sb[A[i]+1] = i; end
        end
        if i == N
            for j in 0:mymex-1
                if sb[j+1] < N-M+1; mymex = j; break; end
            end
        end
    end
    return mymex
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M = gis()
    A::VI = gis()
    ans = solve(N,M,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

