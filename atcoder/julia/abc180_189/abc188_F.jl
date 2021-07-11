
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

function doit(cache::Dict{I,I},X::I,Y::I)::I
    if !haskey(cache,Y)
        if X > Y
            cache[Y] = X - Y
        elseif X == Y
            cache[Y] = 0
        elseif Y % 2 == 0 
            cache[Y] = min(Y-X,1+doit(cache,X,Y÷2))
        else
            cache[Y] = min(Y-X,2+doit(cache,X,(Y+1)÷2),2+doit(cache,X,(Y-1)÷2))
        end
    end
    return cache[Y]
end

function solve(X::I,Y::I)::I
    ####################################################################################################
    ## If X > Y, all we can do is decrease
    ## If X == Y, answer is zero
    ## If X < Y, we think about the operations in reverse
    ## -- Now we have add 1, sub 1, and div by 2 if we are even
    ## -- We contend that if we are going to div by 2, we should do it as soon as possible.  It is cheaper to make adjustments below.
    ####################################################################################################
    cache::Dict{I,I} = Dict{I,I}()
    return doit(cache,X,Y)
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    X,Y = gis()
    ans = solve(X,Y)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

