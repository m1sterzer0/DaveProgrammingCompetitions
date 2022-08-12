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

void checkPalindrome(string s, vector<bool> &sb) {
    int sl = s.size();
    FOR(i,sl) {
        ll l(i),r(i);
        while (sb[i] && l >= 0 && r < sl) {
            if (s[l] != ' ' && s[r] != ' ' && s[l] != s[r]) sb[i] = false; 
            l--; r++;
        }
    }
}

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    string buf;
    getline(cin,buf); auto T = stoll(buf);
    for (ll tt=1;tt<=T;tt++) {
        getline(cin,buf); auto k = stoll(buf);
        vector<string> bd(2*k-1);
        FOR(i,2*k-1) {
            getline(cin,bd[i]);
            while (bd[i].size() < 2*k-1) bd[i].push_back(' ');
        }
        vector<bool> rsb(2*k-1,true), csb(2*k-1,true);
        FOR(i,2*k-1) {
            checkPalindrome(bd[i],rsb);
        }
        FOR(j,2*k-1) {
            buf.clear();
            FOR(i,2*k-1) buf.push_back(bd[i][j]);
            checkPalindrome(buf,csb);
        }
        ll radd(1LL<<61), cadd(1LL<<61);
        FOR(i,2*k-1) if (rsb[i]) radd = min(radd,abs(k-1-i));
        FOR(i,2*k-1) if (csb[i]) cadd = min(cadd,abs(k-1-i));
        auto ans = (k+radd+cadd)*(k+radd+cadd)-k*k;
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

