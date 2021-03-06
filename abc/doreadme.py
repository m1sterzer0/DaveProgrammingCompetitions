import os
from pathlib import Path

solversData = '''
abc160,9557,9427,7443,3417,2721,339
abc161,9784,8263,7740,2865,682,1117
abc162,10333,9924,8304,4199,850,680
abc163,11284,10786,8976,3663,380,107
abc164,11148,10415,9553,1926,497,41
abc165,11225,9994,2514,5554,975,596
abc166,11540,10088,7970,4999,2874,863
abc167,11507,10836,5652,4666,1393,415
abc168,10686,10466,7598,3856,467,85
abc169,11268,7017,5525,4591,1659,818
abc170,10398,9617,8272,2835,1177,388
abc171,10416,10164,5522,5885,4264,624
abc172,10050,9757,3234,3119,470,186
abc173,10564,10269,4890,4555,972,499
abc174,9720,9103,3349,5486,2013,1257
abc175,8513,7207,5316,1103,984,73
abc176,9228,8679,8381,1804,1958,18
abc177,9347,7137,5863,4125,2584,152
abc178,9574,8730,4565,3462,2653,512
abc179,8716,8333,6059,1777,2008,707
abc180,5699,4683,4585,2471,1189,80
abc181,6568,5936,4329,3152,1344,230
abc182,7436,6249,5077,3419,1988,233
abc183,7152,5919,4633,3400,1393,789
abc184,7738,7439,3137,1561,1223,1209
abc185,7332,5993,4326,3882,1006,1994
abc186,6188,5812,4936,3727,979,475
abc187,7110,6187,5025,3316,1250,438
abc188,7698,7524,6135,2510,1795,480
abc189,8747,5436,4158,3270,1008,244
abc190,8916,8636,4698,3619,821,1510
abc191,8128,8215,2257,424,1492,157
abc192,8571,8193,6985,1331,2113,680
abc193,7641,7114,4602,2771,423,110
abc194,7636,6319,4535,2027,1992,215
abc195,7588,4054,4989,2369,811,304
abc196,8435,7725,5503,1517,765,173
abc197,7757,6164,2659,2615,1140,376
abc198,7933,6864,4369,1507,1667,37
abc199,8600,7745,4524,571,550,310
abc200,8475,8136,5846,1696,432,91
abc201,8453,7994,4570,1382,701,110
abc202,8628,8317,5871,2442,772,25
abc203,8295,8124,5878,775,596,184
abc204,8710,8644,3776,2970,687,329
abc205,8692,8356,7366,3336,327,282
abc206,8961,8693,6267,2788,636,209
abc207,8260,6577,4502,285,521,126
abc208,7805,7352,6165,1711,341,43
abc209,8642,8395,5267,3517,247,163
abc210,8339,8256,4869,1011,811,60
abc211,8468,8114,3933,3158,530,144
abc212,7785,6536,4921,2839,1122,152,248,46
abc213,7599,7667,4327,3431,1193,201,49,22
abc214,8306,7169,4828,1346,551,406,32,14
abc215,8363,7060,5851,3216,1192,518,186,17
abc216,6997,6554,5356,1969,1847,898,393,9
abc217,8119,8096,7544,3115,2445,444,376,59
abc218,8653,8436,2385,3491,2408,669,235,48
abc219,7203,6901,4665,1859,681,93,181,10
abc220,7185,6418,5778,3286,867,1411,116,19
abc221,8032,7002,4553,2803,991,287,26,39
abc222,7082,7010,4237,2540,995,415,200,5
abc223,6699,5301,3651,1511,693,535,82,23
abc224,6433,5373,4027,1069,650,559,118,26
abc225,6417,5730,4092,2636,684,75,85,21
abc226,6377,4674,3350,2825,1255,289,133,36
abc227,5315,4760,2494,659,169,166,202,17
abc228,5570,5260,4422,1890,842,249,59,56
abc229,5751,5525,4582,2673,1929,438,279,10
'''
problemRatingData = '''
abc160,0,0,62,879,1036,2048
abc161,0,0,4,991,1760,1528
abc162,0,0,34,757,1662,1764
abc163,0,0,125,960,2037,2470
abc164,0,0,0,1232,1877,2683
abc165,0,0,1136,600,1620,1843
abc166,0,0,233,694,1062,1668
abc167,0,0,595,754,1442,1961
abc168,0,0,178,804,1896,2478
abc169,0,349,597,732,1353,1698
abc170,0,0,15,1033,1502,1968
abc171,0,0,560,498,778,1795
abc172,0,0,930,963,1880,2216
abc173,0,0,653,720,1607,1892
abc174,0,0,902,486,1227,1495
abc175,0,0,417,1491,1554,2512
abc176,0,0,0,1248,1204,2912
abc177,0,108,386,732,1057,2291
abc178,0,0,653,875,1054,1877
abc179,0,0,261,1251,1175,1713
abc180,0,0,0,752,1256,2419
abc181,0,0,248,600,1193,2009
abc182,0,0,274,701,1098,2121
abc183,0,0,329,662,1288,1586
abc184,0,0,782,1276,1418,1423
abc185,0,0,373,490,1468,1053
abc186,0,0,0,436,1461,1833
abc187,0,0,137,650,1358,1895
abc188,0,0,0,933,1170,1865
abc189,0,249,565,769,1526,2154
abc190,0,0,472,722,1645,1321
abc191,0,0,1063,1953,1323,2333
abc192,0,0,0,1425,1135,1783
abc193,0,0,378,866,1948,2475
abc194,0,0,386,1078,1088,2197
abc195,0,483,235,945,1609,2068
abc196,0,0,202,1277,1650,2274
abc197,0,0,809,831,1379,1945
abc198,0,0,413,1224,1161,2769
abc199,0,0,436,1804,1814,2065
abc200,0,0,138,1217,1955,2556
abc201,0,0,439,1317,1694,2484
abc202,0,0,130,966,1638,2905
abc203,0,0,54,1622,1750,2252
abc204,0,0,629,832,1710,2044
abc205,0,0,0,713,2025,2088
abc206,0,0,60,879,1745,2221
abc207,0,0,397,2074,1820,2398
abc208,0,0,0,1190,2024,2772
abc209,0,0,264,686,2153,2307
abc210,0,0,357,1507,1618,2632
abc211,0,0,559,755,1823,2350
abc212,0,0,205,775,1410,2332,2150,2741
abc213,0,0,481,710,1423,2215,2663,2806
abc214,0,0,309,1341,1835,1973,2893,3138
abc215,0,0,76,736,1413,1853,2276,3101
abc216,0,0,0,1039,1084,1541,1963,3295
abc217,0,0,0,802,986,1954,2047,3112
abc218,0,0,1012,715,1004,1753,2217,2805
abc219,0,0,260,1085,1690,2542,2287,3297
abc220,0,0,119,664,1593,1304,2439,3047
abc221,0,0,379,832,1515,2093,2914,2793
abc222,0,0,367,865,1491,1934,2242,3477
abc223,0,0,354,1069,1510,1641,2351,2771
abc224,0,0,301,1376,1657,1733,2374,2887
abc225,0,0,326,778,1678,2612,2566,3061
abc226,0,0,539,706,1327,2086,2373,2813
abc227,0,0,692,1643,2377,2323,2207,3229
abc228,0,0,214,1035,1579,2179,2711,2729
abc229,0,0,165,745,1015,1917,2121,3308
'''

