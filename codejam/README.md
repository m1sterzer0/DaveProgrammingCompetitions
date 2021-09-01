# m1sterzer0 CodeJam Solutions ![Language](https://img.shields.io/badge/language-Python-orange.svg) ![Language](https://img.shields.io/badge/language-Julia-blueviolet.svg) ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) ![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2026-ff69b4.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-0%20%2F%2026-ff69b4.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2026-ff69b4.svg) 

These are CodeJam solutions for the archived CodeJam problems in the current CodeJam Archive competition enviroment.  In some cases for the older problems, these problems differ from the original versions that used the `4-8 minute` timers and non-centralized computes.  

`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below.

## Contest Shortcuts
* CodeJam 2021 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2021) [MySolutions](#Google-Code-Jam-2021)
* CodeJam 2020 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2020) [MySolutions](#Google-Code-Jam-2020)
* CodeJam 2019 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2019) [MySolutions](#Google-Code-Jam-2019)
* CodeJam 2018 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2018) [MySolutions](#Google-Code-Jam-2018)
* CodeJam 2017 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2017) [MySolutions](#Google-Code-Jam-2017)
* CodeJam 2016 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2016) [MySolutions](#Google-Code-Jam-2016)
* CodeJam 2015 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2015) [MySolutions](#Google-Code-Jam-2015)
* CodeJam 2014 [OfficialSite](https://codingcompetitions.withgoogle.com/codejam/archive/2014) [MySolutions](#Google-Code-Jam-2014)
* [Code Jam 2013](https://codingcompetitions.withgoogle.com/codejam/archive/2013)
* [Code Jam 2012](https://codingcompetitions.withgoogle.com/codejam/archive/2012)
* [Code Jam 2011](https://codingcompetitions.withgoogle.com/codejam/archive/2011)
* [Code Jam 2010](https://codingcompetitions.withgoogle.com/codejam/archive/2010)
* [Code Jam 2009](https://codingcompetitions.withgoogle.com/codejam/archive/2009)
* [Code Jam 2008](https://codingcompetitions.withgoogle.com/codejam/archive/2008)

## Google Code Jam 2021
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2027-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-14%20%2F%2027-yellow.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2027-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Reversort](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a)|[Julia](./julia/2021/Qual_Reversort.jl) | Easy | |
|Qual| [Moons and Umbrellas](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1145)|[Julia](./julia/2021/Qual_MoonsAndUmbrellas.jl) | Medium-Easy | DP |
|Qual| [Reversort Engineering](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d12d7)|[Julia](./julia/2021/Qual_ReversortEngineering.jl) | Easy | |
|Qual| [Median Sort](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1284)|[Julia](./julia/2021/Qual_MedianSort.jl) | Medium-Easy | Interactive |
|Qual| [Cheating Detection](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1155)|[Julia](./julia/2021/Qual_CheatingDetection.jl) | Medium | Bayes Theorem |
|1A| [Append Sort](https://codingcompetitions.withgoogle.com/codejam/round/000000000043585d/00000000007549e5)|[Julia](./julia/2021/1A_AppendSort.jl) | Easy| |
|1A| [Prime Time](https://codingcompetitions.withgoogle.com/codejam/round/000000000043585d/00000000007543d8)|[Julia](./julia/2021/1A_PrimeTime.jl) | Medium | Prime Numbers |
|1A| [Hacked Exam](https://codingcompetitions.withgoogle.com/codejam/round/000000000043585d/0000000000754750)|[Julia](./julia/2021/1A_HackedExam.jl) | Medium-Hard | Equation Solving |
|1B| [Broken Clock](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435baf/00000000007ae694)|[Julia](./julia/2021/1B_BrokenClock.jl) | Easy | Modular Equations |
|1B| [Subtransmutation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435baf/00000000007ae4aa)|[Julia](./julia/2021/1B_Subtransmutation.jl) | Medium-Easy | GCD, Simulation |
|1B| [Digit Blocks](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435baf/00000000007ae37b)|[Julia](./julia/2021/1B_DigitBlocks.jl) | Medium | Expected Value DP |
|1C| [Closest Pick](https://codingcompetitions.withgoogle.com/codejam/round/00000000004362d7/00000000007c0f00)|[Julia](./julia/2021/1C_ClosestPick.jl) | Easy | |
|1C| [Roaring Years](https://codingcompetitions.withgoogle.com/codejam/round/00000000004362d7/00000000007c0f01)|[Julia](./julia/2021/1C_RoaringYears.jl) | Medium-Easy | Binary Search |
|1C| [Double or NOTing](https://codingcompetitions.withgoogle.com/codejam/round/00000000004362d7/00000000007c1139)|[Julia](./julia/2021/1C_DoubleOrNOTing.jl) | Medium | |
|2| [Minimum Sort](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dc51c)| | Easy | |
|2| [Matrygons](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dbf06)| | Medium-Easy | |
|2| [Hidden Pancakes](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dc20c)| | Medium-Easy | |
|2| [Retiling](https://codingcompetitions.withgoogle.com/codejam/round/0000000000435915/00000000007dc2de)| | Medium-Hard | Min-cost Max Flow |
|3| [Build-A-Pair](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436142/0000000000813aa8)| | Easy | |
|3| [Square Free](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436142/0000000000813e1a)| | Medium | |
|3| [Fence Design](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436142/0000000000813bc7)| | Hard | |
|3| [Binary Search Game](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436142/0000000000813e1b)| | Hard | |
|WF| [Cutting Cake](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436329/000000000084fba1)| | Medium | |
|WF| [Slide Circuits](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436329/000000000084f7b2)| | Medium | |
|WF| [Ropes](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436329/000000000084fad0)| | Hard | |
|WF| [Divisible Divisions](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436329/000000000084fb3a)| | Medium| |
|WF| [Infinitree](https://codingcompetitions.withgoogle.com/codejam/round/0000000000436329/000000000084fc01)| | Very Hard| |

## Google Code Jam 2020
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2027-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-27%20%2F%2027-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2027-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Vestigium](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/000000000020993c)|[Julia](./julia/2020/Qual_Vestigium.jl) |Easy | |
|Qual| [Nesting Depth](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/0000000000209a9f)|[Julia](./julia/2020/Qual_NestingDepth.jl) | Easy | |
|Qual| [Parenting Partnering Returns](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/000000000020bdf9)|[Julia](./julia/2020/Qual_ParentingPartneringReturns.jl) |Easy | |
|Qual| [ESAb ATAd](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/0000000000209a9e)|[Julia](./julia/2020/Qual_ESAbATAd.jl) |Easy |Interactive |
|Qual| [Indicium](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/0000000000209aa0)|[Julia](./julia/2020/Qual_Indicium.jl) |Medium | |
|1A| [Pattern Matching](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd74/00000000002b3034)|[Julia](./julia/2020/1A_PatternMatching.jl) | Easy| |
|1A| [Pascal Walk](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd74/00000000002b1353)|[Julia](./julia/2020/1A_PascalWalk.jl) | Easy | Pascal's Triangle |
|1A| [Square Dance](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd74/00000000002b1355)|[Julia](./julia/2020/1A_SquareDance.jl) | Medium | Linked Lists|
|1B| [Expogo](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef2/00000000002d5b62)|[Julia](./julia/2020/1B_Expogo.jl) | Easy | |
|1B| [Blindfolded Bullseye](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef2/00000000002d5b63)|[Julia](./julia/2020/1B_BlindfoldedBullseye.jl) | Medium-Easy | Interactive, Binary Search|
|1B| [Join the Ranks](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef2/00000000002d5b64)|[Julia](./julia/2020/1B_JoinTheRanks.jl) | Medium-Hard | |
|1C| [Overexcited Fan](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef4/0000000000317409)|[Julia](./julia/2020/1C_OverexcitedFan.jl) | Easy | |
|1C| [Overrandomized](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef4/00000000003179a1)|[Julia](./julia/2020/1C_Overrandomized.jl) | Medium-Easy | Benford's Law |
|1C| [Oversized Pancake Choppers](https://codingcompetitions.withgoogle.com/codejam/round/000000000019fef4/00000000003172d1)|[Julia](./julia/2020/1C_OversizedPancakeChoppers.jl) | Hard | |
|2| [Incremental House of Pancakes](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ffb9/00000000003384ea)|[Julia](./julia/2020/2_IncrementalHouseofPancakes.jl) | Easy | |
|2| [Security Update](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ffb9/000000000033871f)|[Julia](./julia/2020/2_SecurityUpdate.jl) | Medium-Easy | |
|2| [Wormhole in One](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ffb9/00000000003386d0)|[Julia](./julia/2020/2_WormholeInOne.jl) | Medium | |
|2| [Emacs++](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ffb9/000000000033893b)|[Julia](./julia/2020/2_Emacspp.jl) | Very Hard | |
|3| [Naming Compromise](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff7e/00000000003774db)|[Julia](./julia/2020/3_NamingCompromise.jl) | Easy | Edit-Distance |
|3| [Thermometers](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff7e/000000000037776b)|[Julia](./julia/2020/3_Thermometers.jl) | Hard | |
|3| [Pen Testing](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff7e/0000000000377630)|[Julia](./julia/2020/3_PenTesting.jl) | Hard | Interactive, EV DP |
|3| [Recalculating](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff7e/00000000003775e9)|[Julia](./julia/2020/3_Recalculating.jl) | Very Hard | Rolling-Hash |
|WF| [Pack the Slopes](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff31/00000000003b4f31)|[Julia](./julia/2020/WF_PackTheSlopes.jl) | Medium | Heavy-Light Decomp, Segment Tree |
|WF| [Adjacent and Consecutive](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff31/00000000003b53ce)|[Julia](./julia/2020/WF_AdjacentAndConsecutive.jl) | Hard | |
|WF| [Hexacoin Jam](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff31/00000000003b4bc5)|[Julia](./julia/2020/WF_HexacoinJam.jl) | Very Hard | |
|WF| [Musical Cords](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff31/00000000003b532b)|[Julia](./julia/2020/WF_MusicalChords.jl) | Very Hard | |
|WF| [Replace All](https://codingcompetitions.withgoogle.com/codejam/round/000000000019ff31/00000000003b4bc4)|[Julia](./julia/2020/WF_ReplaceAll.jl) | Medium | Kosaraju, Floyd-Warshall, Max-Flow |

## Google Code Jam 2019
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2027-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-26%20%2F%2027-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2027-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Foregone Solution](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051705/0000000000088231)|[Julia](./julia/2019/Qual_ForegoneSolution.jl) | Easy | |
|Qual| [You Can Go Your Own Way](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051705/00000000000881da)|[Julia](./julia/2019/Qual_YouCanGoYourOwnWay.jl) | Easy | |
|Qual| [Cryptoanagrams](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051705/000000000008830b)|[Julia](./julia/2019/Qual_Cryptoanagrams.jl) | Easy | |
|Qual| [Dat Bae](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051705/00000000000881de)|[Julia](./julia/2019/Qual_DatBae.jl) | Medium-Easy | Interactive |
|1A| [Pylons](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051635/0000000000104e03)|[Julia](./julia/2019/1A_Pylons.jl) | Medium-Easy | |
|1A| [Golf Gophers](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051635/0000000000104f1a)|[Julia](./julia/2019/1A_GolfGophers.jl) | Easy | Interactive |
|1A| [Alien Rhyme](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051635/0000000000104e05)|[Julia](./julia/2019/1A_AlienRhyme.jl) | Easy | |
|1B| [Manhattan Crepe Cart](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051706/000000000012295c)|[Julia](./julia/2019/1B_Expogo.jl) | Easy | |
|1B| [Draupnir](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051706/0000000000122837)|[Julia](./julia/2019/1B_BlindfoldedBullseye.jl) | Medium-Easy | Interactive |
|1B| [Fair Fight](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051706/0000000000122838)|[Julia](./julia/2019/1B_JoinTheRanks.jl) | Medium-Hard | |
|1C| [Robot Programming Strategy](https://codingcompetitions.withgoogle.com/codejam/round/00000000000516b9/0000000000134c90)|[Julia](./julia/2019/1C_RobotProgrammingStrategy.jl) | Easy | |
|1C| [Power Arrangers](https://codingcompetitions.withgoogle.com/codejam/round/00000000000516b9/0000000000134e91)|[Julia](./julia/2019/1C_PowerArrangers.jl) | Medium-Easy | Interactive |
|1C| [Bacterial Tactics](https://codingcompetitions.withgoogle.com/codejam/round/00000000000516b9/0000000000134cdf)|[Julia](./julia/2019/1C_BacterialTactics.jl) | Hard | |
|2| [New Elements: Part 1](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051679/0000000000146183)|[Julia](./julia/2019/2_NewElementsPart1.jl) | Easy | |
|2| [Pottery Lottery](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051679/00000000001461c8)|[Julia](./julia/2019/2_PotteryLottery.jl) | Medium | Interactive |
|2| [New Elements: Part 2](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051679/0000000000146184)|[Julia](./julia/2019/2_NewElementsPart2.jl) | Medium-Hard | |
|2| [Contransmutation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051679/0000000000146185)|[Julia](./julia/2019/2_Contransmutation.jl) | Hard | |
|3| [Zillionim](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051707/0000000000158f1a)|[Julia](./julia/2019/3_Zillionim.jl) | Medium | Interactive |
|3| [Pancake Pyramid](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051707/00000000001591be)|[Julia](./julia/2019/3_PancakePyramid.jl) | Medium | |
|3| [Datacenter Duplex](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051707/0000000000158f1c)|[Julia](./julia/2019/3_DatacenterDuplex.jl) | Hard | |
|3| [Napkin Folding](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051707/0000000000159170)|[*Julia](./julia/2019/3_NapkinFolding_broken.jl) | Very Hard | Broken for Large |
|WF| [Board Meeting](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77c)|[Julia](./julia/2019/WF_BoardMeeting.jl) | Medium | |
|WF| [Sorting Permutation Unit](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77d)|[Julia](./julia/2019/WF_SortingPermutationUnit.jl) | Medium | |
|WF| [Won't sum? Must Now](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77e)|[Julia](./julia/2019/WF_WontSumMustNow.jl) | Very Hard | |
|WF| [Juggle Struggle: Part 1](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77f)|[Julia](./julia/2019/WF_JuggleStrugglePart1.jl) | Hard | |
|WF| [Juggle Struggle: Part 2](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c933)|[Julia](./julia/2019/WF_JuggleStrugglePart2.jl) | Very Hard | |
|WF| [Go To Considered Helpful](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c934)|[Julia](./julia/2019/WF_GoToConsideredHelpful.jl) | Medium | |

## Google Code Jam 2018
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2026-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-26%20%2F%2026-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2026-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Saving The Universe Again](https://codingcompetitions.withgoogle.com/codejam/round/00000000000000cb/0000000000007966)|[Julia](./julia/2018/Qual_SavingTheUniverseAgain.jl) | Easy | |
|Qual| [Trouble Sort](https://codingcompetitions.withgoogle.com/codejam/round/00000000000000cb/00000000000079cb)|[Julia](./julia/2018/Qual_TroubleSort.jl) | Easy | |
|Qual| [Go, Gopher!](https://codingcompetitions.withgoogle.com/codejam/round/00000000000000cb/0000000000007a30)|[Julia](./julia/2018/Qual_GoGopher.jl) | Easy | Interactive |
|Qual| [Cubic UFO](https://codingcompetitions.withgoogle.com/codejam/round/00000000000000cb/00000000000079cc)|[Julia](./julia/2018/Qual_CubicUFO.jl) | Medium-Easy | |
|1A| [Waffle Choppers](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007883/000000000003005a)|[Julia](./julia/2018/1A_WaffleChoppers.jl) | Easy | |
|1A| [Bit Party](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007883/000000000002fff6)|[Julia](./julia/2018/1A_BitParty.jl) | Medium-Easy | |
|1A| [Edgy Baking](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007883/000000000002fff7)|[Julia](./julia/2018/1A_EdgyBaking.jl) | Medium | |
|1B| [Rounding Error](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007764/0000000000036601)|[Julia](./julia/2018/1B_RoundingError.jl) | Medium-Easy | |
|1B| [Mysterious Road Signs](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007764/000000000003675b)|[Julia](./julia/2018/1B_MysteriousRoadSigns.jl) | Medium | |
|1B| [Transmutation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007764/000000000003675c)|[Julia](./julia/2018/1B_Transmutation.jl) | Hard | |
|1C| [A Whole New Word](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007765/000000000003e064)|[Julia](./julia/2018/1C_AWholeNewWord.jl) | Easy | |
|1C| [Lollipop Shop](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007765/000000000003e068)|[Julia](./julia/2018/1C_LollipopShop.jl) | Easy | Interactive |
|1C| [Ant Stack](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007765/000000000003e0a8)|[Julia](./julia/2018/1C_AntStack.jl) | Medium | |
|2| [Falling Balls](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/00000000000459f2)|[Julia](./julia/2018/2_FallingBalls.jl) | Easy | |
|2| [Graceful Chansaw Jugglers](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/00000000000459f3)|[Julia](./julia/2018/2_GracefulChainsawJugglers.jl) | Medium-Easy | |
|2| [Costume Change](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/0000000000045875)|[Julia](./julia/2018/2_CostumeChange.jl) | Medium | |
|2| [Gridception](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007706/00000000000459f4)|[Julia](./julia/2018/2_Gridception.jl) | Medium-Hard | |
|3| [Field Trip](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007707/000000000004b7fe)|[Julia](./julia/2018/3_FieldTrip.jl) | Easy | |
|3| [Name-Preserving Network](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007707/000000000004ba29)|[Julia](./julia/2018/3_NamePreservingNetwork.jl) | Hard | Interactive |
|3| [Raise the Roof](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007707/000000000004b90d)|[Julia](./julia/2018/3_RaiseTheRoof.jl) | Hard | |
|3| [Fence Construction](https://codingcompetitions.withgoogle.com/codejam/round/0000000000007707/000000000004b90e)|[Julia](./julia/2018/3_FenceConstruction.jl) | Hard | |
|WF| [Jurisdiction Restrictions](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77c)|[Julia](./julia/2018/WF_JurisdictionRestrictions.jl) | Medium | |
|WF| [Two-Tiling](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77d)|[Julia](./julia/2018/WF_TwoTiling.jl) | Very Hard | |
|WF| [Go, Gophers!](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77e)|[Julia](./julia/2018/WF_GoGophers.jl) | Hard | Interactive |
|WF| [Swordmaster](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c77f)|[Julia](./julia/2018/WF_Swordmaster.jl) | Very Hard | |
|WF| [The Cartesian Job](https://codingcompetitions.withgoogle.com/codejam/round/0000000000051708/000000000016c933)|[Julia](./julia/2018/WF_TheCartesianJob.jl) | Very Hard | |

## Google Code Jam 2017
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2027-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-27%20%2F%2027-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2027-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Oversized Pancake Flipper](https://codingcompetitions.withgoogle.com/codejam/round/00000000002017f7/0000000000201847)|[Julia](./julia/2017/Qual_OversizedPancakeFlipper.jl) | Easy | |
|Qual| [Tidy Numbers](https://codingcompetitions.withgoogle.com/codejam/round/00000000002017f7/0000000000201878)|[Julia](./julia/2017/Qual_TidyNumbers.jl) | Easy | |
|Qual| [Bathroom Stalls](https://codingcompetitions.withgoogle.com/codejam/round/00000000002017f7/0000000000201905)|[Julia](./julia/2017/Qual_BathroomStalls.jl) | Easy | |
|Qual| [Fashion Show](https://codingcompetitions.withgoogle.com/codejam/round/00000000002017f7/00000000002017f8)|[Julia](./julia/2017/Qual_FashionShow.jl) | Medium | |
|1A| [Alphabet Cake](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201843/0000000000201875)|[Julia](./julia/2017/1A_AlphabetCake.jl) | Easy | |
|1A| [Ratatouille](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201843/00000000002018fe)|[Julia](./julia/2017/1A_Ratatouille.jl) | Easy | |
|1A| [Play the Dragon](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201843/00000000002017f3)|[Julia](./julia/2017/1A_PlayTheDragon.jl) | Hard | |
|1B| [Steed 2: Cruise Control](https://codingcompetitions.withgoogle.com/codejam/round/000000000020187f/000000000020190e)|[Julia](./julia/2017/1B_Steed2CruiseControl.jl) | Easy | |
|1B| [Stable Neigh-bors](https://codingcompetitions.withgoogle.com/codejam/round/000000000020187f/0000000000201804)|[Julia](./julia/2017/1B_StableNeighbors.jl) | Medium | |
|1B| [Pony Express](https://codingcompetitions.withgoogle.com/codejam/round/000000000020187f/000000000020184d)|[Julia](./julia/2017/1B_PonyExpress.jl) | Medium-Easy | |
|1C| [Ample Syrup](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201842/0000000000201874)|[Julia](./julia/2017/1C_AmpleSyrup.jl) | Easy | |
|1C| [Parenting Partnering](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201842/00000000002018fd)|[Julia](./julia/2017/1C_ParentingPartnering.jl) | Easy | |
|1C| [Core Training](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201842/00000000002017f2)|[Julia](./julia/2017/1C_CoreTraining.jl) | Hard | |
|2| [Fresh Chocolate](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201900/00000000002017f4)|[Julia](./julia/2017/2_FreshChocolate.jl) | Easy | |
|2| [Roller Coaster Scheduling](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201900/0000000000201845)|[Julia](./julia/2017/2_RollerCoasterScheduling.jl) | Medium-Easy | |
|2| [Beaming With Joy](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201900/0000000000201876)|[Julia](./julia/2017/2_BeamingWithJoy.jl) | Medium | |
|2| [Shoot the Turrets](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201900/0000000000201901)|[Julia](./julia/2017/2_ShootTheTurrets.jl) | Hard | |
|3| [Googlements](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201902/00000000002017f6)|[Julia](./julia/2017/3_Googlements.jl) | Easy | |
|3| [Good News and Bad News](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201902/0000000000201846)|[Julia](./julia/2017/3_GoodNewsAndBadNews.jl) | Medium | |
|3| [Mountain Tour](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201902/0000000000201877)|[Julia](./julia/2017/3_MountainTour.jl) | Hard | |
|3| [Slate Modern](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201902/0000000000201903)|[Julia](./julia/2017/3_SlateModern.jl) | Hard | |
|WF| [Dice Straight](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/00000000002017fc)|[Julia](./julia/2017/WF_DiceStriaght.jl) | Medium | |
|WF| [Operation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/000000000020184a)|[Julia](./julia/2017/WF_Operation.jl) | Medium | |
|WF| [Spanning Planning](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/000000000020187a)|[Julia](./julia/2017/WF_SpanningPlanning.jl) | Medium | |
|WF| [Omnicircumnavigation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/000000000020190a)|[Julia](./julia/2017/WF_Omnicircumnavigation.jl) | Hard | |
|WF| [Stack Management](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/00000000002017fd)|[Julia](./julia/2017/WF_StackManagement.jl) | Very Hard | |
|WF| [Teleporters](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201909/000000000020184b)|[Julia](./julia/2017/WF_Teleporters.jl) | Very Hard | |

## Google Code Jam 2016
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2026-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-26%20%2F%2026-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2026-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Counting Sheep](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bee/0000000000201c8a)|[Julia](./julia/2016/Qual_CountingSheep.jl) | Easy | |
|Qual| [Revenge of the Pancakes](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bee/0000000000201d17)|[Julia](./julia/2016/Qual_RevengeOfThePancakes.jl) | Easy | |
|Qual| [Coin Jam](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bee/0000000000201b6d)|[Julia](./julia/2016/Qual_CoinJam.jl) | Easy | |
|Qual| [Fractiles](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bee/0000000000201bf1)|[Julia](./julia/2016/Qual_Fractiles.jl) | Easy | |
|1A| [The Last Word](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf2/0000000000201c8d)|[Julia](./julia/2016/1A_TheLastWord.jl) | Easy | |
|1A| [Rank and File](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf2/0000000000201d1a)|[Julia](./julia/2016/1A_RankAndFile.jl) | Easy | |
|1A| [BFFs](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf2/0000000000201b6f)|[Julia](./julia/2016/1A_BFFs.jl) | Medium-Easy | |
|1B| [Getting the Digits](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b6c/0000000000201bf0)|[Julia](./julia/2016/1B_GettingTheDigits.jl) | Easy | |
|1B| [Close Match](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b6c/0000000000201c8c)|[Julia](./julia/2016/1B_CloseMatch.jl) | Medium-Easy | |
|1B| [Technobabble](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b6c/0000000000201d19)|[Julia](./julia/2016/1B_Technobabble.jl) | Medium | |
|1C| [Senate Evacuation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bef/0000000000201c8b)|[Julia](./julia/2016/1C_SenateEvacuation.jl) | Easy | |
|1C| [Slides!](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bef/0000000000201d18)|[Julia](./julia/2016/1C_Slides.jl) | Easy | |
|1C| [Fashion Police](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bef/0000000000201b6e)|[Julia](./julia/2016/1C_FashionPolice.jl) | Medium-Hard | |
|2| [Rather Perplexing Showdown](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201c91/0000000000201d1e)|[Julia](./julia/2016/2_RatherPerplexingShowdown.jl) | Easy | |
|2| [Red Tape Committee](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201c91/0000000000201b71)|[Julia](./julia/2016/2_RedTapeCommittee.jl) | Medium-Easy | |
|2| [The Gardner of Seville](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201c91/0000000000201bf5)|[Julia](./julia/2016/2_TheGardnerOfSeville.jl) | Hard | |
|2| [Freeform Factory](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201c91/0000000000201c92)|[Julia](./julia/2016/2_FreeformFactory.jl) | Hard | |
|3| [Teaching Assistant](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf3/0000000000201c90)|[Julia](./julia/2016/3_TeachingAssistant.jl) | Easy | |
|3| [Forest University](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf3/0000000000201d1c)|[Julia](./julia/2016/3_ForestUniversity.jl) | Medium | |
|3| [Rebel Against The Empire](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf3/0000000000201b70)|[Julia](./julia/2016/3_RebelAgainstTheEmpire.jl) | Hard | |
|3| [Go++](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201bf3/0000000000201bf4)|[Julia](./julia/2016/3_Gopp.jl) | Hard | |
|WF| [Family Hotel](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b72/0000000000201c93)|[Julia](./julia/2016/WF_FamilyHotel.jl) | Medium | |
|WF| [Integeregex](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b72/0000000000201bf8)|[Julia](./julia/2016/WF_Integeregex.jl) | Medium | |
|WF| [Gallery of Pillars](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b72/0000000000201d21)|[Julia](./julia/2016/WF_GalleryOfPillars.jl) | Hard | |
|WF| [Map Reduce](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b72/0000000000201b73)|[Julia](./julia/2016/WF_MapReduce.jl) | Hard | |
|WF| [Radioactive Islands](https://codingcompetitions.withgoogle.com/codejam/round/0000000000201b72/0000000000201bf9)|[Julia](./julia/2016/WF_RadioactiveIslands.jl) | Very Hard | |

## Google Code Jam 2015
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2028-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-28%20%2F%2028-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2028-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Standing Ovation](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433515/0000000000433738)|[Julia](./julia/2015/Qual_StandingOvation.jl) | Easy | |
|Qual| [Infinite House of Pancakes](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433515/0000000000433827)|[Julia](./julia/2015/Qual_InfiniteHouseOfPancakes.jl) | Easy | |
|Qual| [Dijkstra](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433515/0000000000433a60)|[Julia](./julia/2015/Qual_Dijkstra.jl) | Easy | |
|Qual| [Ominous Omino](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433515/0000000000433c83)|[Julia](./julia/2015/Qual_OminousOmino.jl) | Medium | |
|1A| [Mushroom Monster](https://codingcompetitions.withgoogle.com/codejam/round/00000000004336e9/0000000000433792)|[Julia](./julia/2015/1A_MushroomMonster.jl) | Easy | |
|1A| [Haircut](https://codingcompetitions.withgoogle.com/codejam/round/00000000004336e9/0000000000433602)|[Julia](./julia/2015/1A_Haircut.jl) | Easy | |
|1A| [Logging](https://codingcompetitions.withgoogle.com/codejam/round/00000000004336e9/0000000000433d3a)|[Julia](./julia/2015/1A_Logging.jl) | Medium-Hard | |
|1B| [Counter Culture](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433551/0000000000433a0f)|[Julia](./julia/2015/1B_CounterCulture.jl) | Medium-Easy | |
|1B| [Noisy Neighbors](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433551/0000000000433516)|[Julia](./julia/2015/1B_NoisyNeighbors.jl) | Medium | |
|1B| [Hiking Deer](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433551/0000000000433739)|[Julia](./julia/2015/1B_HikingDeer.jl) | Hard | |
|1C| [Brattleship](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433b4d/0000000000433de4)|[Julia](./julia/2015/1C_Brattleship.jl) | Easy | |
|1C| [Typewriter Monkey](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433b4d/0000000000433411)|[Julia](./julia/2015/1C_TypewriterMonkey.jl) | Medium | |
|1C| [Less Money, More Problems](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433b4d/0000000000433650)|[Julia](./julia/2015/1C_LessMoneyMoreProblems.jl) | Medium | |
|2| [Pegman](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433651/0000000000433552)|[Julia](./julia/2015/2_Pegman.jl) | Easy | |
|2| [Kiddie Pool](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433651/0000000000433a10)|[Julia](./julia/2015/2_KiddiePool.jl) | Medium-Hard | |
|2| [Bilingual](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433651/0000000000433517)|[Julia](./julia/2015/2_Bilingual.jl) | Hard | |
|2| [Drum Decorator](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433651/000000000043373a)|[Julia](./julia/2015/2_DrumDecorator.jl) | Hard | |
|3| [Fairland](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433c84/00000000004336ea)|[Julia](./julia/2015/3_Fairland.jl) | Easy | |
|3| [Smoothing Window](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433c84/0000000000433793)|[Julia](./julia/2015/3_SmoothingWindow.jl) | Easy | |
|3| [Runaway Quail](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433c84/0000000000433603)|[Julia](./julia/2015/3_RunawayQuail.jl) | Hard | |
|3| [Log Set](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433c84/0000000000433d3b)|[Julia](./julia/2015/3_LogSet.jl) | Medium-Hard | |
|3| [River Flow](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433c84/0000000000433918)|[Julia](./julia/2015/3_RiverFlow.jl) | Hard | |
|WF| [Campinatorics](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/0000000000433c8d)|[Julia](./julia/2015/WF_Campinatorics.jl) | Medium-Easy | |
|WF| [Costly Binary Search](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/0000000000433a69)|[Julia](./julia/2015/WF_CostlyBinarySearch.jl) | Medium | |
|WF| [Pretty Good Proportion](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/00000000004336f3)|[Julia](./julia/2015/WF_PrettyGoodProportion.jl) | Medium-Hard | |
|WF| [Taking Over The World](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/000000000043379b)|[Julia](./julia/2015/WF_TakingOverTheWorld.jl) | Very Hard | |
|WF| [Merlin QA](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/000000000043360e)|[Julia](./julia/2015/WF_MerlinQA.jl) | Hard | |
|WF| [Crane Truck](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433831/0000000000433d47)|[Julia](./julia/2015/WF_CraneTruck.jl) | Very Hard | |

## Google Code Jam 2014
![PythonProgress](https://img.shields.io/badge/PythonProgress-0%20%2F%2027-red.svg) ![JuliaProgress](https://img.shields.io/badge/JuliaProgress-27%20%2F%2027-green.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-0%20%2F%2027-red.svg)
| Round | Title | Solutions | Difficulty | Note |
|------ | ----- | --------- | ---------- | ---- |
|Qual| [Magic Trick](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432add/0000000000433092)|[Julia](./julia/2014/Qual_MagicTrick.jl) | Easy | |
|Qual| [Cookie Clicker Alpha](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432add/00000000004331a3)|[Julia](./julia/2014/Qual_CookieClickerAlpha.jl) | Easy | |
|Qual| [Deceitful War](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432add/0000000000432fec)|[Julia](./julia/2014/Qual_DeceitfulWar.jl) | Easy | |
|Qual| [Minesweeper Master](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432add/0000000000432e04)|[Julia](./julia/2014/Qual_MinesweeperMaster.jl) | Easy | |
|1A| [Charging Chaos](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433162/00000000004334c7)|[Julia](./julia/2014/1A_ChargingChaos.jl) | Easy | |
|1A| [Full Binary Tree](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433162/0000000000432a8f)|[Julia](./julia/2014/1A_FullBinaryTree.jl) | Easy | |
|1A| [Proper Shuffle](https://codingcompetitions.withgoogle.com/codejam/round/0000000000433162/0000000000432ade)|[Julia](./julia/2014/1A_ProperShuffle.jl) | Medium-Hard | |
|1B| [The Repeater](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432a90/0000000000432adf)|[Julia](./julia/2014/1B_TheRepeater.jl) | Easy | |
|1B| [New Lottery Game](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432a90/0000000000433096)|[Julia](./julia/2014/1B_NewLotteryGame.jl) | Medium | |
|1B| [The Bored Traveling Salesman](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432a90/00000000004331a4)|[Julia](./julia/2014/1B_BoredTravelingSalesman.jl) | Hard | |
|1C| [Part Elf](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432cd8/0000000000433163)|[Julia](./julia/2014/1C_PartElf.jl) | Easy | |
|1C| [Reordering Train Cars](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432cd8/00000000004334c8)|[Julia](./julia/2014/1C_ReorderingTrainCars.jl) | Medium | |
|1C| [Enclosure](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432cd8/0000000000432a91)|[Julia](./julia/2014/1C_Enclosure.jl) | Hard | |
|2| [Data Packing](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432fed/0000000000432b8d)|[Julia](./julia/2014/2_DataPacking.jl) | Easy | |
|2| [Up and Down](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432fed/000000000043333d)|[Julia](./julia/2014/2_UpAndDown.jl) | Easy | |
|2| [Don't Break The Nile](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432fed/0000000000433109)|[Julia](./julia/2014/2_DontBreakTheNile.jl) | Medium-Hard | |
|2| [Trie Sharding](https://codingcompetitions.withgoogle.com/codejam/round/0000000000432fed/0000000000432f41)|[Julia](./julia/2014/2_TrieSharding.jl) | Hard | |
|3| [Magical, Marvelous Tour](https://codingcompetitions.withgoogle.com/codejam/round/000000000043371f/000000000043380e)|[Julia](./julia/2014/3_MagicalMarvelousTour.jl) | Easy | |
|3| [Last Hit](https://codingcompetitions.withgoogle.com/codejam/round/000000000043371f/0000000000433a3e)|[Julia](./julia/2014/3_LastHit.jl) | Easy | |
|3| [Crime House](https://codingcompetitions.withgoogle.com/codejam/round/000000000043371f/00000000004331cb)|[Julia](./julia/2014/3_CrimeHouse.jl) | Hard | |
|3| [Willow](https://codingcompetitions.withgoogle.com/codejam/round/000000000043371f/00000000004336d0)|[Julia](./julia/2014/3_Willow.jl) | Very Hard | |
|WF| [Checkerboard Matrix](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/0000000000433537)|[Julia](./julia/2014/WF_CheckerboardMatrix.jl) | Medium-Easy | |
|WF| [Power Swapper](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/00000000004339f3)|[Julia](./julia/2014/WF_PowerSwapper.jl) | Medium-Easy | |
|WF| [Symmetric Trees](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/00000000004334f8)|[Julia](./julia/2014/WF_SymmetricTrees.jl) | Medium | |
|WF| [Paradox Sort](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/0000000000433720)|[Julia](./julia/2014/WF_ParadoxSort.jl) | Medium-Hard | |
|WF| [Alergy Testing](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/000000000043380f)|[Julia](./julia/2014/WF_AlergyTesting.jl) | Very Hard | |
|WF| [ARAM](https://codingcompetitions.withgoogle.com/codejam/round/000000000043363b/0000000000433a40)|[Julia](./julia/2014/WF_ARAM.jl) | Very Hard | |
