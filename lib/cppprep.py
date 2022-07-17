import argparse
import os.path
import shutil
from pathlib import Path


def mkCppStarterFile(fn,type) :

    header = '''#include <bits/stdc++.h>
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
'''
    body = '''int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
}
'''
    googlebody = '''int main (int argc, char **argv) {
    if (argc > 1) { freopen(argv[1],"r",stdin); }
    ios_base::sync_with_stdio(false);cin.tie(0);
    // PROGRAM STARTS HERE
    ll T; cin >> T;
    for (ll tt=1;tt<=T;tt++) {
        printf("Case #%lld: %lld\\n",tt,0);
    }
}
'''
    vars = [header,body]
    if type == "cj" or type == "ks" : vars = [header,googlebody]
    with open(fn,'wt') as fp :
        for v in vars :
            print(v,file=fp) 
 
def mkGoLaunchJson(fn) :
    ttt = '''
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}",
            "args": [ "${fileDirname}/junk.in" ]
        }
    ]
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkGoGitignore(fn) :
    with open(fn,'wt') as fp : 
        print("*.in\n*.out\n*.exe\n*.prof", file=fp)

def parseCLArgs() :
    clargparse = argparse.ArgumentParser()
    clargparse.add_argument( '--dir', required=True, action='store', default='', help='Parent Directory for the preparations')
    clargparse.add_argument( '--type', action='store', choices=["cj","ks","cf","abc","arc","agc"], default='def', help='Contest type for template customization')
    clargparse.add_argument( '--cname', required=True, action='store', help='Name of the contest')
    clargparse.add_argument( '--plist', required=True, action='store', help='String with comma separated list of problem names')
    clargs = clargparse.parse_args()
    if not os.path.exists(clargs.dir) : raise Exception(f"Directory '{clargs.dir}' does not exist.  Exiting...")
    return clargs

if __name__ == "__main__" :
    clargs = parseCLArgs()
    xx = clargs.plist.split(',')
    ## Prepend contest name to cf, abc, arc, and agc problems
    if clargs.type in ("cf","abc","arc","agc","cj"): 
        probList = [(clargs.cname,clargs.cname+"_"+x) for x in xx]
    else :
        probList = [(clargs.cname,x) for x in xx]

    if not os.path.exists(f"{clargs.dir}/.vscode") :
        os.mkdir(f"{clargs.dir}/.vscode")
        Path(f"{clargs.dir}/.vscode/launch.json").touch()
    
    if not os.path.exists(f"{clargs.dir}/.gitignore") :
        Path(f"{clargs.dir}/.gitignore").touch()
        mkGoGitignore(f"{clargs.dir}/.gitignore")

    for (d,prob) in probList :
        if not os.path.exists(f"{clargs.dir}/{d}") : os.mkdir(f"{clargs.dir}/{d}")
        if not os.path.exists(f"{clargs.dir}/{d}/.vscode") :
            os.mkdir(f"{clargs.dir}/{d}/.vscode")
            mkGoLaunchJson(f"{clargs.dir}/{d}/.vscode/launch.json")
            Path(f"{clargs.dir}/{d}/junk.in").touch()
            ##mkGoSettingsJson(f"{clargs.dir}/{d}/.vscode/settings.json")
        if not os.path.exists(f"{clargs.dir}/{d}/{prob}.cpp") : 
            mkCppStarterFile(f"{clargs.dir}/{d}/{prob}.cpp",clargs.type)

