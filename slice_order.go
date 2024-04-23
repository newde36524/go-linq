package linq

type OrderEntity[S ~[]E, E any] struct {
	data  S
	order *Decorator[S]
	where *Decorator[S]
}

func (s *OrderEntity[S, E]) Select() *SelectEntity[S, E] {
	return &SelectEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
	}
}

func (s *OrderEntity[S, E]) Skip(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
		skip:  n,
		take:  0,
	}
}

func (s *OrderEntity[S, E]) Where(whereFunc func(a E) bool) *WhereEntity[S, E] {
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

func (s *OrderEntity[S, E]) Take(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
		skip:  0,
		take:  n,
	}
}

func (s *OrderEntity[S, E]) do() S {
	v := s.where.Build()(append((S)(nil), s.data...))
	return s.order.Build()(v)
}

func (s *OrderEntity[S, E]) First() E {
	return s.At(0)
}

func (s *OrderEntity[S, E]) Last() E {
	return s.At(len(s.data) - 1)
}

func (s *OrderEntity[S, E]) At(i int) E {
	return s.do()[i]
}

func (s *OrderEntity[S, E]) ToList() S {
	return s.do()
}
