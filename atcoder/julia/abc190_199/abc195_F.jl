
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

function solve(A::I,B::I)::I
    masterpr::VI = [2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71]
    pr::VI = [ x for x in masterpr if x <= B-A ]
    dp::VI = fill(0,2^length(pr))
    virtprime::I = 0
    for x in A:B
        mask::I = 0
        for i in 1:length(pr); if x % pr[i] == 0; mask |= 1 << (i-1); end; end
        if mask == 0 ##e.g. x is prime, or is product of prime factors all above 71
            virtprime += 1
        else
            for i::I in 1:2^length(pr)-1
                if mask == i
                    dp[i+1] += 1
                elseif mask & i == mask
                    childmask::I = i ⊻ mask
                    dp[i+1] += dp[childmask+1]
                end
            end
        end
    end
    return 2^(virtprime) + 2^(virtprime) * sum(dp)
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    A,B = gis()
    ans = solve(A,B)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

