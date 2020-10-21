/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {

	root := &TrieNode{}

	//构建字典树
	for _, word := range words {
		node := root

		for _, s := range word {

			if node.children[s-'a'] == nil {
				node.children[s-'a'] = &TrieNode{}
			}
			node = node.children[s-'a']

		}

		node.word = word //***存储完整单词，在使用后置为“”，方便去重复
	}

	result := make([]string, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, root, i, j, &result)
		}
	}

	return result
}

func dfs(board [][]byte, node *TrieNode, i, j int, result *[]string) {

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	target := board[i][j]

	if target == '#' { //已经被使用
		return
	}

	if node.children[target-'a'] == nil { //字典树中没有目标
		return
	}

	node = node.children[target-'a']

	//防止重复
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '#' //标记使用

	dfs(board, node, i-1, j, result) //上
	dfs(board, node, i+1, j, result) //下
	dfs(board, node, i, j-1, result) //左
	dfs(board, node, i, j+1, result) //右

	board[i][j] = target

}

// @lc code=end

