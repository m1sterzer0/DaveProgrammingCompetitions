import os
import sys
import numpy as np
    
    
def solve(inp):
    r, c, k = inp[:3]
    items = np.zeros((r + 1, c + 1), dtype=np.int64)
    rrr = inp[3::3]
    ccc = inp[4::3]
    vvv = inp[5::3]
    for r_, c_, v in zip(rrr, ccc, vvv):
        items[r_][c_] = v
    
    dp = np.zeros((r + 1, c + 1, 4), dtype=np.int64)
    
    for i in range(1, r + 1):
        for j in range(1, c + 1):
            up = dp[i - 1][j][3]
            for w in range(4):
                dp[i][j][w] = max(dp[i][j - 1][w], up)
            v = items[i][j]
            if v == 0:
                continue
            for w in range(2, -1, -1):
                dp[i][j][w + 1] = max(dp[i][j][w + 1], dp[i][j][w] + v)
    
    return dp[-1][-1][3]
    
    
if sys.argv[-1] == 'ONLINE_JUDGE':
    from numba.pycc import CC
    cc = CC('my_module')
    cc.export('solve', '(i8[:],)')(solve)
    cc.compile()
    exit()
else :
    # noinspection PyUnresolvedReferences
    from my_module import solve
    inp = np.fromstring(sys.stdin.read(), dtype=np.int64, sep=' ')
    ans = solve(inp)
    print(ans)
    
#if os.name == 'posix':
#else:
#    from numba import njit
#    solve = njit('(i8[:],)', cache=True)(solve)
#    print('compiled', file=sys.stderr)

#from numba import njit
#solve = njit('(i8[:],)', cache=True)(solve)
#print('compiled', file=sys.stderr)
