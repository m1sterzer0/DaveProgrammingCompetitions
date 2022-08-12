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

ll sb[2050][12];
ll lev[2050];

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll P; cin >> P; auto n = 1LL << P;
        vector<ll> G(n); FOR(i,n) { ll x; cin >> x; G[i] = P-x; }
        vector<ll> C(n-1); FOR(i,n-1) cin >> C[i]; C.push_back(0);
        reverse(C.begin(),C.end());
        reverse(G.begin(),G.end());
        ll inf = 1LL << 61;
        FOR(i,2*n) FOR(j,P+1) sb[i][j] = inf;
        FOR(i,2*n) lev[i] = 0;
        for (ll i=2*n-1;i>=1;i--) {
            if (i>=n) { 
                for(ll j=G[i-n];j<=P;j++) sb[i][j] = 0;
            } else {
                lev[i] = 1+lev[2*i];
                FOR(j,P+1-lev[i]) {
                    sb[i][j] = min(inf,min(C[i]+sb[2*i][j+1]+sb[2*i+1][j+1],sb[2*i][j]+sb[2*i+1][j]));
                    //printf("DBG sb[%lld][%lld] = %lld\n",i,j,sb[i][j]);
                } 
            }
        }
        auto ans = sb[1][0];
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

