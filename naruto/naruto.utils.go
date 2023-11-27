package naruto

import "fmt"

func RouteHandlerKeyGen(method string, path string, uniqueId int) string {
	return (method + "_" + path + "_" + fmt.Sprint(uniqueId))
}

func RoutePathFixer(path string) string {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return path
}
