package bean

import "reflect"


func defaultBeanProvider(v any) BeanProvider {
	return func() reflect.Value {
		return reflect.ValueOf(v)
	}
}

// 对外暴露依赖注入的能力，name为空字符串时表示默认使用类型注入
func Inject(obj any, name string) error {
	return instance.Inject(defaultBeanProvider(obj), name)
}

// 对外暴露依赖注入的能力，name为空字符串时表示默认使用类型注入
func DeepInject(provider BeanProvider, name string) error {
	return instance.Inject(provider, name)
}

// 对外暴露自动装配的能力，name为空字符串时表示默认使用类型自动装配
func Autowise[T any](obj *T, name string) error {
	return instance.Autowise(obj, name)
}
