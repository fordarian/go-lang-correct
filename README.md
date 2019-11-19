# go-lang-correct

> Tckb ds xfcnj dcnhtxftnt gjlj,ysq ntrcn nj 'nj inerf lkz dfc!

Пакет пытается определить, не ошибся ли автор текста раскладкой. Если да - можно исправить.

**Usage:**

> go get github.com/fordarian/go-lang-correct

```golang
package main

import (
    "fmt"

    "github.com/fordarian/go-lang-correct"
)

func main() {
    str := "'nj zdyj cnhfiyfz jib,rf!"
    // просто меняем раскладку
    fmt.Println(lang.Correct(str)) // это явно страшная ошибка!

    // попробуем определить, насколько вероятно что текст испорчен:
    // если == 0 - неясно
    // если > 0 - скорее всего раскладка не та
    // если < 0 - скорее всего раскладка правильная

    score := lang.ScoreText(str)
    fmt.Println(score) // 4
}

```

**Как это работает:**

На оценку строки функцией `ScoreText` влияют:
1) Явное совпадение слов с некоторыми шаблонами
2) Наличие\отсутствие типичных окончаний, которые "странно" смотрятся в другой раскладке
3) Количество идущих подряд гласных\согласных букв и знаков препинания (не всех, только тех что на англ. раскладке под русскими символами, вроде Ю, Б и т.д.). 
4) Начинаются ли слова "нетипичных" букв, например "ъ"
5) Содержит ли строка ссылки (http*)

Наиболее явные признаки дают +-10 к счету, это почти полная гарантия что определено корректно.
Средние признаки +-5 и незначительные +-2 балла. Строка разбивается на отдельные слова (разделитель - пробел), возвращается средний рейтинг из всех слов. 

Перед обработкой из строки удаляются все повторения символов более двух раз подряд. Для определения это не сильно важно, но зато убирает ложные срабатывания на фразах вроде "даааа нууууу" или "ааааааааа!!".

**Лицензия**

MIT
