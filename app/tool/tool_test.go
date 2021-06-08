package tool_test

import (
	"test/app/tool"
	"test/test"
	"testing"
)

func TestCreateTaskA(t *testing.T) {
	for _, v := range test.TestList() {
		result := tool.CreateTaskA(v.Value)
		if result.Sentence != v.Result.Sentence {
			t.Errorf("Expected: %v Got: %v", v.Result.Sentence, result.Sentence)
		}
		for key, value := range v.Result.Values {
			if result.Values[key] != value {
				t.Errorf("Expected: %v Got: %v", value, result.Values[key])
			}
		}
	}
}

func TestCreateTaskB(t *testing.T) {
	for _, v := range test.TestList() {
		result := tool.CreateTaskB(v.Value)
		if result.Sentence != v.Result.Sentence {
			t.Errorf("Expected: %v Got: %v", v.Result.Sentence, result.Sentence)
		}
		for key, value := range v.Result.Values {
			if result.Values[key] != value {
				t.Errorf("Expected: %v Got: %v", value, result.Values[key])
			}
		}
	}
}

func TestCreateTaskC(t *testing.T) {
	for _, v := range test.TestList() {
		result := tool.CreateTaskC(v.Value)
		if result.Sentence != v.Result.Sentence {
			t.Errorf("Expected: %v Got: %v", v.Result.Sentence, result.Sentence)
		}
		for key, value := range v.Result.Values {
			if result.Values[key] != value {
				t.Errorf("Expected: %v Got: %v", value, result.Values[key])
			}
		}
	}
}

func BenchmarkCreateTaskA(b *testing.B) {
	for _, v := range test.TestList() {
		for i := 0; i < b.N; i++ {
			tool.CreateTaskA(v.Value)
		}
	}
}

func BenchmarkCreateTaskB(b *testing.B) {
	for _, v := range test.TestList() {
		for i := 0; i < b.N; i++ {
			tool.CreateTaskB(v.Value)
		}
	}
}

func BenchmarkCreateTaskC(b *testing.B) {
	for _, v := range test.TestList() {
		for i := 0; i < b.N; i++ {
			tool.CreateTaskC(v.Value)
		}
	}
}
