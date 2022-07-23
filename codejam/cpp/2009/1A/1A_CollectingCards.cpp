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
vector<vector<double>> twodf(ll N, ll M, double v) { vector<vector<double>> res(N); for (auto &vv : res) vv.resize(M,v); return res; }

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    auto comb = twodf(41,41,0.00);
    for (ll i=0;i<=40;i++) for (ll j=0;j<=i;j++) comb[i][j] = (j==0 || j==i) ? 1.0 : comb[i-1][j-1]+comb[i-1][j];
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll C,N; cin >> C >> N;
        vector<double> dp(C+1);
        auto denom = comb[C][N];
        for (ll i=1;i<=C;i++) {
            auto coeff = denom-comb[C-i][N];
            auto num = comb[C][N];
            for (ll j=1; j<=i; j++) {if (j > N) break; num += comb[i][j] * comb[C-i][N-j] * dp[i-j]; }
            dp[i] = num/coeff;
        }
        printf("Case #%lld: %.17g\n",tt,dp[C]);
    }
}

