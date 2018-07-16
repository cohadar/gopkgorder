package order

import "strings"

func (g Graph) GetOrderedPackages() (ret [][]string, err error) {
	for {
		bottom := g.extractBottom()
		if len(bottom) == 0 {
			if len(g) == 0 {
				break
			} else {
				panic("graph not acyclic")
			}
		}
		if out := g.removeBottom(bottom); len(out) > 0 {
			ret = append(ret, out)
		}
	}
	return
}

func (g Graph) extractBottom() (ret []string) {
	for name, subMap := range g {
		if len(subMap) == 0 {
			ret = append(ret, name)
		}
	}
	return ret
}

func (g Graph) removeBottom(bottom []string) (ret []string) {
	for _, name := range bottom {
		delete(g, name)
	}
	for _, subMap := range g {
		for _, name := range bottom {
			if _, ok := subMap[name]; ok {
				delete(subMap, name)
			}
		}
	}
	for _, root := range roots(bottom) {
		if !g.hasRoot(root) {
			// exlude root packages like 'github.com'
			if !strings.ContainsRune(root, '.') {
				ret = append(ret, root)
			}
		}
	}
	return ret
}

func roots(bottom []string) (ret []string) {
	m := make(map[string]bool)
	for _, b := range bottom {
		m[extractRoot(b)] = true
	}
	for root, _ := range m {
		ret = append(ret, root)
	}
	return
}

func (g Graph) hasRoot(root string) bool {
	for nameA, subMap := range g {
		if root == extractRoot(nameA) {
			return true
		}
		for nameB, _ := range subMap {
			if root == extractRoot(nameB) {
				return true
			}
		}
	}
	return false
}

func extractRoot(path string) string {
	return strings.FieldsFunc(path, func(x rune) bool { return x == '/' })[0]
}
