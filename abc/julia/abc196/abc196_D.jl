
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

function solveit(bd::I,A::I,B::I,H::I,W::I,cache::VI)
    key::I = bd | A << 16 | B << 20
    if cache[key] == -1
        if A == 0 && B == 0
            cache[key] = 1
        else
            i = 0; ans = 0
            while bd & (1<<i) != 0; i += 1; end
            if B != 0
                ans += solveit(bd | (1<<i),A,B-1,H,W,cache)
            end
            if A != 0 && i % W != W-1 && bd & (1<<(i+1)) == 0
                ans += solveit(bd | (1<<i) | (1 << (i+1)),A-1,B,H,W,cache)
            end
            if A != 0 && i+W < H*W && bd & (1<<(i+W)) == 0
                ans += solveit(bd | (1<<i) | (1 << (i+W)),A-1,B,H,W,cache)
            end
            cache[key] = ans
        end
    end
    return cache[key]
end

function solve(H::I,W::I,A::I,B::I)
    cache::VI = fill(-1,2^25)
    return solveit(0,A,B,H,W,cache)
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    H,W,A,B = gis()
    ans = solve(H,W,A,B)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

