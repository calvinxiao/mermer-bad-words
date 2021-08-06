# murmur-bad-words
Calculate bad word probability in base36 with murmur hash

## Analysis

Len of murmur hash of uuid is normally 13, using base36, total number is:

```
>>> 36 ** 13
170581728179578208256L
```

about `2^46`, a lot

So how many times can a bad word in English happend in a hash

> It's very mathematical/algorithmic, but I will try...

Let's say a bad word has len=3, for example `ass`, we need to know how many times this bad word
happend to be at the end of a string which len is `3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13`

- len(str) = 3, answer is 1
- len(str) = 4, str can be `[0-9a-z]ass`, answer is 36, 
- len(str) = 5, str can be `[0-9a-z]ass`, answer is 36 * 36

You can see the pattern, for a bad word has len=A, a string has len=B where B >= A, in base36 there are 
`36 ** (B - A)` ways that string ends with this bad word.

so 

- len(str) = 10, answer is 36 ** 7
- len(str) = 11, answer is 36 ** 8
- len(str) = 12, answer is 36 ** 9
- len(str) = 13, answer is 36 ** 10

Also, if `w` is at the end of `s[:N]`, then `w` is in `s`, I don't know how to probably prove it,
but it looks very close to correct.

But there is one case to consider, if `w` is at the end of str with len=13, we need to minus the count
where `w` is at the end of string with len=12 and len=11. However it depends on `w`, if `w` is like `"aaa"`, 
We don't need to minus those, but normally a bad word is a normal English word.

Consider a bad word list [https://www.cs.cmu.edu/~biglou/resources/bad-words.txt](https://www.cs.cmu.edu/~biglou/resources/bad-words.txt)

```
python3 count.py < badwords.txt
```