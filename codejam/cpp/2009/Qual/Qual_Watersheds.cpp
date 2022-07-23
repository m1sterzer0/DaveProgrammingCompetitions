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

struct dsu {
    private :
        ll _n;
        vi parentOrSize;
    public :
        dsu() : _n(0) {}
        dsu(ll n) : _n(n),parentOrSize(n,-1) {}
        ll leader(ll a) {
            ll x = parentOrSize[a];
            if (x < 0) { return a; }
            return parentOrSize[a] = leader(x);
        }
        ll merge(ll a, ll b) {
            ll x = leader(a); ll y = leader(b);
            if (x == y) { return x; }
            if (parentOrSize[y] < parentOrSize[x]) { swap(x,y); }
            parentOrSize[x] += parentOrSize[y];
            return parentOrSize[y] = x;
        }
        bool same(ll a, ll b) {
            return leader(a) == leader(b);
        }
        ll size(ll a) {
            auto x = leader(a);
            return -parentOrSize[x];
        }
        vector<vector<ll>> groups() {
            vi leaderBuf(_n,0), leaders;
            FOR(i,_n) { if (i == leader(i)) { leaders.push_back(i);} }
            ll nlead = len(leaders);
            vi lkup(_n,-1); FOR(i,nlead) { lkup[leaders[i]]=i; }
            vector<vector<ll>> ans(nlead);
            FOR(i,_n) { ans[lkup[leader(i)]].push_back(i); }
            return ans;
        }
};

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll gr[110][110];
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll H,W; cin >> H >> W;
        FOR(i,H) FOR(j,W) cin >> gr[i][j];
        dsu du(H*W);
        FOR(i,H) {
            FOR(j,W) {
                ll best = gr[i][j]; ll ii = -1; ll jj = -1;
                if (i-1 >= 0 && gr[i-1][j] < best) { best = gr[i-1][j]; ii = i-1; jj = j; }
                if (j-1 >= 0 && gr[i][j-1] < best) { best = gr[i][j-1]; ii = i; jj = j-1; }
                if (j+1  < W && gr[i][j+1] < best) { best = gr[i][j+1]; ii = i; jj = j+1; }
                if (i+1  < H && gr[i+1][j] < best) { best = gr[i+1][j]; ii = i+1; jj = j; }
                if (ii != -1) { du.merge(W*ii+jj,W*i+j); }
            }
        }
        cout << "Case #" << tt << ":\n";
        vector<char> lookup(H*W,'.'); char c = 'a';
        FOR(i,H*W) {
            auto l = du.leader(i);
            if (lookup[l] == '.') { lookup[l] = c; c++; }
            cout << lookup[l] << ((i%W==W-1) ? '\n' : ' ');
        }
    }
}

