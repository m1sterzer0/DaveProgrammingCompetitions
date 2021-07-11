
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

function ntt(a::VI,invert::Bool,p::I,root::I,rootpw::I)
    n::I = length(a)
    rootinv::I = invmod(root,p)
    j::I = 0
    for i::I in 1:n-1
        bit::I = n >> 1
        while (j & bit != 0); j ⊻= bit; bit >>= 1; end
        j ⊻= bit
        if i < j; ip1::I = i+1; jp1::I = j+1; (a[ip1],a[jp1]) = (a[jp1],a[ip1]); end
    end

    len::I = 2
    while len <= n
        wlen::I = invert ? rootinv : root; i::I = len
        while i < rootpw; wlen = wlen * wlen % p; i <<= 1; end
        for i::I in 0:len:n-1
            w::I = 1; i1::I = i+1; i2::I = i1+len>>1
            for j::I in 0:len>>1-1
                i1j::I = i1+j; i2j::I = i2+j
                u::I = a[i1j]; v::I = a[i2j] * w % p
                a[i1j] = u + v < p ? u + v : u + v - p
                a[i2j] = u - v >= 0 ? u - v : u - v + p
                w = w * wlen % p
            end
        end
        len <<= 1
    end

    if invert
        n1 = invmod(n,p)
        for i in 1:n; a[i] = a[i] * n1 % p; end
    end
end

function fft(a::Vector{Complex{Float64}},invert::Bool)
    n::I = length(a)
    j::I = 0
    for i::I in 1:n-1
        bit::I = n >> 1
        while (j & bit != 0); j ⊻= bit; bit >>= 1; end
        j ⊻= bit
        if i < j; ip1::I = i+1; jp1::I = j+1; (a[ip1],a[jp1]) = (a[jp1],a[ip1]); end
    end

    len::I = 2
    while len <= n
        wlen::Complex{Float64} = exp(2*pi/len*(invert ? -1 : 1)*im)
        for i::I in 1:len:n
            w::Complex{Float64} = 1
            hlen::I = len>>1
            for j::I in 0:hlen-1
                u::Complex{Float64} = a[i+j]
                v::Complex{Float64} = w * a[i+j+hlen]
                a[i+j] = u+v
                a[i+j+hlen] = u-v
                w *= wlen
            end
        end
        len <<= 1
    end
    if invert; for i in 1:n; a[i] /= n; end; end
end

function convolventt(a::VI,b::VI,p::I,root::I,rootpw::I)::VI
    n = 1; while n < length(a)+length(b); n *= 2; end
    aa::VI = fill(0,n); aa[1:length(a)] .= a
    bb::VI = fill(0,n); bb[1:length(b)] .= b
    ntt(aa,false,p,root,rootpw)
    ntt(bb,false,p,root,rootpw)
    cc::VI = [aa[i]*bb[i] % p for i in 1:n]
    ntt(cc,true,p,root,rootpw)
    return cc
end

function convolvefft(a::VI,b::VI)::VI
    n = 1; while n < length(a)+length(b); n *= 2; end
    aa::Vector{Complex{Float64}} = fill(Complex{Float64}(0.00),n)
    bb::Vector{Complex{Float64}} = fill(Complex{Float64}(0.00),n)
    aa[1:length(a)] .= a
    bb[1:length(b)] .= b
    fft(aa,false); fft(bb,false)
    cc::Vector{Complex{Float64}} = [aa[i]*bb[i] for i in 1:n]
    fft(cc,true)
    dd::VI = [round(Int64,abs(x)) for x in cc]
    return dd
end

function solve(S::String,T::String)::I
    Trev::String = reverse(T)
    As::VI    = [x == '0' ? 0 : 1 for x in S]
    At::VI    = [x == '0' ? 0 : 1 for x in Trev]
    Asinv::VI = [1-x for x in As]
    Atinv::VI = [1-x for x in At]
    conv1::VI = convolvefft(As,At)
    conv2::VI = convolvefft(Asinv,Atinv)
    conv3::VI = conv1 + conv2
    bestmatch::I = maximum(conv3[length(T):length(S)])
    return length(T) - bestmatch
end

function gencase(Smin::I,Smax::I,Tmin::I,Tmax::I)
    soneprob::F = rand()
    toneprob::F = rand()
    slen = rand(Smin:Smax)
    tlen = rand(min(slen,Tmin):min(slen,Tmax))
    S = join([rand() < soneprob ? '1' : '0' for i in 1:slen],"")
    T = join([rand() < toneprob ? '1' : '0' for i in 1:tlen],"")
    return (S,T)
end

function test(ntc::I,Smin::I,Smax::I,Tmin::I,Tmax::I)
    for ttt in 1:ntc
        (S,T) = gencase(Smin,Smax,Tmin,Tmax)
        ans = solve(S,T); println("Case #$ttt: $ans")
    end
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    S = gs()
    T = gs()
    ans = solve(S,T)
    println(ans)
end

Random.seed!(8675309)
main()
#test(10,900000,1000000,500000,1000000)

#a = [1, 2, 1, 0]
#aa = [Complex{Float64}(x) for x in a]
#fft(aa,false)
#println("$aa")
#fft(aa,true)
#println("$aa")
#bb = [round(Int64,abs(x)) for x in aa]
#println("$bb")


#using Profile, StatProfilerHTML
#Profile.clear()
#@profile test(1,100, 100, 100, 100)
#Profile.clear()
#@profilehtml test(10,900000,1000000,500000,1000000)

