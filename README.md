# utils
A set of small packages that aren't worth their own repositories but are still useful and used across the projects

## buffer
Ordinary buffer with pre-allocated (but dynamically-growing up to upper limit) storage for values. The only difference between it and, to say, `bytes.Buffer`, is that it contains markers. This isn't a lot of code neither difficult, but this found its application all across the indigo by itself (primarily http parser with header keys and values) and some other cross-projects.

## constraint
A small set of generic annotations.

## ft
Functional Tools: map, reduce, summer. Yet another fp-tools library without even tests.

## mapconv
Contains `Keys` and `Copy` map operations. Ain't much but it's honest work.

## pool
Simple thread-unsafe object-pool. Wrote own implementation as too lazy to benchmark the one from std-library and compare them (thread-safety is anyway not needed, so considered by me to be an overkill). 

## uf
Small package with common and widely used (at least in indigo and some other side projects) unsafe functions `B2S` and `S2B`. The implementation is taken from fasthttp.
