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
        ll N; cin >> N;
        vi V1(N),V2(N); FOR(i,N) { cin >> V1[i]; } FOR(i,N) { cin >> V2[i]; }
        auto lt = [&](const int &a, const int &b) -> bool { return a < b; };
        auto gt = [&](const int &a, const int &b) -> bool { return a > b; };
        sort(V1.begin(),V1.end(),lt);
        sort(V2.begin(),V2.end(),gt);
        ll ans(0); FOR(i,N) { ans += V1[i] * V2[i]; }
        printf("Case #%lld: %lld\n",tt,ans);
    }
}

