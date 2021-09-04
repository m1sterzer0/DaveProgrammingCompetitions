import os
from pathlib import Path

difficulty_data = '''
abc160,9557,9427,7443,3417,2721,339
abc161,9784,8263,7740,2865,682,1117
abc162,10333,9924,8304,4199,850,680
abc163,11284,10786,8976,3663,380,107
abc164,11148,10415,9553,1926,497,41
abc165,11225,9994,2514,5554,975,596
abc166,11540,10088,7970,4999,2874,863
abc167,11507,10836,5652,4666,1393,415
abc168,10686,104666,7598,3856,467,85
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
notes["abc164_E"] = "Interesting Dijkstra"
notes["abc164_F"] = "Bitwise, Logical constraint solving"

def parseDifficulty(s) :
    probsByContest = {}
    numprobs = 0
    pyprogress = 0
    goprogress = 0
    difficulty = {}
    alph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    rows = s.split('\n')
    for r in rows :
        terms = r.split(',')
        if len(terms) < 3 : continue
        problems = []
        contest = terms[0]
        for (i,d) in enumerate(terms[1:]) :
            problem = f"{contest}_{alph[i]}"
            numprobs += 1
            if os.path.exists(f"python/{contest}/{problem}.py")       : pyprogress += 1
            if os.path.exists(f"go/{contest}/{problem}/{problem}.go") : goprogress += 1
            difficulty[problem] = d
            problems.append(problem)
        probsByContest[contest] = problems
    contests = [x for x in probsByContest]
    contests.sort()
    return contests,probsByContest,numprobs,pyprogress,goprogress,difficulty

def doHeader(numprobs,pyprogress,goprogress,fp) :
    p1 = f"# m1sterzer0 Atcoder ABC Solutions ![Language](https://img.shields.io/badge/language-Python-orange.svg) ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) ![PythonProgress](https://img.shields.io/badge/PythonProgress-{pyprogress}%20%2F%20{numprobs}-ff69b4.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-{goprogress}%20%2F%20{numprobs}-ff69b4.svg)"
    p2 = "These are the solutions for the Atcoder beginner contests (ABC).  I find the Atcoder beginner contests very educational.  There is a large difficulty spread in the problems so that you can ramp, and the 500/600 point solutions often introduce new (even advanced) concepts in a reasonably straightforward way."
    p3 = "`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below."
    print(p1,file=fp)
    print("",file=fp)
    print(p2,file=fp)
    print("",file=fp)
    print(p3,file=fp)

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

def doSolutions(contests,probsByContest,difficulty,notes,fp) :
    for contest in contests[::-1] :
        print(f"## {contest} Solutions",file=fp)
        print("| Contest | Problem | Num Correct | Solutions | Notes |",file=fp)
        print("| ------- | ------- | ----------: | --------- | ----- |",file=fp)
        for prob in probsByContest[contest] :
            sol = ""
            if os.path.exists(f"python/{contest}/{prob}.py") :
                sol += f" [pypy3](./python/{contest}/{prob}.py)"
            if os.path.exists(f"go/{contest}/{prob}/{prob}.go") :
                sol += f" [go](./go/{contest}/{prob}/{prob}.go)"
            pnotes = notes[prob] if prob in notes else ""
            print(f"| [{contest}](http:/atcoder.jp/contests/{contest}) | [{prob}](http:/atcoder.jp/contests/{contest}/tasks/{prob}) | {difficulty[prob]} | {sol} | {pnotes} |",file=fp)
        print(f"",file=fp)


if __name__ == "__main__" :
    contests,probsByContest,numprobs,pyprogress,goprogress,difficulty = parseDifficulty(difficulty_data)
    with open("README.md","wt") as fp :
        doHeader(numprobs,pyprogress,goprogress,fp)
        doTable(contests,fp)
        doSolutions(contests,probsByContest,difficulty,notes,fp)
    Path(f"README.md").touch()
