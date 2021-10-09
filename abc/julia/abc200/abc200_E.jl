
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

function solve(N::I,K::I)::String
    dp::Array{I,2} = fill(0,3,3N)
    ## Do first level explicitly
    for i in 1:N; dp[1,i] = 1; end
    ## Do second & third level
    deriv::VI = fill(0,3N+1)
    for l in (2,3)
        fill!(deriv,0)
        for i in 1:(l==2 ? N : 2N)
            deriv[i+1] += dp[l-1,i]
            deriv[i+N+1] -= dp[l-1,i]
        end
        cur::I = 0
        for i in 1:(l == 2 ? 2N : 3N)
            cur += deriv[i]
            dp[l,i] = cur
        end
    end


    ## Find the sum
    cum = 0; mysum = 0
    for i in 3:3N; 
        newcum::I = cum + dp[3,i]
        if newcum >= K; K -= cum; mysum = i; break; else; cum = newcum; end
    end
    ## Find the first element
    cum = 0; firstdig = 0
    for i in 1:N
        newcum::I = cum + dp[2,mysum-i]
        if newcum >= K; K -= cum; firstdig = i; break; else; cum = newcum; end
    end
    ## Find the second element
    cum = 0; seconddig = 0
    for i in 1:N
        newcum::I = cum + dp[1,mysum-firstdig-i]
        if newcum >= K; K -= cum; seconddig = i; break; else; cum = newcum; end
    end
    thirddig = mysum-firstdig-seconddig
    return "$firstdig $seconddig $thirddig"
end 

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,K = gis()
    ans = solve(N,K)
    println(ans)
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