notes = {}
notes["abc160_E"] = "Max Heap or simple sorting"
notes["abc160_F"] = "Tree DP, Rerooting, Combinatorics"
notes["abc161_E"] = "Simple DP"
notes["abc161_F"] = "GCD, Factors"
notes["abc162_E"] = "Divisors, Inclusion/Exclusion, Modular Math"
notes["abc162_F"] = "DP of DP"
notes["abc163_E"] = "DP, Greedy"
notes["abc163_F"] = "Tree DP"
notes["abc164_D"] = "Simple Digit DP"
notes["abc164_E"] = "Interesting Dijkstra"
notes["abc164_F"] = "Bitwise, Logical constraint solving"
notes["abc165_C"] = "Brute force search.  Stars and Bars.  Combination iteration"
notes["abc165_E"] = "Interesting matching construction"
notes["abc165_F"] = "Tree DP, DFS w/ rollback"
notes["abc166_E"] = "Simple DP"
notes["abc166_F"] = "Greedy game theory"
notes["abc167_E"] = "Simple Combinatorics"
notes["abc167_F"] = "Bracket Sequences"
notes["abc168_E"] = "Dot product, GCD, Combinatorics, Special Cases"
notes["abc168_F"] = ("Coordinate compression, BFS, Bookkeeping.  Note: big swing on constant factors dependent on subtle " +
                    "algorithm choices. My first coding of pypy was too slow.  Using byte arrays (array.array -- factor of 2) " +
                    "and slightly better algorithm (1 coord per segment vs. 3 -- factor of 9) made it much faster.")
