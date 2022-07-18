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
        ll N,M; cin >> N >> M;
        vi malted(M,-1);
        vi uncnt(M,0);
        vector<vector<int>> gr(N);
        vi sb(N,0);
        queue<ll> mq;
        FOR(i,M) {
            ll t; cin >> t;
            FOR(j,t) {
                ll x,y; cin >> x >> y;
                if (y == 0) { uncnt[i]++; gr[x-1].push_back(i); } else { malted[i] = x-1; }
            }
            if (uncnt[i] == 0) { mq.push(malted[i]); }
        }
        bool good(true);
        while (!mq.empty()) {
            auto x = mq.front(); mq.pop(); 
            if (sb[x] == 1) { continue; }
            sb[x] = 1;
            for (auto &i : gr[x]) {
                uncnt[i]--;
                if (uncnt[i] == 0) {
                    if (malted[i] == -1) { good = false; break; }
                    mq.push(malted[i]);
                }
            }
            if (!good) { break; }
        }
        cout << "Case #" << tt << ": ";
        if (!good) {
            cout << "IMPOSSIBLE\n";
        } else {
            cout << sb[0];
            for(ll i=1;i<N;i++) { cout << " " << sb[i]; }
            cout << '\n';
        }
    }
}

