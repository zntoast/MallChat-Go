package utils

import (
	"reflect"
	"testing"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func TestExtractFieldToMap(t *testing.T) {
	people := []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	tests := []struct {
		name      string
		elements  []Person
		extractor func(Person) (int, string)
		want      map[int]string
	}{
		{
			name:     "happy path",
			elements: people,
			extractor: func(p Person) (int, string) {
				return p.ID, p.Name
			},
			want: map[int]string{
				1: "Alice",
				2: "Bob",
				3: "Charlie",
			},
		},
		{
			name:     "empty slice",
			elements: []Person{},
			extractor: func(p Person) (int, string) {
				return p.ID, p.Name
			},
			want: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractFieldToMap(tt.elements, tt.extractor)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractFieldToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractToSlice(t *testing.T) {
	people := []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	tests := []struct {
		name      string
		elements  []Person
		extractor func(Person) string
		want      []string
	}{
		{
			name:     "happy path",
			elements: people,
			extractor: func(p Person) string {
				return p.Name
			},
			want: []string{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "empty slice",
			elements: []Person{},
			extractor: func(p Person) string {
				return p.Name
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractToSlice(tt.elements, tt.extractor)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	people := []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	tests := []struct {
		name      string
		elements  []Person
		predicate func(Person) bool
		want      []Person
	}{
		{
			name:     "happy path",
			elements: people,
			predicate: func(p Person) bool {
				return p.Age > 28
			},
			want: []Person{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Charlie", Age: 35},
			},
		},
		{
			name:     "no match",
			elements: people,
			predicate: func(p Person) bool {
				return p.Age < 20
			},
			want: []Person{},
		},
		{
			name:     "empty slice",
			elements: []Person{},
			predicate: func(p Person) bool {
				return p.Age > 28
			},
			want: []Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.elements, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     []int
	}{
		{
			name:     "happy path",
			elements: []int{1, 2, 3, 2, 1, 4},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "all unique",
			elements: []int{1, 2, 3, 4, 5},
			want:     []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			elements: []int{1, 1, 1, 1},
			want:     []int{1},
		},
		{
			name:     "empty slice",
			elements: []int{},
			want:     []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.elements)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

type person struct {
	ID   int
	Name string
}

type complexKey struct {
	Category string
	Code     int
}

func TestUniqueBy(t *testing.T) {
	tests := []struct {
		name         string
		input        []person
		keyExtractor func(person) interface{}
		want         []person
	}{
		{
			name:         "空切片测试",
			input:        []person{},
			keyExtractor: func(p person) interface{} { return p.ID },
			want:         nil,
		},
		{
			name: "全唯一元素",
			input: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			keyExtractor: func(p person) interface{} { return p.ID },
			want: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
		},
		{
			name: "完全重复元素",
			input: []person{
				{ID: 1, Name: "Alice"},
				{ID: 1, Name: "Bob"},
				{ID: 1, Name: "Charlie"},
			},
			keyExtractor: func(p person) interface{} { return p.ID },
			want: []person{
				{ID: 1, Name: "Alice"},
			},
		},
		{
			name: "混合重复情况",
			input: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 1, Name: "Alice"}, // ID重复
				{ID: 3, Name: "Charlie"},
				{ID: 2, Name: "Bobson"}, // ID重复
			},
			keyExtractor: func(p person) interface{} { return p.ID },
			want: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
		},
		{
			name: "复合键测试",
			input: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 1, Name: "Alice"},  // 完全重复
				{ID: 1, Name: "Alicia"}, // ID相同但名称不同
			},
			keyExtractor: func(p person) interface{} {
				return complexKey{
					Category: p.Name[:3],
					Code:     p.ID,
				}
			},
			want: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 1, Name: "Alicia"},
			},
		},
		{
			name: "零值处理",
			input: []person{
				{ID: 0, Name: ""},     // 零值键
				{ID: 0, Name: "John"}, // 重复零值键
				{ID: 1, Name: ""},     // 不同键
			},
			keyExtractor: func(p person) interface{} { return p.ID },
			want: []person{
				{ID: 0, Name: ""},
				{ID: 1, Name: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueBy(tt.input, tt.keyExtractor)

			// 深度比较结果
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueBy() = %v, want %v", got, tt.want)
			}

			// 额外检查切片容量是否紧缩
			if cap(got) != len(got) && len(got) > 0 {
				t.Errorf("切片容量未紧缩，实际容量 %d，期望容量 %d", cap(got), len(got))
			}

			// 检查元素顺序
			if len(got) != len(tt.want) {
				return // 前面已经比较过内容
			}
			for i := range got {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("元素顺序不一致，位置 %d: got %v, want %v",
						i, got[i], tt.want[i])
				}
			}
		})
	}
}
