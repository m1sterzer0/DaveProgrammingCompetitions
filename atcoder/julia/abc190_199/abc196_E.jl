
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

function sample(arr::VI,w::VF,sw::F)
    x = sw*rand()
    for i in 1:length(arr)
        x -= w[i]
        if x <= 0; return arr[i]; end
    end
    return arr[end]
end

function gencase(Nmin::I,Nmax::I,Amin::I,Amax::I,Qmin::I,Qmax::I,Xmin::I,Xmax::I)
    N = rand(Nmin:Nmax)
    A = rand(Amin:Amax,N)
    Q = rand(Qmin:Qmax)
    X = rand(Xmin:Xmax,Q)
    twts::VF = rand(3); sw::F = sum(twts)
    T = [sample([1,2,3],twts,sw) for i in 1:N]
    return (N,A,T,Q,X)
end

function solveBruteForce(N::I,A::VI,T::VI,Q::I,X::VI)::VI
    ans::VI = []
    for x in X
        for i in 1:N
            if T[i] == 1; x += A[i]; end
            if T[i] == 2; x = max(x,A[i]); end
            if T[i] == 3; x = min(x,A[i]); end
        end
        push!(ans,x)
    end
    return ans
end

function test(ntc::I,Nmin::I,Nmax::I,Amin::I,Amax::I,Qmin::I,Qmax::I,Xmin::I,Xmax::I,check::Bool=true)
    pass = 0
    for ttt in 1:ntc
        (N,A,T,Q,X) = gencase(Nmin,Nmax,Amin,Amax,Qmin,Qmax,Xmin,Xmax)
        ans2 = solve(N,A,T,Q,X)
        if check
            ans1 = solveBruteForce(N,A,T,Q,X)
            if ans1 == ans2
                 pass += 1
            else
                print("ERROR: ttt:$ttt ans1:$ans1 ans2:$ans2\n")
                #print("ERROR: ttt:$ttt\n")
                ans1 = solveBruteForce(N,A,T,Q,X)
                ans2 = solve(N,A,T,Q,X)
            end
       else
           print("Case $ttt: $ans2\n")
       end
    end
    if check; print("$pass/$ntc passed\n"); end
end


function solve(N::I,A::VI,T::VI,Q::I,X::VI)::VI
    lb::I=-10^18; ub::I=10^18; adder::I=0
    for i in 1:N
        if T[i] == 1;      lb += A[i]; ub += A[i]; adder += A[i]
        elseif T[i] == 2;  lb = max(lb,A[i]); ub = max(ub,A[i])
        else               ub = min(ub,A[i]); lb = min(lb,A[i])
        end
    end
    ans::VI = []
    for x in X; push!(ans,max(lb,min(ub,x+adder))); end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    A::VI = fill(0,N)
    T::VI = fill(0,N)
    for i in 1:N; A[i],T[i] = gis(); end
    Q = gi()
    X::VI = gis()
    ans = solve(N,A,T,Q,X)
    for x in ans; println(x); end
end

Random.seed!(8675309)
main()
#test(1000,1,1000,-10^9,10^9,1,1000,-10^9,10^9)
#test(1000,1,10,-10,10,1,10,-10,10)

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

