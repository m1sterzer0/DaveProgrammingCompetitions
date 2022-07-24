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

struct treenode { double w; string f; ll t1; ll t0; };

void addTokens(string &buf, vector<string> &tokens) {
    size_t p = 0; auto n = buf.size();
    while (p < n) {
        while (p < n && buf[p] == ' ') p++; if (p == n) break;
        if (buf[p] == '(') { tokens.push_back("("); p++; continue; }
        if (buf[p] == ')') { tokens.push_back(")"); p++; continue; }
        auto st = p; while (p < n && buf[p] != '(' && buf[p] != ')' && buf[p] != ' ') p++;
        tokens.push_back(buf.substr(st,p-st));
    }
}
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
    string buf;
    getline(cin,buf); ll T = stoll(buf);
    for (ll tt=1;tt<=T;tt++) {
        getline(cin,buf); ll L = stoll(buf);
        vector<string> tokens;
        FOR(i,L) { getline(cin,buf); addTokens(buf,tokens); }
        fflush(stdout);
        vector<treenode> tree;
        function<ll(ll)> parsetree;
        parsetree = [&](ll idx) -> ll {
            auto tidx = (ll) tree.size();
            tree.push_back({1.0,"",-1,-1});
            tree[tidx].w = stod(tokens[idx+1]);
            if (tokens[idx+2] == ")") { return idx+2; }
            tree[tidx].f = tokens[idx+2];
            tree[tidx].t1 = (ll) tree.size();
            auto last1 = parsetree(idx+3);
            tree[tidx].t0 = (ll) tree.size();
            auto last2 = parsetree(last1+1);
            return last2+1;
        };
        parsetree(0);
        function<double(ll)> evaltree;
        set<string> features;
        evaltree = [&](ll tidx) -> double {
            auto tt = tree[tidx];
            auto ans = tt.w;
            ll c = features.count(tt.f);
            return ans * (tt.f == "" ? 1.0 : c > 0 ? evaltree(tt.t1) : evaltree(tt.t0));
        };
        getline(cin,buf); ll A = stoll(buf);
        printf("Case #%lld:\n",tt);
        FOR(i,A) {
            getline(cin,buf); auto sarr = splitString(buf); ll n = len(sarr);
            features.clear(); for (ll i = 2; i < n; i++) { features.insert(sarr[i]); }
            printf("%.17g\n",evaltree(0));
        }
    }
}

