benchcmp-vcs
=====================

benchcmp-vcs is a wrapper of benchcmp.

[![Build Status](https://travis-ci.org/pocke/benchcmp-vcs.svg?branch=master)](https://travis-ci.org/pocke/benchcmp-vcs)
[![Coverage Status](https://coveralls.io/repos/pocke/benchcmp-vcs/badge.svg?branch=master)](https://coveralls.io/r/pocke/benchcmp-vcs?branch=master)



Installation
-----------------

```sh
go get -u golang.org/x/tools/cmd/benchcmp
go get github.com/pocke/benchcmp-vcs
```


Usage
-----------

```sh
$ benchcmp-vcs
old revision: 1811c66af7ba48ab8f34b8bda2476ff3198f0ace
new revision: b71cc3ac54e6f045690e30a5d6ba48df1d30f0e5

benchmark          old ns/op     new ns/op     delta
BenchmarkBuild     24716         16021         -35.18%

benchmark          old allocs     new allocs     delta
BenchmarkBuild     305            5              -98.36%

benchmark          old bytes     new bytes     delta
BenchmarkBuild     3053          2749          -9.96%
```

Why should use benchcmp-vcs?
------------------------------

### If you use raw benchcmp.

```sh
$ go test -run=NONE -bench=. -benchmem > new.txt
$ git co HEAD~
$ go test -run=NONE -bench=. -benchmem > old.txt
$ go co master
$ benchcmp old.txt new.txt
$ rm old.txt new.txt
```

Many commands....

### If you use benchcmp-vcs

```sh
$ benchcmp-vcs
```

Very Simple!


Command line options
----------------------

*TODO*

Not Implemented...



License
-------------

Copyright &copy; 2015 Masataka Kuwabara
Licensed [MIT][mit]
[MIT]: http://www.opensource.org/licenses/mit-license.php
