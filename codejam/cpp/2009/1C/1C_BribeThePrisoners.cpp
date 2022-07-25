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

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll P,Q; cin >> P >> Q; vi QQ(Q); FOR(i,Q) { cin >> QQ[i]; QQ[i]--; }
        map <pair<ll,ll>,ll> cache;
        function<ll(ll,ll,ll,ll)> solve;
        solve = [&](ll pl,ll pr,ll ql,ll qr) -> ll {
            auto key = make_pair(pl,pr);
            auto it = cache.find(key);
            if (it != cache.end()) { return it->second; }
            ll adder = 1LL << 62;
            for (ll qidx = ql; qidx <= qr; qidx++) {
                ll cand = 0; ll qr2 = qidx-1; ll ql2 = qidx+1;
                if (ql <= qr2) { cand += solve(pl,QQ[qidx]-1,ql,qr2); }
                if (ql2 <= qr) { cand += solve(QQ[qidx]+1,pr,ql2,qr); }
                adder = min(adder,cand);
            }
            ll res = pr-pl+adder;
            cache[key] = res;
            return res;
        };
        ll ans = solve(0,P-1,0,Q-1);
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

