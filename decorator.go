package linq

// Decorator x
type Decorator[T any] struct {
	components []Component[T]
}

// NewDecorator 创建一个新的实例
func NewDecorator[T any]() *Decorator[T] {
	return new(Decorator[T])
}

// New 新实例
func (a *Decorator[T]) New(app *Decorator[T]) *Decorator[T] {
	return &Decorator[T]{components: app.components}
}

// Middleware 中间件
type Middleware[T any] func(o T) T

// Component 组件
type Component[T any] func(middle Middleware[T]) Middleware[T]

// Use 使用中间件
func (app *Decorator[T]) Use(middleware Component[T]) {
	app.components = append(app.components, middleware)
}

// Build 创建中间件
func (app *Decorator[T]) Build() Middleware[T] {
	var middleware Middleware[T] = func(o T) T {
		return o
	}
	for _, m := range app.components {
		middleware = m(middleware)
	}
	return middleware
}
