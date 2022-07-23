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
    string buf; 
    ll T; getline(cin,buf); T = stoll(buf);
    for (ll tt=1;tt<=T;tt++) {
        string S; getline(cin,S);
        string ref = "welcome to code jam";
        ll N = len(ref);
        vi dp(N+1),ndp(N+1); dp[0] = 1;
        for (auto &c : S) {
            ndp[0] = dp[0];
            FOR(i,N) ndp[i+1] = c == ref[i] ? (dp[i+1] + dp[i]) % 10000 : dp[i+1];
            swap(ndp,dp);
        }
        printf("Case #%lld: %04lld\n",tt,dp[N]);
    }
}

