package go_utils

func GenerateSubsequences(inputSequence []int) map[string][]int {
	var sequences map[string][]int = make(map[string][]int)
	var recur func(currentSequence []int)
	recur = func(currentSequence []int) {
		if len(currentSequence) == 1 {
			return
		}

		for i := 0; i < len(currentSequence); i++ {
			var newSequence = DupArray(currentSequence)
			newSequence = RemoveArrayElement(newSequence, i)
			hashId := ArrayToString(newSequence)
			if _, ok := sequences[hashId]; ok {
				// already exists; we're done
			} else {
				sequences[hashId] = newSequence
				recur(newSequence)
			}
		}
	}
	hashId := ArrayToString(inputSequence)
	sequences[hashId] = inputSequence
	recur(inputSequence)
	//fmt.Printf("%v\n", sequences)
	return sequences
}

func GetKeys(hash map[string]interface{}) []string {
	keys := make([]string, len(hash))

	var index int
	for key := range hash {
		keys[index] = key
		index++
	}
	return keys
}

func MapToArray(hash map[string]bool) []string {
	arr := make([]string, 0, len(hash))
	for key := range hash {
		arr = append(arr, key)
	}
	return arr
}
