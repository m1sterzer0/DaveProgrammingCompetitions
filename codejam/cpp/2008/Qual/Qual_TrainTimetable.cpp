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

struct event { ll typ; ll time; };
bool eventlt (const event &a, const event &b) { return a.time < b.time || a.time == b.time && a.typ < b.typ; }

int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        ll ta,NA,NB; cin >> ta >> NA >> NB;
        vector<event> events;
        FOR(i,NA) {
            string s1,s2; cin >> s1 >> s2;
            ll t1 = ll(s1[4]-'0') + 10*ll(s1[3]-'0') + 60*ll(s1[1]-'0') + 600 * ll(s1[0]-'0');
            ll t2 = ll(s2[4]-'0') + 10*ll(s2[3]-'0') + 60*ll(s2[1]-'0') + 600 * ll(s2[0]-'0');
            events.push_back({2,t1}); events.push_back({1,t2+ta});
        }
        FOR(i,NB) {
            string s1,s2; cin >> s1 >> s2;
            ll t1 = ll(s1[4]-'0') + 10*ll(s1[3]-'0') + 60*ll(s1[1]-'0') + 600 * ll(s1[0]-'0');
            ll t2 = ll(s2[4]-'0') + 10*ll(s2[3]-'0') + 60*ll(s2[1]-'0') + 600 * ll(s2[0]-'0');
            events.push_back({3,t1}); events.push_back({0,t2+ta});
        }
        sort(events.begin(),events.end(),eventlt);
        ll ansa(0),ansb(0),availa(0),availb(0);
        for (auto e : events) {
            if (e.typ == 0) { availa++; continue; }
            if (e.typ == 1) { availb++; continue; }
            if (e.typ == 2) { if (availa == 0) { ansa++; } else { availa--; } continue; }
            if (e.typ == 3) { if (availb == 0) { ansb++; } else { availb--; } continue; }
        }
        printf("Case #%lld: %lld %lld\n",tt,ansa,ansb);
    }
}

