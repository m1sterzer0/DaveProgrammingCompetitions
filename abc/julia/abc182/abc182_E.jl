
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

function solve(H::I,W::I,N::I,M::I,A::VI,B::VI,C::VI,D::VI)
    gr::Array{Char,2} = fill('.',H,W)
    sb::Array{I,2} = fill(0,H,W)
    for i in 1:N; gr[A[i],B[i]] = 'L'; sb[A[i],B[i]] = 1; end
    for i in 1:M; gr[C[i],D[i]] = 'x'; end
    for i in 1:H
        on::Bool = false
        for j in 1:W;
            if gr[i,j] == 'L'; on = true; elseif gr[i,j] == 'x'; on = false; end
            if on; sb[i,j] |= 1; end
        end
        on = false
        for j in W:-1:1;
            if gr[i,j] == 'L'; on = true; elseif gr[i,j] == 'x'; on = false; end
            if on; sb[i,j] |= 1; end
        end
    end
    for j in 1:W
        on::Bool = false
        for i in 1:H;
            if gr[i,j] == 'L'; on = true; elseif gr[i,j] == 'x'; on = false; end
            if on; sb[i,j] |= 1; end
        end
        on = false
        for i in H:-1:1;
            if gr[i,j] == 'L'; on = true; elseif gr[i,j] == 'x'; on = false; end
            if on; sb[i,j] |= 1; end
        end
    end
    return sum(sb)
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    H,W,N,M = gis()
    A::VI = fill(0,N)
    B::VI = fill(0,N)
    C::VI = fill(0,M)
    D::VI = fill(0,M)
    for i in 1:N; A[i],B[i] = gis(); end
    for i in 1:M; C[i],D[i] = gis(); end
    ans = solve(H,W,N,M,A,B,C,D)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

