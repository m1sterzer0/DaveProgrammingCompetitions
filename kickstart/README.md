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
| [2021_D](https://codingcompetitions.withgoogle.com/kickstart/round/00000000004361e3); [sol](#2021_D-Solutions) | [2021_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000043585c); [sol](#2021_E-Solutions) | [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae); [sol](#2021_F-Solutions) | | |

## 2021_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Trash Bins](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000887c32) | 6349 |  |  |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Festival](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000887dba) | 1192 |  |  |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Star Trappers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000888d45) | 176 |  |  |
| [2021_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae) | [Graph Travel](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000435bae/0000000000888764) | 497 |  |  |

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
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Retype](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043adc7) | 0 |  |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Boring Numbers](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043b0c6) | 0 |  |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Rugby](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043b027) | 0 |  |  |
| [2020_H](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49) | [Friends](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff49/000000000043aee7) | 0 |  |  |

## 2020_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Kick_Start](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414bfb) | 0 |  |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Maximum Coins](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414a23) | 0 |  |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Combination Lock](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000414a24) | 0 |  |  |
| [2020_G](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069) | [Merge Cards](https://codingcompetitions.withgoogle.com/kickstart/round/00000000001a0069/0000000000415054) | 0 |  |  |

## 2020_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [ATM Queue](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4ed8) | 0 |  |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Metal Harvest](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4b8b) | 0 |  |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Painters' Duel](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f47fb) | 0 |  |  |
| [2020_F](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48) | [Yeetzhee](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4dea) | 0 |  |  |

## 2020_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Longest Arithmetic](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bf4ed) | 0 |  |  |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [High Buildings](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bef73) | 0 |  |  |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Toys](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bede9) | 0 |  |  |
| [2020_E](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47) | [Golden Stone](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff47/00000000003bef29) | 0 |  |  |

## 2020_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Record Breaker](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000387171) | 0 |  |  |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Alien Piano](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000387174) | 0 |  |  |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Beauty of tree](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000386edd) | 0 |  |  |
| [2020_D](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08) | [Locked Doors](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff08/0000000000386d5c) | 0 |  |  |

## 2020_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Countdown](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003380d2) | 0 |  |  |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Stable Wall](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003379bb) | 0 |  |  |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Perfect Subarray](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/00000000003381cb) | 0 |  |  |
| [2020_C](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43) | [Candies](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/0000000000337b4d) | 0 |  |  |

## 2020_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Bike Tour](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d82e6) | 0 |  |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Bus Routes](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d83bf) | 0 |  |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Robot Path Decoding](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d83dc) | 0 |  |  |
| [2020_B](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8) | [Wandering Robot](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc8/00000000002d8565) | 0 |  |  |

## 2020_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Allocation](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3f56) | 0 |  |  |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Plates](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d40bb) | 0 |  |  |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Workout](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3f5b) | 0 |  |  |
| [2020_A](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7) | [Bundling](https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ffc7/00000000001d3ff3) | 0 |  |  |

## 2019_H Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [H-index](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a274e) | 0 |  |  |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [Diagonal Puzzle](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a2835) | 0 |  |  |
| [2019_H](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd) | [Elevanagram](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edd/00000000001a286d) | 0 |  |  |

## 2019_G Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [Book Reading](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fd0d) | 0 |  |  |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [The Equation](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fe36) | 0 |  |  |
| [2019_G](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02) | [Shifts](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e02/000000000018fd5e) | 0 |  |  |

## 2019_F Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Flattening](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/000000000018666c) | 0 |  |  |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Teach Me](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/00000000001864bc) | 0 |  |  |
| [2019_F](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc) | [Spectating Villages](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edc/000000000018666b) | 0 |  |  |

## 2019_E Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Cherries Mesh](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/0000000000170721) | 0 |  |  |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Code-Eat Switcher](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/00000000001707b8) | 0 |  |  |
| [2019_E](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb) | [Street Checkers](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050edb/00000000001707b9) | 0 |  |  |

## 2019_D Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [X or What?](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161426) | 0 |  |  |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [Latest Guests](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161427) | 0 |  |  |
| [2019_D](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061) | [Food Stalls](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000051061/0000000000161476) | 0 |  |  |

## 2019_C Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Wiggle Walk](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac) | 0 |  |  |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Circuit Board](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aae) | 0 |  |  |
| [2019_C](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2) | [Catch Some](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150a0d) | 0 |  |  |

## 2019_B Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Building Palindromes](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/0000000000119866) | 0 |  |  |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Energy Stones](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/00000000001198c3) | 0 |  |  |
| [2019_B](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda) | [Diverse Subarray](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050eda/00000000001198c1) | 0 |  |  |

## 2019_A Solutions
| Contest | Problem | Num Correct | Solutions | Notes |
| ------- | ------- | ----------: | --------- | ----- |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Training](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/00000000000698d6) | 0 |  |  |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Parcels](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/000000000006987d) | 0 |  |  |
| [2019_A](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01) | [Contention](https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050e01/0000000000069881) | 0 |  |  |

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