notes["abc169_E"] = "Median, Casework"
notes["abc169_F"] = "DP, Subsets"
notes["abc170_D"] = "Sort and Sieve"
notes["abc170_E"] = "Lots of minheaps or lots of multisets"
notes["abc170_F"] = "Clever Dijkstra.  Somewhat time challenged in both go and python (there is likely more optimization to be done)."
notes["abc171_E"] = "Simple xor problem"
notes["abc171_F"] = "Counting. Strings of length N with given subsequence"
notes["abc172_E"] = "Counting. Inclusion/Exclusion"
notes["abc172_F"] = "NIM. Relationship between xor and addition. Bitwise."
notes["abc173_E"] = "Either lots of casework or smarter brute force (I did the casework)."
notes["abc173_F"] = "Counting.  Connected components on tree."
notes["abc174_E"] = "Simple binary search"
notes["abc174_F"] = "Range set query. Offline processing. Fenwick tree."
notes["abc175_D"] = "Permutation Loops.  Bookeeping."
notes["abc175_E"] = "DP on a grid"
notes["abc175_F"] = "Interesting Dijkstra"
notes["abc176_D"] = "0/1 BFS"
notes["abc176_E"] = "Coordinate compression, maps"
notes["abc176_F"] = "Tricky DP (need to be efficient)"
notes["abc177_E"] = "Factor sieve"
notes["abc177_F"] = "Tricky lazy segtree"
notes["abc178_E"] = "45 deg Rotation for L1 -> Linf norm"
notes["abc178_F"] = "Easy greedy assignment problem"
notes["abc179_D"] = "Receiving DP"
notes["abc179_E"] = "Typical 'find the sequence loop' problem"
notes["abc179_F"] = "Another lazy segtree problem (other approaches too)"
notes["abc180_E"] = "Permutation to subset dp conversion"
notes["abc180_F"] = "Graph creation, DP"
notes["abc181_E"] = "Prefix Sum, Suffix sum, Binary Search"
notes["abc181_F"] = "Binary search, DSU"
notes["abc182_E"] = "Simple grid problem"
notes["abc182_F"] = "Money denominations, recursive, caching"
notes["abc183_E"] = "Chess, Grid DP"
notes["abc183_F"] = "Nice augmented DSU problem"
notes["abc184_D"] = "Expected Value DP"
notes["abc184_E"] = "Grid, BFS, Teleporters"
notes["abc184_F"] = "Meet in the middle, Subset sums"
notes["abc185_E"] = "Edit distance"
notes["abc185_F"] = "Simple segtree problem"
notes["abc186_E"] = "Extended euclidean algorithm"
notes["abc186_F"] = "segtree or fenwick, chess"
notes["abc187_E"] = "Simple tree DP"
notes["abc187_F"] = "Cliques, subset DP"
notes["abc188_E"] = "DAG DP"
notes["abc188_F"] = "Reverse problem.  Recursive functions with memorization."
notes["abc189_C"] = "Max area under histogram"
notes["abc189_E"] = "Bookeeping with reflections, translations, and rotations"
notes["abc189_F"] = "Expected value w/ simple algebra"
notes["abc190_E"] = "Dijkstra, then Permutaion/Subset DP"
notes["abc190_F"] = "Inversion counting.  Fenwick"
notes["abc191_C"] = "Rectilinear polygon side counting"
notes["abc191_D"] = "Counting lattice points in a circle"
notes["abc191_E"] = "Loop around a slightly modified Dijkstra"
notes["abc191_F"] = "GCD. Clever DP"
notes["abc192_D"] = "Binary search.  Easier with BigInt"
notes["abc192_E"] = "Loop around a slightly modified Dijkstra (again)"
notes["abc192_F"] = "Good DP"
notes["abc193_E"] = "CRT problem (linear congruences)"
notes["abc193_F"] = "Max Flow"
notes["abc194_D"] = "Coupon collector"
notes["abc194_E"] = "Simple sliding window"
notes["abc194_F"] = "Digit DP"
notes["abc195_E"] = "DP"
notes["abc195_F"] = "Subset DP"
notes["abc196_D"] = "Counting floor tilings. Brute force"
notes["abc196_E"] = "Function Composition"
notes["abc196_F"] = "Convolution"
notes["abc197_E"] = "Simple O(N) DP"
notes["abc197_F"] = "Graph synthesis then Dijkstra"
notes["abc198_D"] = "Alphametic addition solver"
notes["abc198_E"] = "Tree DFS w/ entry and exit modification"
notes["abc198_F"] = "Burnside's Lemma"
notes["abc199_D"] = "Smart exhaustive search"
notes["abc199_E"] = "Subset DP"
notes["abc199_F"] = "Matrix exponentiation"
notes["abc200_D"] = "Subset sums.  Birthday paradox."
notes["abc200_E"] = "Counting DP w/ Cumulative Sums"
notes["abc200_F"] = "Wraparound trick. Simple inversion counting."
notes["abc201_D"] = "Simple minimax DP"
notes["abc201_E"] = "Counting DP w/ Cumulative Sums"
notes["abc201_F"] = "Wraparound trick. Simple inversion counting."
notes["abc202_E"] = "Tree DFS, cumulative data structure"
notes["abc202_F"] = "Complex convex hull DP"
notes["abc203_D"] = "Binary search, Cumulative 2D Sums"
notes["abc203_E"] = "Coordinate compression.  Serial processing"
notes["abc203_F"] = "DP"
notes["abc204_E"] = "Good Dijkstra problem with a twist"
notes["abc204_F"] = "Floor tiling.  Matrix exponentiation"
notes["abc205_E"] = "Complementary counting"
notes["abc205_F"] = "Max flow"
notes["abc206_E"] = "Simple inclusion/exclusion DP"
notes["abc206_F"] = "DP, Nimbers"
notes["abc207_D"] = "Pointset equivalence modulo rotations/translations"
notes["abc207_E"] = "DP"
notes["abc207_F"] = "Tree DP.  Convolution."
notes["abc208_D"] = "Floyd-Warshall"
notes["abc208_E"] = "Digit DP"
notes["abc208_F"] = "DP. Lagrange Interpolation."
notes["abc209_E"] = "Game on directed graph.  Dependency counting"
notes["abc209_F"] = "Counting permutations sastifying pairwise constraints.  DP"
notes["abc210_D"] = "DP"
notes["abc210_E"] = "Faster Kruskal MST"
notes["abc210_F"] = "Two-sat with one-hot trick"
notes["abc211_E"] = "DP on 8x8 graphs"
notes["abc211_F"] = "Lazy segtree. Tricky 2 pass. Ordered rectangle processing"
notes["abc212_E"] = "DP. Complementary counting."
notes["abc212_F"] = "DSU. Offline query processing."
notes["abc212_G"] = "Sum of orders of elements of cyclic group (or OEIS)"
notes["abc212_H"] = "Hadamard transform"
notes["abc213_E"] = "0/1 BFS"
notes["abc213_F"] = "Suffix array. Longest Common Prefix. "
notes["abc213_G"] = "Graph counting. Subset DP"
notes["abc213_H"] = "Divide & Conquer FFT"
notes["abc214_D"] = "DSU on tree"
notes["abc214_E"] = "Minheap. Sorted interval processing."
notes["abc214_F"] = "Subsequence DP"
notes["abc214_G"] = "Combinatorics. DSU. Convolution. Permutation Cycles."
notes["abc214_H"] = "Kosaraju. Min cost flow. Hard."
notes["abc215_E"] = "Subset DP"
notes["abc215_F"] = "Bin Search. Array suffix max/min."
notes["abc215_G"] = "Combinatorics. N^3/2 DP"
notes["abc215_H"] = "Hall's marriage theorem.  Sum over subsets/supersets.  Mobius Transform"
notes["abc216_D"] = "Straightforward greedy matching w/ some bookkeeping"
notes["abc216_E"] = "Efficient counting"
notes["abc216_F"] = "Simple DP"
notes["abc216_G"] = "Fenwick tree plus stack.  Interval processing."
notes["abc216_H"] = "LGV (Lindstrom-Gessel-Viennot) lemma.  Determinants"
notes["abc217_E"] = "Deque + Minheap.  Online queries."
notes["abc217_F"] = "Straightforward N^2 DP"
notes["abc217_G"] = "Inclusion-exclusion calculation.  Combinatorics. (official solution is more elegant, but harder to see)"
notes["abc217_H"] = "Slope Trick.  Minheap + Maxheap"
notes["abc218_C"] = "Point set equivalence modulo translation and 90 degree rotations"
notes["abc218_E"] = "Complementary counting with modified Kruskal's"
notes["abc218_F"] = "BFS N+1 times"
notes["abc218_G"] = "Median data structure with additions and removals (RB multiset or 4 priority queues). Tree DP. Minimax."
notes["abc218_H"] = "Nice divide and conquer problem"
notes["abc219_E"] = "Polygon construction, DSU"
notes["abc219_F"] = "Stepping a shifted pattern, Equivalence classes"
notes["abc219_G"] = "Large tree, sqrt(N) strategy"
notes["abc219_H"] = "Hard DP.  Pypy3 solution is TLE (same alg as Go solution).  Need to work on it a bit (probably need to flatten multi-dim dp array)"
notes["abc220_E"] = "Reasonably straightforward O(N) sum calculation.  Cumulative sums."
notes["abc220_F"] = "Double-dfs tree-DP"
notes["abc220_G"] = "Counting isocelese trapezoids. Nice trick."
notes["abc220_H"] = "Meet in the middle. Hadamard like problem (**Need to dig into this more to understand -- solution here is simply transcribed into Go)"
notes["abc221_E"] = "Coordinate compression.  Fenwick tree."
notes["abc221_F"] = "Tree diameter. DFS. Even/Odd Casework. Combinatorics"
notes["abc221_G"] = "Subset sum. Using a large bitset helps. Processing the parents as we find them vs. keeping the full history" + \
                    " and tracing back from the end keeps the space requirement down."
