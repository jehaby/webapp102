package resolver

func (r *Resolver) Brands() ([]*brandResolver, error) {
	// TODO: some cache
	brands, err := r.app.Repo.Brand.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]*brandResolver, 0, len(brands))
	for i, _ := range brands {
		res = append(res, &brandResolver{*brands[i]})
	}
	return res, nil
}
