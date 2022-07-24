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

struct midx { ll i,j,v; };
bool operator < (const midx a, const midx b) { return a.i < b.i || a.i == b.i && (a.j < b.j || a.j == b.j && (a.v < b.v)); }

struct mval {ll l; string s;};
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll W,Q; cin >> W >> Q;
        vector<string> bd(W); FOR(i,W) { cin >> bd[i]; }
        vi QQ(Q); FOR(i,Q) { cin >> QQ[i]; }
        vector<mval> ansarr(251); FOR(i,251) ansarr[i] = {0,""};
        ll numleft = 0; for (auto &q : QQ) if (ansarr[q].l == 0) { ansarr[q].l = -1; numleft++; }
        map<midx,mval> lkup;
        vector<midx> que,nque;
        FOR(i,W) {
            FOR(j,W) {
                if (bd[i][j] != '+' && bd[i][j] != '-') { 
                    ll v = ll(bd[i][j]-'0');
                    midx m = {i,j,v};
                    mval vv = {1,bd[i].substr(j,1)};
                    if (v >= 0 && v <= 250 && ansarr[v].l == -1) { numleft--; ansarr[v] = vv; } 
                    lkup[m] = {1,bd[i].substr(j,1)};
                    que.push_back(m);
                }
            }
        }
        vector<pair<ll,ll>> deltas{{-1,0},{1,0},{0,-1},{0,1}};
        ll rnd = 1;
        while (numleft > 0) {
            rnd++;
            for (auto &m : que) {
                for (auto &p1 : deltas) {
                    auto i1 = m.i+p1.first; auto j1 = m.j+p1.second;
                    if (i1 < 0 || j1 < 0 || i1 >= W || j1 >= W) continue;
                    ll sgn = bd[i1][j1] == '+' ? 1 : -1;
                    for (auto &p2 : deltas) {
                        auto i2 = i1+p2.first; auto j2 = j1+p2.second;
                        if (i2 < 0 || j2 < 0 || i2 >= W || j2 >= W) continue;
                        ll v = bd[i2][j2]-'0';
                        midx m2({i2,j2,m.v+sgn*v});
                        auto it = lkup.find(m2);
                        if (it != lkup.end() && it->second.l < rnd) continue;
                        string path(lkup[m].s); path.push_back(bd[i1][j1]); path.push_back(bd[i2][j2]);
                        if (it != lkup.end() && it->second.l == rnd && path >= it->second.s) continue;
                        if (it == lkup.end()) { nque.push_back(m2); }
                        mval v2({rnd,path});
                        lkup[m2] = v2;
                        if (m2.v >= 0 && m2.v <= 250 && (ansarr[m2.v].l == -1 || ansarr[m2.v].l == rnd && path < ansarr[m2.v].s)) {
                            if (ansarr[m2.v].l == -1) numleft--;
                            ansarr[m2.v] = v2;
                        }
                    }
                }
            }
            swap(que,nque); nque.clear();
        }
        printf("Case #%lld:\n",tt);
        for (auto &qq : QQ) { printf("%s\n",ansarr[qq].s.c_str()); }
    }
}



