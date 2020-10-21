/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {

	step := 0
	hasEndWord := false
	dict := make(map[string][]string)

	for _, word := range wordList {
		if word == endWord {
			hasEndWord = true
		}

		for i := 0; i < len(endWord); i++ {
			dict[word[:i]+"*"+word[i+1:]] = append(dict[word[:i]+"*"+word[i+1:]], word)
		}
	}

	if !hasEndWord {
		return step
	}

	//双向队列
	queueL, queueR := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	vistedL, vistedR := map[string]bool{beginWord: true}, map[string]bool{endWord: true}

	for len(queueL) > 0 && len(queueR) > 0 {

		step++

		nextQueue := map[string]bool{}

		for word, _ := range queueL {

			delete(queueL, word) //pop

			for i := 0; i < len(word); i++ {
				newWord := word[:i] + "*" + word[i+1:]

				for _, v := range dict[newWord] {
					if _, ok := queueR[v]; ok { //命中
						return step + 1
					}

					if _, ok := vistedL[v]; ok {
						continue
					}

					vistedL[v] = true
					nextQueue[v] = true
				}
			}
		}

		queueL = nextQueue

		//交换队列
		queueL, queueR = queueR, queueL
		vistedL, vistedR = vistedR, vistedL

	}

	return 0

}

// @lc code=end
func ladderLength(beginWord string, endWord string, wordList []string) int {

	hasEndWord := false //判断列表中是否存在endWord
	wordLen := len(beginWord)

	hashMap := make(map[string][]string)

	for _, word := range wordList {

		if word == endWord {
			hasEndWord = true
		}

		for i := 0; i < wordLen; i++ {
			hashMap[word[:i]+"*"+word[i+1:]] = append(hashMap[word[:i]+"*"+word[i+1:]], word)
		}

	}

	if !hasEndWord {
		return 0
	}

	step := 0
	queueL, queueR := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	vistedL, vistedR := map[string]bool{beginWord: true}, map[string]bool{endWord: true}

	for len(queueL) > 0 && len(queueR) > 0 {

		step++
		nextL := map[string]bool{}

		for word, _ := range queueL {

			delete(queueL, word) //pop

			for i := 0; i < wordLen; i++ {

				newWord := word[:i] + "*" + word[i+1:]

				for _, w := range hashMap[newWord] {

					if _, ok := queueR[w]; ok {
						step++
						return step
					}

					if _, ok := vistedL[w]; ok { //已访问
						continue
					}

					vistedL[w] = true
					nextL[w] = true

				}
			}

		}

		queueL = nextL

		step++
		nextR := map[string]bool{}

		for word, _ := range queueR {

			delete(queueR, word) //pop

			for i := 0; i < wordLen; i++ {

				newWord := word[:i] + "*" + word[i+1:]

				for _, w := range hashMap[newWord] {

					if _, ok := queueL[w]; ok {
						step++
						return step
					}

					if _, ok := vistedR[w]; ok { //已访问
						continue
					}

					vistedR[w] = true
					nextR[w] = true

				}
			}

		}

		queueR = nextR

	}

	return 0

}

func bfs(beginWord string, endWord string, wordList []string) int {

	hasEndWord := false //判断列表中是否存在endWord

	//创建hasMap，方便查找
	hashMap := make(map[string]bool, len(wordList))
	for _, word := range wordList {

		hashMap[word] = true //把word放入hashMap

		if word == endWord { //查找endword
			hasEndWord = true
		}

	}

	//不存在endWord，返回0
	if !hasEndWord {
		return 0
	}

	//bfs
	queue := []string{beginWord}
	wordLenth := len(beginWord)

	step := 0

	for len(queue) > 0 {

		step++
		nextQueue := []string{}

		for len(queue) > 0 {

			curr := queue[0]
			queue = queue[1:]

			//查找与head相差一个字母的单词
			for i := 0; i < wordLenth; i++ { //字母不同的位置

				//======================= 遍历对比超时 =======================
				// for word, _ := range hashMap {
				// 	if word[i] == curr[i] { //预期不同的位置相同
				// 		continue
				// 	}
				// 	if word[:i] == curr[:i] && word[i+1:] == curr[i+1:] {
				// 		if word == endWord {
				// 			return step + 1
				// 		}
				// 		delete(hashMap, word)
				// 		nextQueue = append(nextQueue, word)
				// 	}
				// }
				//======================= 优化为替换字母A-Z，查找map =======================
				//======================= 其他方案：hashmap优化二维字典，所有位置依次替换为*作为key，详见双向队列 =======================
				//======================= end =======================

				word := []rune(curr)
				for ch := 'a'; ch <= 'z'; ch++ {
					if ch == word[i] { //变化内容与原文相同
						continue
					}

					word[i] = ch
					temp := string(word)

					if temp == endWord {
						return step + 1
					}

					if _, ok := hashMap[temp]; !ok { //原字典中无变化后的单词
						continue
					}

					delete(hashMap, temp)
					nextQueue = append(nextQueue, temp)
				}
			}

		}

		queue = nextQueue

	}

	return 0

}

