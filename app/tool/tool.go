package tool

import (
	"strings"
	"test/app/model"
)

func CreateTaskA(list []string) (task model.Task) {

	newSentence := strings.ToLower(string(list[0]))
	for _, w := range list[1:] {
		lower := strings.ToLower(string(w))
		newSentence = newSentence + "_" + lower
	}
	task.Sentence = newSentence

	task.Values = make(map[string]int)
	for _, l := range task.Sentence {
		if task.Values[string(l)] == 0 {
			task.Values[string(l)] = 1
		} else {
			task.Values[string(l)]++
		}
	}

	return
}

func CreateTaskB(list []string) (task model.Task) {

	sentence := strings.Join(list, "_")

	task.Values = make(map[string]int)
	for _, l := range sentence {
		l := strings.ToLower(string(l))
		task.Sentence += l
		if task.Values[string(l)] == 0 {
			task.Values[string(l)] = 1
		} else {
			task.Values[string(l)]++
		}
	}

	return
}

func CreateTaskC(list []string) (task model.Task) {

	task.Sentence = strings.Join(list, "_")
	task.Sentence = strings.ToLower(task.Sentence)

	task.Values = make(map[string]int)
	for _, l := range task.Sentence {
		if task.Values[string(l)] == 0 {
			task.Values[string(l)] = 1
		} else {
			task.Values[string(l)]++
		}
	}

	return
}
