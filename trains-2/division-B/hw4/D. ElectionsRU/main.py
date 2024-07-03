parties = []
sumVoices = 0
i=0
with open('input.txt','r') as fin:
    for line in fin:
        words = line.split()
        voices = int(words[-1])
        name = " ".join(words[:-1])
        parties.append([name,voices,i])
        sumVoices +=voices
        i+=1

div = sumVoices/450
free =450
for i in range(len(parties)):
    party = parties[i]
    got = party[1]//div
    party.append(got)
    free -=got
    party.append(party[1]%div)

parties.sort(key=lambda x:x[4],reverse=True)
for i in range(int(free)):
    parties[i][3]+=1

parties.sort(key=lambda x:x[2])
for p in parties:
    print(p[0],int(p[3]))


