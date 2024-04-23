package linq

import (
	"sort"
)

type SliceEntity[S ~[]E, E any] struct {
	data  S
	order *Decorator[S]
	where *Decorator[S]
}

func Query[S ~[]E, E any](data S) *SliceEntity[S, E] {
	return &SliceEntity[S, E]{
		data:  data,
		order: NewDecorator[S](),
		where: NewDecorator[S](),
	}
}

func (s *SliceEntity[S, E]) do() S {
	v := s.where.Build()(append((S)(nil), s.data...))
	return s.order.Build()(v)
}

func (s *SliceEntity[S, E]) At(i int) E {
	return s.do()[i]
}

func (s *SliceEntity[S, E]) First() E {
	return s.At(0)
}

func (s *SliceEntity[S, E]) Last() E {
	return s.At(len(s.data) - 1)
}

func (s *SliceEntity[S, E]) ToList() S {
	return s.do()
}

func (s *SliceEntity[S, E]) Where(whereFunc func(a E) bool) *WhereEntity[S, E] {
	where := NewDecorator[S]()
	where.Use(func(middle Middleware[S]) Middleware[S] {
		return func(o S) S {
			var newData S
			for _, v := range o {
				if whereFunc(v) {
					newData = append(newData, v)
				}
			}
			return middle(newData)
		}
	})
	return &WhereEntity[S, E]{
		data:  append((S)(nil), s.data...),
		order: NewDecorator[S](),
		where: where,
	}
}

func (s *SliceEntity[S, E]) Order(orderFunc func(a, b E) bool) *OrderEntity[S, E] {
	order := NewDecorator[S]()
	order.Use(func(middle Middleware[S]) Middleware[S] {
		return func(o S) S {
			o2 := append((S)(nil), o...)
			sort.Slice(o2, func(i, j int) bool {
				return orderFunc(o2[i], o2[j])
			})
			return middle(o2)
		}
	})
	return &OrderEntity[S, E]{
		data:  append((S)(nil), s.data...),
		order: order,
		where: NewDecorator[S](),
	}
}

func (s *SliceEntity[S, E]) Skip(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  n,
		take:  0,
	}
}

func (s *SliceEntity[S, E]) Take(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  0,
		take:  n,
	}
}

func (s *SliceEntity[S, E]) Select() *SelectEntity[S, E] {
	return &SelectEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
	}
}
