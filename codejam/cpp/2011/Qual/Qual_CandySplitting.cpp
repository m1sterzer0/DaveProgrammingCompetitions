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
        ll N; cin >> N; vi C(N); FOR(i,N) cin >> C[i];
        ll xorsum = 0; for (auto c : C) xorsum ^= c;
        if (xorsum != 0) {
            printf("Case #%lld: NO\n",tt);
        } else {
            sort(C.begin(),C.end());
            ll sumc = 0; for (auto c : C) sumc += c;
            ll ans = sumc - C[0];
            printf("Case #%lld: %lld\n",tt,ans);
        }
    }
}

