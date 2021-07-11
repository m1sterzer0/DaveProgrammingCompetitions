
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

function solve(H::I,W::I,A::Array{Char,2})
    v::Array{I,2} = fill(0,H,W)
    for i::I in H:-1:1
        for j::I in W:-1:1
            if i == H && j == W; continue; end
            if (i+j) % 2 == 0
                ## Tashahaki's turn -- maximize 
                if i == H; v[i,j] = v[i,j+1] + (A[i,j+1] == '+' ? 1 : -1); continue; end
                if j == W; v[i,j] = v[i+1,j] + (A[i+1,j] == '+' ? 1 : -1); continue; end
                v[i,j] = max(v[i+1,j] + (A[i+1,j] == '+' ? 1 : -1), v[i,j+1] + (A[i,j+1] == '+' ? 1 : -1))
            else
                ## Aoki's turn -- minimize and reverse affect of signs 
                if i == H; v[i,j] = v[i,j+1] + (A[i,j+1] == '+' ? -1 : 1); continue; end
                if j == W; v[i,j] = v[i+1,j] + (A[i+1,j] == '+' ? -1 : 1); continue; end
                v[i,j] = min(v[i+1,j] + (A[i+1,j] == '+' ? -1 : 1), v[i,j+1] + (A[i,j+1] == '+' ? -1 : 1))
            end
        end
    end
    return v[1,1] > 0 ? "Takahashi" : v[1,1] < 0 ? "Aoki" : "Draw"
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    H,W = gis()
    A::Array{Char,2} = fill('.',H,W)
    for i in 1:H; A[i,:] = [x for x in gs()]; end
    ans = solve(H,W,A)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

