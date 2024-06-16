# breaketernity.go

A Go numerical package to represent numbers as large as 10^^1e308 and as 'small' as 10^-(10^^1e308). Based on Patashu's break_eternity.js ( https://github.com/Patashu/break_eternity.js ).

Now with arbitrary real height and base handling in your favourite hyper 4 operators (tetrate, iterated exponentiation, iterated logarithm, super logarithm, super root) and even in pentate (if you want to do that for some reason)! Using an analytic approximation for bases <= 10, but linear approximation for bases > 10 (but there's options to use the linear approximation everywhere if you need consistent behavior).

The internal representation is as follows: `DFC(sign, layer, mag)` === `sign*10^10^10^ ... (layer times) mag`. So a layer 0 number is just `sign*mag`, a layer 1 number is `sign*10^mag`, a layer 2 number is `sign*10^10^mag`, and so on. If `layer > 0` and `mag < 0`, then the number's exponent is negative, e.g. `sign*10^-10^10^10^ ... mag`.

- sign is -1, 0 or 1.
- layer is a non-negative integer.
- mag is a Number, normalized as follows: if it is above 9e15, log10(mag) it and increment layer. If it is below log10(9e15) (about 15.954) and layer > 0, Math.pow(10, mag) it and decrement layer. At layer 0, sign is extracted from negative mags. Zeroes (`d.sign === 0 || (d.mag === 0 && d.layer === 0)`) become `0, 0, 0` in all fields. Any infinities have both mag and layer as positive Infinity.

Create a Decimal with `D(Decimal or String or Any numeric type)` or with `DFC(sign, layer, mag)`. Then use operations to manipulate the values.

Functions you can call include `abs, neg, round, floor, ceil, trunc, add, sub, mul, div, recip, mod, cmp, cmpabs, max, min, maxabs, minabs, log, log10, ln, pow, root, factorial, gamma, exp, sqrt, tetrate, iteratedexp, iteratedlog, layeradd10, layeradd, slog, ssqrt, lambertw, linear_sroot, pentate` and more! Numeric operators like `+` and `*` do not work - you need to call the equivalent functions instead. Note that all these functions return a pointer to a new Decimal - they do not mutate the original Decimal.

Most functions are available both as a method on Decimal or as an exported function. However, due to Go's limitations on generics, exported functions support DecimalSource as parameters whilst methods require Decimal inputs, example:

```go
var a = Add(238, "1e3")
var b = D(238).Add(D("1e3"))
a.Eq(b) // true
```

Accepted string formats for D() (X, Y, and N represent numbers):

```
M === M
eX === 10^X
MeX === M*10^X
eXeY === 10^(XeY)
MeXeY === M*10^(XeY)
eeX === 10^10^X
eeXeY === 10^10^(XeY)
eeeX === 10^10^10^X
eeeXeY === 10^10^10^(XeY)
eeee... (N es) X === 10^10^10^ ... (N 10^s) X
(e^N)X === 10^10^10^ ... (N 10^s) X
NpX === 10^10^10^ ... (N 10^s) X
FN === 10^10^10^ ... (N 10^s)
XFN === 10^10^10^ ... (N 10^s) X
X^Y === X^Y
X^^N === X^X^X^ ... (N X^s) 1
X^^N;Y === X^X^X^ ... (N X^s) Y
X^^^N === X^^X^^X^^ ... (N X^^s) 1
X^^^N;Y === X^^X^^X^^ ... (N X^^s) Y
```

# Use

The library exports the struct type Decimal, constructed with D() and DFC(), as well as all the operations as both methods and standalone functions.


```go
x = D(123.4567);
y = D("123456.7e-3");
z = D(x);
x.Eq(y) && y.Eq(z) && x.Eq(z) && Eq(x, y); // true
```

Methods that return *Decimal can be chained.

```go
x.Divide(y).Add(z).Multiply(9).Floor();
x.Multiply(D("1.23456780123456789e+9")).Add(D(9876.5432321)).dividedBy(D("4444562598.111772")).Ceil();
```

A list of functions is provided earlier in this readme, or you can read through math.go for a more detailed list.

# Also check out:

- https://github.com/Patashu/break_eternity.js/ break_eternity.js, the JavaScript library this Go package is based on
- https://github.com/Pannoniae/BreakEternity.cs/ BreakEternity.cs, a C# port, for use in Unity or other C# environments

---
SEO: number library, big number, big num, bignumber, bignum, big integer, biginteger, bigint, incremental games, idle games, large numbers, huge numbers, go decimal, go big number


