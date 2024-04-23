package linq

import "sort"

type WhereEntity[S ~[]E, E any] struct {
	data  S
	order *Decorator[S]
	where *Decorator[S]
}

func (s *WhereEntity[S, E]) Where(whereFunc func(a E) bool) *WhereEntity[S, E] {
	s.where.Use(func(middle Middleware[S]) Middleware[S] {
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
		data:  s.data,
		order: s.order,
		where: s.where,
	}
}

func (s *WhereEntity[S, E]) Order(orderFunc func(a, b E) bool) *OrderEntity[S, E] {
	s.order.Use(func(middle Middleware[S]) Middleware[S] {
		return func(o S) S {
			o2 := append((S)(nil), o...)
			sort.Slice(o2, func(i, j int) bool {
				return orderFunc(o2[i], o2[j])
			})
			return middle(o2)
		}
	})
	return &OrderEntity[S, E]{
		order: s.order,
		where: s.where,
		data:  s.data,
	}
}

func (s *WhereEntity[S, E]) Skip(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  n,
		take:  0,
	}
}

func (s *WhereEntity[S, E]) Take(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  0,
		take:  n,
	}
}

func (s *WhereEntity[S, E]) Select() *SelectEntity[S, E] {
	return &SelectEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
	}
}
