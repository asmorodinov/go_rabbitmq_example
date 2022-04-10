package receive

import (
	"container/list"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func FindPath(from, to, lang string, titles bool) (int, []string) {
	if from == to {
		return 0, []string{from}
	}

	visited := make(map[string]bool)
	inQueue := make(map[string]bool)
	nodeQueue := list.New()
	prev := make(map[string]string)

	nodeQueue.PushBack(from)
	inQueue[from] = true

	for nodeQueue.Len() > 0 {
		current := nodeQueue.Front().Value.(string)
		nodeQueue.Remove(nodeQueue.Front())

		if current == to {
			break
		}

		if _, ok := visited[current]; ok {
			continue
		}

		visited[current] = true

		var adjacent []string
		if titles {
			adjacent = getTitles(current, lang)
		} else {
			adjacent = extractLinks(current)
		}

		for _, node := range adjacent {
			if _, ok := visited[node]; ok {
				continue
			}
			if _, ok := inQueue[node]; ok {
				continue
			}

			nodeQueue.PushBack(node)
			prev[node] = current
			inQueue[node] = true
		}
	}

	if _, ok := prev[to]; !ok {
		return -1, nil
	}

	path := make([]string, 0)
	for node := to; node != from; node = prev[node] {
		path = append(path, node)
	}
	path = append(path, from)
	reverse(path)

	return len(path) - 1, path
}
