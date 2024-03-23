n, m = map(int, input().split())
first = set()
second = set()
for i in range(n):
    first.add(int(input()))
for i in range(m):
    second.add(int(input()))

intersect = [*first.intersection(second)]
intersect.sort()
left = [*first.difference(second)]
left.sort()
right = [*second.difference(first)]
right.sort()
print(len(intersect))
print(*intersect)
print(len(left))
print(*left)
print(len(right))
print(*right)
