# K. Медиана объединения-2

Дано N упорядоченных по неубыванию последовательностей целых чисел (т.е. каждый следующий элемент больше либо равен предыдущему), в каждой из последовательностей ровно L элементов. Для каждых двух последовательностей выполняют следующую операцию: объединяют их элементы (в объединенной последовательности каждое число будет идти столько раз, сколько раз оно встречалось суммарно в объединяемых последовательностях), упорядочивают их по неубыванию и смотрят, какой элемент в этой последовательности из 2L элементов окажется на месте номер L (этот элемент называют левой медианой).

Напишите программу, которая для каждой пары последовательностей выведет левую медиану их объединения.

## Формат ввода

Сначала вводятся числа N и L (2 ≤ N ≤ 200, 1 ≤ L ≤ 50000). В следующих N строках задаются параметры, определяющие последовательности.

Каждая последовательность определяется пятью целочисленными параметрами: x1, d1, a, c, m. Элементы последовательности вычисляются по следующим формулам: x1 нам задано, а для всех i от 2 до L: xi = xi–1+di–1. Последовательность di определяется следующим образом: d1 нам задано, а для i ≥ 2 di = ((a*di–1+c) mod m), где mod – операция получения остатка от деления (a*di–1+c) на m.

Для всех последовательностей выполнены следующие ограничения: 1 ≤ m ≤ 40000, 0 ≤ a < m, 0≤c<m, 0≤d1<m. Гарантируется, что все члены всех последовательностей по модулю не превышают 109.

## Формат вывода

В первой строке выведите медиану объединения 1-й и 2-й последовательностей, во второй строке — объединения 1-й и 3-й, и так далее, в (N‑1)-ой строке — объединения 1-й и N-ой последовательностей, далее медиану объединения 2-й и 3-й, 2-й и 4-й, и т.д. до 2-й и N-ой, затем 3-й и 4-й и так далее. В последней строке должна быть выведена медиана объединения (N–1)-й и N-ой последовательностей.

### Пример

Ввод

3 6
1 3 1 0 5
0 2 1 1 100
1 6 8 5 11

Вывод

7
10
9
