line1 = set(map(int, input().split()))
line2 = set(map(int, input().split()))
intersect = list(line1.intersection(line2))
intersect.sort()
print(*intersect)