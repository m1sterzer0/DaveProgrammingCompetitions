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
        ll R,k,N; cin >> R >> k >> N;
        vi G(N); for (auto &g : G) cin >> g;
        vi histn(N,-1), histc(N,-1);
        ll ans(0), nrides(0), curs(0);
        bool processedLoop = false;
        while (nrides < R) {
            ll idx(0), cur(0);
            while (idx<N && cur + G[(curs+idx)%N] <= k) cur += G[(curs+idx)%N], idx++;
            curs = (curs+idx)%N; ans += cur; nrides++;
            if (processedLoop) continue;
            if (histn[curs] == -1) { histn[curs] = nrides; histc[curs]=ans; continue; }
            ll loopsize = nrides-histn[curs];
            ll nloops = (R-nrides) / loopsize;
            ans += nloops * (ans - histc[curs]);
            nrides += nloops * loopsize;
            processedLoop = true; 
        }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

