package api

import (
	"github.com/gin-gonic/gin/binding"
	"reflect"
	"strings"
	"sync"
)

const (
	_ uint8 = iota
	json
	xml
	yaml
	form
	query
)

var cache = &bindingCache{}

// bindingCache 结构体用于缓存绑定信息
type bindingCache struct {
	cache map[string][]uint8
	mux   sync.Mutex
}

// GetBinding 方法用于获取绑定信息
func (e *bindingCache) GetBinding(d interface{}) []binding.Binding {
	bs := e.getBinding(reflect.TypeOf(d).String())
	if bs == nil {
		bs = e.resolve(d)
	}
	gbs := make([]binding.Binding, 0)
	mp := make(map[uint8]binding.Binding)
	for _, b := range bs {
		switch b {
		case json:
			mp[json] = binding.JSON
		case xml:
			mp[xml] = binding.XML
		case yaml:
			mp[yaml] = binding.YAML
		case form:
			mp[form] = binding.Form
		case query:
			mp[query] = binding.Query
		default:
			mp[0] = nil
		}
	}
	for e := range mp {
		gbs = append(gbs, mp[e])
	}
	return gbs
}

// resolve 方法用于解析绑定信息
func (e *bindingCache) resolve(d interface{}) []uint8 {
	bs := make([]uint8, 0)
	qType := reflect.TypeOf(d).Elem()
	var tag reflect.StructTag
	var ok bool

	for i := 0; i < qType.NumField(); i++ {
		tag = qType.Field(i).Tag
		if _, ok = tag.Lookup("json"); ok {
			bs = append(bs, json)
		}
		if _, ok = tag.Lookup("xml"); ok {
			bs = append(bs, xml)
		}
		if _, ok = tag.Lookup("yaml"); ok {
			bs = append(bs, yaml)
		}
		if _, ok = tag.Lookup("form"); ok {
			bs = append(bs, form)
		}
		if _, ok = tag.Lookup("query"); ok {
			bs = append(bs, query)
		}
		if _, ok = tag.Lookup("uri"); ok {
			bs = append(bs, 0)
		}
		if t, ok := tag.Lookup("api"); ok && strings.Index(t, "dive") > -1 {
			qValue := reflect.ValueOf(d)
			bs = append(bs, e.resolve(qValue.Field(i))...)
			continue
		}
		if t, ok := tag.Lookup("validate"); ok && strings.Index(t, "dive") > -1 {
			qValue := reflect.ValueOf(d)
			bs = append(bs, e.resolve(qValue.Field(i))...)
		}
	}
	e.setBinding(reflect.TypeOf(d).String(), bs)
	return bs
}

// getBinding 方法用于从缓存中获取绑定信息
func (e *bindingCache) getBinding(name string) []uint8 {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.cache[name]
}

// setBinding 方法用于在缓存中设置绑定信息
func (e *bindingCache) setBinding(name string, bs []uint8) {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e.cache == nil {
		e.cache = make(map[string][]uint8)
	}
	e.cache[name] = bs
}
