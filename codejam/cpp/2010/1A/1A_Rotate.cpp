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
        ll N,K; cin >> N >> K;
        vector<vector<char>> bd(N); 
        FOR(i,N) { string s; cin >> s; for (auto c : s) { bd[i].push_back(c); } }
        FOR(i,N) { 
            auto jj = N-1;
            for (auto j=jj;j>=0;j--) { if (bd[i][j] != '.') { bd[i][jj] = bd[i][j]; jj--; } }
			for (auto j=jj;j>=0;j--) bd[i][j] = '.';
        }
        vector<pi> dirs{{0,1},{1,0},{1,-1},{1,1}};
        bool red(false),blue(false);
        for (auto &[di,dj] : dirs) {
            FOR(i,N) {
                FOR(j,N) {
                    ll rcnt(0),bcnt(0),ii(i),jj(j);
                    FOR(k,K) {
                        if (ii < 0 || ii >= N || jj < 0 || jj >= N || bd[ii][jj] == '.') break;
                        if (bd[ii][jj] == 'R') rcnt++; else bcnt++;
                        ii += di; jj += dj;
                    }
                    if (rcnt == K) red = true;
                    if (bcnt == K) blue = true;
                }
            }
        }
        string ans = red && blue ? "Both" : red ? "Red" : blue ? "Blue" : "Neither";
        printf("Case #%lld: %s\n",tt,ans.c_str());
    }
}

