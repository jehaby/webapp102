package resolver

func (r *Resolver) Manufacturers() ([]*manufacturerResolver, error) {
	manufacturers, err := r.app.Repo.Manufacturer.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]*manufacturerResolver, 0, len(manufacturers))
	for i, _ := range manufacturers {
		res = append(res, &manufacturerResolver{manufacturers[i]})
	}
	return res, nil
}
