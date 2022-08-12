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
        ll opos(1),ot(0),bpos(1),bt(0);
        ll n; cin >> n;
        FOR(i,n) {
            string x; ll newpos; cin >> x >> newpos;
            if (x[0] == 'O') { ot = max(bt+1,ot+abs(newpos-opos)+1); opos = newpos; }
            else             { bt = max(ot+1,bt+abs(newpos-bpos)+1); bpos = newpos; }
        }
        ll ans = max(ot,bt);
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

