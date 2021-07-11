
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


## After the each operation, the plane is in one of 8 orientations (D_2*4 group) and is shifted
## Keep track of what happens to (0,0), (0,1), and (1,0), and answer questions based on that.
## We can do offline processing, but the history vector here isn't horrible
rot90clk(pt::PI) = (pt[2],-pt[1])
rot90cclk(pt::PI) = (-pt[2],pt[1])
trans(pt::PI,sh::PI) = (pt[1]+sh[1],pt[2]+sh[2])
flipy(pt::PI) = (pt[1],-pt[2])
flipx(pt::PI) = (-pt[1],pt[2])

function solve(N::I,X::VI,Y::VI,M::I,op::VPI,Q::I,A::VI,B::VI)
    ans::VS = []
    hist1::VPI = []
    hist2::VPI = []
    hist3::VPI = []
    orig = (0,0); xpt = (1,0); ypt = (0,1)
    for i in 1:M
        if op[i][1] == 1;     
            orig = rot90clk(orig)
            xpt  = rot90clk(xpt)
            ypt  = rot90clk(ypt)
        elseif op[i][1] == 2;
            orig = rot90cclk(orig)
            xpt = rot90cclk(xpt)
            ypt = rot90cclk(ypt)
        elseif op[i][1] == 3
            orig = trans(flipx(orig),(2*op[i][2],0))
            xpt =  trans(flipx(xpt), (2*op[i][2],0))
            ypt =  trans(flipx(ypt), (2*op[i][2],0))
        else
            orig = trans(flipy(orig),(0,2*op[i][2]))
            xpt =  trans(flipy(xpt), (0,2*op[i][2]))
            ypt =  trans(flipy(ypt), (0,2*op[i][2]))
        end
        #print("DBG: orig:$orig xpt:$xpt ypt:$ypt\n")
        push!(hist1,orig)
        push!(hist2,xpt)
        push!(hist3,ypt)
    end
    for i in 1:Q
        if A[i] == 0
            push!(ans,"$(X[B[i]]) $(Y[B[i]])")
        else
            lorig::PI,lx::PI,ly::PI = hist1[A[i]],hist2[A[i]],hist3[A[i]]
            x = lorig[1] + (lx[1]-lorig[1]) * X[B[i]] + (ly[1]-lorig[1]) * Y[B[i]]
            y = lorig[2] + (lx[2]-lorig[2]) * X[B[i]] + (ly[2]-lorig[2]) * Y[B[i]]
            push!(ans,"$x $y")
        end
    end
    return ans
end

function main(infn="")
    global infile
    infile = (infn != "") ? open(infn,"r") : length(ARGS) > 0 ? open(ARGS[1],"r") : stdin
    N = gi()
    X::VI = fill(0,N)
    Y::VI = fill(0,N)
    for i in 1:N; X[i],Y[i] = gis(); end
    M = gi()
    op::VPI = []
    for i in 1:M
        xx = gis()
        if xx[1] in (1,2); 
            push!(op,(xx[1],0))
        else
            push!(op,(xx[1],xx[2]))
        end
    end
    Q = gi()
    A::VI = fill(0,Q)
    B::VI = fill(0,Q)
    for i in 1:Q; A[i],B[i] = gis(); end
    ans = solve(N,X,Y,M,op,Q,A,B)
    for l in ans; println(l); end
end

Random.seed!(8675309)
main()

#using Profile, StatProfilerHTML
#Profile.clear()
#@profile main("B.in")
#Profile.clear()
#@profilehtml main("B.in")

