package linq

import "sort"

type SkipTakeEntity[S ~[]E, E any] struct {
	data  S
	order *Decorator[S]
	where *Decorator[S]
	skip  int
	take  int
}

func (s *SkipTakeEntity[S, E]) Skip(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		data:  append((S)(nil), s.data[n:]...),
		order: s.order,
		where: s.where,
		skip:  0,
		take:  s.take,
	}
}

func (s *SkipTakeEntity[S, E]) Take(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		data:  append((S)(nil), s.data[s.skip:s.skip+n]...),
		order: s.order,
		where: s.where,
		skip:  0,
		take:  0,
	}
}

func (s *SkipTakeEntity[S, E]) Where(whereFunc func(a E) bool) *SkipTakeEntity[S, E] {
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
	return &SkipTakeEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
		skip:  s.skip,
		take:  s.take,
	}
}

func (s *SkipTakeEntity[S, E]) Order(orderFunc func(a, b E) bool) *SkipTakeEntity[S, E] {
	s.order.Use(func(middle Middleware[S]) Middleware[S] {
		return func(o S) S {
			o2 := append((S)(nil), o...)
			sort.Slice(o2, func(i, j int) bool {
				return orderFunc(o2[i], o2[j])
			})
			return middle(o2)
		}
	})
	return &SkipTakeEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
		skip:  s.skip,
		take:  s.take,
	}
}

func (s *SkipTakeEntity[S, E]) First() E {
	return s.At(0)
}

func (s *SkipTakeEntity[S, E]) Last() E {
	return s.At(len(s.data) - 1)
}

func (s *SkipTakeEntity[S, E]) At(i int) E {
	return s.do()[i]
}

func (s *SkipTakeEntity[S, E]) ToList() S {
	return s.do()
}

func (s *SkipTakeEntity[S, E]) Select() *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
		skip:  s.skip,
		take:  s.take,
	}
}

func (s *SkipTakeEntity[S, E]) do() S {
	v := s.where.Build()(append((S)(nil), s.data...))
	return s.order.Build()(v)
}
