# STC - **ST**atistics **C**li

Calculate basic statistics from a bash pipeline. By default it does not need
to save any data, and therefore can handle arbitrarily sized streams of data.

```bash
$ python3 -c "import math; print('\n'.join(str(math.sin(i / 150) / i) for i in range(1, 3000)))" \
| stc
lines: 2999 sum: 1.5447562772859111 mean: 0.0005150904559139416 max: 0.006666617284060357 min: -0.0014482241838785694
```

Optionally can make calculations that require all of the data in memory.

```bash
$ python3 -c "import math; print('\n'.join(str(math.sin(i / 150) / i) for i in range(1, 3000)))" \
| stc -a
lines: 2999 sum: 1.5447562772859111 mean: 0.0005150904559139416 max: 0.006666617284060357 min: -0.0014482241838785694 median: 5.745873438259939e-05 mode: (-0.0014482241838785694 1x) variance: 3.1624275041922932e-06 stddev: 0.0017783215412833229 pct1: -0.001441104853002076 pct5: -0.0012715119850872374 pct10: -0.000780694754377271 pct25: -0.00037166704970215526 pct75: 0.00047223240104556915 pct95: 0.005609806565385977 pct99: 0.006622311026502041 pct99.9: 0.0066662222311110264
```

Graph data in the terminal.

```bash
$ python3 -c "import math; print('\n'.join(str(math.sin(i / 150) / i) for i in range(1, 3000)))" \
| stc -g -xdim 100 -ydim 50
 6.67e-03 |*
 6.50e-03 |**
 6.34e-03 | **
 6.18e-03 |  **
 6.02e-03 |   *
 5.86e-03 |   **
 5.69e-03 |    *
 5.53e-03 |    **
 5.37e-03 |     *
 5.21e-03 |     *
 5.04e-03 |     **
 4.88e-03 |      *
 4.72e-03 |      *
 4.56e-03 |      **
 4.39e-03 |       *
 4.23e-03 |       *
 4.07e-03 |       **
 3.91e-03 |        *
 3.75e-03 |        *
 3.58e-03 |        **
 3.42e-03 |         *
 3.26e-03 |         *
 3.10e-03 |         *
 2.93e-03 |         **
 2.77e-03 |          *
 2.61e-03 |          *
 2.45e-03 |          **
 2.28e-03 |           *
 2.12e-03 |           *
 1.96e-03 |           *
 1.80e-03 |           **
 1.64e-03 |            *
 1.47e-03 |            *
 1.31e-03 |            **
 1.15e-03 |             *
 9.86e-04 |             *
 8.24e-04 |             *                       ****
 6.62e-04 |             **                    ***  ***
 4.99e-04 |              *                   **      **
 3.37e-04 |              *                  **        **                     *********
 1.75e-04 |              **                **          **                  ***       ***                   *****
 0.00e+00 |---------------*---------------**------------***--------------***-----------***---------------***----
-1.62e-04 |               **             **               **            **               ***          ****
-3.25e-04 |                *            **                 **         ***                  ****     ***
-4.87e-04 |                *            *                   ***     ***                       *******
-6.49e-04 |                **          **                     *******
-8.11e-04 |                 *         **
-9.74e-04 |                  *       **
-1.14e-03 |                  **      *
-1.30e-03 |                   **   **
-1.46e-03 |                      *
lines: 2999 sum: 1.5447562772859111 mean: 0.0005150904559139416 max: 0.006666617284060357 min: -0.0014482241838785694
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
