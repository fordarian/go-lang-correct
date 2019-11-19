# go-lang-correct

> Tckb ds xfcnj dcnhtxftnt gjlj,ysq ntrcn nj 'nj inerf lkz dfc!

Пакет пытается определить, не ошибся ли автор текста раскладкой. Если да - можно исправить.

*** Usage: ***
```golang
package main

import (
    "fmt"

    "github.com/fordarian/go-lang-correct"
)

func main() {
    str := "'nj zdyj cnhfiyfz jib,rf!"
	// просто меняем раскладку
	fmt.Println(lang.Correct(str))

	// попробуем определить, насколько вероятно что текст испорчен:
	// если == 0 - неясно
	// если > 0 - скорее всего раскладка не та
	// если < 0 - скорее всего раскладка правильная

	score := lang.ScoreText(str)
	fmt.Println(score) // 4
}

```