notes["abc221_H"] = "Tricky DP to see."
notes["abc222_E"] = "Tree DP plus subset sum"
notes["abc222_F"] = "Another tree DP"
notes["abc222_G"] = "Totient function. Elementary group theory"
notes["abc222_H"] = "Generating functions. Lagrange inversion theorem. P-recursive"
notes["abc223_E"] = "Rectangle placement.  Simple casework."
notes["abc223_F"] = "Parenthesis sequence.  Lazy Segtree."
notes["abc223_G"] = "Tree DP.  Rerooting"
notes["abc223_H"] = "2^n xor linear algebra.  Basis finding"
notes["abc224_E"] = "Simple DP-lite after sorting"
notes["abc224_F"] = "Another simple DP"
notes["abc224_G"] = "Expected value"
notes["abc224_H"] = "Linear Programming. Duality. Min-cost flow problem conversion."
notes["abc225_E"] = "Geometry.  Sorting angles."
notes["abc225_F"] = "Tricky string DP."
notes["abc225_G"] = "Tricky max flow problem (hard to see)."
notes["abc225_H"] = "Generating functions. Convolution. Also hard to see."
notes["abc226_E"] = "Counting rings.  1:1 correspondance."
notes["abc226_F"] = "Integer partitions.  Combinatorics."
notes["abc226_G"] = "Mildly tricky greedy problem."
notes["abc226_H"] = "Modular fraction DP.  Not as bad as it looks."
notes["abc227_E"] = "Tricky string DP.  O(N^5)."
notes["abc227_F"] = "Traversal DP.  Also O(N^5)."
notes["abc227_G"] = "Sieving.  Prime numbers"
notes["abc227_H"] = "Interesting linear algebra problem. Construction. Non-unique solutions."
notes["abc228_E"] = "Modulo arithmetic. Fermat's little theorem. Fast exponentiation."
notes["abc228_F"] = "2-d cumulative sums.  Sliding window min."
notes["abc228_G"] = "Tricky DP with subsets."
notes["abc228_H"] = "Convex hull trick."
notes["abc229_E"] = "Reverse query order. DSU"
notes["abc229_F"] = "Simple DP."
notes["abc229_G"] = "Binary search.  Sliding window."
notes["abc229_H"] = "Partisan game theory.  Surreal numbers."

