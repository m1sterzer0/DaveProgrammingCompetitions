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
        ll P,K,L; cin >> P >> K >> L;
        vi freq(L); FOR(i,L) { cin >> freq[i]; }
        sort(freq.begin(),freq.end());
        reverse(freq.begin(),freq.end());
        ll ans = 0;
        FOR(i,L) { ans += freq[i] * (1 + i / K); }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

