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
        ll n,A,B,C,D,x0,y0,M; cin >> n >> A >> B >> C >> D >> x0 >> y0 >> M;
        vi sb(9,0);
        ll x,y; x = x0; y = y0;
        FOR(i,n) { sb[3 * (x%3) + (y%3)]++; x = (A*x+B) % M; y = (C*y+D) % M; }
        ll ans = 0;
        for (ll i = 0; i < 9; i++ ) {
            //if (sb[i] == 0) { continue; }
            for(ll j = i; j < 9; j++) {
                //if (sb[j] == 0) { continue; }
                for (ll k = j; k < 9; k++) {
                    //if (sb[k] == 0) { continue; }
                    if ((i/3+j/3+k/3) % 3 != 0) { continue; }
                    if ((i%3+j%3+k%3) % 3 != 0) { continue; }
                    if (i == j && j == k ) {
                        ans += sb[i] * (sb[i]-1) * (sb[i]-2) / 6;
                    } else if (i == j) {
                        ans += sb[i] * (sb[i]-1) / 2 * sb[k];
                    } else if (j == k) {
                        ans += sb[i] * sb[j] * (sb[j]-1) / 2;
                    } else {
                        ans += sb[i] * sb[j] * sb[k];
                    }
                }
            }
        }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