def parseDifficulty(s1,s2) :
    probsByContest = {}
    numprobs = 0
    pyprogress = 0
    goprogress = 0
    numright = {}
    diffrating = {}
    alph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    rows1 = s1.split('\n')
    rows2 = s2.split('\n')

    for (r1,r2) in zip(rows1,rows2) :
        terms1 = r1.split(',')
        terms2 = r2.split(',')
        if len(terms1) < 3 : continue
        problems = []
        contest = terms1[0]
        for (i,(nr,rat)) in enumerate(zip(terms1[1:],terms2[1:])) :
            problem = f"{contest}_{alph[i]}"
            numprobs += 1
            if os.path.exists(f"python/{contest}/{problem}.py")       : pyprogress += 1
            if os.path.exists(f"go/{contest}/{problem}/{problem}.go") : goprogress += 1
            numright[problem] = nr
            diffrating[problem] = rat
            problems.append(problem)
        probsByContest[contest] = problems
    contests = [x for x in probsByContest]
    contests.sort()
    ##print(numright)
    ##print(diffrating)
    return contests,probsByContest,numprobs,pyprogress,goprogress,numright,diffrating

def doHeader(numprobs,pyprogress,goprogress,fp) :
    p1 = f"# m1sterzer0 Atcoder ABC Solutions ![Language](https://img.shields.io/badge/language-Python-orange.svg) ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) ![PythonProgress](https://img.shields.io/badge/PythonProgress-{pyprogress}%20%2F%20{numprobs}-ff69b4.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-{goprogress}%20%2F%20{numprobs}-ff69b4.svg)"
    p2 = "These are the solutions for the Atcoder beginner contests (ABC).  I find the Atcoder beginner contests very educational.  There is a large difficulty spread in the problems so that you can ramp, and the 500/600 point solutions often introduce new (even advanced) concepts in a reasonably straightforward way."
    p3 = "`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below."
    p4 = "`NOTE`: Problem difficulty ratings in the table below were obtained from kenkoooo.com."

    print(p1,file=fp)
    print("",file=fp)
    print(p2,file=fp)
    print("",file=fp)
    print(p3,file=fp)
    print("",file=fp)
    print(p4,file=fp)

