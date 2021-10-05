# m1sterzer0 Atcoder ABC Solutions ![Language](https://img.shields.io/badge/language-Python-orange.svg) ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) ![PythonProgress](https://img.shields.io/badge/PythonProgress-274%20%2F%20392-ff69b4.svg) ![GolangProgress](https://img.shields.io/badge/GolangProgress-392%20%2F%20392-ff69b4.svg)

These are the solutions for the Atcoder beginner contests (ABC).  I find the Atcoder beginner contests very educational.  There is a large difficulty spread in the problems so that you can ramp, and the 500/600 point solutions often introduce new (even advanced) concepts in a reasonably straightforward way.

`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below.

`NOTE`: Problem difficulty ratings in the table below were obtained from kenkoooo.com.
## Contest Shortcuts
|     |     |     |     |     |
| --- | --- | --- | --- | --- |
| [abc160](https://atcoder.jp/contests/abc160); [sol](#abc160-Solutions) | [abc161](https://atcoder.jp/contests/abc161); [sol](#abc161-Solutions) | [abc162](https://atcoder.jp/contests/abc162); [sol](#abc162-Solutions) | [abc163](https://atcoder.jp/contests/abc163); [sol](#abc163-Solutions) | [abc164](https://atcoder.jp/contests/abc164); [sol](#abc164-Solutions) |
| [abc165](https://atcoder.jp/contests/abc165); [sol](#abc165-Solutions) | [abc166](https://atcoder.jp/contests/abc166); [sol](#abc166-Solutions) | [abc167](https://atcoder.jp/contests/abc167); [sol](#abc167-Solutions) | [abc168](https://atcoder.jp/contests/abc168); [sol](#abc168-Solutions) | [abc169](https://atcoder.jp/contests/abc169); [sol](#abc169-Solutions) |
| [abc170](https://atcoder.jp/contests/abc170); [sol](#abc170-Solutions) | [abc171](https://atcoder.jp/contests/abc171); [sol](#abc171-Solutions) | [abc172](https://atcoder.jp/contests/abc172); [sol](#abc172-Solutions) | [abc173](https://atcoder.jp/contests/abc173); [sol](#abc173-Solutions) | [abc174](https://atcoder.jp/contests/abc174); [sol](#abc174-Solutions) |
| [abc175](https://atcoder.jp/contests/abc175); [sol](#abc175-Solutions) | [abc176](https://atcoder.jp/contests/abc176); [sol](#abc176-Solutions) | [abc177](https://atcoder.jp/contests/abc177); [sol](#abc177-Solutions) | [abc178](https://atcoder.jp/contests/abc178); [sol](#abc178-Solutions) | [abc179](https://atcoder.jp/contests/abc179); [sol](#abc179-Solutions) |
| [abc180](https://atcoder.jp/contests/abc180); [sol](#abc180-Solutions) | [abc181](https://atcoder.jp/contests/abc181); [sol](#abc181-Solutions) | [abc182](https://atcoder.jp/contests/abc182); [sol](#abc182-Solutions) | [abc183](https://atcoder.jp/contests/abc183); [sol](#abc183-Solutions) | [abc184](https://atcoder.jp/contests/abc184); [sol](#abc184-Solutions) |
| [abc185](https://atcoder.jp/contests/abc185); [sol](#abc185-Solutions) | [abc186](https://atcoder.jp/contests/abc186); [sol](#abc186-Solutions) | [abc187](https://atcoder.jp/contests/abc187); [sol](#abc187-Solutions) | [abc188](https://atcoder.jp/contests/abc188); [sol](#abc188-Solutions) | [abc189](https://atcoder.jp/contests/abc189); [sol](#abc189-Solutions) |
| [abc190](https://atcoder.jp/contests/abc190); [sol](#abc190-Solutions) | [abc191](https://atcoder.jp/contests/abc191); [sol](#abc191-Solutions) | [abc192](https://atcoder.jp/contests/abc192); [sol](#abc192-Solutions) | [abc193](https://atcoder.jp/contests/abc193); [sol](#abc193-Solutions) | [abc194](https://atcoder.jp/contests/abc194); [sol](#abc194-Solutions) |
| [abc195](https://atcoder.jp/contests/abc195); [sol](#abc195-Solutions) | [abc196](https://atcoder.jp/contests/abc196); [sol](#abc196-Solutions) | [abc197](https://atcoder.jp/contests/abc197); [sol](#abc197-Solutions) | [abc198](https://atcoder.jp/contests/abc198); [sol](#abc198-Solutions) | [abc199](https://atcoder.jp/contests/abc199); [sol](#abc199-Solutions) |
| [abc200](https://atcoder.jp/contests/abc200); [sol](#abc200-Solutions) | [abc201](https://atcoder.jp/contests/abc201); [sol](#abc201-Solutions) | [abc202](https://atcoder.jp/contests/abc202); [sol](#abc202-Solutions) | [abc203](https://atcoder.jp/contests/abc203); [sol](#abc203-Solutions) | [abc204](https://atcoder.jp/contests/abc204); [sol](#abc204-Solutions) |
| [abc205](https://atcoder.jp/contests/abc205); [sol](#abc205-Solutions) | [abc206](https://atcoder.jp/contests/abc206); [sol](#abc206-Solutions) | [abc207](https://atcoder.jp/contests/abc207); [sol](#abc207-Solutions) | [abc208](https://atcoder.jp/contests/abc208); [sol](#abc208-Solutions) | [abc209](https://atcoder.jp/contests/abc209); [sol](#abc209-Solutions) |
| [abc210](https://atcoder.jp/contests/abc210); [sol](#abc210-Solutions) | [abc211](https://atcoder.jp/contests/abc211); [sol](#abc211-Solutions) | [abc212](https://atcoder.jp/contests/abc212); [sol](#abc212-Solutions) | [abc213](https://atcoder.jp/contests/abc213); [sol](#abc213-Solutions) | [abc214](https://atcoder.jp/contests/abc214); [sol](#abc214-Solutions) |
| [abc215](https://atcoder.jp/contests/abc215); [sol](#abc215-Solutions) | [abc216](https://atcoder.jp/contests/abc216); [sol](#abc216-Solutions) | [abc217](https://atcoder.jp/contests/abc217); [sol](#abc217-Solutions) | [abc218](https://atcoder.jp/contests/abc218); [sol](#abc218-Solutions) | [abc219](https://atcoder.jp/contests/abc219); [sol](#abc219-Solutions) |
| [abc220](https://atcoder.jp/contests/abc220); [sol](#abc220-Solutions) | [abc221](https://atcoder.jp/contests/abc221); [sol](#abc221-Solutions) | | | |

## abc221 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_A](http:/atcoder.jp/contests/abc221/tasks/abc221_A) | 8032 | 0 |  [go](./go/abc221/abc221_A/abc221_A.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_B](http:/atcoder.jp/contests/abc221/tasks/abc221_B) | 7002 | 0 |  [go](./go/abc221/abc221_B/abc221_B.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_C](http:/atcoder.jp/contests/abc221/tasks/abc221_C) | 4553 | 379 |  [go](./go/abc221/abc221_C/abc221_C.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_D](http:/atcoder.jp/contests/abc221/tasks/abc221_D) | 2803 | 832 |  [go](./go/abc221/abc221_D/abc221_D.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_E](http:/atcoder.jp/contests/abc221/tasks/abc221_E) | 991 | 1515 |  [go](./go/abc221/abc221_E/abc221_E.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_F](http:/atcoder.jp/contests/abc221/tasks/abc221_F) | 287 | 2093 |  [go](./go/abc221/abc221_F/abc221_F.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_G](http:/atcoder.jp/contests/abc221/tasks/abc221_G) | 26 | 2914 |  [go](./go/abc221/abc221_G/abc221_G.go) |  |
| [abc221](http:/atcoder.jp/contests/abc221) | [abc221_H](http:/atcoder.jp/contests/abc221/tasks/abc221_H) | 39 | 2793 |  [go](./go/abc221/abc221_H/abc221_H.go) |  |

## abc220 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_A](http:/atcoder.jp/contests/abc220/tasks/abc220_A) | 7185 | 0 |  [go](./go/abc220/abc220_A/abc220_A.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_B](http:/atcoder.jp/contests/abc220/tasks/abc220_B) | 6418 | 0 |  [go](./go/abc220/abc220_B/abc220_B.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_C](http:/atcoder.jp/contests/abc220/tasks/abc220_C) | 5778 | 119 |  [go](./go/abc220/abc220_C/abc220_C.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_D](http:/atcoder.jp/contests/abc220/tasks/abc220_D) | 3286 | 664 |  [go](./go/abc220/abc220_D/abc220_D.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_E](http:/atcoder.jp/contests/abc220/tasks/abc220_E) | 867 | 1593 |  [go](./go/abc220/abc220_E/abc220_E.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_F](http:/atcoder.jp/contests/abc220/tasks/abc220_F) | 1411 | 1304 |  [go](./go/abc220/abc220_F/abc220_F.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_G](http:/atcoder.jp/contests/abc220/tasks/abc220_G) | 116 | 2439 |  [go](./go/abc220/abc220_G/abc220_G.go) |  |
| [abc220](http:/atcoder.jp/contests/abc220) | [abc220_H](http:/atcoder.jp/contests/abc220/tasks/abc220_H) | 19 | 3047 |  [go](./go/abc220/abc220_H/abc220_H.go) |  |

## abc219 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_A](http:/atcoder.jp/contests/abc219/tasks/abc219_A) | 7203 | 0 |  [pypy3](./python/abc219/abc219_A.py) [go](./go/abc219/abc219_A/abc219_A.go) |  |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_B](http:/atcoder.jp/contests/abc219/tasks/abc219_B) | 6901 | 0 |  [pypy3](./python/abc219/abc219_B.py) [go](./go/abc219/abc219_B/abc219_B.go) |  |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_C](http:/atcoder.jp/contests/abc219/tasks/abc219_C) | 4665 | 260 |  [pypy3](./python/abc219/abc219_C.py) [go](./go/abc219/abc219_C/abc219_C.go) |  |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_D](http:/atcoder.jp/contests/abc219/tasks/abc219_D) | 1859 | 1085 |  [pypy3](./python/abc219/abc219_D.py) [go](./go/abc219/abc219_D/abc219_D.go) |  |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_E](http:/atcoder.jp/contests/abc219/tasks/abc219_E) | 681 | 1690 |  [pypy3](./python/abc219/abc219_E.py) [go](./go/abc219/abc219_E/abc219_E.go) | Polygon construction, DSU |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_F](http:/atcoder.jp/contests/abc219/tasks/abc219_F) | 93 | 2542 |  [pypy3](./python/abc219/abc219_F.py) [go](./go/abc219/abc219_F/abc219_F.go) | Stepping a shifted pattern, Equivalence classes |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_G](http:/atcoder.jp/contests/abc219/tasks/abc219_G) | 181 | 2287 |  [pypy3](./python/abc219/abc219_G.py) [go](./go/abc219/abc219_G/abc219_G.go) | Large tree, sqrt(N) strategy |
| [abc219](http:/atcoder.jp/contests/abc219) | [abc219_H](http:/atcoder.jp/contests/abc219/tasks/abc219_H) | 10 | 3297 |  [pypy3](./python/abc219/abc219_H.py) [go](./go/abc219/abc219_H/abc219_H.go) | Hard DP.  Pypy3 solution is TLE (same alg as Go solution).  Need to work on it a bit (probably need to flatten multi-dim dp array) |

## abc218 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_A](http:/atcoder.jp/contests/abc218/tasks/abc218_A) | 8653 | 0 |  [pypy3](./python/abc218/abc218_A.py) [go](./go/abc218/abc218_A/abc218_A.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_B](http:/atcoder.jp/contests/abc218/tasks/abc218_B) | 8436 | 0 |  [pypy3](./python/abc218/abc218_B.py) [go](./go/abc218/abc218_B/abc218_B.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_C](http:/atcoder.jp/contests/abc218/tasks/abc218_C) | 2385 | 1012 |  [pypy3](./python/abc218/abc218_C.py) [go](./go/abc218/abc218_C/abc218_C.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_D](http:/atcoder.jp/contests/abc218/tasks/abc218_D) | 3491 | 715 |  [pypy3](./python/abc218/abc218_D.py) [go](./go/abc218/abc218_D/abc218_D.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_E](http:/atcoder.jp/contests/abc218/tasks/abc218_E) | 2408 | 1004 |  [pypy3](./python/abc218/abc218_E.py) [go](./go/abc218/abc218_E/abc218_E.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_F](http:/atcoder.jp/contests/abc218/tasks/abc218_F) | 669 | 1753 |  [pypy3](./python/abc218/abc218_F.py) [go](./go/abc218/abc218_F/abc218_F.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_G](http:/atcoder.jp/contests/abc218/tasks/abc218_G) | 235 | 2217 |  [pypy3](./python/abc218/abc218_G.py) [go](./go/abc218/abc218_G/abc218_G.go) |  |
| [abc218](http:/atcoder.jp/contests/abc218) | [abc218_H](http:/atcoder.jp/contests/abc218/tasks/abc218_H) | 48 | 2805 |  [pypy3](./python/abc218/abc218_H.py) [go](./go/abc218/abc218_H/abc218_H.go) |  |

## abc217 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_A](http:/atcoder.jp/contests/abc217/tasks/abc217_A) | 8119 | 0 |  [pypy3](./python/abc217/abc217_A.py) [go](./go/abc217/abc217_A/abc217_A.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_B](http:/atcoder.jp/contests/abc217/tasks/abc217_B) | 8096 | 0 |  [pypy3](./python/abc217/abc217_B.py) [go](./go/abc217/abc217_B/abc217_B.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_C](http:/atcoder.jp/contests/abc217/tasks/abc217_C) | 7544 | 0 |  [pypy3](./python/abc217/abc217_C.py) [go](./go/abc217/abc217_C/abc217_C.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_D](http:/atcoder.jp/contests/abc217/tasks/abc217_D) | 3115 | 802 |  [pypy3](./python/abc217/abc217_D.py) [go](./go/abc217/abc217_D/abc217_D.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_E](http:/atcoder.jp/contests/abc217/tasks/abc217_E) | 2445 | 986 |  [pypy3](./python/abc217/abc217_E.py) [go](./go/abc217/abc217_E/abc217_E.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_F](http:/atcoder.jp/contests/abc217/tasks/abc217_F) | 444 | 1954 |  [pypy3](./python/abc217/abc217_F.py) [go](./go/abc217/abc217_F/abc217_F.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_G](http:/atcoder.jp/contests/abc217/tasks/abc217_G) | 376 | 2047 |  [pypy3](./python/abc217/abc217_G.py) [go](./go/abc217/abc217_G/abc217_G.go) |  |
| [abc217](http:/atcoder.jp/contests/abc217) | [abc217_H](http:/atcoder.jp/contests/abc217/tasks/abc217_H) | 59 | 3112 |  [pypy3](./python/abc217/abc217_H.py) [go](./go/abc217/abc217_H/abc217_H.go) |  |

## abc216 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_A](http:/atcoder.jp/contests/abc216/tasks/abc216_A) | 6997 | 0 |  [pypy3](./python/abc216/abc216_A.py) [go](./go/abc216/abc216_A/abc216_A.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_B](http:/atcoder.jp/contests/abc216/tasks/abc216_B) | 6554 | 0 |  [pypy3](./python/abc216/abc216_B.py) [go](./go/abc216/abc216_B/abc216_B.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_C](http:/atcoder.jp/contests/abc216/tasks/abc216_C) | 5356 | 0 |  [pypy3](./python/abc216/abc216_C.py) [go](./go/abc216/abc216_C/abc216_C.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_D](http:/atcoder.jp/contests/abc216/tasks/abc216_D) | 1969 | 1039 |  [pypy3](./python/abc216/abc216_D.py) [go](./go/abc216/abc216_D/abc216_D.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_E](http:/atcoder.jp/contests/abc216/tasks/abc216_E) | 1847 | 1084 |  [pypy3](./python/abc216/abc216_E.py) [go](./go/abc216/abc216_E/abc216_E.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_F](http:/atcoder.jp/contests/abc216/tasks/abc216_F) | 898 | 1541 |  [pypy3](./python/abc216/abc216_F.py) [go](./go/abc216/abc216_F/abc216_F.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_G](http:/atcoder.jp/contests/abc216/tasks/abc216_G) | 393 | 1963 |  [pypy3](./python/abc216/abc216_G.py) [go](./go/abc216/abc216_G/abc216_G.go) |  |
| [abc216](http:/atcoder.jp/contests/abc216) | [abc216_H](http:/atcoder.jp/contests/abc216/tasks/abc216_H) | 9 | 3295 |  [pypy3](./python/abc216/abc216_H.py) [go](./go/abc216/abc216_H/abc216_H.go) |  |

## abc215 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_A](http:/atcoder.jp/contests/abc215/tasks/abc215_A) | 8363 | 0 |  [pypy3](./python/abc215/abc215_A.py) [go](./go/abc215/abc215_A/abc215_A.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_B](http:/atcoder.jp/contests/abc215/tasks/abc215_B) | 7060 | 0 |  [pypy3](./python/abc215/abc215_B.py) [go](./go/abc215/abc215_B/abc215_B.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_C](http:/atcoder.jp/contests/abc215/tasks/abc215_C) | 5851 | 76 |  [pypy3](./python/abc215/abc215_C.py) [go](./go/abc215/abc215_C/abc215_C.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_D](http:/atcoder.jp/contests/abc215/tasks/abc215_D) | 3216 | 736 |  [pypy3](./python/abc215/abc215_D.py) [go](./go/abc215/abc215_D/abc215_D.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_E](http:/atcoder.jp/contests/abc215/tasks/abc215_E) | 1192 | 1413 |  [pypy3](./python/abc215/abc215_E.py) [go](./go/abc215/abc215_E/abc215_E.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_F](http:/atcoder.jp/contests/abc215/tasks/abc215_F) | 518 | 1853 |  [pypy3](./python/abc215/abc215_F.py) [go](./go/abc215/abc215_F/abc215_F.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_G](http:/atcoder.jp/contests/abc215/tasks/abc215_G) | 186 | 2276 |  [pypy3](./python/abc215/abc215_G.py) [go](./go/abc215/abc215_G/abc215_G.go) |  |
| [abc215](http:/atcoder.jp/contests/abc215) | [abc215_H](http:/atcoder.jp/contests/abc215/tasks/abc215_H) | 17 | 3101 |  [pypy3](./python/abc215/abc215_H.py) [go](./go/abc215/abc215_H/abc215_H.go) |  |

## abc214 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_A](http:/atcoder.jp/contests/abc214/tasks/abc214_A) | 8306 | 0 |  [pypy3](./python/abc214/abc214_A.py) [go](./go/abc214/abc214_A/abc214_A.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_B](http:/atcoder.jp/contests/abc214/tasks/abc214_B) | 7169 | 0 |  [pypy3](./python/abc214/abc214_B.py) [go](./go/abc214/abc214_B/abc214_B.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_C](http:/atcoder.jp/contests/abc214/tasks/abc214_C) | 4828 | 309 |  [pypy3](./python/abc214/abc214_C.py) [go](./go/abc214/abc214_C/abc214_C.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_D](http:/atcoder.jp/contests/abc214/tasks/abc214_D) | 1346 | 1341 |  [pypy3](./python/abc214/abc214_D.py) [go](./go/abc214/abc214_D/abc214_D.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_E](http:/atcoder.jp/contests/abc214/tasks/abc214_E) | 551 | 1835 |  [pypy3](./python/abc214/abc214_E.py) [go](./go/abc214/abc214_E/abc214_E.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_F](http:/atcoder.jp/contests/abc214/tasks/abc214_F) | 406 | 1973 |  [pypy3](./python/abc214/abc214_F.py) [go](./go/abc214/abc214_F/abc214_F.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_G](http:/atcoder.jp/contests/abc214/tasks/abc214_G) | 32 | 2893 |  [pypy3](./python/abc214/abc214_G.py) [go](./go/abc214/abc214_G/abc214_G.go) |  |
| [abc214](http:/atcoder.jp/contests/abc214) | [abc214_H](http:/atcoder.jp/contests/abc214/tasks/abc214_H) | 14 | 3138 |  [pypy3](./python/abc214/abc214_H.py) [go](./go/abc214/abc214_H/abc214_H.go) |  |

## abc213 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_A](http:/atcoder.jp/contests/abc213/tasks/abc213_A) | 7599 | 0 |  [pypy3](./python/abc213/abc213_A.py) [go](./go/abc213/abc213_A/abc213_A.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_B](http:/atcoder.jp/contests/abc213/tasks/abc213_B) | 7667 | 0 |  [pypy3](./python/abc213/abc213_B.py) [go](./go/abc213/abc213_B/abc213_B.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_C](http:/atcoder.jp/contests/abc213/tasks/abc213_C) | 4327 | 481 |  [pypy3](./python/abc213/abc213_C.py) [go](./go/abc213/abc213_C/abc213_C.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_D](http:/atcoder.jp/contests/abc213/tasks/abc213_D) | 3431 | 710 |  [pypy3](./python/abc213/abc213_D.py) [go](./go/abc213/abc213_D/abc213_D.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_E](http:/atcoder.jp/contests/abc213/tasks/abc213_E) | 1193 | 1423 |  [pypy3](./python/abc213/abc213_E.py) [go](./go/abc213/abc213_E/abc213_E.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_F](http:/atcoder.jp/contests/abc213/tasks/abc213_F) | 201 | 2215 |  [pypy3](./python/abc213/abc213_F.py) [go](./go/abc213/abc213_F/abc213_F.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_G](http:/atcoder.jp/contests/abc213/tasks/abc213_G) | 49 | 2663 |  [pypy3](./python/abc213/abc213_G.py) [go](./go/abc213/abc213_G/abc213_G.go) |  |
| [abc213](http:/atcoder.jp/contests/abc213) | [abc213_H](http:/atcoder.jp/contests/abc213/tasks/abc213_H) | 22 | 2806 |  [pypy3](./python/abc213/abc213_H.py) [go](./go/abc213/abc213_H/abc213_H.go) |  |

## abc212 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_A](http:/atcoder.jp/contests/abc212/tasks/abc212_A) | 7785 | 0 |  [pypy3](./python/abc212/abc212_A.py) [go](./go/abc212/abc212_A/abc212_A.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_B](http:/atcoder.jp/contests/abc212/tasks/abc212_B) | 6536 | 0 |  [pypy3](./python/abc212/abc212_B.py) [go](./go/abc212/abc212_B/abc212_B.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_C](http:/atcoder.jp/contests/abc212/tasks/abc212_C) | 4921 | 205 |  [pypy3](./python/abc212/abc212_C.py) [go](./go/abc212/abc212_C/abc212_C.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_D](http:/atcoder.jp/contests/abc212/tasks/abc212_D) | 2839 | 775 |  [pypy3](./python/abc212/abc212_D.py) [go](./go/abc212/abc212_D/abc212_D.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_E](http:/atcoder.jp/contests/abc212/tasks/abc212_E) | 1122 | 1410 |  [pypy3](./python/abc212/abc212_E.py) [go](./go/abc212/abc212_E/abc212_E.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_F](http:/atcoder.jp/contests/abc212/tasks/abc212_F) | 152 | 2332 |  [pypy3](./python/abc212/abc212_F.py) [go](./go/abc212/abc212_F/abc212_F.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_G](http:/atcoder.jp/contests/abc212/tasks/abc212_G) | 248 | 2150 |  [pypy3](./python/abc212/abc212_G.py) [go](./go/abc212/abc212_G/abc212_G.go) |  |
| [abc212](http:/atcoder.jp/contests/abc212) | [abc212_H](http:/atcoder.jp/contests/abc212/tasks/abc212_H) | 46 | 2741 |  [pypy3](./python/abc212/abc212_H.py) [go](./go/abc212/abc212_H/abc212_H.go) |  |

## abc211 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_A](http:/atcoder.jp/contests/abc211/tasks/abc211_A) | 8468 | 0 |  [pypy3](./python/abc211/abc211_A.py) [go](./go/abc211/abc211_A/abc211_A.go) |  |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_B](http:/atcoder.jp/contests/abc211/tasks/abc211_B) | 8114 | 0 |  [pypy3](./python/abc211/abc211_B.py) [go](./go/abc211/abc211_B/abc211_B.go) |  |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_C](http:/atcoder.jp/contests/abc211/tasks/abc211_C) | 3933 | 559 |  [pypy3](./python/abc211/abc211_C.py) [go](./go/abc211/abc211_C/abc211_C.go) |  |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_D](http:/atcoder.jp/contests/abc211/tasks/abc211_D) | 3158 | 755 |  [pypy3](./python/abc211/abc211_D.py) [go](./go/abc211/abc211_D/abc211_D.go) |  |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_E](http:/atcoder.jp/contests/abc211/tasks/abc211_E) | 530 | 1823 |  [pypy3](./python/abc211/abc211_E.py) [go](./go/abc211/abc211_E/abc211_E.go) |  |
| [abc211](http:/atcoder.jp/contests/abc211) | [abc211_F](http:/atcoder.jp/contests/abc211/tasks/abc211_F) | 144 | 2350 |  [pypy3](./python/abc211/abc211_F.py) [go](./go/abc211/abc211_F/abc211_F.go) |  |

## abc210 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_A](http:/atcoder.jp/contests/abc210/tasks/abc210_A) | 8339 | 0 |  [pypy3](./python/abc210/abc210_A.py) [go](./go/abc210/abc210_A/abc210_A.go) |  |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_B](http:/atcoder.jp/contests/abc210/tasks/abc210_B) | 8256 | 0 |  [pypy3](./python/abc210/abc210_B.py) [go](./go/abc210/abc210_B/abc210_B.go) |  |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_C](http:/atcoder.jp/contests/abc210/tasks/abc210_C) | 4869 | 357 |  [pypy3](./python/abc210/abc210_C.py) [go](./go/abc210/abc210_C/abc210_C.go) |  |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_D](http:/atcoder.jp/contests/abc210/tasks/abc210_D) | 1011 | 1507 |  [pypy3](./python/abc210/abc210_D.py) [go](./go/abc210/abc210_D/abc210_D.go) |  |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_E](http:/atcoder.jp/contests/abc210/tasks/abc210_E) | 811 | 1618 |  [pypy3](./python/abc210/abc210_E.py) [go](./go/abc210/abc210_E/abc210_E.go) |  |
| [abc210](http:/atcoder.jp/contests/abc210) | [abc210_F](http:/atcoder.jp/contests/abc210/tasks/abc210_F) | 60 | 2632 |  [pypy3](./python/abc210/abc210_F.py) [go](./go/abc210/abc210_F/abc210_F.go) |  |

## abc209 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_A](http:/atcoder.jp/contests/abc209/tasks/abc209_A) | 8642 | 0 |  [pypy3](./python/abc209/abc209_A.py) [go](./go/abc209/abc209_A/abc209_A.go) |  |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_B](http:/atcoder.jp/contests/abc209/tasks/abc209_B) | 8395 | 0 |  [pypy3](./python/abc209/abc209_B.py) [go](./go/abc209/abc209_B/abc209_B.go) |  |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_C](http:/atcoder.jp/contests/abc209/tasks/abc209_C) | 5267 | 264 |  [pypy3](./python/abc209/abc209_C.py) [go](./go/abc209/abc209_C/abc209_C.go) |  |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_D](http:/atcoder.jp/contests/abc209/tasks/abc209_D) | 3517 | 686 |  [pypy3](./python/abc209/abc209_D.py) [go](./go/abc209/abc209_D/abc209_D.go) |  |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_E](http:/atcoder.jp/contests/abc209/tasks/abc209_E) | 247 | 2153 |  [pypy3](./python/abc209/abc209_E.py) [go](./go/abc209/abc209_E/abc209_E.go) |  |
| [abc209](http:/atcoder.jp/contests/abc209) | [abc209_F](http:/atcoder.jp/contests/abc209/tasks/abc209_F) | 163 | 2307 |  [pypy3](./python/abc209/abc209_F.py) [go](./go/abc209/abc209_F/abc209_F.go) |  |

## abc208 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_A](http:/atcoder.jp/contests/abc208/tasks/abc208_A) | 7805 | 0 |  [pypy3](./python/abc208/abc208_A.py) [go](./go/abc208/abc208_A/abc208_A.go) |  |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_B](http:/atcoder.jp/contests/abc208/tasks/abc208_B) | 7352 | 0 |  [pypy3](./python/abc208/abc208_B.py) [go](./go/abc208/abc208_B/abc208_B.go) |  |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_C](http:/atcoder.jp/contests/abc208/tasks/abc208_C) | 6165 | 0 |  [pypy3](./python/abc208/abc208_C.py) [go](./go/abc208/abc208_C/abc208_C.go) |  |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_D](http:/atcoder.jp/contests/abc208/tasks/abc208_D) | 1711 | 1190 |  [pypy3](./python/abc208/abc208_D.py) [go](./go/abc208/abc208_D/abc208_D.go) |  |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_E](http:/atcoder.jp/contests/abc208/tasks/abc208_E) | 341 | 2024 |  [pypy3](./python/abc208/abc208_E.py) [go](./go/abc208/abc208_E/abc208_E.go) |  |
| [abc208](http:/atcoder.jp/contests/abc208) | [abc208_F](http:/atcoder.jp/contests/abc208/tasks/abc208_F) | 43 | 2772 |  [pypy3](./python/abc208/abc208_F.py) [go](./go/abc208/abc208_F/abc208_F.go) |  |

## abc207 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_A](http:/atcoder.jp/contests/abc207/tasks/abc207_A) | 8260 | 0 |  [pypy3](./python/abc207/abc207_A.py) [go](./go/abc207/abc207_A/abc207_A.go) |  |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_B](http:/atcoder.jp/contests/abc207/tasks/abc207_B) | 6577 | 0 |  [pypy3](./python/abc207/abc207_B.py) [go](./go/abc207/abc207_B/abc207_B.go) |  |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_C](http:/atcoder.jp/contests/abc207/tasks/abc207_C) | 4502 | 397 |  [pypy3](./python/abc207/abc207_C.py) [go](./go/abc207/abc207_C/abc207_C.go) |  |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_D](http:/atcoder.jp/contests/abc207/tasks/abc207_D) | 285 | 2074 |  [pypy3](./python/abc207/abc207_D.py) [go](./go/abc207/abc207_D/abc207_D.go) |  |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_E](http:/atcoder.jp/contests/abc207/tasks/abc207_E) | 521 | 1820 |  [pypy3](./python/abc207/abc207_E.py) [go](./go/abc207/abc207_E/abc207_E.go) |  |
| [abc207](http:/atcoder.jp/contests/abc207) | [abc207_F](http:/atcoder.jp/contests/abc207/tasks/abc207_F) | 126 | 2398 |  [pypy3](./python/abc207/abc207_F.py) [go](./go/abc207/abc207_F/abc207_F.go) |  |

## abc206 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_A](http:/atcoder.jp/contests/abc206/tasks/abc206_A) | 8961 | 0 |  [pypy3](./python/abc206/abc206_A.py) [go](./go/abc206/abc206_A/abc206_A.go) |  |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_B](http:/atcoder.jp/contests/abc206/tasks/abc206_B) | 8693 | 0 |  [pypy3](./python/abc206/abc206_B.py) [go](./go/abc206/abc206_B/abc206_B.go) |  |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_C](http:/atcoder.jp/contests/abc206/tasks/abc206_C) | 6267 | 60 |  [pypy3](./python/abc206/abc206_C.py) [go](./go/abc206/abc206_C/abc206_C.go) |  |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_D](http:/atcoder.jp/contests/abc206/tasks/abc206_D) | 2788 | 879 |  [pypy3](./python/abc206/abc206_D.py) [go](./go/abc206/abc206_D/abc206_D.go) |  |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_E](http:/atcoder.jp/contests/abc206/tasks/abc206_E) | 636 | 1745 |  [pypy3](./python/abc206/abc206_E.py) [go](./go/abc206/abc206_E/abc206_E.go) |  |
| [abc206](http:/atcoder.jp/contests/abc206) | [abc206_F](http:/atcoder.jp/contests/abc206/tasks/abc206_F) | 209 | 2221 |  [pypy3](./python/abc206/abc206_F.py) [go](./go/abc206/abc206_F/abc206_F.go) |  |

## abc205 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_A](http:/atcoder.jp/contests/abc205/tasks/abc205_A) | 8692 | 0 |  [pypy3](./python/abc205/abc205_A.py) [go](./go/abc205/abc205_A/abc205_A.go) |  |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_B](http:/atcoder.jp/contests/abc205/tasks/abc205_B) | 8356 | 0 |  [pypy3](./python/abc205/abc205_B.py) [go](./go/abc205/abc205_B/abc205_B.go) |  |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_C](http:/atcoder.jp/contests/abc205/tasks/abc205_C) | 7366 | 0 |  [pypy3](./python/abc205/abc205_C.py) [go](./go/abc205/abc205_C/abc205_C.go) |  |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_D](http:/atcoder.jp/contests/abc205/tasks/abc205_D) | 3336 | 713 |  [pypy3](./python/abc205/abc205_D.py) [go](./go/abc205/abc205_D/abc205_D.go) |  |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_E](http:/atcoder.jp/contests/abc205/tasks/abc205_E) | 327 | 2025 |  [pypy3](./python/abc205/abc205_E.py) [go](./go/abc205/abc205_E/abc205_E.go) |  |
| [abc205](http:/atcoder.jp/contests/abc205) | [abc205_F](http:/atcoder.jp/contests/abc205/tasks/abc205_F) | 282 | 2088 |  [pypy3](./python/abc205/abc205_F.py) [go](./go/abc205/abc205_F/abc205_F.go) |  |

## abc204 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_A](http:/atcoder.jp/contests/abc204/tasks/abc204_A) | 8710 | 0 |  [pypy3](./python/abc204/abc204_A.py) [go](./go/abc204/abc204_A/abc204_A.go) |  |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_B](http:/atcoder.jp/contests/abc204/tasks/abc204_B) | 8644 | 0 |  [pypy3](./python/abc204/abc204_B.py) [go](./go/abc204/abc204_B/abc204_B.go) |  |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_C](http:/atcoder.jp/contests/abc204/tasks/abc204_C) | 3776 | 629 |  [pypy3](./python/abc204/abc204_C.py) [go](./go/abc204/abc204_C/abc204_C.go) |  |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_D](http:/atcoder.jp/contests/abc204/tasks/abc204_D) | 2970 | 832 |  [pypy3](./python/abc204/abc204_D.py) [go](./go/abc204/abc204_D/abc204_D.go) |  |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_E](http:/atcoder.jp/contests/abc204/tasks/abc204_E) | 687 | 1710 |  [pypy3](./python/abc204/abc204_E.py) [go](./go/abc204/abc204_E/abc204_E.go) |  |
| [abc204](http:/atcoder.jp/contests/abc204) | [abc204_F](http:/atcoder.jp/contests/abc204/tasks/abc204_F) | 329 | 2044 |  [pypy3](./python/abc204/abc204_F.py) [go](./go/abc204/abc204_F/abc204_F.go) |  |

## abc203 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_A](http:/atcoder.jp/contests/abc203/tasks/abc203_A) | 8295 | 0 |  [pypy3](./python/abc203/abc203_A.py) [go](./go/abc203/abc203_A/abc203_A.go) |  |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_B](http:/atcoder.jp/contests/abc203/tasks/abc203_B) | 8124 | 0 |  [pypy3](./python/abc203/abc203_B.py) [go](./go/abc203/abc203_B/abc203_B.go) |  |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_C](http:/atcoder.jp/contests/abc203/tasks/abc203_C) | 5878 | 54 |  [pypy3](./python/abc203/abc203_C.py) [go](./go/abc203/abc203_C/abc203_C.go) |  |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_D](http:/atcoder.jp/contests/abc203/tasks/abc203_D) | 775 | 1622 |  [pypy3](./python/abc203/abc203_D.py) [go](./go/abc203/abc203_D/abc203_D.go) |  |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_E](http:/atcoder.jp/contests/abc203/tasks/abc203_E) | 596 | 1750 |  [pypy3](./python/abc203/abc203_E.py) [go](./go/abc203/abc203_E/abc203_E.go) |  |
| [abc203](http:/atcoder.jp/contests/abc203) | [abc203_F](http:/atcoder.jp/contests/abc203/tasks/abc203_F) | 184 | 2252 |  [pypy3](./python/abc203/abc203_F.py) [go](./go/abc203/abc203_F/abc203_F.go) |  |

## abc202 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_A](http:/atcoder.jp/contests/abc202/tasks/abc202_A) | 8628 | 0 |  [pypy3](./python/abc202/abc202_A.py) [go](./go/abc202/abc202_A/abc202_A.go) |  |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_B](http:/atcoder.jp/contests/abc202/tasks/abc202_B) | 8317 | 0 |  [pypy3](./python/abc202/abc202_B.py) [go](./go/abc202/abc202_B/abc202_B.go) |  |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_C](http:/atcoder.jp/contests/abc202/tasks/abc202_C) | 5871 | 130 |  [pypy3](./python/abc202/abc202_C.py) [go](./go/abc202/abc202_C/abc202_C.go) |  |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_D](http:/atcoder.jp/contests/abc202/tasks/abc202_D) | 2442 | 966 |  [pypy3](./python/abc202/abc202_D.py) [go](./go/abc202/abc202_D/abc202_D.go) |  |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_E](http:/atcoder.jp/contests/abc202/tasks/abc202_E) | 772 | 1638 |  [pypy3](./python/abc202/abc202_E.py) [go](./go/abc202/abc202_E/abc202_E.go) |  |
| [abc202](http:/atcoder.jp/contests/abc202) | [abc202_F](http:/atcoder.jp/contests/abc202/tasks/abc202_F) | 25 | 2905 |  [pypy3](./python/abc202/abc202_F.py) [go](./go/abc202/abc202_F/abc202_F.go) |  |

## abc201 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_A](http:/atcoder.jp/contests/abc201/tasks/abc201_A) | 8453 | 0 |  [pypy3](./python/abc201/abc201_A.py) [go](./go/abc201/abc201_A/abc201_A.go) |  |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_B](http:/atcoder.jp/contests/abc201/tasks/abc201_B) | 7994 | 0 |  [pypy3](./python/abc201/abc201_B.py) [go](./go/abc201/abc201_B/abc201_B.go) |  |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_C](http:/atcoder.jp/contests/abc201/tasks/abc201_C) | 4570 | 439 |  [pypy3](./python/abc201/abc201_C.py) [go](./go/abc201/abc201_C/abc201_C.go) |  |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_D](http:/atcoder.jp/contests/abc201/tasks/abc201_D) | 1382 | 1317 |  [pypy3](./python/abc201/abc201_D.py) [go](./go/abc201/abc201_D/abc201_D.go) |  |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_E](http:/atcoder.jp/contests/abc201/tasks/abc201_E) | 701 | 1694 |  [pypy3](./python/abc201/abc201_E.py) [go](./go/abc201/abc201_E/abc201_E.go) |  |
| [abc201](http:/atcoder.jp/contests/abc201) | [abc201_F](http:/atcoder.jp/contests/abc201/tasks/abc201_F) | 110 | 2484 |  [pypy3](./python/abc201/abc201_F.py) [go](./go/abc201/abc201_F/abc201_F.go) |  |

## abc200 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_A](http:/atcoder.jp/contests/abc200/tasks/abc200_A) | 8475 | 0 |  [pypy3](./python/abc200/abc200_A.py) [go](./go/abc200/abc200_A/abc200_A.go) |  |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_B](http:/atcoder.jp/contests/abc200/tasks/abc200_B) | 8136 | 0 |  [pypy3](./python/abc200/abc200_B.py) [go](./go/abc200/abc200_B/abc200_B.go) |  |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_C](http:/atcoder.jp/contests/abc200/tasks/abc200_C) | 5846 | 138 |  [pypy3](./python/abc200/abc200_C.py) [go](./go/abc200/abc200_C/abc200_C.go) |  |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_D](http:/atcoder.jp/contests/abc200/tasks/abc200_D) | 1696 | 1217 |  [pypy3](./python/abc200/abc200_D.py) [go](./go/abc200/abc200_D/abc200_D.go) |  |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_E](http:/atcoder.jp/contests/abc200/tasks/abc200_E) | 432 | 1955 |  [pypy3](./python/abc200/abc200_E.py) [go](./go/abc200/abc200_E/abc200_E.go) |  |
| [abc200](http:/atcoder.jp/contests/abc200) | [abc200_F](http:/atcoder.jp/contests/abc200/tasks/abc200_F) | 91 | 2556 |  [pypy3](./python/abc200/abc200_F.py) [go](./go/abc200/abc200_F/abc200_F.go) |  |

## abc199 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_A](http:/atcoder.jp/contests/abc199/tasks/abc199_A) | 8600 | 0 |  [go](./go/abc199/abc199_A/abc199_A.go) |  |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_B](http:/atcoder.jp/contests/abc199/tasks/abc199_B) | 7745 | 0 |  [go](./go/abc199/abc199_B/abc199_B.go) |  |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_C](http:/atcoder.jp/contests/abc199/tasks/abc199_C) | 4524 | 436 |  [go](./go/abc199/abc199_C/abc199_C.go) |  |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_D](http:/atcoder.jp/contests/abc199/tasks/abc199_D) | 571 | 1804 |  [go](./go/abc199/abc199_D/abc199_D.go) |  |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_E](http:/atcoder.jp/contests/abc199/tasks/abc199_E) | 550 | 1814 |  [go](./go/abc199/abc199_E/abc199_E.go) |  |
| [abc199](http:/atcoder.jp/contests/abc199) | [abc199_F](http:/atcoder.jp/contests/abc199/tasks/abc199_F) | 310 | 2065 |  [go](./go/abc199/abc199_F/abc199_F.go) |  |

## abc198 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_A](http:/atcoder.jp/contests/abc198/tasks/abc198_A) | 7933 | 0 |  [go](./go/abc198/abc198_A/abc198_A.go) |  |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_B](http:/atcoder.jp/contests/abc198/tasks/abc198_B) | 6864 | 0 |  [go](./go/abc198/abc198_B/abc198_B.go) |  |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_C](http:/atcoder.jp/contests/abc198/tasks/abc198_C) | 4369 | 413 |  [go](./go/abc198/abc198_C/abc198_C.go) |  |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_D](http:/atcoder.jp/contests/abc198/tasks/abc198_D) | 1507 | 1224 |  [go](./go/abc198/abc198_D/abc198_D.go) |  |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_E](http:/atcoder.jp/contests/abc198/tasks/abc198_E) | 1667 | 1161 |  [go](./go/abc198/abc198_E/abc198_E.go) |  |
| [abc198](http:/atcoder.jp/contests/abc198) | [abc198_F](http:/atcoder.jp/contests/abc198/tasks/abc198_F) | 37 | 2769 |  [go](./go/abc198/abc198_F/abc198_F.go) |  |

## abc197 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_A](http:/atcoder.jp/contests/abc197/tasks/abc197_A) | 7757 | 0 |  [go](./go/abc197/abc197_A/abc197_A.go) |  |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_B](http:/atcoder.jp/contests/abc197/tasks/abc197_B) | 6164 | 0 |  [go](./go/abc197/abc197_B/abc197_B.go) |  |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_C](http:/atcoder.jp/contests/abc197/tasks/abc197_C) | 2659 | 809 |  [go](./go/abc197/abc197_C/abc197_C.go) |  |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_D](http:/atcoder.jp/contests/abc197/tasks/abc197_D) | 2615 | 831 |  [go](./go/abc197/abc197_D/abc197_D.go) |  |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_E](http:/atcoder.jp/contests/abc197/tasks/abc197_E) | 1140 | 1379 |  [go](./go/abc197/abc197_E/abc197_E.go) |  |
| [abc197](http:/atcoder.jp/contests/abc197) | [abc197_F](http:/atcoder.jp/contests/abc197/tasks/abc197_F) | 376 | 1945 |  [go](./go/abc197/abc197_F/abc197_F.go) |  |

## abc196 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_A](http:/atcoder.jp/contests/abc196/tasks/abc196_A) | 8435 | 0 |  [go](./go/abc196/abc196_A/abc196_A.go) |  |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_B](http:/atcoder.jp/contests/abc196/tasks/abc196_B) | 7725 | 0 |  [go](./go/abc196/abc196_B/abc196_B.go) |  |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_C](http:/atcoder.jp/contests/abc196/tasks/abc196_C) | 5503 | 202 |  [go](./go/abc196/abc196_C/abc196_C.go) |  |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_D](http:/atcoder.jp/contests/abc196/tasks/abc196_D) | 1517 | 1277 |  [go](./go/abc196/abc196_D/abc196_D.go) |  |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_E](http:/atcoder.jp/contests/abc196/tasks/abc196_E) | 765 | 1650 |  [go](./go/abc196/abc196_E/abc196_E.go) |  |
| [abc196](http:/atcoder.jp/contests/abc196) | [abc196_F](http:/atcoder.jp/contests/abc196/tasks/abc196_F) | 173 | 2274 |  [go](./go/abc196/abc196_F/abc196_F.go) |  |

## abc195 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_A](http:/atcoder.jp/contests/abc195/tasks/abc195_A) | 7588 | 0 |  [go](./go/abc195/abc195_A/abc195_A.go) |  |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_B](http:/atcoder.jp/contests/abc195/tasks/abc195_B) | 4054 | 483 |  [go](./go/abc195/abc195_B/abc195_B.go) |  |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_C](http:/atcoder.jp/contests/abc195/tasks/abc195_C) | 4989 | 235 |  [go](./go/abc195/abc195_C/abc195_C.go) |  |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_D](http:/atcoder.jp/contests/abc195/tasks/abc195_D) | 2369 | 945 |  [go](./go/abc195/abc195_D/abc195_D.go) |  |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_E](http:/atcoder.jp/contests/abc195/tasks/abc195_E) | 811 | 1609 |  [go](./go/abc195/abc195_E/abc195_E.go) |  |
| [abc195](http:/atcoder.jp/contests/abc195) | [abc195_F](http:/atcoder.jp/contests/abc195/tasks/abc195_F) | 304 | 2068 |  [go](./go/abc195/abc195_F/abc195_F.go) |  |

## abc194 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_A](http:/atcoder.jp/contests/abc194/tasks/abc194_A) | 7636 | 0 |  [go](./go/abc194/abc194_A/abc194_A.go) |  |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_B](http:/atcoder.jp/contests/abc194/tasks/abc194_B) | 6319 | 0 |  [go](./go/abc194/abc194_B/abc194_B.go) |  |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_C](http:/atcoder.jp/contests/abc194/tasks/abc194_C) | 4535 | 386 |  [go](./go/abc194/abc194_C/abc194_C.go) |  |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_D](http:/atcoder.jp/contests/abc194/tasks/abc194_D) | 2027 | 1078 |  [go](./go/abc194/abc194_D/abc194_D.go) |  |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_E](http:/atcoder.jp/contests/abc194/tasks/abc194_E) | 1992 | 1088 |  [go](./go/abc194/abc194_E/abc194_E.go) |  |
| [abc194](http:/atcoder.jp/contests/abc194) | [abc194_F](http:/atcoder.jp/contests/abc194/tasks/abc194_F) | 215 | 2197 |  [go](./go/abc194/abc194_F/abc194_F.go) |  |

## abc193 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_A](http:/atcoder.jp/contests/abc193/tasks/abc193_A) | 7641 | 0 |  [go](./go/abc193/abc193_A/abc193_A.go) |  |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_B](http:/atcoder.jp/contests/abc193/tasks/abc193_B) | 7114 | 0 |  [go](./go/abc193/abc193_B/abc193_B.go) |  |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_C](http:/atcoder.jp/contests/abc193/tasks/abc193_C) | 4602 | 378 |  [go](./go/abc193/abc193_C/abc193_C.go) |  |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_D](http:/atcoder.jp/contests/abc193/tasks/abc193_D) | 2771 | 866 |  [go](./go/abc193/abc193_D/abc193_D.go) |  |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_E](http:/atcoder.jp/contests/abc193/tasks/abc193_E) | 423 | 1948 |  [go](./go/abc193/abc193_E/abc193_E.go) |  |
| [abc193](http:/atcoder.jp/contests/abc193) | [abc193_F](http:/atcoder.jp/contests/abc193/tasks/abc193_F) | 110 | 2475 |  [go](./go/abc193/abc193_F/abc193_F.go) |  |

## abc192 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_A](http:/atcoder.jp/contests/abc192/tasks/abc192_A) | 8571 | 0 |  [go](./go/abc192/abc192_A/abc192_A.go) |  |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_B](http:/atcoder.jp/contests/abc192/tasks/abc192_B) | 8193 | 0 |  [go](./go/abc192/abc192_B/abc192_B.go) |  |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_C](http:/atcoder.jp/contests/abc192/tasks/abc192_C) | 6985 | 0 |  [go](./go/abc192/abc192_C/abc192_C.go) |  |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_D](http:/atcoder.jp/contests/abc192/tasks/abc192_D) | 1331 | 1425 |  [go](./go/abc192/abc192_D/abc192_D.go) |  |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_E](http:/atcoder.jp/contests/abc192/tasks/abc192_E) | 2113 | 1135 |  [go](./go/abc192/abc192_E/abc192_E.go) |  |
| [abc192](http:/atcoder.jp/contests/abc192) | [abc192_F](http:/atcoder.jp/contests/abc192/tasks/abc192_F) | 680 | 1783 |  [go](./go/abc192/abc192_F/abc192_F.go) |  |

## abc191 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_A](http:/atcoder.jp/contests/abc191/tasks/abc191_A) | 8128 | 0 |  [go](./go/abc191/abc191_A/abc191_A.go) |  |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_B](http:/atcoder.jp/contests/abc191/tasks/abc191_B) | 8215 | 0 |  [go](./go/abc191/abc191_B/abc191_B.go) |  |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_C](http:/atcoder.jp/contests/abc191/tasks/abc191_C) | 2257 | 1063 |  [go](./go/abc191/abc191_C/abc191_C.go) |  |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_D](http:/atcoder.jp/contests/abc191/tasks/abc191_D) | 424 | 1953 |  [go](./go/abc191/abc191_D/abc191_D.go) |  |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_E](http:/atcoder.jp/contests/abc191/tasks/abc191_E) | 1492 | 1323 |  [go](./go/abc191/abc191_E/abc191_E.go) |  |
| [abc191](http:/atcoder.jp/contests/abc191) | [abc191_F](http:/atcoder.jp/contests/abc191/tasks/abc191_F) | 157 | 2333 |  [go](./go/abc191/abc191_F/abc191_F.go) |  |

## abc190 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_A](http:/atcoder.jp/contests/abc190/tasks/abc190_A) | 8916 | 0 |  [go](./go/abc190/abc190_A/abc190_A.go) |  |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_B](http:/atcoder.jp/contests/abc190/tasks/abc190_B) | 8636 | 0 |  [go](./go/abc190/abc190_B/abc190_B.go) |  |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_C](http:/atcoder.jp/contests/abc190/tasks/abc190_C) | 4698 | 472 |  [go](./go/abc190/abc190_C/abc190_C.go) |  |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_D](http:/atcoder.jp/contests/abc190/tasks/abc190_D) | 3619 | 722 |  [go](./go/abc190/abc190_D/abc190_D.go) |  |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_E](http:/atcoder.jp/contests/abc190/tasks/abc190_E) | 821 | 1645 |  [go](./go/abc190/abc190_E/abc190_E.go) | Dijkstra, then Permutaion/Subset DP |
| [abc190](http:/atcoder.jp/contests/abc190) | [abc190_F](http:/atcoder.jp/contests/abc190/tasks/abc190_F) | 1510 | 1321 |  [go](./go/abc190/abc190_F/abc190_F.go) | Inversion counting.  Fenwick |

## abc189 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_A](http:/atcoder.jp/contests/abc189/tasks/abc189_A) | 8747 | 0 |  [go](./go/abc189/abc189_A/abc189_A.go) |  |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_B](http:/atcoder.jp/contests/abc189/tasks/abc189_B) | 5436 | 249 |  [go](./go/abc189/abc189_B/abc189_B.go) |  |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_C](http:/atcoder.jp/contests/abc189/tasks/abc189_C) | 4158 | 565 |  [go](./go/abc189/abc189_C/abc189_C.go) | Max area under histogram |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_D](http:/atcoder.jp/contests/abc189/tasks/abc189_D) | 3270 | 769 |  [go](./go/abc189/abc189_D/abc189_D.go) |  |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_E](http:/atcoder.jp/contests/abc189/tasks/abc189_E) | 1008 | 1526 |  [go](./go/abc189/abc189_E/abc189_E.go) | Bookeeping with reflections, translations, and rotations |
| [abc189](http:/atcoder.jp/contests/abc189) | [abc189_F](http:/atcoder.jp/contests/abc189/tasks/abc189_F) | 244 | 2154 |  [go](./go/abc189/abc189_F/abc189_F.go) | Expected value w/ simple algebra |

## abc188 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_A](http:/atcoder.jp/contests/abc188/tasks/abc188_A) | 7698 | 0 |  [go](./go/abc188/abc188_A/abc188_A.go) |  |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_B](http:/atcoder.jp/contests/abc188/tasks/abc188_B) | 7524 | 0 |  [go](./go/abc188/abc188_B/abc188_B.go) |  |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_C](http:/atcoder.jp/contests/abc188/tasks/abc188_C) | 6135 | 0 |  [go](./go/abc188/abc188_C/abc188_C.go) |  |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_D](http:/atcoder.jp/contests/abc188/tasks/abc188_D) | 2510 | 933 |  [go](./go/abc188/abc188_D/abc188_D.go) |  |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_E](http:/atcoder.jp/contests/abc188/tasks/abc188_E) | 1795 | 1170 |  [go](./go/abc188/abc188_E/abc188_E.go) | DAG DP |
| [abc188](http:/atcoder.jp/contests/abc188) | [abc188_F](http:/atcoder.jp/contests/abc188/tasks/abc188_F) | 480 | 1865 |  [go](./go/abc188/abc188_F/abc188_F.go) | Reverse problem.  Recursive functions with memorization. |

## abc187 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_A](http:/atcoder.jp/contests/abc187/tasks/abc187_A) | 7110 | 0 |  [go](./go/abc187/abc187_A/abc187_A.go) |  |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_B](http:/atcoder.jp/contests/abc187/tasks/abc187_B) | 6187 | 0 |  [go](./go/abc187/abc187_B/abc187_B.go) |  |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_C](http:/atcoder.jp/contests/abc187/tasks/abc187_C) | 5025 | 137 |  [go](./go/abc187/abc187_C/abc187_C.go) |  |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_D](http:/atcoder.jp/contests/abc187/tasks/abc187_D) | 3316 | 650 |  [go](./go/abc187/abc187_D/abc187_D.go) |  |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_E](http:/atcoder.jp/contests/abc187/tasks/abc187_E) | 1250 | 1358 |  [go](./go/abc187/abc187_E/abc187_E.go) | Simple tree DP |
| [abc187](http:/atcoder.jp/contests/abc187) | [abc187_F](http:/atcoder.jp/contests/abc187/tasks/abc187_F) | 438 | 1895 |  [go](./go/abc187/abc187_F/abc187_F.go) | Cliques, subset DP |

## abc186 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_A](http:/atcoder.jp/contests/abc186/tasks/abc186_A) | 6188 | 0 |  [go](./go/abc186/abc186_A/abc186_A.go) |  |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_B](http:/atcoder.jp/contests/abc186/tasks/abc186_B) | 5812 | 0 |  [go](./go/abc186/abc186_B/abc186_B.go) |  |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_C](http:/atcoder.jp/contests/abc186/tasks/abc186_C) | 4936 | 0 |  [go](./go/abc186/abc186_C/abc186_C.go) |  |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_D](http:/atcoder.jp/contests/abc186/tasks/abc186_D) | 3727 | 436 |  [go](./go/abc186/abc186_D/abc186_D.go) |  |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_E](http:/atcoder.jp/contests/abc186/tasks/abc186_E) | 979 | 1461 |  [go](./go/abc186/abc186_E/abc186_E.go) | Extended euclidean algorithm |
| [abc186](http:/atcoder.jp/contests/abc186) | [abc186_F](http:/atcoder.jp/contests/abc186/tasks/abc186_F) | 475 | 1833 |  [go](./go/abc186/abc186_F/abc186_F.go) | segtree or fenwick, chess |

## abc185 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_A](http:/atcoder.jp/contests/abc185/tasks/abc185_A) | 7332 | 0 |  [go](./go/abc185/abc185_A/abc185_A.go) |  |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_B](http:/atcoder.jp/contests/abc185/tasks/abc185_B) | 5993 | 0 |  [go](./go/abc185/abc185_B/abc185_B.go) |  |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_C](http:/atcoder.jp/contests/abc185/tasks/abc185_C) | 4326 | 373 |  [go](./go/abc185/abc185_C/abc185_C.go) |  |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_D](http:/atcoder.jp/contests/abc185/tasks/abc185_D) | 3882 | 490 |  [go](./go/abc185/abc185_D/abc185_D.go) |  |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_E](http:/atcoder.jp/contests/abc185/tasks/abc185_E) | 1006 | 1468 |  [go](./go/abc185/abc185_E/abc185_E.go) | Edit distance |
| [abc185](http:/atcoder.jp/contests/abc185) | [abc185_F](http:/atcoder.jp/contests/abc185/tasks/abc185_F) | 1994 | 1053 |  [go](./go/abc185/abc185_F/abc185_F.go) | Simple segtree problem |

## abc184 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_A](http:/atcoder.jp/contests/abc184/tasks/abc184_A) | 7738 | 0 |  [go](./go/abc184/abc184_A/abc184_A.go) |  |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_B](http:/atcoder.jp/contests/abc184/tasks/abc184_B) | 7439 | 0 |  [go](./go/abc184/abc184_B/abc184_B.go) |  |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_C](http:/atcoder.jp/contests/abc184/tasks/abc184_C) | 3137 | 782 |  [go](./go/abc184/abc184_C/abc184_C.go) |  |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_D](http:/atcoder.jp/contests/abc184/tasks/abc184_D) | 1561 | 1276 |  [go](./go/abc184/abc184_D/abc184_D.go) |  |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_E](http:/atcoder.jp/contests/abc184/tasks/abc184_E) | 1223 | 1418 |  [go](./go/abc184/abc184_E/abc184_E.go) | Grid, BFS, Teleporters |
| [abc184](http:/atcoder.jp/contests/abc184) | [abc184_F](http:/atcoder.jp/contests/abc184/tasks/abc184_F) | 1209 | 1423 |  [go](./go/abc184/abc184_F/abc184_F.go) | Meet in the middle, Subset sums |

## abc183 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_A](http:/atcoder.jp/contests/abc183/tasks/abc183_A) | 7152 | 0 |  [go](./go/abc183/abc183_A/abc183_A.go) |  |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_B](http:/atcoder.jp/contests/abc183/tasks/abc183_B) | 5919 | 0 |  [go](./go/abc183/abc183_B/abc183_B.go) |  |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_C](http:/atcoder.jp/contests/abc183/tasks/abc183_C) | 4633 | 329 |  [go](./go/abc183/abc183_C/abc183_C.go) |  |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_D](http:/atcoder.jp/contests/abc183/tasks/abc183_D) | 3400 | 662 |  [go](./go/abc183/abc183_D/abc183_D.go) |  |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_E](http:/atcoder.jp/contests/abc183/tasks/abc183_E) | 1393 | 1288 |  [go](./go/abc183/abc183_E/abc183_E.go) | Chess, Grid DP |
| [abc183](http:/atcoder.jp/contests/abc183) | [abc183_F](http:/atcoder.jp/contests/abc183/tasks/abc183_F) | 789 | 1586 |  [go](./go/abc183/abc183_F/abc183_F.go) | Nice augmented DSU problem |

## abc182 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_A](http:/atcoder.jp/contests/abc182/tasks/abc182_A) | 7436 | 0 |  [pypy3](./python/abc182/abc182_A.py) [go](./go/abc182/abc182_A/abc182_A.go) |  |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_B](http:/atcoder.jp/contests/abc182/tasks/abc182_B) | 6249 | 0 |  [pypy3](./python/abc182/abc182_B.py) [go](./go/abc182/abc182_B/abc182_B.go) |  |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_C](http:/atcoder.jp/contests/abc182/tasks/abc182_C) | 5077 | 274 |  [pypy3](./python/abc182/abc182_C.py) [go](./go/abc182/abc182_C/abc182_C.go) |  |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_D](http:/atcoder.jp/contests/abc182/tasks/abc182_D) | 3419 | 701 |  [pypy3](./python/abc182/abc182_D.py) [go](./go/abc182/abc182_D/abc182_D.go) |  |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_E](http:/atcoder.jp/contests/abc182/tasks/abc182_E) | 1988 | 1098 |  [pypy3](./python/abc182/abc182_E.py) [go](./go/abc182/abc182_E/abc182_E.go) | Simple grid problem |
| [abc182](http:/atcoder.jp/contests/abc182) | [abc182_F](http:/atcoder.jp/contests/abc182/tasks/abc182_F) | 233 | 2121 |  [pypy3](./python/abc182/abc182_F.py) [go](./go/abc182/abc182_F/abc182_F.go) | Money denominations, recursive, caching |

## abc181 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_A](http:/atcoder.jp/contests/abc181/tasks/abc181_A) | 6568 | 0 |  [pypy3](./python/abc181/abc181_A.py) [go](./go/abc181/abc181_A/abc181_A.go) |  |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_B](http:/atcoder.jp/contests/abc181/tasks/abc181_B) | 5936 | 0 |  [pypy3](./python/abc181/abc181_B.py) [go](./go/abc181/abc181_B/abc181_B.go) |  |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_C](http:/atcoder.jp/contests/abc181/tasks/abc181_C) | 4329 | 248 |  [pypy3](./python/abc181/abc181_C.py) [go](./go/abc181/abc181_C/abc181_C.go) |  |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_D](http:/atcoder.jp/contests/abc181/tasks/abc181_D) | 3152 | 600 |  [pypy3](./python/abc181/abc181_D.py) [go](./go/abc181/abc181_D/abc181_D.go) |  |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_E](http:/atcoder.jp/contests/abc181/tasks/abc181_E) | 1344 | 1193 |  [pypy3](./python/abc181/abc181_E.py) [go](./go/abc181/abc181_E/abc181_E.go) | Prefix Sum, Suffix sum, Binary Search |
| [abc181](http:/atcoder.jp/contests/abc181) | [abc181_F](http:/atcoder.jp/contests/abc181/tasks/abc181_F) | 230 | 2009 |  [pypy3](./python/abc181/abc181_F.py) [go](./go/abc181/abc181_F/abc181_F.go) | Binary search, DSU |

## abc180 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_A](http:/atcoder.jp/contests/abc180/tasks/abc180_A) | 5699 | 0 |  [pypy3](./python/abc180/abc180_A.py) [go](./go/abc180/abc180_A/abc180_A.go) |  |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_B](http:/atcoder.jp/contests/abc180/tasks/abc180_B) | 4683 | 0 |  [pypy3](./python/abc180/abc180_B.py) [go](./go/abc180/abc180_B/abc180_B.go) |  |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_C](http:/atcoder.jp/contests/abc180/tasks/abc180_C) | 4585 | 0 |  [pypy3](./python/abc180/abc180_C.py) [go](./go/abc180/abc180_C/abc180_C.go) |  |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_D](http:/atcoder.jp/contests/abc180/tasks/abc180_D) | 2471 | 752 |  [pypy3](./python/abc180/abc180_D.py) [go](./go/abc180/abc180_D/abc180_D.go) |  |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_E](http:/atcoder.jp/contests/abc180/tasks/abc180_E) | 1189 | 1256 |  [pypy3](./python/abc180/abc180_E.py) [go](./go/abc180/abc180_E/abc180_E.go) | Permutation to subset dp conversion |
| [abc180](http:/atcoder.jp/contests/abc180) | [abc180_F](http:/atcoder.jp/contests/abc180/tasks/abc180_F) | 80 | 2419 |  [pypy3](./python/abc180/abc180_F.py) [go](./go/abc180/abc180_F/abc180_F.go) | Graph creation, DP |

## abc179 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_A](http:/atcoder.jp/contests/abc179/tasks/abc179_A) | 8716 | 0 |  [pypy3](./python/abc179/abc179_A.py) [go](./go/abc179/abc179_A/abc179_A.go) |  |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_B](http:/atcoder.jp/contests/abc179/tasks/abc179_B) | 8333 | 0 |  [pypy3](./python/abc179/abc179_B.py) [go](./go/abc179/abc179_B/abc179_B.go) |  |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_C](http:/atcoder.jp/contests/abc179/tasks/abc179_C) | 6059 | 261 |  [pypy3](./python/abc179/abc179_C.py) [go](./go/abc179/abc179_C/abc179_C.go) |  |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_D](http:/atcoder.jp/contests/abc179/tasks/abc179_D) | 1777 | 1251 |  [pypy3](./python/abc179/abc179_D.py) [go](./go/abc179/abc179_D/abc179_D.go) |  |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_E](http:/atcoder.jp/contests/abc179/tasks/abc179_E) | 2008 | 1175 |  [pypy3](./python/abc179/abc179_E.py) [go](./go/abc179/abc179_E/abc179_E.go) | Typical 'find the sequence loop' problem |
| [abc179](http:/atcoder.jp/contests/abc179) | [abc179_F](http:/atcoder.jp/contests/abc179/tasks/abc179_F) | 707 | 1713 |  [pypy3](./python/abc179/abc179_F.py) [go](./go/abc179/abc179_F/abc179_F.go) | Another lazy segtree problem (other approaches too) |

## abc178 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_A](http:/atcoder.jp/contests/abc178/tasks/abc178_A) | 9574 | 0 |  [pypy3](./python/abc178/abc178_A.py) [go](./go/abc178/abc178_A/abc178_A.go) |  |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_B](http:/atcoder.jp/contests/abc178/tasks/abc178_B) | 8730 | 0 |  [pypy3](./python/abc178/abc178_B.py) [go](./go/abc178/abc178_B/abc178_B.go) |  |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_C](http:/atcoder.jp/contests/abc178/tasks/abc178_C) | 4565 | 653 |  [pypy3](./python/abc178/abc178_C.py) [go](./go/abc178/abc178_C/abc178_C.go) |  |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_D](http:/atcoder.jp/contests/abc178/tasks/abc178_D) | 3462 | 875 |  [pypy3](./python/abc178/abc178_D.py) [go](./go/abc178/abc178_D/abc178_D.go) |  |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_E](http:/atcoder.jp/contests/abc178/tasks/abc178_E) | 2653 | 1054 |  [pypy3](./python/abc178/abc178_E.py) [go](./go/abc178/abc178_E/abc178_E.go) | 45 deg Rotation for L1 -> Linf norm |
| [abc178](http:/atcoder.jp/contests/abc178) | [abc178_F](http:/atcoder.jp/contests/abc178/tasks/abc178_F) | 512 | 1877 |  [pypy3](./python/abc178/abc178_F.py) [go](./go/abc178/abc178_F/abc178_F.go) | Easy greedy assignment problem |

## abc177 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_A](http:/atcoder.jp/contests/abc177/tasks/abc177_A) | 9347 | 0 |  [pypy3](./python/abc177/abc177_A.py) [go](./go/abc177/abc177_A/abc177_A.go) |  |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_B](http:/atcoder.jp/contests/abc177/tasks/abc177_B) | 7137 | 108 |  [pypy3](./python/abc177/abc177_B.py) [go](./go/abc177/abc177_B/abc177_B.go) |  |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_C](http:/atcoder.jp/contests/abc177/tasks/abc177_C) | 5863 | 386 |  [pypy3](./python/abc177/abc177_C.py) [go](./go/abc177/abc177_C/abc177_C.go) |  |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_D](http:/atcoder.jp/contests/abc177/tasks/abc177_D) | 4125 | 732 |  [pypy3](./python/abc177/abc177_D.py) [go](./go/abc177/abc177_D/abc177_D.go) |  |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_E](http:/atcoder.jp/contests/abc177/tasks/abc177_E) | 2584 | 1057 |  [pypy3](./python/abc177/abc177_E.py) [go](./go/abc177/abc177_E/abc177_E.go) | Factor sieve |
| [abc177](http:/atcoder.jp/contests/abc177) | [abc177_F](http:/atcoder.jp/contests/abc177/tasks/abc177_F) | 152 | 2291 |  [pypy3](./python/abc177/abc177_F.py) [go](./go/abc177/abc177_F/abc177_F.go) | Tricky lazy segtree |

## abc176 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_A](http:/atcoder.jp/contests/abc176/tasks/abc176_A) | 9228 | 0 |  [pypy3](./python/abc176/abc176_A.py) [go](./go/abc176/abc176_A/abc176_A.go) |  |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_B](http:/atcoder.jp/contests/abc176/tasks/abc176_B) | 8679 | 0 |  [pypy3](./python/abc176/abc176_B.py) [go](./go/abc176/abc176_B/abc176_B.go) |  |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_C](http:/atcoder.jp/contests/abc176/tasks/abc176_C) | 8381 | 0 |  [pypy3](./python/abc176/abc176_C.py) [go](./go/abc176/abc176_C/abc176_C.go) |  |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_D](http:/atcoder.jp/contests/abc176/tasks/abc176_D) | 1804 | 1248 |  [pypy3](./python/abc176/abc176_D.py) [go](./go/abc176/abc176_D/abc176_D.go) |  |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_E](http:/atcoder.jp/contests/abc176/tasks/abc176_E) | 1958 | 1204 |  [pypy3](./python/abc176/abc176_E.py) [go](./go/abc176/abc176_E/abc176_E.go) | Coordinate compression, maps |
| [abc176](http:/atcoder.jp/contests/abc176) | [abc176_F](http:/atcoder.jp/contests/abc176/tasks/abc176_F) | 18 | 2912 |  [pypy3](./python/abc176/abc176_F.py) [go](./go/abc176/abc176_F/abc176_F.go) | Tricky DP (need to be efficient) |

## abc175 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_A](http:/atcoder.jp/contests/abc175/tasks/abc175_A) | 8513 | 0 |  [pypy3](./python/abc175/abc175_A.py) [go](./go/abc175/abc175_A/abc175_A.go) |  |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_B](http:/atcoder.jp/contests/abc175/tasks/abc175_B) | 7207 | 0 |  [pypy3](./python/abc175/abc175_B.py) [go](./go/abc175/abc175_B/abc175_B.go) |  |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_C](http:/atcoder.jp/contests/abc175/tasks/abc175_C) | 5316 | 417 |  [pypy3](./python/abc175/abc175_C.py) [go](./go/abc175/abc175_C/abc175_C.go) |  |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_D](http:/atcoder.jp/contests/abc175/tasks/abc175_D) | 1103 | 1491 |  [pypy3](./python/abc175/abc175_D.py) [go](./go/abc175/abc175_D/abc175_D.go) |  |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_E](http:/atcoder.jp/contests/abc175/tasks/abc175_E) | 984 | 1554 |  [pypy3](./python/abc175/abc175_E.py) [go](./go/abc175/abc175_E/abc175_E.go) | DP on a grid |
| [abc175](http:/atcoder.jp/contests/abc175) | [abc175_F](http:/atcoder.jp/contests/abc175/tasks/abc175_F) | 73 | 2512 |  [pypy3](./python/abc175/abc175_F.py) [go](./go/abc175/abc175_F/abc175_F.go) |  |

## abc174 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_A](http:/atcoder.jp/contests/abc174/tasks/abc174_A) | 9720 | 0 |  [pypy3](./python/abc174/abc174_A.py) [go](./go/abc174/abc174_A/abc174_A.go) |  |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_B](http:/atcoder.jp/contests/abc174/tasks/abc174_B) | 9103 | 0 |  [pypy3](./python/abc174/abc174_B.py) [go](./go/abc174/abc174_B/abc174_B.go) |  |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_C](http:/atcoder.jp/contests/abc174/tasks/abc174_C) | 3349 | 902 |  [pypy3](./python/abc174/abc174_C.py) [go](./go/abc174/abc174_C/abc174_C.go) |  |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_D](http:/atcoder.jp/contests/abc174/tasks/abc174_D) | 5486 | 486 |  [pypy3](./python/abc174/abc174_D.py) [go](./go/abc174/abc174_D/abc174_D.go) |  |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_E](http:/atcoder.jp/contests/abc174/tasks/abc174_E) | 2013 | 1227 |  [pypy3](./python/abc174/abc174_E.py) [go](./go/abc174/abc174_E/abc174_E.go) | Simple binary search |
| [abc174](http:/atcoder.jp/contests/abc174) | [abc174_F](http:/atcoder.jp/contests/abc174/tasks/abc174_F) | 1257 | 1495 |  [pypy3](./python/abc174/abc174_F.py) [go](./go/abc174/abc174_F/abc174_F.go) | Range set query. Offline processing. Fenwick tree. |

## abc173 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_A](http:/atcoder.jp/contests/abc173/tasks/abc173_A) | 10564 | 0 |  [pypy3](./python/abc173/abc173_A.py) [go](./go/abc173/abc173_A/abc173_A.go) |  |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_B](http:/atcoder.jp/contests/abc173/tasks/abc173_B) | 10269 | 0 |  [pypy3](./python/abc173/abc173_B.py) [go](./go/abc173/abc173_B/abc173_B.go) |  |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_C](http:/atcoder.jp/contests/abc173/tasks/abc173_C) | 4890 | 653 |  [pypy3](./python/abc173/abc173_C.py) [go](./go/abc173/abc173_C/abc173_C.go) |  |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_D](http:/atcoder.jp/contests/abc173/tasks/abc173_D) | 4555 | 720 |  [pypy3](./python/abc173/abc173_D.py) [go](./go/abc173/abc173_D/abc173_D.go) |  |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_E](http:/atcoder.jp/contests/abc173/tasks/abc173_E) | 972 | 1607 |  [pypy3](./python/abc173/abc173_E.py) [go](./go/abc173/abc173_E/abc173_E.go) | Either lots of casework or smarter brute force (I did the casework). |
| [abc173](http:/atcoder.jp/contests/abc173) | [abc173_F](http:/atcoder.jp/contests/abc173/tasks/abc173_F) | 499 | 1892 |  [pypy3](./python/abc173/abc173_F.py) [go](./go/abc173/abc173_F/abc173_F.go) | Counting.  Connected components on tree. |

## abc172 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_A](http:/atcoder.jp/contests/abc172/tasks/abc172_A) | 10050 | 0 |  [pypy3](./python/abc172/abc172_A.py) [go](./go/abc172/abc172_A/abc172_A.go) |  |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_B](http:/atcoder.jp/contests/abc172/tasks/abc172_B) | 9757 | 0 |  [pypy3](./python/abc172/abc172_B.py) [go](./go/abc172/abc172_B/abc172_B.go) |  |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_C](http:/atcoder.jp/contests/abc172/tasks/abc172_C) | 3234 | 930 |  [pypy3](./python/abc172/abc172_C.py) [go](./go/abc172/abc172_C/abc172_C.go) |  |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_D](http:/atcoder.jp/contests/abc172/tasks/abc172_D) | 3119 | 963 |  [pypy3](./python/abc172/abc172_D.py) [go](./go/abc172/abc172_D/abc172_D.go) |  |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_E](http:/atcoder.jp/contests/abc172/tasks/abc172_E) | 470 | 1880 |  [pypy3](./python/abc172/abc172_E.py) [go](./go/abc172/abc172_E/abc172_E.go) | Counting. Inclusion/Exclusion |
| [abc172](http:/atcoder.jp/contests/abc172) | [abc172_F](http:/atcoder.jp/contests/abc172/tasks/abc172_F) | 186 | 2216 |  [pypy3](./python/abc172/abc172_F.py) [go](./go/abc172/abc172_F/abc172_F.go) | NIM. Relationship between xor and addition. Bitwise. |

## abc171 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_A](http:/atcoder.jp/contests/abc171/tasks/abc171_A) | 10416 | 0 |  [pypy3](./python/abc171/abc171_A.py) [go](./go/abc171/abc171_A/abc171_A.go) |  |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_B](http:/atcoder.jp/contests/abc171/tasks/abc171_B) | 10164 | 0 |  [pypy3](./python/abc171/abc171_B.py) [go](./go/abc171/abc171_B/abc171_B.go) |  |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_C](http:/atcoder.jp/contests/abc171/tasks/abc171_C) | 5522 | 560 |  [pypy3](./python/abc171/abc171_C.py) [go](./go/abc171/abc171_C/abc171_C.go) |  |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_D](http:/atcoder.jp/contests/abc171/tasks/abc171_D) | 5885 | 498 |  [pypy3](./python/abc171/abc171_D.py) [go](./go/abc171/abc171_D/abc171_D.go) |  |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_E](http:/atcoder.jp/contests/abc171/tasks/abc171_E) | 4264 | 778 |  [pypy3](./python/abc171/abc171_E.py) [go](./go/abc171/abc171_E/abc171_E.go) | Simple xor problem |
| [abc171](http:/atcoder.jp/contests/abc171) | [abc171_F](http:/atcoder.jp/contests/abc171/tasks/abc171_F) | 624 | 1795 |  [pypy3](./python/abc171/abc171_F.py) [go](./go/abc171/abc171_F/abc171_F.go) | Counting. Strings of length N with given subsequence |

## abc170 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_A](http:/atcoder.jp/contests/abc170/tasks/abc170_A) | 10398 | 0 |  [pypy3](./python/abc170/abc170_A.py) [go](./go/abc170/abc170_A/abc170_A.go) |  |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_B](http:/atcoder.jp/contests/abc170/tasks/abc170_B) | 9617 | 0 |  [pypy3](./python/abc170/abc170_B.py) [go](./go/abc170/abc170_B/abc170_B.go) |  |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_C](http:/atcoder.jp/contests/abc170/tasks/abc170_C) | 8272 | 15 |  [pypy3](./python/abc170/abc170_C.py) [go](./go/abc170/abc170_C/abc170_C.go) |  |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_D](http:/atcoder.jp/contests/abc170/tasks/abc170_D) | 2835 | 1033 |  [pypy3](./python/abc170/abc170_D.py) [go](./go/abc170/abc170_D/abc170_D.go) |  |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_E](http:/atcoder.jp/contests/abc170/tasks/abc170_E) | 1177 | 1502 |  [pypy3](./python/abc170/abc170_E.py) [go](./go/abc170/abc170_E/abc170_E.go) | Lots of minheaps or lots of multisets |
| [abc170](http:/atcoder.jp/contests/abc170) | [abc170_F](http:/atcoder.jp/contests/abc170/tasks/abc170_F) | 388 | 1968 |  [pypy3](./python/abc170/abc170_F.py) [go](./go/abc170/abc170_F/abc170_F.go) | Clever Dijkstra.  Somewhat time challenged in both go and python (there is likely more optimization to be done). |

## abc169 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_A](http:/atcoder.jp/contests/abc169/tasks/abc169_A) | 11268 | 0 |  [pypy3](./python/abc169/abc169_A.py) [go](./go/abc169/abc169_A/abc169_A.go) |  |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_B](http:/atcoder.jp/contests/abc169/tasks/abc169_B) | 7017 | 349 |  [pypy3](./python/abc169/abc169_B.py) [go](./go/abc169/abc169_B/abc169_B.go) |  |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_C](http:/atcoder.jp/contests/abc169/tasks/abc169_C) | 5525 | 597 |  [pypy3](./python/abc169/abc169_C.py) [go](./go/abc169/abc169_C/abc169_C.go) |  |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_D](http:/atcoder.jp/contests/abc169/tasks/abc169_D) | 4591 | 732 |  [pypy3](./python/abc169/abc169_D.py) [go](./go/abc169/abc169_D/abc169_D.go) |  |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_E](http:/atcoder.jp/contests/abc169/tasks/abc169_E) | 1659 | 1353 |  [pypy3](./python/abc169/abc169_E.py) [go](./go/abc169/abc169_E/abc169_E.go) | Median, Casework |
| [abc169](http:/atcoder.jp/contests/abc169) | [abc169_F](http:/atcoder.jp/contests/abc169/tasks/abc169_F) | 818 | 1698 |  [pypy3](./python/abc169/abc169_F.py) [go](./go/abc169/abc169_F/abc169_F.go) | DP, Subsets |

## abc168 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_A](http:/atcoder.jp/contests/abc168/tasks/abc168_A) | 10686 | 0 |  [pypy3](./python/abc168/abc168_A.py) [go](./go/abc168/abc168_A/abc168_A.go) |  |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_B](http:/atcoder.jp/contests/abc168/tasks/abc168_B) | 10466 | 0 |  [pypy3](./python/abc168/abc168_B.py) [go](./go/abc168/abc168_B/abc168_B.go) |  |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_C](http:/atcoder.jp/contests/abc168/tasks/abc168_C) | 7598 | 178 |  [pypy3](./python/abc168/abc168_C.py) [go](./go/abc168/abc168_C/abc168_C.go) |  |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_D](http:/atcoder.jp/contests/abc168/tasks/abc168_D) | 3856 | 804 |  [pypy3](./python/abc168/abc168_D.py) [go](./go/abc168/abc168_D/abc168_D.go) |  |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_E](http:/atcoder.jp/contests/abc168/tasks/abc168_E) | 467 | 1896 |  [pypy3](./python/abc168/abc168_E.py) [go](./go/abc168/abc168_E/abc168_E.go) | Dot product, GCD, Combinatorics, Special Cases |
| [abc168](http:/atcoder.jp/contests/abc168) | [abc168_F](http:/atcoder.jp/contests/abc168/tasks/abc168_F) | 85 | 2478 |  [pypy3](./python/abc168/abc168_F.py) [go](./go/abc168/abc168_F/abc168_F.go) | Coordinate compression, BFS, Bookkeeping.  Note: big swing on constant factors dependent on subtle algorithm choices. My first coding of pypy was too slow.  Using byte arrays (array.array -- factor of 2) and slightly better algorithm (1 coord per segment vs. 3 -- factor of 9) made it much faster. |

## abc167 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_A](http:/atcoder.jp/contests/abc167/tasks/abc167_A) | 11507 | 0 |  [pypy3](./python/abc167/abc167_A.py) [go](./go/abc167/abc167_A/abc167_A.go) |  |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_B](http:/atcoder.jp/contests/abc167/tasks/abc167_B) | 10836 | 0 |  [pypy3](./python/abc167/abc167_B.py) [go](./go/abc167/abc167_B/abc167_B.go) |  |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_C](http:/atcoder.jp/contests/abc167/tasks/abc167_C) | 5652 | 595 |  [pypy3](./python/abc167/abc167_C.py) [go](./go/abc167/abc167_C/abc167_C.go) |  |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_D](http:/atcoder.jp/contests/abc167/tasks/abc167_D) | 4666 | 754 |  [pypy3](./python/abc167/abc167_D.py) [go](./go/abc167/abc167_D/abc167_D.go) |  |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_E](http:/atcoder.jp/contests/abc167/tasks/abc167_E) | 1393 | 1442 |  [pypy3](./python/abc167/abc167_E.py) [go](./go/abc167/abc167_E/abc167_E.go) | Simple Combinatorics |
| [abc167](http:/atcoder.jp/contests/abc167) | [abc167_F](http:/atcoder.jp/contests/abc167/tasks/abc167_F) | 415 | 1961 |  [pypy3](./python/abc167/abc167_F.py) [go](./go/abc167/abc167_F/abc167_F.go) | Bracket Sequences |

## abc166 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_A](http:/atcoder.jp/contests/abc166/tasks/abc166_A) | 11540 | 0 |  [pypy3](./python/abc166/abc166_A.py) [go](./go/abc166/abc166_A/abc166_A.go) |  |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_B](http:/atcoder.jp/contests/abc166/tasks/abc166_B) | 10088 | 0 |  [pypy3](./python/abc166/abc166_B.py) [go](./go/abc166/abc166_B/abc166_B.go) |  |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_C](http:/atcoder.jp/contests/abc166/tasks/abc166_C) | 7970 | 233 |  [pypy3](./python/abc166/abc166_C.py) [go](./go/abc166/abc166_C/abc166_C.go) |  |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_D](http:/atcoder.jp/contests/abc166/tasks/abc166_D) | 4999 | 694 |  [pypy3](./python/abc166/abc166_D.py) [go](./go/abc166/abc166_D/abc166_D.go) |  |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_E](http:/atcoder.jp/contests/abc166/tasks/abc166_E) | 2874 | 1062 |  [pypy3](./python/abc166/abc166_E.py) [go](./go/abc166/abc166_E/abc166_E.go) | Simple DP |
| [abc166](http:/atcoder.jp/contests/abc166) | [abc166_F](http:/atcoder.jp/contests/abc166/tasks/abc166_F) | 863 | 1668 |  [pypy3](./python/abc166/abc166_F.py) [go](./go/abc166/abc166_F/abc166_F.go) | Greedy game theory |

## abc165 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_A](http:/atcoder.jp/contests/abc165/tasks/abc165_A) | 11225 | 0 |  [pypy3](./python/abc165/abc165_A.py) [go](./go/abc165/abc165_A/abc165_A.go) |  |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_B](http:/atcoder.jp/contests/abc165/tasks/abc165_B) | 9994 | 0 |  [pypy3](./python/abc165/abc165_B.py) [go](./go/abc165/abc165_B/abc165_B.go) |  |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_C](http:/atcoder.jp/contests/abc165/tasks/abc165_C) | 2514 | 1136 |  [pypy3](./python/abc165/abc165_C.py) [go](./go/abc165/abc165_C/abc165_C.go) |  |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_D](http:/atcoder.jp/contests/abc165/tasks/abc165_D) | 5554 | 600 |  [pypy3](./python/abc165/abc165_D.py) [go](./go/abc165/abc165_D/abc165_D.go) |  |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_E](http:/atcoder.jp/contests/abc165/tasks/abc165_E) | 975 | 1620 |  [pypy3](./python/abc165/abc165_E.py) [go](./go/abc165/abc165_E/abc165_E.go) | Interesting matching construction |
| [abc165](http:/atcoder.jp/contests/abc165) | [abc165_F](http:/atcoder.jp/contests/abc165/tasks/abc165_F) | 596 | 1843 |  [pypy3](./python/abc165/abc165_F.py) [go](./go/abc165/abc165_F/abc165_F.go) | Tree DP, DFS w/ rollback |

## abc164 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_A](http:/atcoder.jp/contests/abc164/tasks/abc164_A) | 11148 | 0 |  [pypy3](./python/abc164/abc164_A.py) [go](./go/abc164/abc164_A/abc164_A.go) |  |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_B](http:/atcoder.jp/contests/abc164/tasks/abc164_B) | 10415 | 0 |  [pypy3](./python/abc164/abc164_B.py) [go](./go/abc164/abc164_B/abc164_B.go) |  |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_C](http:/atcoder.jp/contests/abc164/tasks/abc164_C) | 9553 | 0 |  [pypy3](./python/abc164/abc164_C.py) [go](./go/abc164/abc164_C/abc164_C.go) |  |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_D](http:/atcoder.jp/contests/abc164/tasks/abc164_D) | 1926 | 1232 |  [pypy3](./python/abc164/abc164_D.py) [go](./go/abc164/abc164_D/abc164_D.go) |  |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_E](http:/atcoder.jp/contests/abc164/tasks/abc164_E) | 497 | 1877 |  [pypy3](./python/abc164/abc164_E.py) [go](./go/abc164/abc164_E/abc164_E.go) | Interesting Dijkstra |
| [abc164](http:/atcoder.jp/contests/abc164) | [abc164_F](http:/atcoder.jp/contests/abc164/tasks/abc164_F) | 41 | 2683 |  [pypy3](./python/abc164/abc164_F.py) [go](./go/abc164/abc164_F/abc164_F.go) | Bitwise, Logical constraint solving |

## abc163 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_A](http:/atcoder.jp/contests/abc163/tasks/abc163_A) | 11284 | 0 |  [pypy3](./python/abc163/abc163_A.py) [go](./go/abc163/abc163_A/abc163_A.go) |  |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_B](http:/atcoder.jp/contests/abc163/tasks/abc163_B) | 10786 | 0 |  [pypy3](./python/abc163/abc163_B.py) [go](./go/abc163/abc163_B/abc163_B.go) |  |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_C](http:/atcoder.jp/contests/abc163/tasks/abc163_C) | 8976 | 125 |  [pypy3](./python/abc163/abc163_C.py) [go](./go/abc163/abc163_C/abc163_C.go) |  |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_D](http:/atcoder.jp/contests/abc163/tasks/abc163_D) | 3663 | 960 |  [pypy3](./python/abc163/abc163_D.py) [go](./go/abc163/abc163_D/abc163_D.go) |  |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_E](http:/atcoder.jp/contests/abc163/tasks/abc163_E) | 380 | 2037 |  [pypy3](./python/abc163/abc163_E.py) [go](./go/abc163/abc163_E/abc163_E.go) | DP, Greedy |
| [abc163](http:/atcoder.jp/contests/abc163) | [abc163_F](http:/atcoder.jp/contests/abc163/tasks/abc163_F) | 107 | 2470 |  [pypy3](./python/abc163/abc163_F.py) [go](./go/abc163/abc163_F/abc163_F.go) | Tree DP |

## abc162 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_A](http:/atcoder.jp/contests/abc162/tasks/abc162_A) | 10333 | 0 |  [pypy3](./python/abc162/abc162_A.py) [go](./go/abc162/abc162_A/abc162_A.go) |  |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_B](http:/atcoder.jp/contests/abc162/tasks/abc162_B) | 9924 | 0 |  [pypy3](./python/abc162/abc162_B.py) [go](./go/abc162/abc162_B/abc162_B.go) |  |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_C](http:/atcoder.jp/contests/abc162/tasks/abc162_C) | 8304 | 34 |  [pypy3](./python/abc162/abc162_C.py) [go](./go/abc162/abc162_C/abc162_C.go) |  |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_D](http:/atcoder.jp/contests/abc162/tasks/abc162_D) | 4199 | 757 |  [pypy3](./python/abc162/abc162_D.py) [go](./go/abc162/abc162_D/abc162_D.go) |  |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_E](http:/atcoder.jp/contests/abc162/tasks/abc162_E) | 850 | 1662 |  [pypy3](./python/abc162/abc162_E.py) [go](./go/abc162/abc162_E/abc162_E.go) | Divisors, Inclusion/Exclusion, Modular Math |
| [abc162](http:/atcoder.jp/contests/abc162) | [abc162_F](http:/atcoder.jp/contests/abc162/tasks/abc162_F) | 680 | 1764 |  [pypy3](./python/abc162/abc162_F.py) [go](./go/abc162/abc162_F/abc162_F.go) | DP of DP |

## abc161 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_A](http:/atcoder.jp/contests/abc161/tasks/abc161_A) | 9784 | 0 |  [pypy3](./python/abc161/abc161_A.py) [go](./go/abc161/abc161_A/abc161_A.go) |  |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_B](http:/atcoder.jp/contests/abc161/tasks/abc161_B) | 8263 | 0 |  [pypy3](./python/abc161/abc161_B.py) [go](./go/abc161/abc161_B/abc161_B.go) |  |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_C](http:/atcoder.jp/contests/abc161/tasks/abc161_C) | 7740 | 4 |  [pypy3](./python/abc161/abc161_C.py) [go](./go/abc161/abc161_C/abc161_C.go) |  |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_D](http:/atcoder.jp/contests/abc161/tasks/abc161_D) | 2865 | 991 |  [pypy3](./python/abc161/abc161_D.py) [go](./go/abc161/abc161_D/abc161_D.go) |  |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_E](http:/atcoder.jp/contests/abc161/tasks/abc161_E) | 682 | 1760 |  [pypy3](./python/abc161/abc161_E.py) [go](./go/abc161/abc161_E/abc161_E.go) | Simple DP |
| [abc161](http:/atcoder.jp/contests/abc161) | [abc161_F](http:/atcoder.jp/contests/abc161/tasks/abc161_F) | 1117 | 1528 |  [pypy3](./python/abc161/abc161_F.py) [go](./go/abc161/abc161_F/abc161_F.go) | GCD, Factors |

## abc160 Solutions
| Contest | Problem | Num Correct | Diff Rating | Solutions | Notes |
| ------- | ------- | ----------: | ----------: | --------- | ----- |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_A](http:/atcoder.jp/contests/abc160/tasks/abc160_A) | 9557 | 0 |  [pypy3](./python/abc160/abc160_A.py) [go](./go/abc160/abc160_A/abc160_A.go) |  |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_B](http:/atcoder.jp/contests/abc160/tasks/abc160_B) | 9427 | 0 |  [pypy3](./python/abc160/abc160_B.py) [go](./go/abc160/abc160_B/abc160_B.go) |  |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_C](http:/atcoder.jp/contests/abc160/tasks/abc160_C) | 7443 | 62 |  [pypy3](./python/abc160/abc160_C.py) [go](./go/abc160/abc160_C/abc160_C.go) |  |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_D](http:/atcoder.jp/contests/abc160/tasks/abc160_D) | 3417 | 879 |  [pypy3](./python/abc160/abc160_D.py) [go](./go/abc160/abc160_D/abc160_D.go) |  |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_E](http:/atcoder.jp/contests/abc160/tasks/abc160_E) | 2721 | 1036 |  [pypy3](./python/abc160/abc160_E.py) [go](./go/abc160/abc160_E/abc160_E.go) | Max Heap or simple sorting |
| [abc160](http:/atcoder.jp/contests/abc160) | [abc160_F](http:/atcoder.jp/contests/abc160/tasks/abc160_F) | 339 | 2048 |  [pypy3](./python/abc160/abc160_F.py) [go](./go/abc160/abc160_F/abc160_F.go) | Tree DP, Rerooting, Combinatorics |

