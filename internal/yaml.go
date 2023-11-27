package internal

import (
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Yaml map[string]any

func Unmarshal(data []byte) (Yaml, error) {
	_yaml := make(Yaml)
	err := yaml.Unmarshal(data, &_yaml)
	return _yaml, err
}

func GetIndex(key string, callback func(index int)) bool {
	if strings.HasPrefix(key, "{") {
		i := strings.Trim(strings.TrimRight(strings.TrimLeft(key, "{"), "}"), " ")
		pos, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		callback(pos)
		return true
	}
	return false
}

func (mapper Yaml) Get(key string) any {
	segments := strings.Split(key, ".")
	var ref any
	for index, segment := range segments {
		if GetIndex(segment, func(index int) { ref = ref.([]any)[index] }) {
			continue
		}
		if index == 0 {
			ref = mapper[segment].(Yaml)
			continue
		}
		ref = ref.(Yaml)[segment]
	}
	return ref
}

func (mapper *Yaml) Set(key string, value any) {
	segments := strings.Split(key, ".")
	ref := (*mapper)[segments[0]]
	for i := 1; i < len(segments)-1; i++ {
		segment := segments[i]
		if GetIndex(segment, func(index int) { ref = ref.([]any)[index] }) {
			continue
		}
		ref = ref.(Yaml)[segment]
	}
	segment := segments[len(segments)-1]
	if GetIndex(segment, func(index int) { ref.([]any)[index] = value }) {
		return
	}
	ref.(Yaml)[segments[len(segments)-1]] = value
}

func (mapper *Yaml) Marshall() ([]byte, error) {
	return yaml.Marshal(mapper)
}
