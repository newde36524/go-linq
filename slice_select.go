package linq

type SelectEntity[S ~[]E, E any] struct {
	data  S
	order *Decorator[S]
	where *Decorator[S]
}

func (s *SelectEntity[S, E]) Select() *SelectEntity[S, E] {
	return &SelectEntity[S, E]{
		data:  s.data,
		order: s.order,
		where: s.where,
	}
}

func (s *SelectEntity[S, E]) do() S {
	v := s.where.Build()(append((S)(nil), s.data...))
	return s.order.Build()(v)
}

func (s *SelectEntity[S, E]) First() E {
	return s.At(0)
}

func (s *SelectEntity[S, E]) Last() E {
	return s.At(len(s.data) - 1)
}

func (s *SelectEntity[S, E]) At(i int) E {
	return s.do()[i]
}

func (s *SelectEntity[S, E]) ToList() S {
	return s.do()
}

func (s *SelectEntity[S, E]) ToMap(cmp func(E) (key string)) map[string]S {
	result := make(map[string]S, len(s.data))
	for _, v := range s.data {
		key := cmp(v)
		result[key] = append(result[key], v)
	}
	return result
}

func (s *SelectEntity[S, E]) Skip(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  n,
		take:  0,
	}
}

func (s *SelectEntity[S, E]) Take(n int) *SkipTakeEntity[S, E] {
	return &SkipTakeEntity[S, E]{
		order: s.order,
		data:  s.data,
		where: s.where,
		skip:  0,
		take:  n,
	}
}
