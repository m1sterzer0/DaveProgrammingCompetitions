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
        double f,R,t,r,g; cin >> f >> R >> t >> r >> g;
        double ans = 0.0;
        if (2*f >= g) {
            ans = 1.0;
        } else {
            double num = 0.00;
            double num2 = 0.00;
            double denom = 0.25 * PI * R * R;
            double inner = R-t-f;
            double inner2 = inner*inner;
            auto sliver = [&](double x1, double y1, double x2, double y2) -> double {
                auto ang = atan2(y2,x2)-atan2(y1,x1);
                return 0.5*inner2*(ang-sin(ang));
            };
            for (double x=r+f; x < inner; x+=2*r+g) {
                for (double y=r+f; x*x+y*y < inner2; y+=2*r+g) {
                    double x2 = x-f+g-f;
                    double y2 = y-f+g-f;
                    if (x2*x2+y2*y2 <= inner2) { num += (x2-x)*(y2-y); continue; }  //Square
                    bool c1 = x*x+y2*y2 < inner2;
                    bool c2 = x2*x2+y*y < inner2;
                    if (c1 && c2) { 
                        auto x3 = sqrt(inner2-y2*y2);
                        auto y3 = sqrt(inner2-x2*x2);
                        auto adder1 = (x2-x)*(y3-y) + 0.5*(x2-x+x3-x)*(y2-y3);
                        auto adder2 = sliver(x2,y3,x3,y2);
                        //printf("DBG Case1 x:%.9g y:%.9g adder1:%.9g adder2:%.9g\n",x,y,adder1,adder2);
                        num += adder1; num2 += adder2;
                    } else if (!c1 && !c2) {
                        auto x3 = sqrt(inner2-y*y);
                        auto y3 = sqrt(inner2-x*x);
                        auto adder1 = 0.5*(x3-x)*(y3-y); // Triangle
                        auto adder2 = sliver(x3,y,x,y3);
                        //printf("DBG Case2 x:%.9g y:%.9g adder1:%.9g adder2:%.9g\n",x,y,adder1,adder2);
                        num += adder1; num2 += adder2;
                    } else if (!c1) {
                        auto y3 = sqrt(inner2-x*x);
                        auto y4 = sqrt(inner2-x2*x2);
                        auto adder1 = 0.5*(y4-y+y3-y)*(x2-x); //Trapezoid
                        auto adder2 = sliver(x2,y4,x,y3);
                        //printf("DBG Case3 x:%.9g y:%.9g adder1:%.9g adder2:%.9g\n",x,y,adder1,adder2);
                        num += adder1; num2 += adder2;
                    } else {
                        auto x3 = sqrt(inner2-y*y);
                        auto x4 = sqrt(inner2-y2*y2);
                        auto adder1 = 0.5*(x4-x+x3-x)*(y2-y); //Trapezoid
                        auto adder2 = sliver(x3,y,x4,y2);
                        //printf("DBG Case4 x:%.9g y:%.9g x2:%.9g y2:%.9g x3:%.9g x4:%.9g adder1:%.9g adder2:%.9g\n",x,y,x2,y2,x3,x4,adder1,adder2);
                        num += adder1; num2 += adder2;
                    }
                }
            }
            num += num2; ans = (denom-num)/denom;
        }
        printf("Case #%lld: %.9g\n",tt,ans);
    }
}

