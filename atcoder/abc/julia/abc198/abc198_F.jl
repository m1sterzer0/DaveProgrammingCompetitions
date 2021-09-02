
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

#################################################################################################################
## Use Burnside's lemma
## 1 Identity
## 6 Face rotations by +/- 90 degrees.  1-4-1 fixes this.
## 3 Face rotations by +/- 180 degrees. 2-2-1-1 fixes these
## 6 180 degree rotations around 2 opposite edges.  2-2-2 fixes these.
## 8 +/- 120deg rotations around 2 opposite corners. 3-3 fixes these.
#################################################################################################################

function mycomb(x::I,n::I,mm::I)
    ans::I = 1
    for i in 0:n-1; ans = ans * ((x-i)%mm) % mm; end
    for i in n:-1:2; ans = ans * invmod(i,mm) % mm; end
    return ans
end

function solve(S::I)::I
    mm::I = 998244353
    S -= 6  ## Change positive now to non-negative
    ans::I = mycomb(S+5,5,mm) ## Stars and Bars

    ## 3-3 case
    if S % 3 == 0 
        adder::I = 8 * ((S÷3 + 1) % mm) % mm
        ans = (ans + adder) % mm
    end

    ## 2-2-2 case
    if S % 2 == 0 
        adder = 6 * mycomb(S÷2+2,2,mm) % mm
        ans = (ans + adder) % mm
    end

    ## 4-1-1 case
    for a in 0:3; for b in 0:3
        n = S-a-b
        if n < 0 || n % 4 != 0; continue; end
        adder = 6 * mycomb(n÷4+2,2,mm) % mm
        ans = (ans + adder) % mm
    end; end
    
    ## 2-2-1-1 case
    for a in 0:1; for b in 0:1
        n= S-a-b
        if n < 0 || n % 2 != 0; continue; end
        adder = 3 * mycomb(n÷2+3,3,mm) % mm
        ans = (ans + adder) % mm
    end; end

    ans = ans * invmod(24,mm) % mm
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    S = gi()
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

