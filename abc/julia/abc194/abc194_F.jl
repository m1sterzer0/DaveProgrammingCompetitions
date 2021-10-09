
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

function solve(N::String,K::I)::I
    n::I = length(N)
    mm::I = 10^9+7
    dp::VI = fill(0,16)
    ldp::VI = fill(0,16)
    seen::Set{Char} = Set{Char}()
    for i::I in 1:n
        (ldp,dp) = (dp,ldp); dp .= 0
        if i == 1
            for c in "123456789ABCDEF"
                if N[i] == c; break; end
                dp[1] += 1
            end
        else
            ## Case 1, prefix is less than N already -- we can add anything
            for j in 1:16
                dp[j] = (dp[j] + j * ldp[j] % mm) % mm
                if j < 16; dp[j+1] = (dp[j+1] + (16-j) * ldp[j] % mm) % mm; end
            end
            ## Case 2, we are starting a new number -- can add start with any digit except 0
            dp[1] = (dp[1] + 15) % mm
            ## Case 3, we use i-1 prefix of N, and then choose something smaller than ith element
            for c in "0123456789ABCDEF"
                if N[i] == c; break; end
                j = c ∈ seen ? length(seen) : length(seen)+1
                dp[j] = (dp[j]+1) % mm
            end
        end
        push!(seen,N[i])
    end
    return (dp[K] + (length(seen) == K ? 1 : 0)) % mm
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    xx = gss()
    N = xx[1]
    K = parse(Int64,xx[2])
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

