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

char combine[256][256];
bool opposed[256][256];


int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll C; cin >> C; vector<string> CC(C); FOR(i,C) cin >> CC[i];
        ll D; cin >> D; vector<string> DD(D); FOR(i,D) cin >> DD[i];
        ll N; string S; cin >> N >> S;
        FOR(i,256) FOR(j,256) { combine[i][j] = 0; opposed[i][j] = false; }
        for (auto &c : CC) { combine[c[0]][c[1]] = c[2]; combine[c[1]][c[0]] = c[2]; }
        for (auto &d : DD) { opposed[d[0]][d[1]] = true; opposed[d[1]][d[0]] = true; }
        vector<char> ansarr;
        for (char c : S) {
            ansarr.push_back(c);
            while (ansarr.size() >= 2 && combine[ansarr[ansarr.size()-1]][ansarr[ansarr.size()-2]] != 0) {
                char x = combine[ansarr[ansarr.size()-1]][ansarr[ansarr.size()-2]];
                ansarr.pop_back(); ansarr.pop_back(); ansarr.push_back(x);
            }
            char l = ansarr.back();
            for (auto x : ansarr) if (opposed[l][x]) { ansarr.clear(); break; }
        }
        cout << "Case #" << tt << ": [";
        auto n = len(ansarr);
        FOR(i,n) { if (i > 0) cout << ", "; cout << ansarr[i]; }
        cout << "]\n";
    }
}

