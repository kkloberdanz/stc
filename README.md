# STC - **ST**atistics **C**li

Calculate basic statistics from a bash pipeline. By default it does not need
to save any data, and therefore can handle arbitrarily sized streams of data.

```bash
$ cat ~/src/quotes2.txt | grep DOGEUSD | awk '{ print $6 }' | stc
lines: 98509 sum: 20959.062716800796 mean: 0.21276292234009883 max: 0.2301995 min: 0.19782
```

Optionally can make calculations that require all of the data in memory

```bash
$ cat ~/src/quotes2.txt | grep DOGEUSD | awk '{ print $6 }' | stc -a
lines: 98509 sum: 20959.062716800796 mean: 0.21276292234009883 max: 0.2301995 min: 0.19782 median: 0.2111222 mode: (0.2045479 1404x) variance: 3.9813301595625433e-05 stddev: 0.006309778252492352 pct1: 0.2013022 pct5: 0.2037395 pct10: 0.2058326 pct25: 0.2085667 pct75: 0.217 pct95: 0.2241104 pct99: 0.2256 pct99.9: 0.22851
```

Graph data in the terminal

```bash
$ cat ~/src/quotes2.txt | grep DOGEUSD | awk '{ print $6 }' | stc -a -g -xdim 80 -ydim 40
 2.30e-01 |                                                                           *
 2.29e-01 |                                                                           *
 2.29e-01 |*                                                                          *
 2.28e-01 |*                                                                          **
 2.27e-01 |*                                                                    *     **
 2.26e-01 |*                                                                    *     **
 2.25e-01 |*                                                                    **    ****
 2.25e-01 |*                                                                    ** *  *****
 2.24e-01 |** **                                                                ** * ******
 2.23e-01 |******                                                               ** ********
 2.22e-01 |*******                                                              ****** * ***
 2.21e-01 |*** ***                                                              ******    *
 2.20e-01 |* *  **                                                              *****
 2.20e-01 |* *   ***                                                            *** *
 2.19e-01 |*     *****                                                          *** *
 2.18e-01 |*     *****                            *                             * *
 2.17e-01 |      ** **                            *                             * *
 2.16e-01 |      ** **                            *                             *
 2.16e-01 |      ** **       *                    *                             *
 2.15e-01 |      ** ***  * ****                  ***                            *
 2.14e-01 |          **  * *****    *            ****                *         **
 2.13e-01 |          **  *******   **           ***** *           *  *         **
 2.12e-01 |           *  *******   **           ********          *  **        *
 2.12e-01 |           *  ****  * ****    *   *****  ****          * ***     * **
 2.11e-01 |           *  ****  * *****   **  *****  ****   *      * ******  ****
 2.10e-01 |           * ** *   * ** **  **** ***    ****   *      * ******* ***
 2.09e-01 |           **** *   **** **  ****   *     *** * **     ********* **
 2.08e-01 |           ***      ****  ** * **           ******     ****** ****
 2.08e-01 |           ***      **    ** *  **          ******     ***     **
 2.07e-01 |           ***       *    ** *  ***          *****  ** **      **
 2.06e-01 |           ***       *    ** *   *           ***** *** **
 2.05e-01 |           ***       *    ** *               ***** *** *
 2.04e-01 |           ***             ***                   *** * *
 2.03e-01 |            **             ***                   **  ***
 2.03e-01 |            **             **                    **  ***
 2.02e-01 |            **             **                    **  **
 2.01e-01 |            **             **                    **
 2.00e-01 |            **             *                      *
 1.99e-01 |             *             *                      *
 1.99e-01 |             *             *                      *
 1.98e-01 |                                                  *
lines: 98509 sum: 20959.062716800796 mean: 0.21276292234009883 max: 0.2301995 min: 0.19782 median: 0.2111222 mode: (0.2045479 1404x) variance: 3.9813301595625433e-05 stddev: 0.006309778252492352 pct1: 0.2013022 pct5: 0.2037395 pct10: 0.2058326 pct25: 0.2085667 pct75: 0.217 pct95: 0.2241104 pct99: 0.2256 pct99.9: 0.22851
```

Usage:

```
$ stc -h
Usage of stc:
  -a	enables statistics that require memory allocation
  -g	graph the data in the terminal
  -xdim int
    	character length of x axis (default 68)
  -ydim int
    	character length of y axis (default 20)
```
