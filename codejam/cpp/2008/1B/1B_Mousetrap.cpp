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

struct ds {
    private :
        ll _n,sz;
        vi arr;
    public :
        ds() : _n(0),sz(0) {}
        ds(ll n) {
            _n=n;
            sz = 1; while(sz<n) { sz <<= 1; }
            arr.resize(2*sz,0);
            for (ll i = sz; i < sz+n; i++ ) { arr[i] = 1; }
            for (ll i = sz-1; i >= 1; i-- ) { arr[i] = arr[2*i] + arr[2*i+1]; }
        }
        ll _query(ll idx, ll n) {
            if (idx >= sz) return idx-sz;
            if (arr[2*idx] >= n) return _query(2*idx,n);
            return _query(2*idx+1,n-arr[2*idx]);
        }
        ll query(ll n) { return _query(1,n); }
        void use(ll n) {
            ll idx = sz+n;
            while (idx > 0) { arr[idx] -= 1; idx >>= 1; }
        }
};

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll K,n; cin >> K >> n;
        vi d(n); FOR(i,n) { cin >> d[i]; }
        ll unused(K), curs(0);
        auto deck = ds(K);
        vi ansarr(K,0);
        for (ll i = 1; i <= K; i++ ) {
            ll v = (curs+i) % unused;
            if (v == 0) { v = unused; }
            auto x = deck.query(v);
            deck.use(x);
            ansarr[x] = i;
            curs = v-1; 
            unused -= 1;
        }
        vi ans(n); FOR(i,n) { ans[i] = ansarr[d[i]-1]; }
        cout << "Case #" << tt << ": " << ans[0];
        for (ll i = 1; i < n; i++) { cout << " " << ans[i]; }
        cout << '\n';
    }
}

