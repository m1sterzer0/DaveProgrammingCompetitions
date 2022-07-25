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
        ll N; cin >> N;
        ll sumx(0),sumy(0),sumz(0),sumvx(0),sumvy(0),sumvz(0);
        FOR(i,N) {
            ll x,y,z,vx,vy,vz; cin >> x >> y >> z >> vx >> vy >> vz; 
            sumx += x; sumy += y; sumz += z; sumvx += vx; sumvy += vy; sumvz += vz;
        }
        double tmin;
        if (sumvx == 0 && sumvy == 0 && sumvz == 0) {
            tmin = 0.0;
        } else {
            double num   = double(sumx*sumvx+sumy*sumvy+sumz*sumvz);
            double denom = -1.0*double(sumvx*sumvx+sumvy*sumvy+sumvz*sumvz);
            tmin = max(0.00,num/denom);
        }
        double x = (sumx + tmin * sumvx) / double(N);
        double y = (sumy + tmin * sumvy) / double(N);
        double z = (sumz + tmin * sumvz) / double(N);
        double dmin = sqrt(x*x+y*y+z*z);
        printf("Case #%lld: %.17g %.17g\n",tt,dmin,tmin);
    }
}