def doTable(contests,fp) :
    print('## Contest Shortcuts\n|     |     |     |     |     |\n| --- | --- | --- | --- | --- |',file=fp)
    ptr = 0
    while (ptr < len(contests)) :
        print("|",end='',file=fp)
        for _ in range(5) :
            elem = f" [{contests[ptr]}](https://atcoder.jp/contests/{contests[ptr]}); [sol](#{contests[ptr]}-Solutions) |" if ptr < len(contests) else " |"
            ptr += 1
            print(elem,end='',file=fp)
        print("\n",end='',file=fp)
    print("",file=fp)

def doSolutions(contests,probsByContest,numright,diffrating,notes,fp) :
    for contest in contests[::-1] :
        print(f"## {contest} Solutions",file=fp)
        print("| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |",file=fp)
        print("| ------- | ------- | ----------: | ----------: | --------- | ----- |",file=fp)
        for prob in probsByContest[contest] :
            sol = ""
            if os.path.exists(f"python/{contest}/{prob}.py") :
                sol += f" [pypy3](./python/{contest}/{prob}.py)"
            if os.path.exists(f"go/{contest}/{prob}/{prob}.go") :
                sol += f" [go](./go/{contest}/{prob}/{prob}.go)"
            pnotes = notes[prob] if prob in notes else ""
            print(f"| [{contest}](http:/atcoder.jp/contests/{contest}) | [{prob}](http:/atcoder.jp/contests/{contest}/tasks/{prob}) | {numright[prob]} | {diffrating[prob]} | {sol} | {pnotes} |",file=fp)
        print(f"",file=fp)


if __name__ == "__main__" :
    contests,probsByContest,numprobs,pyprogress,goprogress,numright,diffrating = parseDifficulty(solversData,problemRatingData)
    with open("README.md","wt") as fp :
        doHeader(numprobs,pyprogress,goprogress,fp)
        doTable(contests,fp)
        doSolutions(contests,probsByContest,numright,diffrating,notes,fp)
    Path(f"README.md").touch()
