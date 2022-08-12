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
        ll N; cin >> N;
        vi X1(N),X2(N),Y1(N),Y2(N);
        FOR(i,N) cin >> X1[i] >> Y1[i] >> X2[i] >> Y2[i];
        auto overlap = [&](ll i,ll j) -> bool { return X2[i] >= X1[j] && X2[j] >= X1[i] && Y2[i] >= Y1[j] && Y2[j] >= Y1[i]; }; 
        auto adjew = [&](ll i,ll j) -> bool { return (X2[i]+1 == X1[j] || X2[j]+1 == X1[i]) && Y2[i] >= Y1[j] && Y2[j] >= Y1[i]; };
        auto adjns = [&](ll i,ll j) -> bool { return (Y2[i]+1 == Y1[j] || Y2[j]+1 == Y1[i]) && X2[i] >= X1[j] && X2[j] >= X1[i]; };
        auto adjcorner = [&](ll i,ll j) -> bool { return (X2[i]+1 == X1[j] && Y1[i]-1 == Y2[j]) || (X2[j]+1 == X1[i] && Y1[j]-1 == Y2[i]); };
        vector<bool> visited(N,false); deque<ll> q; ll ans = 0;
        FOR(i,N) {
            if (visited[i]) continue;
            visited[i] = true; q.push_back(i); ll xmax(-1),ymax(-1),minxpy(1LL<<60);
            while (!q.empty()) {
                auto idx = q.front(); q.pop_front();
                xmax = max(xmax,X2[idx]); ymax = max(ymax,Y2[idx]); minxpy = min(minxpy,X1[idx]+Y1[idx]);
                FOR(j,N) {
                    if (visited[j]) continue;
                    if (overlap(idx,j) || adjew(idx,j) || adjns(idx,j) || adjcorner(idx,j)) { visited[j] = true; q.push_back(j); }
                }
            }
            ans = max(ans,xmax+ymax-minxpy+1);
        }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

