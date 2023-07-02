# utils
A set of small packages that aren't worth their own repositories but are still useful and used across the projects

## arena
Small thread-unsafe self-made implementation of an arena. The point of it is to store many defragmented elements (one by one in other words). In indigo, for example, header keys and values are stored in this manner. 

## constraint
A small set of generic annotations.

## ft
Functional Tools: map, reduce. Yet another fp-tools library without even tests.

## mapconv
Contains `Keys` and `Copy` map operations. Ain't much but it's honest work.

## pool
Simple thread-unsafe object-pool. Wrote own implementation as too lazy to benchmark the one from std-library. 

## uf
A set of unsafe features, like a hacky conversion between string and byte-slice.
