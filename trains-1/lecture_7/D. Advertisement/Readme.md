# D. Реклама

Фирма NNN решила транслировать свой рекламный ролик в супермаркете XXX. Однако денег, запланированных на рекламную кампанию, хватило лишь на две трансляции ролика в течение одного рабочего дня.

Фирма NNN собрала информацию о времени прихода и времени ухода каждого покупателя в некоторый день. Менеджер по рекламе предположил, что и на следующий день покупатели будут приходить и уходить ровно в те же моменты времени.

Помогите ему определить моменты времени, когда нужно включить трансляцию рекламных роликов, чтобы как можно большее количество покупателей прослушало ролик целиком от начала до конца хотя бы один раз. Ролик длится ровно 5 единиц времен. Трансляции роликов не должны пересекаться, то есть начало второй трансляции должно быть хотя бы на 5 единиц времени позже, чем начало первой.

Если трансляция ролика включается, например, в момент времени 10, то покупатели, пришедшие в супермаркет в момент времени 10 (или раньше) и уходящие из супермаркета в момент 15 (или позднее) успеют его прослушать целиком, а, например, покупатель, пришедший в момент времени 11, равно как и покупатель, уходящий в момент 14 - не успеют. Если покупатель успевает услышать только конец первой трансляции ролика (не сначала) и начало второй трансляции (не до конца), то считается, что он не услышал объявления. Если покупатель успевает услышать обе трансляции ролика, то при подсчете числа людей, прослушавших ролик, он все равно учитывается всего один раз (фирме важно именно количество различных людей, услышавших ролик).

## Формат ввода

В первой строке входного файла вводится число N - количество покупателей (1 ≤ N ≤ 2000). В следующих N строках записано по паре натуральных чисел - время прихода и время ухода каждого из них. Все значения времени - натуральные числа, не превышающие 109. Время ухода человека из супермаркета всегда строго больше времени его прихода в супермаркет.

## Формат вывода

Выведите через пробел три числа: количество покупателей, которые прослушают ролик целиком от начала до конца хотя бы один раз, и моменты времени, когда должна начинаться трансляция ролика. Моменты времени должны быть выведены в возрастающем порядке и должны быть натуральными числами, не превышающими 2·109. Если вариантов ответа несколько, выведите любой из них.

### Пример 1

Ввод

4
1 11
1 3
6 15
1 6

Вывод

3 1 6

### Пример 2

Ввод

1
1 10

Вывод

1 3 25

### Пример 3

Ввод

3
1 10
11 20
21 30

Вывод

2 1 22

## Примечания

1. Трансляция роликов начинается в моменты времени 1 и 6. Первое объявление успевают прослушать покупатели номер 1 и 4, второе - 1 и 3. Когда бы ни начиналась трансляция объявления, 2-й покупатель не сможет его прослушать, так как находится в супермаркете менее 5 минут. Приведенный ответ является не единственным верным ответом на этот тест.

2. Объявление, трансляция которого начинается в момент 3, единственный покупатель обязательно услышит. Вторую трансляцию (раз она оплачена) мы можем сделать когда угодно, например, в 25 минут в пустом супермаркете (впрочем, мы не можем начать трансляцию второго объявления, например, в момент 7 - т.к. к этому моменту еще не закончится первая трансляция)

3. Объявление услышат лишь 2 из 3-х покупателей.
