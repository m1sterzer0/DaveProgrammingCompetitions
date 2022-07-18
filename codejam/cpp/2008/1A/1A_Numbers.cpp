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

tuple<ll,ll,ll,ll> matexp(ll a11, ll a12, ll a21, ll a22, ll n, ll m) {
    ll b11(1),b12(0),b21(0),b22(1);
    while (n > 0) {
        if (n & 1 == 1) {
            ll c11 = (a11*b11+a12*b21) % m;
            ll c12 = (a11*b12+a12*b22) % m;
            ll c21 = (a21*b11+a22*b21) % m;
            ll c22 = (a21*b12+a22*b22) % m;
            b11 = c11; b12 = c12; b21 = c21; b22 = c22;
        }
        ll c11 = (a11*a11+a12*a21) % m;
        ll c12 = (a11*a12+a12*a22) % m;
        ll c21 = (a21*a11+a22*a21) % m;
        ll c22 = (a21*a12+a22*a22) % m;
        a11 = c11; a12 = c12; a21 = c21; a22 = c22;
        n >>= 1;
    }
    return tuple<ll,ll,ll,ll>(b11,b12,b21,b22);
} 

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N; cin >> N;
        ll a11,a12,a21,a22;
        tie(a11,a12,a21,a22) = matexp(3,5,1,3,N,1000);
        ll ans = (2*a11 + 1000 - 1) % 1000;
        printf("Case #%lld: %03lld\n",tt,ans);
    }
}

