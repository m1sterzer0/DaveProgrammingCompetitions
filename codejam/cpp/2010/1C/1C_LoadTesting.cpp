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
        ll L,P,C; cin >> L >> P >> C;
        ll ival(1),x(L);
        while (C*x < P) ival++,x*=C;
        ll ans(0),n(1);
        while (n<ival) ans++,n*=2;
        printf("Case #%lld: %lld\n",tt,ans);
    }
}
