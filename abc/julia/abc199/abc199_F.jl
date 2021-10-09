
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

function eyeInt(n::I)::Array{I,2}
    a::Array{I,2} = fill(0,n,n)
    for i in 1:n; a[i,i] = 1; end
    return a
end

function modmatmul(a::Array{I,2},b::Array{I,2},mm::I)::Array{I,2}
    m::I,x::I = size(a)
    x2::I,n::I = size(b)
    @assert x == x2
    res::Array{I,2} = fill(0,m,n)
    for i::I in 1:m
        for j::I in 1:n
            for k::I in 1:x
                res[i,j] = (res[i,j]+ (a[i,k]*b[k,j] % mm)) % mm
            end
        end
    end
    return res
end

function modmatmul(a::Array{I,2},b::VI,mm::I)::VI
    m::I,x::I = size(a)
    x2::I = length(b)
    @assert x == x2
    res::VI = fill(0,m)
    for i::I in 1:m
        for k::I in 1:x
            res[i] = (res[i] + (a[i,k]*b[k] % mm)) % mm
        end
    end
    return res
end


function modmatpow(a::Array{I,2},p::I,mm::I)::Array{I,2}
    amul::Array{I,2} = copy(a)
    res::Array{I,2} = eyeInt(size(a)[1])
    while p > 0
        if p & 1 > 0; res = modmatmul(amul,res,mm); end
        amul = modmatmul(amul,amul,mm)
        p >>= 1
    end
    return res
end

function solve(N,M,K,A,X,Y)
    mm::I = 1_000_000_007
    mat::Array{I,2} = fill(0,N,N)
    ## Calculate (2M)^K*expected value, and then divide by 2M afterwards
    for eg in 1:M
        (x::I,y::I) = X[eg],Y[eg]
        for i in 1:N
            mat[i,i] += 1
            if i != x && i != y; mat[i,i] += 1; elseif i == x; mat[i,y] += 1; else; mat[i,x] += 1; end
        end
    end
    matp::Array{I,2} = modmatpow(mat,K,mm)
    res = modmatmul(matp,A,mm)
    ff::I = powermod(invmod(2M,mm),K,mm)
    for i in 1:N; res[i] = (res[i]*ff) % mm; end
    return res
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N,M,K = gis()
    A = gis()
    X::VI = fill(0,M)
    Y::VI = fill(0,M)
    for i in 1:M; X[i],Y[i] = gis(); end
    ans = solve(N,M,K,A,X,Y)
    for i in ans; println(i); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

