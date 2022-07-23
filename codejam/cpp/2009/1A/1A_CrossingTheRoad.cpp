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

struct dnode { ll d,i,j; };
bool gt(dnode a, dnode b) { return a.d > b.d; } // stl crap
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N,M; cin >> N >> M;
        auto S = twodi(N,M,0); auto W = twodi(N,M,0); auto T = twodi(N,M,0);
        FOR(i,N) FOR(j,M) cin >> S[i][j] >> W[i][j] >> T[i][j];
        ll inf = 1LL << 62;
        auto darr = twodi(2*N,2*M,inf);
        priority_queue<dnode,vector<dnode>,function<bool(dnode,dnode)>> mh(gt);
        mh.push({0,2*N-1,0});
        while (!mh.empty()) {
            auto d = mh.top().d; auto i = mh.top().i; auto j = mh.top().j; mh.pop();
            if (darr[i][j] != inf) continue;
            darr[i][j] = d; auto s = S[i/2][j/2]; auto w = W[i/2][j/2]; auto t = T[i/2][j/2];
            t = t % (s+w) - (s+w);
            auto nstime = (d-t) % (s+w) < s  ? d+1 : t + (d-t)/(s+w)*(s+w) + (s+w) + 1;
            auto ewtime = (d-t) % (s+w) >= s ? d+1 : t + (d-t)/(s+w)*(s+w) + s + 1;
            if ((i&1) == 0) { if (i-1 >= 0) mh.push({d+2,i-1,j}); mh.push({nstime,i+1,j});  }
            if ((i&1) == 1) { mh.push({nstime,i-1,j}); if (i+1 < 2*N) mh.push({d+2,i+1,j}); }
            if ((j&1) == 0) { if (j-1 >= 0) mh.push({d+2,i,j-1}); mh.push({ewtime,i,j+1});  }
            if ((j&1) == 1) { mh.push({ewtime,i,j-1}); if (j+1 < 2*M) mh.push({d+2,i,j+1}); }
        }
        printf("Case #%lld: %lld\n",tt,darr[0][2*M-1]);
    }
}

