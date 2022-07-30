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

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll D,I,M,N; cin >> D >> I >> M >> N;
        vi A(N); FOR(i,N) cin >> A[i];
        vi dp(256,0),ndp(256,0);
        for (auto a : A) {
            FOR(i,256) ndp[i] = dp[i] + D;
            FOR(st,256) {
                FOR(en,256) {
                    if (M == 0) {
                        if (st != en) continue;
                        ndp[en] = min(ndp[en],dp[st]+abs(en-a));
                    } else {
                        ll inserts = max(0LL,(abs(en-st)+M-1)/M-1);
                        ndp[en] = min(ndp[en],dp[st]+inserts*I+abs(en-a));
                    }
                }
            }
            swap(dp,ndp);
        }
        ll ans = *min_element(dp.begin(),dp.end());
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

