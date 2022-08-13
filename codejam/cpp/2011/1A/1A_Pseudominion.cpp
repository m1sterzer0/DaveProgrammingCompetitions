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
struct state { ll t,nc,tidx,c1idx,c2idx; };
bool operator<(const state a, const state b) { 
    return a.t < b.t || a.t == b.t && (
           a.nc < b.nc || a.nc == b.nc && ( 
           a.tidx < b.tidx || a.tidx == b.tidx && (
           a.c1idx < b.c1idx || a.c1idx == b.c1idx && a.c2idx < b.c2idx)));
}
struct card { ll idx,c,s,t; };
bool operator<(const card &a, const card &b) { return a.s < b.s; } 
ll c0sum[81][81];

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        vector<card> cards;
        ll N; cin >> N;
        FOR(i,N) { ll c,s,t; cin >> c >> s >> t; cards.push_back(card{i,c,s,t}); }
        ll M; cin >> M;
        FOR(i,M) { ll c,s,t; cin >> c >> s >> t; cards.push_back(card{i+N,c,s,t}); }
        vector<card> T,C0,C1,C2;
        for (auto &c : cards) {
            if (c.t > 0) T.push_back(c); else if (c.c == 0) C0.push_back(c); else if (c.c == 1) C1.push_back(c); else C2.push_back(c);
        }
        FOR(j,81) {
            vector<card> ca;
            for (auto &c : C0) { if (c.idx < j) ca.push_back(c); }
            sort(ca.begin(),ca.end());
            reverse(ca.begin(),ca.end());
            ll s = 0;
            c0sum[0][j] = 0; FOR(i,ca.size()) { s += ca[i].s; c0sum[i+1][j] = s; }
            for (ll i = ca.size()+1;i<=80;i++) { c0sum[i][j] = c0sum[ca.size()][j]; }
        }
        map<state,ll> dict;
        function<void(state)> dfs;
        dfs = [&](state s) {
            if (dict.count(s) > 0) return;
            if (s.t == 0) { dict[s] = 0; return; }
            ll v = 0;
            if (s.tidx+1 < T.size() && T[s.tidx+1].idx < s.nc) {
                card cc = T[s.tidx+1];
                state ns = state{s.t-1+cc.t,min(N+M,s.nc+cc.c),s.tidx+1,s.c1idx,s.c2idx};
                dfs(ns); v = dict[ns]+cc.s;
            } else {
                v = c0sum[min(80LL,s.t)][s.nc];
                if (s.c2idx+1 < C2.size() && C2[s.c2idx+1].idx < s.nc) {
                    card cc = C2[s.c2idx+1];
                    state ns1 = state{s.t,s.nc,s.tidx,s.c1idx,s.c2idx+1};
                    state ns2 = state{s.t-1+cc.t,min(N+M,s.nc+cc.c),s.tidx,s.c1idx,s.c2idx+1};
                    dfs(ns1); v = max(v,dict[ns1]);
                    dfs(ns2); v = max(v,cc.s+dict[ns2]);
                } else if (s.c1idx+1 < C1.size() && C1[s.c1idx+1].idx < s.nc) {
                    card cc = C1[s.c1idx+1];
                    state ns1 = state{s.t,s.nc,s.tidx,s.c1idx+1,s.c2idx};
                    state ns2 = state{s.t-1+cc.t,min(N+M,s.nc+cc.c),s.tidx,s.c1idx+1,s.c2idx};
                    dfs(ns1); v = max(v,dict[ns1]);
                    dfs(ns2); v = max(v,cc.s+dict[ns2]);
                }
            }
            dict[s] = v;
        };
        state start = state{1,N,-1,-1,-1};
        dfs(start);
        ll ans = dict[start];
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

