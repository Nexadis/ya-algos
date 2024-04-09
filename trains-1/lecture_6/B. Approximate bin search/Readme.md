# B. Приближенный двоичный поиск

Для каждого из чисел второй последовательности найдите ближайшее к нему в первой.

## Формат ввода

В первой строке входных данных содержатся числа N и K (). Во второй строке задаются N чисел первого массива, отсортированного по неубыванию, а в третьей строке – K чисел второго массива. Каждое число в обоих массивах по модулю не превосходит 2⋅109.

## Формат вывода

Для каждого из K чисел выведите в отдельную строку число из первого массива, наиболее близкое к данному. Если таких несколько, выведите меньшее из них.

### Пример 1

Ввод

5 5
1 3 5 7 9
2 4 8 1 6

Вывод

1
3
7
1
5

### Пример 2

Ввод

6 11
1 1 4 4 8 120
1 2 3 4 5 6 7 8 63 64 65

Вывод

1
1
4
4
4
4
8
8
8
8
120

### Пример 3

Ввод

10 10
-5 1 1 3 5 5 8 12 13 16
0 3 7 -17 23 11 0 11 15 7

Вывод

1
3
8
-5
16
12
1
12
16
8