#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef vector<ll> vi;
typedef pair<ll,ll> pi;
#define FOR(i,a) for (ll i = 0; i < (a); i++)
#define len(x) (ll) x.size()
const ll INF = 1LL << 62;
const ll MOD = 100003;
//const ll MOD = 998244353;
const double PI = 4*atan(double(1.0));
vector<vector<ll>> twodi(ll N, ll M, ll v) { vector<vector<ll>> res(N); for (auto &vv : res) vv.resize(M,v); return res; }
ll dp[501][501];
ll comb[501][501];
ll ansarr[501];

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    // BEGIN PREWORK
    FOR(i,501) FOR(j,i+1) comb[i][j] = (j==0||j==i) ? 1 : (comb[i-1][j-1]+comb[i-1][j]) % MOD;
    FOR(i,501) FOR(j,501) dp[i][j] = 0;
    for (ll i=2; i<=500; i++) {
        dp[i][1] = 1;
        for (ll k=2;k<i;k++) {
            for (ll l=1;l<k;l++) dp[i][k] += (comb[i-k-1][k-l-1] * dp[k][l]) % MOD;
            dp[i][k] %= MOD;
        }
    }
    for (ll i=0;i<=500;i++) {
        ansarr[i] = 0;
        for (ll j=0;j<=500;j++) ansarr[i] += dp[i][j];
        ansarr[i] %= MOD;
    }
    // END PREWORK
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N; cin >> N;
        auto ans = ansarr[N];
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

