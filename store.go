package main

var store = make(map[string]interface{})

func get(key string) interface{} {
	return store[key]
}

func set(key string, value interface{}) interface{} {
	store[key] = value
	return store[key]
}

func patch(key string, json map[string]interface{}) interface{} {
	old := store[key]
	value := merge(old.(map[string]interface{}), json)
	return set(key, value)
}

func del(key string) interface{} {
	delete(store, key)
	return nil
}

func merge(values ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, json := range values {
		for key, value := range json {
			result[key] = value
		}
	}
	return result
}
