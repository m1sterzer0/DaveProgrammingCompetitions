
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

function solveit(v::I,idx::I,N::I,A::VI,cache::Dict{PI,I})
    if !haskey(cache,(v,idx))
        if idx == 1
            cache[(v,idx)] = 1
        elseif idx == N
            nn = v ÷ A[idx]
            subtot = nn * A[idx]
            if subtot == v
                cache[(v,idx)] = 1
            else 
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache) + solveit(A[idx]-v+subtot,idx-1,N,A,cache)
            end
        else
            nn = v ÷ A[idx]
            subtot = nn * A[idx]
            if subtot == v
                cache[(v,idx)] = 1
            elseif nn == A[idx+1] ÷ A[idx] - 1
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache)
            else
                cache[(v,idx)] = solveit(v-subtot,idx-1,N,A,cache) + solveit(A[idx]-v+subtot,idx-1,N,A,cache)
            end
        end
    end
    return cache[(v,idx)]
end


function solve(N::I,X::I,A::VI)
    cache::Dict{PI,I} = Dict{PI,I}()
    return solveit(X,N,N,A,cache)
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

