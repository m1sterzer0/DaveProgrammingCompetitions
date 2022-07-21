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

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll M,V; cin >> M >> V;
        vi G(M+1,-1), C(M+1,-1), I(M+1,-1);
        for (ll i = 1; i <= (M-1)/2; i++) cin >> G[i] >> C[i];
        for (ll i = (M-1)/2+1; i <= M; i++ ) cin >> I[i];
        for (ll i = (M-1)/2; i >= 1; i-- ) I[i] = G[i] == 0 ? I[2*i] | I[2*i+1] : I[2*i] & I[2*i+1];
        ll inf = 1LL << 61;
        function<ll(ll)> dfs;
        dfs = [&](ll n) -> ll {
            if (V == I[n]) return 0;
            if (G[n] == -1) return inf;
            auto a = dfs(2*n); auto b = dfs(2*n+1);
            return G[n] != V ? min(a,b) : C[n] == 0 ? min(inf,a+b) : min(inf,1+min(a,b));
        };
        auto ans = dfs(1);
        if (ans == inf) {
            printf("Case #%lld: IMPOSSIBLE\n",tt);
        }  else {
            printf("Case #%lld: %lld\n",tt,ans);
        }
    }
}

