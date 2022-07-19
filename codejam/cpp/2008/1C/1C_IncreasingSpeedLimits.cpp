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

template <class S> struct segtree {
    private:
        ll n, sz, log;
        S(*op)(S,S);
        S e;
        std::vector<S> d;
        void update(ll i) { d[i] = op(d[2*i],d[2*i+1]); }
    public:
        segtree() {}
        segtree(ll nin, S(*opin)(S,S), const S ein) : segtree(vector<S>(nin,ein),opin,ein) {}
        segtree(const vector<S> &v, S(*opin)(S,S), const S &ein) {
            n = len(v); sz = 1; log = 0; op = opin; e = ein;
            while (sz < n) { sz <<= 1; log++; }
            d.resize(2*sz,e);
            for (ll i=0;i<n;i++) { d[sz+i] = v[i]; }
            for (ll i=n-1;i>0;i--) { update(i); }
        }
        void set(ll p, const S &x) {
            ll idx = sz+p; d[idx] = x; idx >>= 1;
            while (idx > 0) { update(idx); idx >>= 1; }
        }
        S get(ll p) { return d[p+sz]; }
        S prod(ll l, ll r) {
            r++;
            auto sml = e; auto smr = e; l += sz; r += sz;
            while (l < r) {
                if (l & 1) { sml = op(sml,d[l]); l++; }
                if (r & 1) { r-=1; smr = op(d[r],smr); }
                l >>= 1; r >>= 1;
            }
            return op(sml,smr);
        }
        S allprod() { return d[1]; }
};

ll stop (ll a, ll b) { return (a+b)%MOD; }

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll n,m,X,Y,Z; cin >> n >> m >> X >> Y >> Z;
        vi A(m); FOR(i,m) { cin >> A[i]; }
        vi S(n);
        FOR(i,n) { S[i] = A[i%m]; A[i%m] = (X*A[i%m]+Y*(i+1)) % Z; }
        // Coordinate compression
        vi S2;
        copy(S.begin(),S.end(),back_inserter(S2));
        sort(S2.begin(),S2.end());
        auto it = unique(S2.begin(),S2.end());
        S2.resize(distance(S2.begin(),it));
        map<ll,ll> lkup;
        FOR(i,len(S2)) { lkup[S2[i]]=i; }
        vi S3(n); FOR(i,n) { S3[i] = lkup[S[i]]; }
        // Seg tree stuff
        auto st = segtree<ll>(len(S2),&stop,0LL);
        ll ans = 0;
        for (auto &s : S3) {
            ll lans = 1;
            if (s > 0) { lans += st.prod(0,s-1); lans %= MOD; }
            ans += lans;
            st.set(s,(st.get(s)+lans)%MOD);
        }
        ans %= MOD;
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

