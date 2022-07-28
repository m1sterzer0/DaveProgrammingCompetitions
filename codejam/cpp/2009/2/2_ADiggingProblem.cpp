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
struct Dnode {ll d,i,j,l,r;};
struct Dstate {ll i,j,l,r;};
bool operator <(const Dstate a, const Dstate b) { return a.i < b.i || a.i == b.i && (a.j < b.j || a.j == b.j && (a.l < b.l || a.l == b.l && (a.r < b.r))); }
bool operator <(const Dnode a, const Dnode b) { return a.d > b.d; }

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll R,C,F; cin >> R >> C >> F; vector<string> grid(R); FOR(i,R) cin >> grid[i];
        vector<string> G;
        FOR(i,R) { G.push_back("#"+grid[i]+"#"); }; G.push_back(""); FOR(i,C+2) G[R].push_back('#');
        vector<vi> fall(R+1); FOR(i,R+1) fall[i].resize(C+2,0);
        for (ll i=R-1; i>=0; i--) FOR(j,C+2) fall[i][j] = (G[i+1][j] == '#') ? 0 : 1 + fall[i+1][j];
        map<Dstate,bool> dmap;
        priority_queue<Dnode> mh;
        mh.push({0,0,1,-1,-1});
        vi buf;
        ll ans = -1;
        while (!mh.empty()) {
            auto xx = mh.top(); auto d = xx.d; auto i = xx.i; auto j = xx.j; auto lup = xx.l; auto rup = xx.r; mh.pop();
            if (i == R-1) {ans = d; break; }
            if (dmap.count({i,j,lup,rup}) > 0) continue;
            dmap[{i,j,lup,rup}] = true;
            auto lx = j; while ((lup <= lx-1 && lx-1 <= rup || G[i][lx-1] == '.') && G[i+1][lx-1] == '#') lx--;
            auto rx = j; while ((lup <= rx+1 && rx+1 <= rup || G[i][rx+1] == '.') && G[i+1][rx+1] == '#') rx++;
            // Check for jumping off the ends
            buf.clear(); buf.push_back(lx-1); buf.push_back(rx+1);
            for (auto &x : buf) {
 			    if (G[i][x] == '#' && (x < lup || x > rup) || G[i+1][x] == '#') continue;
			    auto fdist = 1 + fall[i+1][x];
			    if (fdist > F) continue;
			    mh.push({d,i+fdist,x,-1,-1});               
            }
            vi holes;
            if (lx == rx) continue;
            for (auto ii=lx;ii<=rx;ii++) {
                for (auto jj=ii;jj<=rx;jj++) {
                    holes.clear();
                    if (ii == jj) { holes.push_back(ii); } 
                    else { if (ii > lx) holes.push_back(ii); if (jj < rx) holes.push_back(jj); }
                    for (auto &x : holes) {
                        auto fdist = 1 + fall[i+1][x];
                        if (fdist > F) continue;
                        if (fdist == 1) mh.push({d+jj-ii+1,i+fdist,x,ii,jj}); 
                        else if (ii == jj) mh.push({d+jj-ii+1,i+fdist,x,-1,-1});
                    }
                }
            }
        }
        if (ans == -1) printf("Case #%lld: No\n",tt); else printf("Case #%lld: Yes %lld\n",tt,ans);
	}
}
