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
    string buf;
    getline(cin,buf); ll T = atoi(buf.c_str());
    for (ll tt=1;tt<=T;tt++) {
        getline(cin,buf); ll N = atoi(buf.c_str());
        vector<string> S(N); FOR(i,N) { getline(cin,S[i]); }
        map<string,ll> d; FOR(i,N) { d[S[i]] = i; }
        getline(cin,buf); ll Q = atoi(buf.c_str());
        vector<string> QQ(Q); FOR(i,Q) { getline(cin,QQ[i]); }
        vector<ll> dp(N),ndp(N);
        const ll inf = 1LL << 61;
        FOR(i,Q) {
            auto bidx = d[QQ[i]];
            auto mindp = *min_element(dp.begin(),dp.end());
            FOR(j,N) { ndp[j] = j == bidx ? inf : dp[j] == mindp ? mindp : mindp+1; }
            swap(dp,ndp);
        }
        auto ans = *min_element(dp.begin(),dp.end());
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

