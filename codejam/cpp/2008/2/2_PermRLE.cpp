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

ll darr[16][16];
ll earr[16][16];
ll dp [1LL<<16][16];

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll K; string S; cin >> K >> S;
        ll N = len(S);
        FOR(i,K) {
            FOR(j,K) {
                if (i == j) continue;
                ll cnt(0), endtax(0);
                for (ll k=0;k<N;k+=K) { if (S[k+i] != S[k+j]) cnt++; if (k+K<N && S[k+j] != S[k+K+i]) endtax++; }
                darr[i][j] = cnt; earr[i][j] = endtax;
            }            
        }
        ll ans = 1LL << 61; ll last = (1 << K)-1;
        FOR(i,K) {
            for (ll bm=1;bm<=last;bm++) {
                if ((bm & (1 << i)) == 0) continue;
                if ((bm^(1LL << i)) == 0) { dp[bm][i] = 0; continue; }
                FOR(j,K) {
                    if (i==j || (bm & (1LL<<j)) == 0) continue;
                    if ((bm ^ (1LL<<i) ^ (1LL<<j)) == 0) { dp[bm][j] = darr[i][j]; continue; }
                    ll res = 1LL<<61;
                    FOR(k,K) {
                        if (i==k || j==k || (bm & (1<<k)) == 0) continue;
                        res = min(res,dp[bm ^ (1LL<<j)][k] + darr[k][j]);
                    }
                    dp[bm][j] = res;
                }

            }
            FOR(j,K) {
                if (i==j) continue;
                ans = min(ans,dp[last][j]+earr[i][j]);
            }
        }
        ans++; // Have to count first group
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

