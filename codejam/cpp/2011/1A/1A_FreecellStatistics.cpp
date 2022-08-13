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
        ll N,Pd,Pg; cin >> N >> Pd >> Pg;
        string ans = "Possible";
        if (Pg == 0 && Pd != 0) { ans = "Broken"; }
        if (Pg == 100 && Pd != 100) { ans = "Broken"; }
        if (Pd > 0 && Pd < 100) {
            ll g = gcd(Pd,100LL); ll denom = 100LL / g;
            if (denom > N) { ans = "Broken"; }
        }
        printf("Case #%lld: %s\n",tt,ans.c_str());
    }
}

