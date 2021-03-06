import argparse
import os.path
from pathlib import Path
import shutil

def mkGoStarterFile(fn) :
    ttt = '''
package main
import (
	"bufio"
    "fmt"
	"os"
	"strconv"
)
var wrtr = bufio.NewWriterSize(os.Stdout, 10000000)
var rdr = bufio.NewScanner(os.Stdin)
func gs() string  { rdr.Scan(); return rdr.Text() }
func gi() int     { i,e := strconv.Atoi(gs()); if e != nil {panic(e)}; return i }
func gi2() (int,int) { return gi(),gi() }
func gi3() (int,int,int) { return gi(),gi(),gi() }
func gi4() (int,int,int,int) { return gi(),gi(),gi(),gi() }
func gis(n int) []int  { res := make([]int,n); for i:=0;i<n;i++ { res[i] = gi() }; return res }
func gf() float64 { f,e := strconv.ParseFloat(gs(),64); if e != nil {panic(e)}; return f }
func gbs() []byte { return []byte(gs()) }
func gfs(n int) []float64  { res := make([]float64,n); for i:=0;i<n;i++ { res[i] = gf() }; return res }
func gss(n int) []string  { res := make([]string,n); for i:=0;i<n;i++ { res[i] = gs() }; return res }
func ia(m int) []int { return make([]int,m) }
func iai(m int,v int) []int { a := make([]int,m); for i:=0;i<m;i++ { a[i] = v }; return a }
func twodi(n int,m int,v int) [][]int {
	r := make([][]int,n); for i:=0;i<n;i++ { x := make([]int,m); for j:=0;j<m;j++ { x[j] = v }; r[i] = x }; return r
}
func fill2(m int) ([]int,[]int) { a,b := ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i] = gi(),gi()}; return a,b }
func fill3(m int) ([]int,[]int,[]int) { a,b,c := ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i] = gi(),gi(),gi()}; return a,b,c }
func fill4(m int) ([]int,[]int,[]int,[]int) { a,b,c,d := ia(m),ia(m),ia(m),ia(m); for i:=0;i<m;i++ {a[i],b[i],c[i],d[i] = gi(),gi(),gi(),gi()}; return a,b,c,d }
func abs(a int) int { if a < 0 { return -a }; return a }
func rev(a []int) { i,j := 0,len(a)-1; for i < j { a[i],a[j] = a[j],a[i]; i++; j-- } }
func max(a,b int) int { if a > b { return a }; return b }
func min(a,b int) int { if a > b { return b }; return a }
func tern(cond bool, a int, b int) int { if cond { return a }; return b }
func terns(cond bool, a string, b string) string { if cond { return a }; return b }
func maxarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa > ans { ans = aa } }; return ans }
func minarr(a []int) int { ans := a[0]; for _,aa := range(a) { if aa < ans { ans = aa } }; return ans }
func sumarr(a []int) int { ans := 0; for _,aa := range(a) { ans += aa }; return ans }
func zeroarr(a []int) { for i:=0; i<len(a); i++ { a[i] = 0 } }
func powmod(a,e,mod int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m % mod }; m = m * m % mod; e >>= 1 }; return res }
func powint(a,e int) int { res, m := 1, a; for e > 0 { if e&1 != 0 { res = res * m }; m = m * m; e >>= 1 }; return res }
func gcd(a,b int) int { for b != 0 { t:=b; b=a%b; a=t }; return a }
func gcdExtended(a,b int) (int,int,int) { if a == 0 { return b,0,1 }; gcd,x1,y1 := gcdExtended(b%a,a); return gcd, y1-(b/a)*x1,x1 }
func modinv(a,m int) (int,bool) { g,x,_ := gcdExtended(a,m); if g != 1 { return 0,false }; return (x % m + m) % m,true  }
func vecintstring(a []int) string { astr := make([]string,len(a)); for i,a := range a { astr[i] = strconv.Itoa(a) }; return strings.Join(astr," ") }
func makefact(n int,mod int) ([]int,[]int) {
	fact,factinv := make([]int,n+1),make([]int,n+1)
	fact[0] = 1; for i:=1;i<=n;i++ { fact[i] = fact[i-1] * i % mod }
	factinv[n] = powmod(fact[n],mod-2,mod); for i:=n-1;i>=0;i-- { factinv[i] = factinv[i+1] * (i+1) % mod }
	return fact,factinv
}
const inf int = 2000000000000000000
const MOD int = 1000000007
func main() {
	//f1, _ := os.Create("cpu.prof"); pprof.StartCPUProfile(f1); defer pprof.StopCPUProfile()
	defer wrtr.Flush()
	infn := ""; if infn == "" && len(os.Args) > 1 {	infn = os.Args[1] }
	if infn != "" {	f, e := os.Open(infn); if e != nil { panic(e) }; rdr = bufio.NewScanner(f) }
	rdr.Split(bufio.ScanWords); rdr.Buffer(make([]byte,1024),1000000000)
    T := gi()
    for tt:=1;tt<=T;tt++ {
	    // PROGRAM STARTS HERE
        fmt.Fprintf(wrtr,"Case #%v: %v\\n",tt,0)
    }
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

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
            "program": "${file}"
        }
    ]
}
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)

def mkGoSettingsJson(fn) :
    ttt = '''
{
    "[go]": {"editor.formatOnSave": false }  
}    
'''
    with open(fn,'wt') as fp : print(ttt, file=fp)


def mkGoGitignore(fn) :
    with open(fn,'wt') as fp : 
        print("*.in\n*.out\n*.exe\n*.prof", file=fp)

def parseCLArgs() :
    clargparse = argparse.ArgumentParser()
    clargparse.add_argument( '--dir', action='store', default='', help='Parent Directory for the preparations')
    clargs = clargparse.parse_args()
    if not clargs.dir  : raise Exception("Need to provide a --dir option.  Exiting...")
    if not os.path.exists(clargs.dir) : raise Exception(f"Directory '{clargs.dir}' does not exist.  Exiting...")
    return clargs

if __name__ == "__main__" :
    clargs = parseCLArgs()
    probList = []
    probList += [(f"2022_Practice1",x) for x in ["SampleProblem","CentauriPrime","Hindex","Hex","MilkTea"]]
    probList += [(f"2022_A",x) for x in ["SpeedTyping","ChallengeNine","PalindromeFreeStrings","InterestingIntegers"]]
    probList += [(f"2022_B",x) for x in ["InfinityArea","PalindromicFactors","UnlockThePadlock","HamiltonianTour"]]
    probList += [(f"2022_C",x) for x in ["NewPassword","RangePartition","AntsOnAStick","PalindromicDeletions"]]
    probList += [(f"2022_D",x) for x in ["ImageLabeler","MaximumGain","TouchbarTyping","SuspectsAndWitnesses"]]
    #probList += [(f"2021_A",x) for x in ["KGoodnessString","LShapedPlots","RabbitHouse","Checksum"]]
    #probList += [(f"2021_B",x) for x in ["IncreasingSubstring","LongestProgression","ConsecutivePrimes","TruckDelivery"]]
    #probList += [(f"2021_C",x) for x in ["SmallerStrings","AlienGenerator","RockPaperScissors","BinaryOperator"]]
    #probList += [(f"2021_D",x) for x in ["ArithmeticSquare","CuttingIntervals","FinalExam","PrimesandQueries"]]
    #probList += [(f"2021_E",x) for x in ["ShuffledAnagrams","BirthdayCake","PalindromicCrossword","IncreasingSequenceCardGame"]]
    #probList += [(f"2021_F",x) for x in ["TrashBins","Festival","StarTrappers","GraphTravel"]]
    #probList += [(f"2021_G",x) for x in ["DogsandCats","StayingHydrated","BananaBunches","SimplePolygon"]]
    #probList += [(f"2021_H",x) for x in ["TransformtheString","Painter","SillySubstitutions","DependentEvents"]]
    #probList += [(f"2020_A",x) for x in ["Allocation","Plates","Workout","Bundling"]]
    #probList += [(f"2020_B",x) for x in ["BikeTour","BusRoutes","RobotPathDecoding","WanderingRobot"]]
    #probList += [(f"2020_C",x) for x in ["Countdown","StableWall","PerfectSubarray","Candies"]]
    #probList += [(f"2020_D",x) for x in ["RecordBreaker","AlienPiano","Beautyoftree","LockedDoors"]]
    #probList += [(f"2020_E",x) for x in ["LongestArithmetic","HighBuildings","Toys","GoldenStone"]]
    #probList += [(f"2020_F",x) for x in ["ATMQueue","MetalHarvest","PaintersDuel","Yeetzhee"]]
    #probList += [(f"2020_G",x) for x in ["KickStart","MaximumCoins","CombinationLock","MergeCards"]]
    #probList += [(f"2020_H",x) for x in ["Retype","BoringNumbers","Rugby","Friends"]] 
    #probList += [(f"2019_A",x) for x in ["Training","Parcels","Contention"]]
    #probList += [(f"2019_B",x) for x in ["BuildingPalindromes","EnergyStones","DiverseSubarray"]]
    #probList += [(f"2019_C",x) for x in ["WiggleWalk","CircuitBoard","CatchSome"]]
    #probList += [(f"2019_D",x) for x in ["XorWhat","LatestGuests","FoodStalls"]]
    #probList += [(f"2019_E",x) for x in ["CherriesMesh","CodeEatSwitcher","StreetCheckers"]]
    #probList += [(f"2019_F",x) for x in ["Flattening","TeachMe","SpectatingVillages"]]
    #probList += [(f"2019_G",x) for x in ["BookReading","TheEquation","Shifts"]]
    #probList += [(f"2019_H",x) for x in ["Hindex","DiagonalPuzzle","Elevanagram"]]
    #probList += [(f"2018_A",x) for x in ["EvenDigits","LuckyDip","ScrambledWords"]]
    #probList += [(f"2018_B",x) for x in ["NoNine","SherlockAndTheBitStrings","KingsCircle"]]
    #probList += [(f"2018_C",x) for x in ["PlanetDistance","FairiesAndWitches","KickstartAlarm"]]
    #probList += [(f"2018_D",x) for x in ["Candies","Paragliding","FunniestWordSearch"]]
    #probList += [(f"2018_E",x) for x in ["BoardGame","MilkTea","Yogurt"]]
    #probList += [(f"2018_F",x) for x in ["CommonAnagrams","SpecializingVillages","PalindromicSequence"]]
    #probList += [(f"2018_G",x) for x in ["ProductTriplets","CombiningClasses","CaveEsacpe"]]
    #probList += [(f"2018_H",x) for x in ["BigButtons","Mural","LetMeCountTheWays"]]
    probList += [(f"2017_A",x) for x in ["SquareCounting","PatternOverlap","TwoCubes"]]
    probList += [(f"2017_B",x) for x in ["MathEncoder","Center","ChristmasTree"]]
    probList += [(f"2017_C",x) for x in ["AmbiguousCipher","XSquared","MagicalThinkingV2","The4MCorporation"]]
    probList += [(f"2017_D",x) for x in ["Sightseeing","SherlockAndMatrixGame","Trash"]]
    #probList += [(f"2017_E",x) for x in ["TrapezoidCounting","CopyPaste","Blackhole"]]
    #probList += [(f"2017_F",x) for x in ["Cake","Kicksort","DanceBattle","CatchThemAll"]]
    #probList += [(f"2017_G",x) for x in ["HugeNumbers","CardsGame","MatrixCutting"]]
    probList += [(f"2016_A",x) for x in ["CountryLeader","Rain","JanesFlowerShop","ClashRoyale"]]
    probList += [(f"2016_B",x) for x in ["SherlockAndParentheses","SherlockAndWatsonGymSecrets","WatsonAndIntervals","SherlockAndPermutationSorting"]]
    probList += [(f"2016_C",x) for x in ["MonsterPath","SafeSquares","Evaluation","Soldiers"]]
    probList += [(f"2016_D",x) for x in ["Vote","Sitting","CodejamonCipher","StretchRope"]]
    probList += [(f"2016_E",x) for x in ["DiwaliLightings","BeautifulNumbers","PartioningNumber","SortingArray"]]
    probList += [(f"2015_A",x) for x in ["GoogolString","Gcube","Gcampus","Gsnake"]]
    #probList += [(f"2015_B",x) for x in ["Travel","Gwheels","Gnumbers","AlbocedeDNA"]]
    #probList += [(f"2015_C",x) for x in ["Granks","Gfiles","Ggames","Gmatrix"]]
    #probList += [(f"2015_D",x) for x in ["DynamicGrid","Gballoon","IPAddressSummarization","VirtualRabbit"]]
    #probList += [(f"2015_E",x) for x in ["LazySpellingBee","RobotRockBand","NotSoRandom","SumsOfSums"]]
    #probList += [(f"2014_A",x) for x in ["Super2048","SevenSegmentDisplay","CutTiles","Addition"]]
    #probList += [(f"2014_B",x) for x in ["PasswordAttacker","NewYearsEve","CardGame","ParenthesesOrder"]]
    #probList += [(f"2014_C",x) for x in ["Minesweeper","TakingMetro","BrokenCalculator","Tetris"]]
    #probList += [(f"2014_D",x) for x in ["CubeIV","GbusCount","SortAScrambledItinerary","ItzChess"]]
    #probList += [(f"2013_A",x) for x in ["Sorting","ReadPhoneNumber","RationalNumberTree","CrossTheMaze","SpaceshipDefence"]]
    #probList += [(f"2013_B",x) for x in ["SudokuChecker","IgnoreAllMyComments","DragonMaze","MeetAndParty","Hex"]]
    


    if not os.path.exists(f"{clargs.dir}/.vscode") :
        os.mkdir(f"{clargs.dir}/.vscode")
        Path(f"{clargs.dir}/.vscode/launch.json").touch()
        mkGoLaunchJson(f"{clargs.dir}/.vscode/launch.json")
        mkGoSettingsJson(f"{clargs.dir}/.vscode/settings.json")
    
    if not os.path.exists(f"{clargs.dir}/.gitignore") :
        Path(f"{clargs.dir}/.gitignore").touch()
        mkGoGitignore(f"{clargs.dir}/.gitignore")

    for (d,prob) in probList :
        if not os.path.exists(f"{clargs.dir}/{d}") : os.mkdir(f"{clargs.dir}/{d}")
        if not os.path.exists(f"{clargs.dir}/{d}/.vscode") :
            os.mkdir(f"{clargs.dir}/{d}/.vscode")
            shutil.copyfile(f"{clargs.dir}/.vscode/launch.json",f"{clargs.dir}/{d}/.vscode/launch.json")
            shutil.copyfile(f"{clargs.dir}/.vscode/settings.json",f"{clargs.dir}/{d}/.vscode/settings.json")


        if not os.path.exists(f"{clargs.dir}/{d}/{prob}") : 
            os.mkdir(f"{clargs.dir}/{d}/{prob}")
            Path(f"{clargs.dir}/{d}/{prob}/{prob}.go").touch()
            Path(f"{clargs.dir}/{d}/{prob}/junk.in").touch()
            mkGoStarterFile(f"{clargs.dir}/{d}/{prob}/{prob}.go")

