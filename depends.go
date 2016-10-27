package gopkgorder

import "strings"

func (g Graph) GetDependencies(root string) (ret []string) {
	m := make(map[string]bool)
	for from, subMap := range g {
		if strings.HasPrefix(from, root) {
			for name, _ := range subMap {
				if !strings.HasPrefix(name, root) {
					m[name] = true
				}
			}
		}
	}
	for name, _ := range m {
		ret = append(ret, name)
	}
	return
}
