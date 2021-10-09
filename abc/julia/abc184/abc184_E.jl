
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

function solve(H::I,W::I,bd::Array{Char,2})::I
    ## Need to find S, G, and all lower case letters
    d::Dict{Char,VPI} = Dict{Char,VPI}()
    for c in "SGabcdefghijklmnopqrstuvwxyz"; d[c] = VPI(); end
    for i in 1:H; for j in 1:W
        if bd[i,j] != '.' && bd[i,j] != '#'; push!(d[bd[i,j]],(i,j)); end
    end; end
    inf::I = 10^18
    dist::Array{I,2} = fill(inf,H,W)
    si,sj = d['S'][1][1],d['S'][1][2]
    dist[si,sj] = 0
    q::VPI = [(si,sj)]
    tele::Dict{Char,Bool} = Dict{Char,Bool}()
    for c in "abcdefghijklmnopqrstuvwxyz"; tele[c] = false; end
    while !isempty(q)
        (i,j) = popfirst!(q)
        #print("DBG: i:$i j:$j dist[i,j]:$(dist[i,j])\n")
        for (ii,jj) in ((i-1,j),(i+1,j),(i,j-1),(i,j+1))
            #print("DBG:    Trying (ii,jj)=($ii,$jj)\n")
            if ii < 1 || ii > H || jj < 1 || jj > W; continue; end
            if bd[ii,jj] == '#' || dist[ii,jj] != inf; continue; end
            dist[ii,jj] = dist[i,j]+1
            #print("DBG:    Found (ii,jj)=($ii,$jj) at dist $(dist[ii,jj])\n")
            push!(q,(ii,jj))
        end
        if 'a' <= bd[i,j] <= 'z' && !tele[bd[i,j]]
            c = bd[i,j]
            tele[c] = true
            for (ii,jj) in d[c]
                if dist[ii,jj] < inf; continue; end
                dist[ii,jj] = dist[i,j] + 1
                push!(q,(ii,jj))
            end
        end
    end
    ei,ej = d['G'][1][1],d['G'][1][2]
    if dist[ei,ej] == inf; return -1; end
    return dist[ei,ej]
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    H,W = gis()
    bd::Array{Char,2} = fill('.',H,W)
    for i in 1:H; bd[i,:] = [x for x in gs()]; end
    ans = solve(H,W,bd)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

