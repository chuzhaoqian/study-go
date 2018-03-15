package maptools

func Merge(map1, map2 map[interface{}]interface{}) map[interface{}]interface{} {

	for key, views := range map2 {
		map1[key] = views
	}

	return map1
}
