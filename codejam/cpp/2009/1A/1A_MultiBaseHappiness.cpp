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
vector<string> splitString(string s) {
    size_t p = 0; auto n = s.size(); vector<string> res; 
    while (p < n) {
        while (p < n && s[p] == ' ') p++; if (p == n) break;
        auto st = p; while (p < n && s[p] != ' ') p++;
        res.push_back(s.substr(st,p-st));
    }
    return res;
}
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll maxval1 = 10000;
    vi buf;
    ll sb[11][10010]; for (ll i=0;i<=10;i++) for (ll j=0; j<=maxval1; j++) sb[i][j] = -2; 
    auto dosum = [&](ll v, ll b) -> ll { ll x = 0; while (v>0) { ll xx = v%b; x += xx*xx; v/=b; } return x; };
    for (ll base=2; base<=10; base++) {
        sb[base][0] = 0; sb[base][1] = 1;
        for (ll i=2; i<=maxval1; i++) {
            if (sb[base][i] >= 0) { continue; }
            buf.clear(); sb[base][i] = -1; buf.push_back(i); ll curs = i;
            while (true) {
                ll x = dosum(curs,base);
                if (sb[base][x] == -2) { sb[base][x] = -1; buf.push_back(x); curs = x; continue; }
                if (sb[base][x] == 0 || sb[base][x] == -1) { for (auto &xx : buf) sb[base][xx] = 0; break; }
                if (sb[base][x] == 1) { for (auto &xx : buf) sb[base][xx] = 1; break; }
            }
        }
    }
    vi ansarr(1<<11,-1);
    ll maxval2 = 100000000LL;
    for (ll bm=4; bm<1<<11; bm+=4) {
        ll start = 2;
        buf.clear(); for (ll i=2;i<=10;i++) if ((bm & (1 << i)) != 0) { buf.push_back(i); start = max(start,ansarr[bm ^ (1<<i)]); }
        for (ll i = start; i <= maxval2; i++ ) {
            bool good = true;
            for (auto &j : buf) {
                ll ii = i <= maxval1 ? i : dosum(i,j);
                if (sb[j][ii] == 0) { good = false; break; }
            }
            if (good) { ansarr[bm] = i; break; }
        }
        if (ansarr[bm] == -1) printf("ERROR: bm:%lld\n",bm);
    }
    string inbuf;
    getline(cin,inbuf); ll T = stoll(inbuf);
    for (ll tt=1;tt<=T;tt++) {
        getline(cin,inbuf); auto ss = splitString(inbuf);
        ll bm = 0;
        for (auto &s : ss) { bm |= 1 << stoi(s); }
        printf("Case #%lld: %lld\n",tt,ansarr[bm]);
    }
}

