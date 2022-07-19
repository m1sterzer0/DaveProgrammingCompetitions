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

ll dp[50][220];
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        string S; cin >> S; ll N = len(S);
        FOR(i,N) { FOR(j,210) { dp[i][j] = 0; } }
        FOR(i,N) {
            ll pv = 1; ll v = 0;
            for (ll j = i; j >= 0; j--) {
                v += pv * (S[j]-'0'); v %= 210; pv *= 10; pv %= 210;
                if (j == 0) { dp[i][v] += 1; continue; }
                FOR(k,210) {
                    dp[i][(k+v)%210] += dp[j-1][k];
                    dp[i][(k+210-v)%210] += dp[j-1][k];
                }
            }
        }
        ll ans = 0;
        FOR(k,210) { if (k%2==0 || k%3==0 || k%5==0 || k%7==0) { ans += dp[N-1][k]; } }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

