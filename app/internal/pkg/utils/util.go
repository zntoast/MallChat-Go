package utils

// ExtractFieldToMap 从元素切片中提取键值对并构建映射
//   - elements: 需要处理的元素切片
//   - keyValueExtractor: 从单个元素中提取键值对的函数
func ExtractFieldToMap[Element any, Key comparable, Value any](
	elements []Element,
	keyValueExtractor func(Element) (Key, Value),
) map[Key]Value {
	result := make(map[Key]Value, len(elements))
	for _, element := range elements {
		key, value := keyValueExtractor(element)
		result[key] = value
	}
	return result
}

// ExtractToSlice 从元素集合中提取特定值构建新切片
//   - elements: 源数据集合，可以是任意类型的切片
//   - extractor: 值提取函数，接收元素返回要提取的值
func ExtractToSlice[Element any, Extracted any](
	elements []Element,
	extractor func(Element) Extracted,
) []Extracted {
	result := make([]Extracted, len(elements))
	for i, element := range elements {
		result[i] = extractor(element)
	}
	return result
}

// Filter 根据条件筛选集合元素
func Filter[Element any](
	elements []Element,
	predicate func(Element) bool,
) []Element {
	if len(elements) == 0 {
		return nil
	}
	result := make([]Element, 0, len(elements))
	for _, element := range elements {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result[:len(result):len(result)]
}

// Unique 去除重复元素
func Unique[E comparable](elements []E) []E {
	seen := make(map[E]struct{}, len(elements))
	result := make([]E, 0, len(elements))
	for _, e := range elements {
		if _, exists := seen[e]; !exists {
			seen[e] = struct{}{}
			result = append(result, e)
		}
	}
	return result
}

// UniqueBy 根据指定规则去重切片元素，保留首个出现的元素
// param ：
//   - elements: 待去重的元素切片，允许为空
//   - keyExtractor: 从元素中提取唯一标识键的函数
func UniqueBy[Element any, Key comparable](
	elements []Element,
	keyExtractor func(Element) Key,
) []Element {
	if len(elements) == 0 {
		return nil
	}
	seen := make(map[Key]struct{}, len(elements))
	// 预分配结果切片（按最坏情况预分配）
	result := make([]Element, 0, len(elements))

	for _, elem := range elements {
		key := keyExtractor(elem)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, elem)
		}
	}
	return result[:len(result):len(result)]
}

// 辅助函数用于计算最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
