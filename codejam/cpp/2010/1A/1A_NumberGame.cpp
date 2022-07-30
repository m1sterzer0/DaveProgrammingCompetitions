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
        ll A1,A2,B1,B2; cin >> A1 >> A2 >> B1 >> B2;
        ll ans = 0; double phi = (1.0+sqrt(5.0))*0.5;
        for (ll a=A1;a<=A2;a++) {
            ll b = max(0LL,llround((double)a * phi)-5LL);
            while (b*b-a*b-a*a < 0) b++;
            if (b <= B1) ans += B2-B1+1; else if (b <= B2) ans += B2-b+1; else break;
        }
        for (ll b=B1;b<=B2;b++) {
            ll a = max(0LL,llround((double)b * phi)-5LL);
            while (a*a-a*b-b*b < 0) a++;
            if (a <= A1) ans += A2-A1+1; else if (a <= A2) ans += A2-a+1; else break;
        }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

