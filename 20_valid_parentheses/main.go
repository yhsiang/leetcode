package leetcode

var maps = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

type Stack struct {
	Values []string
}

func (s *Stack) Pop() string {
	a := s.Values[len(s.Values)-1]
	s.Values = s.Values[0 : len(s.Values)-1]
	return a
}

func (s *Stack) Push(str string) {
	s.Values = append(s.Values, str)
}

func (s *Stack) isEmpty() bool {
	return len(s.Values) == 0
}

func isValid(s string) bool {
	var stack Stack

	for i := 0; i < len(s); i++ {
		var char = string(s[i])
		v, ok := maps[char]
		if ok {
			var top string
			if stack.isEmpty() {
				top = "#"
			} else {
				top = stack.Pop()
			}
			if top != v {
				return false
			}
		} else {
			stack.Push(char)
		}
	}

	return stack.isEmpty()
}
