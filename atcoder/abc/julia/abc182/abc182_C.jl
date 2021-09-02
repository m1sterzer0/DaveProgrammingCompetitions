
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

function solve(N)
    sn = string(N)
    lsn = length(sn)
    ones::I = 0; twos::I = 0
    for c in sn; if c in "147"; ones += 1; elseif c in "258"; twos += 1; end; end
    if N % 3 == 0; return 0; end
    if N % 3 == 1; return ones >= 1 && lsn > 1 ? 1 : twos >= 2 && lsn > 2 ? 2 : -1; end
    if N % 3 == 2; return twos >= 1 && lsn > 1 ? 1 : ones >= 2 && lsn > 2 ? 2 : -1; end
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    ans = solve(N)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

