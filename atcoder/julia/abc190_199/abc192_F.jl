
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

## dp[numelem][modulus][remainder] == maximum sum
function solve(N::I,X::I,A::VI)
    dp::Array{I,3} = fill(-1,100,100,100)
    for a::I in A
        for numelem::I in N-1:-1:1
            for m::I in 1:N
                for remainder::I in 0:m-1
                    v::I = dp[numelem,m,remainder+1]
                    if v < 0; continue; end
                    newv::I = v + a; newrem = newv % m
                    if newv > dp[numelem+1,m,newrem+1]; dp[numelem+1,m,newrem+1] = newv; end
                end
            end
        end
        for m in 1:N
            myrem::I = a % m
            if a > dp[1,m,myrem+1]; dp[1,m,myrem+1] = a; end
        end
    end
    best::I = 8*10^18
    for m::I in 1:N
        rem::I = X % m
        if dp[m,m,rem+1] < 0; continue; end
        tt::I = (X - dp[m,m,rem+1]) ÷ m
        best = min(best,tt)
    end
    return best
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,X = gis()
    A::VI = gis()
    ans = solve(N,X,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

