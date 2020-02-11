package open_the_lock

// No.752

//https://leetcode-cn.com/problems/open-the-lock/solution/da-kai-zhuan-pan-suo-shi-xian-yu-zheng-li-by-eiger/

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, v := range deadends {dead[v] = true}		// 填充dead set
	if dead["0000"] {return -1}						// 直接死锁
	if target == "0000" {return 0}					// 出发即是终点，特殊

	// BFS --------------------------------------------------------------
	queue := make([]string, 0)						// 构造处理字符串队列
	queue = append(queue, "0000")			// 起点
	visited := make(map[string]int)					// 已访问过的集合，并记录到达该状态花的步数
	visited["0000"] = 0

	//depth := 0
	for len(queue) != 0 {
		cur := queue[0]		// 取出当前待处理的锁状态（或者说无向图的节点）
		queue = queue[1:]	// 出队
		curSlice := []rune(cur)		// 转为切片

		// 获取当前状态下一步的所有（8个）可能状态
		var nexts []string	// 从cur状态出发，可以到达的其他状态.当前状态的下一步总共有4*2 = 8种可能
		for i:=0; i<4; i++ {	// 对每一位进行变动。
			origin := curSlice[i]	// 备份下原始的字符
			// 正向转动转盘
			curSlice[i] = (curSlice[i] - '0' + 1) % 10 + '0'		// '0'~'9'的字符减去'0' 变为整型，来和1作加减，外边再 + '0'又转为字符
			nexts = append(nexts, string(curSlice))
			curSlice[i] = origin		// 恢复原始状态
			// 反向转动转盘
			curSlice[i] = (curSlice[i] - '0' + 9) % 10 + '0'		// 如果是-1会出现负数情况，不好处理。循环左移的技巧
			nexts = append(nexts, string(curSlice))
			curSlice[i] = origin
		}

		// 遍历下一步的所有可能状态
		for _, next := range nexts {
			if _, ok := visited[next]; !ok && !dead[next] {		// 没有访问过，也不是dead
				queue = append(queue, next)						// 入队
				visited[next] = visited[cur] + 1				// 对应到达这个next需要多少步
				if next == target {return visited[next]}		// 如果到达目标，就返回最少需要的步数
			}
		}
	}

	return -1
}





//func openLock(deadends []string, target string) int {
//	deadendsMap := make(map[string]int)
//	for _, v := range deadends {
//		deadendsMap[v] = 1
//	}
//	searchlist := list.New()
//	searchlist.PushBack("BEGIN")
//	if deadendsMap["0000"] == 1 {
//		return -1
//	}
//	if strings.EqualFold(target, "0000") {
//		return 0
//	} else {
//		searchlist.PushBack("0000")
//	}
//	round := 0
//	usedMap := make(map[string]int)
//
//	for searchlist.Len() > 0 {
//		if searchlist.Front().Value=="BEGIN" {
//			round++
//			searchlist.Remove(searchlist.Front())
//		}
//		lockNum := searchlist.Front().Value.(string)
//		fmt.Println(lockNum)
//		searchlist.Remove(searchlist.Front())
//		for i := 0; i < 4; i++ {
//			for j:=-1;j<=1;j+=2  {
//				lNBytes := []rune(lockNum)
//				t:=lNBytes[i]
//				lNBytes[i]=(t-'0') + j + 10) % 10
//				if string(lNBytes) == target {
//					return round
//				} else if deadendsMap[string(lNBytes)] != 1 && usedMap[string(lNBytes)] != 1 {
//					usedMap[string(lNBytes)] = 1
//					searchlist.PushBack(string(lNBytes))
//					fmt.Printf("[lNBytes]%s,[Round]%d, [queue Len]%d \n", string(lNBytes),round,searchlist.Len())
//				}
//			}
//
//		}
//		searchlist.PushBack("BEGIN")
//	}
//	return -1
//}
