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

vector<pi> HopcroftKarp(ll N1, ll N2, vector<vi> &adj) {
    ll mynil = N1+N2;
    vi pairu(N1,mynil),pairv(N2,mynil),dist(N1+N2+1,0);
    ll myinf = 1LL << 62; deque<ll> q;
    auto bfs = [&]() -> bool {
        FOR(u,N1) { if (pairu[u] == mynil) { dist[u] = 0; q.push_back(u); } else { dist[u] = myinf; } }
        dist[mynil] = myinf;
        while (!q.empty()) {
            ll u = q.front(); q.pop_front();
            if (u != mynil && dist[u] < dist[mynil]) {
                for (auto &v : adj[u]) { auto u2 = pairv[v]; if (dist[u2] == myinf) { dist[u2] = dist[u]+1; q.push_back(u2); } }
            }
        }
        return dist[mynil] != myinf;
    };
    function<bool(ll)> dfs;
    dfs = [&](ll u) -> bool {
        if (u == mynil) return true;
        for (auto &v : adj[u]) {
            auto u2 = pairv[v]; if (dist[u2] == dist[u]+1 && dfs(u2)) { pairv[v] = u; pairu[u] = v; return true; }
		}
		dist[u] = myinf; return false;
    };
    while (bfs()) { FOR(u,N1) { if (pairu[u] == mynil) dfs(u); } }
    vector<pi> res;
    FOR(u,N1) if (pairu[u] != mynil) res.push_back(make_pair(u,pairu[u])); 
    return res;
}


int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N,K; cin >> N >> K; vector<vi> SS(N); FOR(i,N) { SS[i].resize(K); FOR(j,K) cin >> SS[i][j]; }
        vector<vi> adj(N);
        FOR(i,N) {
            FOR(j,N) {
                auto good = true;
                FOR(k,K) { if (SS[i][k] >= SS[j][k]) { good = false; break; } }
                if (good) adj[i].push_back(j);
            }
        }
        auto pairs = HopcroftKarp(N,N,adj);
        auto ans = N - len(pairs);
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

