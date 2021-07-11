
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

function solve(S::String)
    sb::VI = fill(0,9)
    for c in S; sb[c-'0'] += 1; end
    ## Length is 1
    if length(S) == 1
        if S == "8"; return "Yes"; else; return "No"; end
    elseif length(S) == 2
        if parse(Int64,S) % 8 == 0 || parse(Int64,reverse(S)) % 8 == 0; return "Yes"; else; return "No"; end
    else
        for v in 104:8:999
            sv = string(v)
            if '0' in sv; continue; end
            if sb[sv[1]-'0'] < (sv[1]==sv[2]==sv[3] ? 3 : (sv[1]==sv[2]||sv[1]==sv[3]) ? 2 : 1); continue; end
            if sb[sv[2]-'0'] < (sv[1]==sv[2]==sv[3] ? 3 : (sv[1]==sv[2]||sv[2]==sv[3]) ? 2 : 1); continue; end
            if sb[sv[3]-'0'] < (sv[1]==sv[2]==sv[3] ? 3 : (sv[2]==sv[3]||sv[1]==sv[3]) ? 2 : 1); continue; end
            return "Yes"
        end
        return "No"
    end
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    S = gs()
    ans = solve(S)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

