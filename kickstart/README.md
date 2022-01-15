# m1sterzer0 Google Kickstart Solutions ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

These are my solutions for Google's Kickstart problems. Please enjoy!

`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below.
## Contest Shortcuts
|     |     |     |     |     |
| --- | --- | --- | --- | --- |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1); [sol](#2013_A-Solutions) | [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7); [sol](#2013_B-Solutions) | [2014_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c); [sol](#2014_A-Solutions) | [2014_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947); [sol](#2014_B-Solutions) | [2014_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f); [sol](#2014_C-Solutions) |
| [2014_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e); [sol](#2014_D-Solutions) | [2015_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f); [sol](#2015_A-Solutions) | [2015_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a); [sol](#2015_B-Solutions) | [2015_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f); [sol](#2015_C-Solutions) | [2015_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac); [sol](#2015_D-Solutions) |
| [2015_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819); [sol](#2015_E-Solutions) | [2016_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2); [sol](#2016_A-Solutions) | [2016_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c); [sol](#2016_B-Solutions) | [2016_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a); [sol](#2016_C-Solutions) | [2016_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8); [sol](#2016_D-Solutions) |
| [2016_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0); [sol](#2016_E-Solutions) | [2017_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97); [sol](#2017_A-Solutions) | [2017_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27); [sol](#2017_B-Solutions) | [2017_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98); [sol](#2017_C-Solutions) | [2017_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77); [sol](#2017_D-Solutions) |
| [2017_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe); [sol](#2017_E-Solutions) | [2017_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29); [sol](#2017_F-Solutions) | [2017_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d); [sol](#2017_G-Solutions) | [2018_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf); [sol](#2018_A-Solutions) | [2018_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4); [sol](#2018_B-Solutions) |
| [2018_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0); [sol](#2018_C-Solutions) | [2018_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1); [sol](#2018_D-Solutions) | [2018_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5); [sol](#2018_E-Solutions) | [2018_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07); [sol](#2018_F-Solutions) | [2018_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066); [sol](#2018_G-Solutions) |
| [2018_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2); [sol](#2018_H-Solutions) | [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01); [sol](#2019_A-Solutions) | [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda); [sol](#2019_B-Solutions) | [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2); [sol](#2019_C-Solutions) | [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061); [sol](#2019_D-Solutions) |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb); [sol](#2019_E-Solutions) | [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc); [sol](#2019_F-Solutions) | [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02); [sol](#2019_G-Solutions) | [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd); [sol](#2019_H-Solutions) | [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7); [sol](#2020_A-Solutions) |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8); [sol](#2020_B-Solutions) | [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43); [sol](#2020_C-Solutions) | [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08); [sol](#2020_D-Solutions) | [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47); [sol](#2020_E-Solutions) | [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48); [sol](#2020_F-Solutions) |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069); [sol](#2020_G-Solutions) | [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49); [sol](#2020_H-Solutions) | [2021_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140); [sol](#2021_A-Solutions) | [2021_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b); [sol](#2021_B-Solutions) | [2021_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44); [sol](#2021_C-Solutions) |
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3); [sol](#2021_D-Solutions) | [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c); [sol](#2021_E-Solutions) | [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae); [sol](#2021_F-Solutions) | [2021_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6); [sol](#2021_G-Solutions) | [2021_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914); [sol](#2021_H-Solutions) |

## 2021_H Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914) | [Transform the String](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914/00000000008da461) | 3609 |  [go](./go/2021_H/TransformtheString/TransformtheString.go) | Simple modular arithmetic |
| [2021_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914) | [Painter](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914/00000000008d9a88) | 2464 |  [go](./go/2021_H/Painter/Painter.go) | Simple sequence problem |
| [2021_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914) | [Silly Substitutions](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914/00000000008d94f5) | 200 |  [go](./go/2021_H/SillySubstitutions/SillySubstitutions.go) | Linked lists and sets |
| [2021_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914) | [Dependent Events](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435914/00000000008d9970) | 63 |  [go](./go/2021_H/DependentEvents/DependentEvents.go) | Binary lifting. Conditional probabilities.  Modular inverses. |

## 2021_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6) | [Dogs and Cats](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6/00000000008b3771) | 6786 |  [go](./go/2021_G/DogsandCats/DogsandCats.go) | Very simple sequence processing |
| [2021_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6) | [Staying Hydrated](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6/00000000008b3a1c) | 1329 |  [go](./go/2021_G/StayingHydrated/StayingHydrated.go) | Sorting.  "Median-like" analysis. |
| [2021_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6) | [Banana Bunches](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6/00000000008b44ef) | 413 |  [go](./go/2021_G/BananaBunches/BananaBunches.go) | Tricky pairing problem. Cumulative sums. O(N^2) |
| [2021_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6) | [Simple Polygon](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004362d6/00000000008b36f9) | 148 |  [go](./go/2021_G/SimplePolygon/SimplePolygon.go) | Polygon construction |

## 2021_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Trash Bins](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000887c32) | 6349 |  [go](./go/2021_F/TrashBins/TrashBins.go) | Simple arithmetic series sums. |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Festival](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000887dba) | 1192 |  [go](./go/2021_F/Festival/Festival.go) | Bookkeeping with dueling multisets OR dueling segment trees |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Star Trappers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000888d45) | 176 |  [go](./go/2021_F/StarTrappers/StarTrappers.go) | Efficient pointInPolygon searching |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Graph Travel](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000888764) | 497 |  [go](./go/2021_F/GraphTravel/GraphTravel.go) | Subset DP |

## 2021_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c) | [Shuffled Anagrams](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c/000000000085a152) | 1594 |  [go](./go/2021_E/ShuffledAnagrams/ShuffledAnagrams.go) | Standard 'construct a derangement with repeats' problem |
| [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c) | [Birthday Cake](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c/000000000085a285) | 64 |  [go](./go/2021_E/BirthdayCake/BirthdayCake.go) | Challenging basic problem.  No fancy data structures or algorithms needed. |
| [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c) | [Palindromic Crossword](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c/0000000000859dcd) | 908 |  [go](./go/2021_E/PalindromicCrossword/PalindromicCrossword.go) | DSU. Palindromes |
| [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c) | [Increasing Sequence Card Game](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c/000000000085a709) | 505 |  [go](./go/2021_E/IncreasingSequenceCardGame/IncreasingSequenceCardGame.go) | DP for mid version.  Harmonic series approximation.  Euler-Maclaurin Summation Formula (for extra credit). |

## 2021_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3) | [Arithmetic Square](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3/000000000082b813) | 6123 |  [go](./go/2021_D/ArithmeticSquare/ArithmeticSquare.go) |  |
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3) | [Cutting Intervals](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3/000000000082b933) | 1468 |  [go](./go/2021_D/CuttingIntervals/CuttingIntervals.go) | Interval sorting. Event processing |
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3) | [Final Exam](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3/000000000082bffc) | 1009 |  [go](./go/2021_D/FinalExam/FinalExam.go) | Searching ordered sets of intervals |
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3) | [Primes and Queries](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3/000000000082bcf4) | 35 |  [go](./go/2021_D/PrimesandQueries/PrimesandQueries.go) | Esoteric 'lifting the exponent' (LTE). Fenwick trees or segment trees |

## 2021_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44) | [Smaller Strings](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ebe5e) | 2657 |  [go](./go/2021_C/SmallerStrings/SmallerStrings.go) |  |
| [2021_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44) | [Alien Generator](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ec1cb) | 4750 |  [go](./go/2021_C/AlienGenerator/AlienGenerator.go) |  |
| [2021_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44) | [Rock Paper Scissors](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ec28e) | 242 |  [go](./go/2021_C/RockPaperScissors/RockPaperScissors.go) | Expected value DP |
| [2021_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44) | [Binary Operator](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435c44/00000000007ec290) | 79 |  [go](./go/2021_C/BinaryOperator/BinaryOperator.go) | Expression trees, random |

## 2021_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b) | [Increasing Substring](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b/000000000077a882) | 6049 |  [go](./go/2021_B/IncreasingSubstring/IncreasingSubstring.go) |  |
| [2021_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b) | [Longest Progression](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b/000000000077a3a5) | 1051 |  [go](./go/2021_B/LongestProgression/LongestProgression.go) | Simple DP. Bookkeeping. |
| [2021_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b) | [Consecutive Primes](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b/000000000077a8e6) | 1389 |  [go](./go/2021_B/ConsecutivePrimes/ConsecutivePrimes.go) | Binary search. Primality testing. |
| [2021_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b) | [Truck Delivery](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435a5b/000000000077a885) | 250 |  [go](./go/2021_B/TruckDelivery/TruckDelivery.go) | Tree DP. Segment tree. |

## 2021_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140) | [K-Goodness String](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140/000000000068cca3) | 13336 |  [go](./go/2021_A/KGoodnessString/KGoodnessString.go) |  |
| [2021_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140) | [L Shaped Plots](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140/000000000068c509) | 4869 |  [go](./go/2021_A/LShapedPlots/LShapedPlots.go) | Grid counting. |
| [2021_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140) | [Rabbit House](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140/000000000068cb14) | 2807 |  [go](./go/2021_A/RabbitHouse/RabbitHouse.go) | Dijkstra or clever BFS for O(N). |
| [2021_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140) | [Checksum](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140/000000000068c2c3) | 116 |  [go](./go/2021_A/Checksum/Checksum.go) | MST |

## 2020_H Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Retype](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043adc7) | 4071 |  [go](./go/2020_H/Retype/Retype.go) |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Boring Numbers](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043b0c6) | 805 |  [go](./go/2020_H/BoringNumbers/BoringNumbers.go) |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Rugby](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043b027) | 320 |  [go](./go/2020_H/Rugby/Rugby.go) |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Friends](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043aee7) | 262 |  [go](./go/2020_H/Friends/Friends.go) |  |

## 2020_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Kick_Start](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414bfb) | 5780 |  [go](./go/2020_G/KickStart/KickStart.go) |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Maximum Coins](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414a23) | 5590 |  [go](./go/2020_G/MaximumCoins/MaximumCoins.go) |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Combination Lock](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414a24) | 606 |  [go](./go/2020_G/CombinationLock/CombinationLock.go) |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Merge Cards](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000415054) | 248 |  [go](./go/2020_G/MergeCards/MergeCards.go) |  |

## 2020_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [ATM Queue](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4ed8) | 3509 |  [go](./go/2020_F/ATMQueue/ATMQueue.go) |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Metal Harvest](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4b8b) | 2448 |  [go](./go/2020_F/MetalHarvest/MetalHarvest.go) |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Painters' Duel](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f47fb) | 346 |  [go](./go/2020_F/PaintersDuel/PaintersDuel.go) |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Yeetzhee](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4dea) | 69 |  [go](./go/2020_F/Yeetzhee/Yeetzhee.go) |  |

## 2020_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Longest Arithmetic](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bf4ed) | 8888 |  [go](./go/2020_E/LongestArithmetic/LongestArithmetic.go) |  |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [High Buildings](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bef73) | 2345 |  [go](./go/2020_E/HighBuildings/HighBuildings.go) | Constructive. Casework |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Toys](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bede9) | 288 |  [go](./go/2020_E/Toys/Toys.go) | Fenwick + Min Heap.  I need to reread the official solutions, as I don't follow their method. |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Golden Stone](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bef29) | 77 |  [go](./go/2020_E/GoldenStone/GoldenStone.go) | Tricky Dijkstra |

## 2020_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Record Breaker](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000387171) | 7799 |  [go](./go/2020_D/RecordBreaker/RecordBreaker.go) |  |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Alien Piano](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000387174) | 3928 |  [go](./go/2020_D/AlienPiano/AlienPiano.go) |  |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Beauty of tree](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000386edd) | 602 |  [go](./go/2020_D/Beautyoftree/Beautyoftree.go) | Tree DP, Stack |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Locked Doors](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000386d5c) | 112 |  [go](./go/2020_D/LockedDoors/LockedDoors.go) | Tree construction, Binary lifting |

## 2020_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Countdown](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003380d2) | 9017 |  [go](./go/2020_C/Countdown/Countdown.go) |  |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Stable Wall](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003379bb) | 2632 |  [go](./go/2020_C/StableWall/StableWall.go) | Topological sort |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Perfect Subarray](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003381cb) | 685 |  [go](./go/2020_C/PerfectSubarray/PerfectSubarray.go) | Sparse prefix sums (pseudo DP) |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Candies](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/0000000000337b4d) | 655 |  [go](./go/2020_C/Candies/Candies.go) | Fenwick tree (clever) |

## 2020_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Bike Tour](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d82e6) | 9267 |  [go](./go/2020_B/BikeTour/BikeTour.go) |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Bus Routes](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d83bf) | 5811 |  [go](./go/2020_B/BusRoutes/BusRoutes.go) |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Robot Path Decoding](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d83dc) | 2733 |  [go](./go/2020_B/RobotPathDecoding/RobotPathDecoding.go) |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Wandering Robot](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d8565) | 135 |  [go](./go/2020_B/WanderingRobot/WanderingRobot.go) | Combinatorics, Logarithms |

## 2020_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Allocation](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3f56) | 10276 |  [go](./go/2020_A/Allocation/Allocation.go) |  |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Plates](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d40bb) | 2941 |  [go](./go/2020_A/Plates/Plates.go) | Simple DP |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Workout](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3f5b) | 1925 |  [go](./go/2020_A/Workout/Workout.go) | Binary Search |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Bundling](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3ff3) | 1032 |  [go](./go/2020_A/Bundling/Bundling.go) | Greedy Trie |

## 2019_H Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [H-index](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a274e) | 877 |  [go](./go/2019_H/Hindex/Hindex.go) |  |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [Diagonal Puzzle](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a2835) | 124 |  [go](./go/2019_H/DiagonalPuzzle/DiagonalPuzzle.go) |  |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [Elevanagram](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a286d) | 184 |  [go](./go/2019_H/Elevanagram/Elevanagram.go) |  |

## 2019_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [Book Reading](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fd0d) | 1067 |  [go](./go/2019_G/BookReading/BookReading.go) |  |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [The Equation](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fe36) | 569 |  [go](./go/2019_G/TheEquation/TheEquation.go) |  |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [Shifts](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fd5e) | 291 |  [go](./go/2019_G/Shifts/Shifts.go) |  |

## 2019_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Flattening](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/000000000018666c) | 320 |  [go](./go/2019_F/Flattening/Flattening.go) |  |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Teach Me](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/00000000001864bc) | 179 |  [go](./go/2019_F/TeachMe/TeachMe.go) |  |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Spectating Villages](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/000000000018666b) | 143 |  [go](./go/2019_F/SpectatingVillages/SpectatingVillages.go) |  |

## 2019_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Cherries Mesh](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/0000000000170721) | 1570 |  [go](./go/2019_E/CherriesMesh/CherriesMesh.go) |  |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Code-Eat Switcher](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/00000000001707b8) | 249 |  [go](./go/2019_E/CodeEatSwitcher/CodeEatSwitcher.go) |  |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Street Checkers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/00000000001707b9) | 435 |  [go](./go/2019_E/StreetCheckers/StreetCheckers.go) |  |

## 2019_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [X or What?](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161426) | 450 |  [go](./go/2019_D/XorWhat/XorWhat.go) |  |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [Latest Guests](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161427) | 110 |  [go](./go/2019_D/LatestGuests/LatestGuests.go) |  |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [Food Stalls](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161476) | 52 |  [go](./go/2019_D/FoodStalls/FoodStalls.go) |  |

## 2019_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Wiggle Walk](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac) | 569 |  [go](./go/2019_C/WiggleWalk/WiggleWalk.go) |  |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Circuit Board](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aae) | 553 |  [go](./go/2019_C/CircuitBoard/CircuitBoard.go) |  |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Catch Some](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150a0d) | 161 |  [go](./go/2019_C/CatchSome/CatchSome.go) |  |

## 2019_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Building Palindromes](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/0000000000119866) | 1011 |  [go](./go/2019_B/BuildingPalindromes/BuildingPalindromes.go) |  |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Energy Stones](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/00000000001198c3) | 65 |  [go](./go/2019_B/EnergyStones/EnergyStones.go) |  |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Diverse Subarray](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/00000000001198c1) | 81 |  [go](./go/2019_B/DiverseSubarray/DiverseSubarray.go) |  |

## 2019_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Training](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/00000000000698d6) | 2112 |  [go](./go/2019_A/Training/Training.go) |  |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Parcels](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/000000000006987d) | 137 |  [go](./go/2019_A/Parcels/Parcels.go) |  |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Contention](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/0000000000069881) | 2 |  [go](./go/2019_A/Contention/Contention.go) |  |

## 2018_H Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2) | [Big Buttons](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2/0000000000051136) | 0 |  |  |
| [2018_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2) | [Mural](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2/000000000005118a) | 0 |  |  |
| [2018_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2) | [Let Me Count The Ways](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee2/0000000000051189) | 0 |  |  |

## 2018_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066) | [Product Triplets](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066/0000000000051187) | 0 |  |  |
| [2018_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066) | [Combining Classes](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066/0000000000051007) | 0 |  |  |
| [2018_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066) | [Cave Escape](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051066/0000000000051135) | 0 |  |  |

## 2018_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07) | [Common Anagrams](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07/00000000000510f2) | 0 |  |  |
| [2018_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07) | [Specializing Villages](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07/0000000000051134) | 0 |  |  |
| [2018_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07) | [Palindromic Sequence](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e07/0000000000051186) | 0 |  |  |

## 2018_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5) | [Board Game](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5/0000000000051184) | 0 |  |  |
| [2018_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5) | [Milk Tea](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5/0000000000051185) | 0 |  |  |
| [2018_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5) | [Yogurt](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff5/00000000000510f1) | 0 |  |  |

## 2018_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1) | [Candies](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1/00000000000510ef) | 0 |  |  |
| [2018_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1) | [Paragliding](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1/0000000000051006) | 0 |  |  |
| [2018_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1) | [Funniest Word Search](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee1/00000000000510f0) | 0 |  |  |

## 2018_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0) | [Planet Distance](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0/0000000000051005) | 0 |  |  |
| [2018_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0) | [Fairies and Witches](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0/0000000000051132) | 0 |  |  |
| [2018_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0) | [Kickstart Alarm](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ee0/0000000000051133) | 0 |  |  |

## 2018_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4) | [No Nine](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4/0000000000051183) | 0 |  |  |
| [2018_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4) | [Sherlock and the Bit Strings](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4/000000000005107b) | 0 |  |  |
| [2018_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4) | [King's Circle](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff4/00000000000510ee) | 0 |  |  |

## 2018_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2018_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf) | [Even Digits](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf/00000000000510ed) | 0 |  |  |
| [2018_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf) | [Lucky Dip](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf/0000000000050e1d) | 0 |  |  |
| [2018_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf) | [Scrambled Words](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edf/0000000000051004) | 0 |  |  |

## 2017_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d) | [Huge Numbers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d/0000000000201c03) | 0 |  |  |
| [2017_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d) | [Cards Game](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d/0000000000201c9c) | 0 |  |  |
| [2017_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d) | [Matrix Cutting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b7d/0000000000201d2b) | 0 |  |  |

## 2017_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29) | [Cake](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29/0000000000201d2a) | 0 |  |  |
| [2017_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29) | [Kicksort](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29/0000000000201b7c) | 0 |  |  |
| [2017_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29) | [Dance Battle](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29/0000000000201c02) | 0 |  |  |
| [2017_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29) | [Catch Them All](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d29/0000000000201c9b) | 0 |  |  |

## 2017_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe) | [Trapezoid Counting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe/0000000000201d24) | 0 |  |  |
| [2017_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe) | [Copy & Paste](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe/0000000000201c96) | 0 |  |  |
| [2017_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe) | [Blackhole](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201bfe/0000000000201b78) | 0 |  |  |

## 2017_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77) | [Sightseeing](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77/0000000000201bfd) | 0 |  |  |
| [2017_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77) | [Sherlock and Matrix Game](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77/0000000000201c95) | 0 |  |  |
| [2017_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77) | [Trash](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201b77/0000000000201d23) | 0 |  |  |

## 2017_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98) | [Ambiguous Cipher](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98/0000000000201d26) | 0 |  |  |
| [2017_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98) | [X Squared](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98/0000000000201b7a) | 0 |  |  |
| [2017_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98) | [Magical Thinking v2](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98/0000000000201c00) | 0 |  |  |
| [2017_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98) | [The 4M Corporation](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c98/0000000000201c99) | 0 |  |  |

## 2017_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27) | [Math Encoder](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27/0000000000201b7b) | 0 |  |  |
| [2017_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27) | [Center](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27/0000000000201c01) | 0 |  |  |
| [2017_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27) | [Christmas Tree](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201d27/0000000000201c9a) | 0 |  |  |

## 2017_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2017_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97) | [Square Counting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97/0000000000201d25) | 0 |  |  |
| [2017_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97) | [Pattern Overlap](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97/0000000000201b79) | 0 |  |  |
| [2017_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97) | [Two Cubes](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c97/0000000000201bff) | 0 |  |  |

## 2016_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2016_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0) | [Diwali lightings](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0/0000000000201d2f) | 0 |  |  |
| [2016_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0) | [Beautiful Numbers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0/0000000000201dba) | 0 |  |  |
| [2016_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0) | [Partioning Number](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0/0000000000201c08) | 0 |  |  |
| [2016_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0) | [Sorting Array](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca0/0000000000201ca1) | 0 |  |  |

## 2016_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2016_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8) | [Vote](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8/0000000000201c06) | 0 |  |  |
| [2016_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8) | [Sitting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8/0000000000201c9e) | 0 |  |  |
| [2016_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8) | [Codejamon Cipher](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8/0000000000201d2e) | 0 |  |  |
| [2016_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8) | [Stretch Rope](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201db8/0000000000201db9) | 0 |  |  |

## 2016_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2016_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a) | [Monster Path](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a/0000000000201ca4) | 0 |  |  |
| [2016_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a) | [Safe Squares](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a/0000000000201d31) | 0 |  |  |
| [2016_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a) | [Evaluation](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a/0000000000201dbc) | 0 |  |  |
| [2016_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a) | [Soldiers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0a/0000000000201c0b) | 0 |  |  |

## 2016_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2016_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c) | [Sherlock and Parentheses](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c/0000000000201ca5) | 0 |  |  |
| [2016_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c) | [Sherlock and Watson Gym Secrets](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c/0000000000201d32) | 0 |  |  |
| [2016_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c) | [Watson and Intervals](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c/0000000000201dbd) | 0 |  |  |
| [2016_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c) | [Sherlock and Permutation Sorting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201c0c/0000000000201c0d) | 0 |  |  |

## 2016_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2016_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2) | [Country Leader](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2/0000000000201d30) | 0 |  |  |
| [2016_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2) | [Rain](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2/0000000000201dbb) | 0 |  |  |
| [2016_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2) | [Jane's Flower Shop](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2/0000000000201c09) | 0 |  |  |
| [2016_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2) | [Clash Royale](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000201ca2/0000000000201ca3) | 0 |  |  |

## 2015_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2015_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819) | [Lazy Spelling Bee](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819/0000000000434b41) | 0 |  |  |
| [2015_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819) | [Robot Rock Band](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819/0000000000434ae0) | 0 |  |  |
| [2015_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819) | [Not So Random](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819/00000000004347b8) | 0 |  |  |
| [2015_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819) | [Sums of Sums](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434819/00000000004348e6) | 0 |  |  |

## 2015_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2015_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac) | [Dynamic Grid](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac/0000000000434c6c) | 0 |  |  |
| [2015_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac) | [gBalloon](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac/0000000000434dfe) | 0 |  |  |
| [2015_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac) | [IP Address Summarization](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac/0000000000434880) | 0 |  |  |
| [2015_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac) | [Virtual Rabbit](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004349ac/0000000000434949) | 0 |  |  |

## 2015_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2015_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f) | [gRanks](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f/0000000000434948) | 0 |  |  |
| [2015_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f) | [gFiles](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f/000000000043474d) | 0 |  |  |
| [2015_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f) | [gGames](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f/0000000000434da1) | 0 |  |  |
| [2015_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f) | [gMatrix](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043487f/0000000000434c0f) | 0 |  |  |

## 2015_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2015_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a) | [Travel](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a/0000000000434ba8) | 0 |  |  |
| [2015_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a) | [gWheels](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a/0000000000434818) | 0 |  |  |
| [2015_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a) | [gNumbers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a/0000000000434b40) | 0 |  |  |
| [2015_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a) | [Albocede DNA](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434a7a/0000000000434add) | 0 |  |  |

## 2015_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2015_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f) | [Googol String](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f/0000000000434adc) | 0 |  |  |
| [2015_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f) | [gCube](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f/00000000004347b7) | 0 |  |  |
| [2015_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f) | [gCampus](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f/00000000004348e5) | 0 |  |  |
| [2015_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f) | [gSnake](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3f/00000000004349ab) | 0 |  |  |

## 2014_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2014_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e) | [Cube IV](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e/0000000000434adb) | 0 |  |  |
| [2014_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e) | [GBus count](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e/00000000004347b6) | 0 |  |  |
| [2014_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e) | [Sort a scrambled itinerary](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e/00000000004348e4) | 0 |  |  |
| [2014_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e) | [Itz Chess](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434b3e/00000000004349aa) | 0 |  |  |

## 2014_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2014_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f) | [Minesweeper](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f/0000000000434c0c) | 0 |  |  |
| [2014_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f) | [Taking Metro](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f/0000000000434cd2) | 0 |  |  |
| [2014_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f) | [Broken Calculator](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f/0000000000434a77) | 0 |  |  |
| [2014_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f) | [Tetris](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9f/0000000000434ba3) | 0 |  |  |

## 2014_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2014_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947) | [Password Attacker](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947/000000000043474c) | 0 |  |  |
| [2014_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947) | [New Years Eve](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947/0000000000434d9d) | 0 |  |  |
| [2014_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947) | [Card Game](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947/0000000000434c0b) | 0 |  |  |
| [2014_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947) | [Parentheses Order](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434947/0000000000434cd1) | 0 |  |  |

## 2014_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2014_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c) | [Super 2048](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c/0000000000434cd0) | 0 |  |  |
| [2014_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c) | [Seven-segment Display](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c/0000000000434c09) | 0 |  |  |
| [2014_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c) | [Cut Tiles](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c/0000000000434ba2) | 0 |  |  |
| [2014_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c) | [Addition](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434d9c/0000000000434a76) | 0 |  |  |

## 2013_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7) | [Sudoku Checker](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7/00000000004347b3) | 0 |  |  |
| [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7) | [Ignore all my comments](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7/0000000000434dfc) | 0 |  |  |
| [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7) | [Dragon Maze](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7/0000000000434c67) | 0 |  |  |
| [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7) | [Meet and party](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7/00000000004348e0) | 0 |  |  |
| [2013_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7) | [Hex](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ad7/00000000004349a6) | 0 |  |  |

## 2013_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1) | [Sorting](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1/0000000000434ad6) | 0 |  |  |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1) | [Read Phone Number](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1/0000000000434813) | 0 |  |  |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1) | [Rational Number Tree](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1/0000000000434b3c) | 0 |  |  |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1) | [Cross the maze](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1/00000000004347b2) | 0 |  |  |
| [2013_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1) | [Spaceship Defence](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000434ba1/00000000004348df) | 0 |  |  |

