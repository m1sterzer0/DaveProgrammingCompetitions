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

vi sieve(ll n) {
    vector<bool> sb(n+1,true);
    sb[0] = false; sb[1] = false;
    for (ll i=4;i<=n;i+=2) { sb[i] = false; }
    for (ll i=3;i*i<=n;i+=2) {
        if (!sb[i]) { continue; }
        for (ll j=i*i;j<=n;j+=2*i) { sb[j] = false; }
    }
    vi ans;
    FOR(i,n+1) { if (sb[i]) { ans.push_back(i); } }
    return ans;
}

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    auto primes = sieve(1000000);
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll A,B,P; cin >> A >> B >> P;
        auto du = dsu(B-A+1);
        vi entries;
        for (auto &p : primes) {
            if (p < P) continue;
            if (p > B-A) break;
            entries.clear();
            ll start = A%p == 0 ? 0 : p-A%p;
            for(ll i = start; i < B-A+1; i += p) entries.push_back(i);
            ll numentries = len(entries);
            for(ll i = 0; i < numentries-1; i++) {
                du.merge(entries[i],entries[i+1]);
            }
        }
        ll ans = 0;
        FOR(i,B-A+1) { if (i == du.leader(i)) { ans++; } }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

