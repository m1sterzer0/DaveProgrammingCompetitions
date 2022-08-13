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

int pat[10000][26];

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N,M; cin >> N >> M; vector<string> D(N); FOR(i,N) cin >> D[i]; vector<string> L(M); FOR(i,M) cin >> L[i];
        vector<vi> bylen(11);
        FOR(i,N) FOR(j,26) pat[i][j] = 0;
        FOR(i,N) {
            auto d = D[i];
            bylen[d.size()].push_back(i);
            FOR(j,d.size()) pat[i][d[j]-'a'] |= (1 << j);
        }
        vector<string> ansarr;
        FOR(lidx,M) {
            vi larr(26); FOR(i,26) larr[i] = L[lidx][i]-'a';
            ll bestidx(0), bestscore(-1);
            function<pair<ll,ll>(const vi&,ll)> dfs;
            dfs = [&](const vi& clist,ll lidx) -> pair<ll,ll> {
                ll l=larr[lidx];
                if (clist.size() == 1) return make_pair(0,clist[0]);
                map<ll,vi> lkup;
                for (auto &idx : clist) {
                    auto p = pat[idx][l];
                    if (lkup.count(p)==0) lkup[p] = vi();
                    lkup[p].push_back(idx); 
                }
                ll best(-1),bestidx(0);
                for (auto const& [key,val] : lkup) {
                    auto [lb,lbs] = dfs(val,lidx+1);
                    if (key == 0 && lkup.size() > 1) lb++;
                    if (lb > best || lb == best && lbs < bestidx) { best = lb; bestidx = lbs; }
                }
                return make_pair(best,bestidx);
            };
            for (ll s = 1; s <= 10; s++) {
                if (bylen[s].size() == 0) { continue; }
                auto [lb,lbs] = dfs(bylen[s],0);
                if (lb > bestscore || lb == bestscore && lbs < bestidx) { bestscore = lb; bestidx = lbs; }
            }
            ansarr.push_back(D[bestidx]);
        }
        cout << "Case #" << tt << ":";
        for (auto &a : ansarr) { cout << ' ' << a; }
        cout << '\n';
    }
}

