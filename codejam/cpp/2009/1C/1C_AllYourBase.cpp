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
        string S; cin >> S; vi vlookup(256,-1); ll numuniq = 0;
        for (auto &c : S) {
            if (vlookup[c] != -1) continue;
            vlookup[c] = numuniq == 0 ? 1 : numuniq == 1 ? 0 : numuniq; numuniq++;
        }
        ll ans = 0; ll pv = 1; ll base = max(2LL,numuniq);
        for (ll i = len(S)-1; i>=0; i--) { ans += pv * vlookup[S[i]]; pv *= base; }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

