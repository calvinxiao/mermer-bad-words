import sys
from collections import Counter

print("{")
list = [word.strip() for word in sys.stdin]
for word in list:
	print("\""+word+"\",")
print("}")
c = Counter((len(word) for word in list))

base = 36
strLen = 13

totalNumberOfHashStrings = base ** strLen

total = 0
for wordLen in c:
	if wordLen == 0 : continue
	if wordLen > strLen : continue
	subs = [0 for i in range(strLen + 1)]
	for subLen in range(wordLen, strLen + 1):
		# bad word is at the end of a string with len = subLen
		subs[subLen] = (base ** (subLen - wordLen)) * c[wordLen]
		excludedStartIndex = subLen - wordLen
		subs[subLen] -= sum(subs[excludedStartIndex:subLen - 1])
	# print(wordLen)
	# print(subs)
	total += 1.0 * sum(subs) 	
print("Total number of occurences of at least a bad word in a 13 string is:", total)

print("Probability of a bad word in a string is:", 1.0 * total / totalNumberOfHashStrings)