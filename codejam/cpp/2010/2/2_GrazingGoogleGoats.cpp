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

struct Line { double m; double b; double cx; double cy; double r; };
struct Pt { double x; double y; };
bool operator<(Line &a, Line &b) { return a.m > b.m; }
Pt rot2d(double x, double y, double ang) { return Pt{x*cos(ang)-y*sin(ang),x*sin(ang)+y*cos(ang)}; }
Pt invertPoint(double x, double y) { return Pt{x/(x*x+y*y),y/(x*x+y*y)}; }
Line invertCircleToLine(double cx, double cy) {
    auto m = -cx/cy; auto p = invertPoint(2.0*cx,2.0*cy); auto b = p.y-m*p.x; return Line{m,b,cx,cy,sqrt(cx*cx+cy*cy)};
}
Pt intersection(Line &l1, Line &l2) { auto x = (l2.b-l1.b)/(l1.m-l2.m); auto y = l1.m*x+l1.b; return Pt{x,y}; }
double circleSegment(double x1, double y1, double x2, double y2, double r) {
    auto ang1 = atan2(y1,x1); auto ang2 = atan2(y2,x2); auto ang = ang2-ang1 >= 0 ? ang2-ang1 : 2.0*PI+ang2-ang1;
    return 0.5*r*r*(ang - sin(ang));
}
double polyArea(vector<Pt> &plist) {
    auto n = plist.size(); auto area = 0.0;
    FOR(i,n-1) area += plist[i].x*plist[i+1].y-plist[i].y*plist[i+1].x;
    area += plist[n-1].x*plist[0].y-plist[n-1].y*plist[0].x;
    area *= 0.5;
    return area;
}
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll N,M; cin >> N >> M;
        vi PX(N),PY(N); FOR(i,N) cin >> PX[i] >> PY[i];
        vi QX(M),QY(M); FOR(i,M) cin >> QX[i] >> QY[i];
        vector<double> ansarr(M,0.00); double eps = 1e-12;
        FOR(idx,M) {
            auto qx = QX[idx]; auto qy = QY[idx];
            vector<double> angs(N); FOR(i,N) angs[i] = atan2(double(PY[i])-qy,double(PX[i])-qx);
            sort(angs.begin(),angs.end());
            vector<double> gaps(N); FOR(i,N) gaps[i] = angs[i]-angs[(N+i-1)%N]; gaps[0] += 2.0*PI;
            double bestgap(0.0); ll bestidx(0);
            FOR(i,N) if (gaps[i] > bestgap) { bestgap = gaps[i]; bestidx = i; }
            if (bestgap <= PI+eps) { ansarr[idx] = 0.00; continue; }
            double rotang = 0.5*PI-(angs[bestidx]-0.5*bestgap);
            vector<Line> lines;
            FOR(i,N) {
                auto p = rot2d(double(PX[i])-qx,double(PY[i])-qy,rotang);
                lines.push_back(invertCircleToLine(p.x,p.y));
            }
            sort(lines.begin(),lines.end());
            vector<Line> st;
            for (auto &l : lines) {
                while(st.size() >= 2) {
                    auto p1 = intersection(st[st.size()-2],st[st.size()-1]);
                    auto p2 = intersection(st[st.size()-1],l);
                    if (p2.x < p1.x) { st.pop_back(); } else break;
                }
                st.push_back(l);
            }
            vector<Pt> poly;
            poly.push_back(Pt{0.00,0.00}); double area = 0.00; double lx = 0.00; double ly = 0.00;
            auto n = st.size();
            FOR(i,n-1) {
                auto cx = st[i].cx; auto cy = st[i].cy; auto r = st[i].r;
                auto p1 = intersection(st[i],st[i+1]);
                auto p2 = invertPoint(p1.x,p1.y);
                poly.push_back(p2);
                auto adder = circleSegment(lx-cx,ly-cy,p2.x-cx,p2.y-cy,r);
                area += adder;
                lx = p2.x; ly = p2.y;
            }
            auto cx = st[n-1].cx; auto cy = st[n-1].cy; auto r = st[n-1].r;
            auto adder = circleSegment(lx-cx,ly-cy,0.0-cx,0.0-cy,r);
            area += adder;
            area += polyArea(poly);
            ansarr[idx] = area;
        }
        cout << "Case #" << tt << ":";
        FOR(i,M) { cout << ' ' << setprecision(17) << ansarr[i]; }
        cout << '\n';
    }
}
