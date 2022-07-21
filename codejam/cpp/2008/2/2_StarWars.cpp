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
        ll N; cin >> N;
        vector<double> X(N), Y(N), Z(N),P(N); FOR(i,N) cin >> X[i] >> Y[i] >> Z[i] >> P[i];
        auto tryit = [&](double m) -> bool {
            double inf = 1e99;
            double a1(-inf),a2(-inf),a3(-inf),a4(-inf),b1(inf),b2(inf),b3(inf),b4(inf);
            FOR(i,N) {
                double d = m * P[i];
                a1 = max(a1,+X[i]+Y[i]-Z[i]-d);
                a2 = max(a2,+X[i]-Y[i]+Z[i]-d);
                a3 = max(a3,-X[i]+Y[i]+Z[i]-d);
                a4 = max(a4,+X[i]+Y[i]+Z[i]-d);
                b1 = min(b1,+X[i]+Y[i]-Z[i]+d);
                b2 = min(b2,+X[i]-Y[i]+Z[i]+d);
                b3 = min(b3,-X[i]+Y[i]+Z[i]+d);
                b4 = min(b4,+X[i]+Y[i]+Z[i]+d);
            }
            if (a1 > b1 || a2 > b2 || a3 > b3 || a4 > b4) return false;
            auto l1 = a1+a2+a3;
            auto r1 = b1+b2+b3;
            if (l1 > b4 || r1 < a4) return false; else return true;
        };
        double l = 0.0; double u = 3.6e6;
        FOR(i,70) {
            auto m = 0.5*(l+u);
            if (tryit(m)) u = m; else l = m;
        }
        printf("Case #%lld: %.17g\n",tt,0.5*(l+u));
    }
}

