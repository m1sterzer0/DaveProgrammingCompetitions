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

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N,M,A; cin >> N >> M >> A;
        if ( A > N*M ) {
            printf("Case #%lld: IMPOSSIBLE\n",tt);
        } else if (A == N*M) {
            printf("Case #%lld: %lld %lld %lld %lld %lld %lld\n",tt,0LL,0LL,0LL,M,N,0LL);
        } else {
            printf("Case #%lld: %lld %lld %lld %lld %lld %lld\n",tt,0LL,0LL,1LL,M,A/M+1LL,M-A%M);
        }
    }
}

