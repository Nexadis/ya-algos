n = int(input())
all = [str(i) for i in range(1,n+1)]
sum = set(all)
while True:
    seq = input()
    if seq == "HELP":
        break
    ans = input()
    if ans == "YES":
        sum.intersection_update(seq.split())
    elif ans == "NO":
        sum.difference_update(seq.split())

print(*sorted(list(sum)))
