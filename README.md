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

output

```
Probability of a bad word in a string is: 0.005944913738782739
```

So the calculated probability of a bad/rude word in a id string is `0.59%`

Running a Go program with the word list shows that 

```
Found rude words: 4520, total: 100000
```

It's about 4.5%, reality is a b*tch.

With two of my favourite rude words here is the output:

```
uuid: 2d700ea9a5d64826b69bd109b0536627, hash: 1ZG8T5MBLCUNT, contains rude word: C*NT
uuid: 249092de939c4d2a9c6a63eaf9d03ff2, hash: 19WECUNTKY7B, contains rude word: C*NT
uuid: 651e46f131864b9e8470bcaddcd9342f, hash: MAV0USCUNTWP, contains rude word: C*NT
uuid: 1a818c8c7e8842c087b856b095735678, hash: 2MVY1ULCUNTI2, contains rude word: C*NT
uuid: bae9854c08d94030972877f1c3e4a05e, hash: QHCUNT2IZ5SO, contains rude word: C*NT
uuid: 83a61dc1a0f9475387a957d55f50699e, hash: Q8R10TFUCKHA, contains rude word: F*CK
uuid: d90ef1e5469243fe91a38888bf603221, hash: 35I93NFUCKBIJ, contains rude word: F*CK
uuid: 3042deb780674d8d8400728713ee99af, hash: 3FPRLSFUCKKTD, contains rude word: F*CK
uuid: 65d679004c71494ca3fff1db709bff90, hash: 7P3FUCKR0QR8, contains rude word: F*CK
uuid: b3085cf1e8794f4ea8bc57a6a050becc, hash: 1FUCK7ZIJW7WZ, contains rude word: F*CK
uuid: 0ab6ae1e1f2641d380975b2d4d06fddf, hash: 2M3FUCK3BYABZ, contains rude word: F*CK
uuid: 8b38d62fc8bc43d690db84fbe105b620, hash: 30TAA120FUCKI, contains rude word: F*CK
uuid: d42be752ef3749c3ab29398d64622c75, hash: 2VJLBHFUCKKTJ, contains rude word: F*CK
uuid: 5d49d81c4ffd4ffd809c789d1f3f414c, hash: L5R7FUCK2XFO, contains rude word: F*CK
```