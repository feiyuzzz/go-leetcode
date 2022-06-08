package hot100

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	htab := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if htab[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != htab[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		} else {
			stack = append(stack, s[i]) // 进栈
		}
	}
	return len(stack) == 0
}
