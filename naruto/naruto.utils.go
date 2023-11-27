package naruto

import "fmt"

func RouteHandlerKeyGen(method string, path string, uniqueId int) string {
	return (method + "_" + path + "_" + fmt.Sprint(uniqueId))
}

func RoutePathFixer(path string) string {
	if len(path) == 0 || path == "/" {
		return "base_path"
	}

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return path
}
