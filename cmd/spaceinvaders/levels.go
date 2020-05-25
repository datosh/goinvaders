package main

import (
	"engine/vec2"
	"log"
)

type Levels struct {
}

func (l *Levels) Load(spaceinvaders *Spaceinvaders, level int) {
	switch level {
	case 1:
		loadLevel1(spaceinvaders)
	case 2:
		loadLevel2(spaceinvaders)
	default:
		log.Panicf("There is no level: %v", level)
	}
}

func (l *Levels) LastLevel() int {
	return 2
}

func loadLevel1(spaceinvaders *Spaceinvaders) {
	enemies := spaceinvaders.enemyController.Enemies
	enemies = enemies[:0]

	controller := spaceinvaders.enemyController

	controller.AddEnemy(NewEnemy(&vec2.T{X: 020, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 120, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 220, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 320, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 420, Y: 30}, EnemyOne))

	controller.AddEnemy(NewEnemy(&vec2.T{X: 020, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 120, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 220, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 320, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 420, Y: 150}, EnemyTwo))
}

func loadLevel2(spaceinvaders *Spaceinvaders) {
	enemies := spaceinvaders.enemyController.Enemies
	enemies = enemies[:0]

	controller := spaceinvaders.enemyController

	controller.AddEnemy(NewEnemy(&vec2.T{X: 020, Y: 30}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 120, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 220, Y: 30}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 320, Y: 30}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 420, Y: 30}, EnemyTwo))

	controller.AddEnemy(NewEnemy(&vec2.T{X: 020, Y: 150}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 120, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 220, Y: 150}, EnemyOne))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 320, Y: 150}, EnemyTwo))
	controller.AddEnemy(NewEnemy(&vec2.T{X: 420, Y: 150}, EnemyOne))
}
