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

struct center {ll i,j; double x,y; };

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N; cin >> N; vector<double> XF(N),YF(N),RF(N); FOR(i,N) cin >> XF[i] >> YF[i] >> RF[i];
        auto testr = [&](double R) -> bool {
            vector<center> centers;
            FOR(i,N) {
                if (R < RF[i]) continue;
                auto x1 = XF[i]; auto y1 = YF[i]; auto r1 = RF[i];
                centers.push_back({i,i,x1,y1});
                for (ll j=i+1;j<N;j++) {
                    auto x2 = XF[j]; auto y2 = YF[j]; auto r2 = RF[j];
                    auto d = sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1));
                    if (R < 0.5*(r1+d+r2)) { continue; }
                    auto vx1 = (x2-x1)/d;
                    auto vy1 = (y2-y1)/d;
                    auto vx2 = -vy1;
                    auto vy2 = vx1;
                    auto u = (d*d-(R-r2)*(R-r2)+(R-r1)*(R-r1))*0.5/d;
                    auto xx = (R-r1)*(R-r1)-u*u;
                    auto v = xx < 0 ? 0.00 : sqrt(xx);
                    auto xa = x1 + u*vx1 + v*vx2;
                    auto ya = y1 + u*vy1 + v*vy2;
                    auto xb = x1 + u*vx1 - v*vx2;
                    auto yb = y1 + u*vy1 - v*vy2;
                    centers.push_back({i,j,xa,ya});
                    centers.push_back({i,j,xb,yb});
                }
            }
            auto nc = len(centers);
            for (ll i=0;i<nc;i++) {
                auto i1 = centers[i].i; auto j1 = centers[i].j; auto xa = centers[i].x; auto ya = centers[i].y;
                for (ll j=i;j<nc;j++) {
                    auto i2 = centers[j].i; auto j2 = centers[j].j; auto xb = centers[j].x; auto yb = centers[j].y;
                    auto good = true;
                    FOR(k,N) {
                        if (k==i1||k==i2||k==j1||k==j2) continue;
                        auto xk = XF[k]; auto yk = YF[k]; auto rk = RF[k];
                        if ((xk-xa)*(xk-xa)+(yk-ya)*(yk-ya) <= (R-rk)*(R-rk)) continue;
						if ((xk-xb)*(xk-xb)+(yk-yb)*(yk-yb) <= (R-rk)*(R-rk)) continue;
						good = false; break;
                    }
                    if (good) return true;
                }
            }
            return false;
        };
        auto l(0.0), r(1200.0);
        FOR(i,60) { auto m = 0.5*(l+r); if (testr(m)) { r = m; } else { l = m; } }
        auto ans = 0.5*(l+r);
        printf("Case #%lld: %.17g\n",tt,ans);
    }
}

