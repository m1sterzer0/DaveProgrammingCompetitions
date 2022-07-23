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

int words[5000][16];
int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll L,D,N; cin >> L >> D >> N;
    vector<string> W(D); FOR(i,D) cin >> W[i];
    vector<string> P(N); FOR(i,N) cin >> P[i];
    FOR(i,D) FOR(j,L) words[i][j] = 1 << (W[i][j]-'a');
    vector<int> pat(L);
    FOR(i,N) {
        FOR(j,L) pat[j] = 0;
        ll idx = 0; bool paren = false;
        for (auto &c : P[i]) {
            if (c == '(') { paren = true; continue; }
            if (c == ')') { idx++; paren = false; continue; }
            pat[idx] |= 1 << (c-'a'); if (!paren) idx++;
        }
        ll ans = D;
        FOR (j,D) FOR(k,L) if ((pat[k] & words[j][k]) == 0) { ans--; break; }
        printf("Case #%lld: %lld\n",i+1,ans);
    }
}

