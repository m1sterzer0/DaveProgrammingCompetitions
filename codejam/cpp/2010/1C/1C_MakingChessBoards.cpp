#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef vector<ll> vi;
typedef pair<ll,ll> pi;
#define FOR(i,a) for (ll i = 0; i < (a); i++)
#define len(x) (ll) x.size()
const ll INF = 1LL << 62;
const ll MOD = 1000000007;
//const ll MOD = 998244353;
const double PI = 4*atan(double(1.0));
vector<vector<ll>> twodi(ll N, ll M, ll v) { vector<vector<ll>> res(N); for (auto &vv : res) vv.resize(M,v); return res; }

struct Square { ll s,i,j; };
bool operator<(const Square &a, const Square &b) { return a.s < b.s || a.s==b.s && (a.i > b.i || a.i == b.i && a.j > b.j); }

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll M,N; cin >> M >> N; vector<vector<bool>> bd(M);
        FOR(i,M) {
            string s; cin >> s;
            for (auto c : s) { 
                ll v = ll(c-'0'); if (v > 9 || v < 0) v = ll(c-'A')+10;
                for (ll bm=8;bm>0;bm>>=1) { 
                    bd[i].push_back((v&bm) != 0);
                }
            }
        }
        vector<vi> dp(M); FOR(i,M) dp[i].resize(N,1);
        for (ll i=M-2;i>=0;i--) {
            for (ll j=N-2;j>=0;j--) {
                if (bd[i][j] == bd[i+1][j] || bd[i][j] == bd[i][j+1] || bd[i][j] != bd[i+1][j+1]) continue;
                dp[i][j] = 1 + min(min(dp[i][j+1],dp[i+1][j]),dp[i+1][j+1]);
            }
        }
        ll maxsquare = min(N,M);
        vi ansarr(maxsquare+1);
        priority_queue<Square> mh;
        for (ll i=0;i<M;i++) for (ll j=0;j<N;j++) mh.push({dp[i][j],i,j});
        while (!mh.empty()) {
            ll s = mh.top().s; ll ii = mh.top().i; ll jj = mh.top().j; mh.pop();
            if (dp[ii][jj] == 0) continue;
            if (dp[ii][jj] < s) { mh.push({dp[ii][jj],ii,jj}); continue; }
            ansarr[s]++;
            ll imin = max(0LL,ii-s+1);
            ll jmin = max(0LL,jj-s+1);
            for (ll i=ii+s-1;i>=imin;i--) {
                for (ll j=jj+s-1;j>=jmin;j--) {
                    if (i >= ii && j >= jj) { dp[i][j] = 0; continue; }
                    if (dp[i][j] <= 1) continue;
                    dp[i][j] = 1 + min(min(dp[i][j+1],dp[i+1][j]),dp[i+1][j+1]);
                }
            }
        }
        ll nans = 0; for (auto x : ansarr) if (x > 0) nans++;
        printf("Case #%lld: %lld\n",tt,nans);
        for (ll i=maxsquare;i>=0;i--) if (ansarr[i] > 0) printf("%lld %lld\n",i,ansarr[i]);
    }
}

