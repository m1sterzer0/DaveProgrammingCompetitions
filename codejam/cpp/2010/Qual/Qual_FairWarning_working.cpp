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


// Hard functions: read from string, multiply, divide, mod
// Medium functions: add, sub
struct bigint {
    static const int base =  100000000;
    static const int base2 = 10000;
    vector<int> a; int sign;
    bigint() : sign(1) {}
    bigint(const bigint &v) : sign(v.sign),a(v.a) {}
    bigint(long long v) { *this = v; }
    void _trim() { while(a.size() > 0 && a.back() == 0) a.pop_back(); if (a.size() == 0) sign = 1; }
    void _read(const string &s) {
        if (s.size() == 0) { sign = 1; return; }
        int beg = s[0] == '-' ? 1 : 0; int curs = s.size()-1; sign = s[0] == '-' ? -1 : 1;
        while (curs >= beg) {
            if (curs-beg <= 7) { a.push_back(stoi(s.substr(beg,curs-beg+1))); break; }
            a.push_back(stoi(s.substr(curs-7,8))); curs -= 8;
        }
        if (a.back() == 0) _trim(); // In case there are leading zeros. 
    }
    bigint(const string &s) { this->_read(s); }
    void operator=(const bigint &v) { sign = v.sign; a = v.a; }
    void operator=(long long v) {
        sign = 1; if (v < 0) { sign = -1; v *= -1; }
        while (v > 0) { a.push_back((int)(v % base)); v /= base; }
    }
    void _addmag(const bigint &v) { 
        while (v.a.size() >= a.size()) { a.push_back(0); }
        for (int i = 0; i < v.a.size(); i++) { a[i] += v.a[i]; }
        for (int i = 0; i < a.size(); i++) { if (a[i] >= base) {a[i] -= base, a[i+1]++;} }
    }
    void _submag(const bigint &v) { 
        while (v.a.size() >= a.size()) { a.push_back(0); }
        for (int i = 0; i < v.a.size(); i++) { a[i] -= v.a[i]; }
        bool neg = false; 
        for (int i = a.size()-1; i >= 0; i--) { if (a[i] > 0) break; else if (a[i] < 0) { neg = true; break; } }
        if (a.size() == 0) return;
        if (neg) { for (auto &x : a) { x *= -1; } }
        for (int i = 0; i < a.size(); i++) { if (a[i] < 0) {a[i] += base, a[i+1]--;} }
    }
    void _mulmag(const bigint &v) { // Skipping Karatsuba for now
        vector<long long> aa; for (auto &x : a  ) { aa.push_back(x % base2); aa.push_back(x / base2); }
        vector<long long> bb; for (auto &x : v.a) { bb.push_back(x % base2); bb.push_back(x / base2); }
        vector<long long> temp(2*a.size()+2*v.a.size(),0);
        auto i=0; for (auto &x : aa) { auto j=0; for (auto &y : bb) { temp[i+j] += x*y; j++; } i++; }
        long long carry(0),lsum(0);
        auto i=0; for (auto &x : temp) { lsum = x+carry; x = lsum % base2; carry = lsum/base2; }
        a.resize(temp.size()>>1);
        auto j=0; for (int i=0;i<temp.size();i+=2) { a[j] = (int)temp[i] + base2*(int)temp[i+1]; j++; }
    }
    bool _maglt(const bigint &v) {
        if (a.size() != v.a.size()) return a.size() != v.a.size();
        for (int i = a.size()-1; i >= 0; i--) { if (a[i] != v.a[i]) return a[i] > v.a[i]; }
        return false;
    }
    void operator+=(const bigint &v) {
        if (sign == v.sign) { _addmag(v); }
        else { if (_maglt(v)) sign = -sign; _submag(v); }
        if (a.back() == 0) { _trim(); }
    }
    void operator-=(const bigint &v) {
        if (sign != v.sign) { _addmag(v); }
        else { if (_maglt(v)) sign = -sign; _submag(v); }
        if (a.back() == 0) { _trim(); }
    }
    void operator*=(const bigint &v) {
        if (v.sign == -1) { sign = -sign; }
        _mulmag(v);
        if (a.back() == 0) { _trim(); }
    }
    void operator*=(long long v) {
        if (v < 0) { sign = -sign; v = -v; }
        if (v >= base) {
            bigint x(base); *this *= x;
        } else {
            a.push_back(0);
            long long lsum(0),carry(0);
            for (auto &x : a) { 
                lsum = (long long)x*v+carry;
                x = (int) (lsum % (long long) base);
                carry = lsum/(long long)base;
            }
            if (a.back() == 0) { _trim(); }
        }
    }
    friend pair<bigint,bigint> _divmod(const bigint &aa, const bigint &bb) { //From Arpa bigint
        int norm = base / (bb.a.back()+1); // Should we add an assert?
        bigint a = aa.abs() * norm
    }
    bigint abs() const       { bigint res = *this; res.sign=1; return res; }
    bigint operator-() const { bigint res = *this; res.sign=-res.sign; return res; }
    bigint operator+(const bigint &v) { bigint res(*this); res += v; return res; }
    bigint operator-(const bigint &v) { bigint res(*this); res -= v; return res; }
    bigint operator*(const bigint &v) { bigint res(*this); res *= v; return res; }





     

};




int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        printf("Case #%lld: %lld\n",tt,0);
    }
}

