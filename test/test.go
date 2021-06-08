package test

import "test/app/model"

type TestTask struct {
	Value  []string
	Result model.Task
}

func TestList() []TestTask {
	list := []TestTask{testTask1(), testTask2(), testTask3()}
	return list
}

func testTask1() TestTask {
	value := []string{"Prueba", "Prueba"}
	mapa := make(map[string]int)
	mapa["_"] = 1
	mapa["a"] = 2
	mapa["b"] = 2
	mapa["e"] = 2
	mapa["p"] = 2
	mapa["r"] = 2
	mapa["u"] = 2
	result := model.Task{Sentence: "prueba_prueba", Values: mapa}
	return TestTask{Value: value, Result: result}
}

func testTask2() TestTask {
	value := []string{"La", "Prueba", "Para", "El", "Test"}
	mapa := make(map[string]int)
	mapa["_"] = 4
	mapa["a"] = 4
	mapa["b"] = 1
	mapa["e"] = 3
	mapa["l"] = 2
	mapa["p"] = 2
	mapa["r"] = 2
	mapa["s"] = 1
	mapa["t"] = 2
	mapa["u"] = 1
	result := model.Task{Sentence: "la_prueba_para_el_test", Values: mapa}
	return TestTask{Value: value, Result: result}
}

func testTask3() TestTask {
	value := []string{"Nos", "preguntamos,", "¿Quién", "soy", "yo", "para", "ser", "brillante,", "hermoso,", "talentoso,", "fabuloso?", "En", "realidad,", "¿quién", "eres", "tu", "para", "no", "serlo?"}
	mapa := make(map[string]int)
	mapa[","] = 5
	mapa["?"] = 2
	mapa["_"] = 18
	mapa["a"] = 10
	mapa["b"] = 2
	mapa["d"] = 2
	mapa["e"] = 10
	mapa["f"] = 1
	mapa["g"] = 1
	mapa["h"] = 1
	mapa["i"] = 4
	mapa["l"] = 6
	mapa["m"] = 2
	mapa["n"] = 8
	mapa["o"] = 12
	mapa["p"] = 3
	mapa["q"] = 2
	mapa["r"] = 9
	mapa["s"] = 9
	mapa["t"] = 5
	mapa["u"] = 5
	mapa["y"] = 2
	mapa["¿"] = 2
	mapa["é"] = 2
	result := model.Task{Sentence: "nos_preguntamos,_¿quién_soy_yo_para_ser_brillante,_hermoso,_talentoso,_fabuloso?_en_realidad,_¿quién_eres_tu_para_no_serlo?", Values: mapa}
	return TestTask{Value: value, Result: result}
}
