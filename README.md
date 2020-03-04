Faster English Porter2 Stemmer for Go
=====================================

This is a reworked fork of the English Porter2 stemmer from Dmitry Chestnykn's
https://github.com/dchest/stemmer/ package, optimised as much for performance as possible
while preserving the original algorithm's shape.

The porter2 stemming algorithm is described here:
http://snowball.tartarus.org/algorithms/english/stemmer.html

This fork operates on bytes rather than runes internally, so may have issues with
Unicode safety, though tests with UTF-8 chars are passing. Most of the
algorithm is just matching on ASCII characters anyway, so it's highly unlikely
to cause problems.

**Very Warning!** Unlike the upstream repo, the version of the `Stem()`
function in this package _mutates the incoming byte slice_, so please remember
to take care of your memory if this is a problem, or use the `StemString()`
function also provided by this package. **Double Warning!**: I may change these
to `Stem()/StemBytes()` instead.

Also, I haven't exactly wrung every last drop out of this, I tapped out after
getting it nearly 90% faster. If you decide you need _even more speed_, I'd
love to hear about what crazy tricks you pull to drag more performance out of
this! It's fast enough for me for now though, so I've stopped.


Expectation Management
----------------------

I have prepared this fork to suit my own strange needs, and will continue to
hack on it as required.

If you would like to take advantage of this stemmer's performance improvements,
I strongly recommend either forking or vendoring as I will not guarantee any
stability, and may even decide to trade some accuracy for more speed at some
point (but will endeavour to hide this behind a flag if possible).

I endeavour to respond to issues as quickly as I can, but I make no promises.
Pull requests are unlikely to be accepted without a conversation prior to
commencement.


Silly Benchmarks Game
---------------------

Here is the output of `benchcmp` after running on my i7-8550U @ 1.8GHz:

    benchmark           old ns/op     new ns/op     delta
    BenchmarkStem-8     1721          230           -86.64%

    benchmark           old allocs     new allocs     delta
    BenchmarkStem-8     0              0              +0.00%

    benchmark           old bytes     new bytes     delta
    BenchmarkStem-8     12            0             -100.00%


Tests
-----

Included `test_output.txt` and `test_voc.txt` are from the referenced original
implementations, used only when running tests with `go test`.


License
-------

2-clause BSD-like (see LICENSE and AUTHORS files).
