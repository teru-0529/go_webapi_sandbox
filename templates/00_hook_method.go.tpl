{{- $alias := .Aliases.Table .Table.Name}}

// Insert hook with contexte user
func {{$alias.UpSingular}}InsertHook(ctx context.Context, exec boil.ContextExecutor, m *{{$alias.UpSingular}}) error {
	id := conn.GetUser(ctx)
	m.CreatedUser = id
	m.ModifiedUser = id
	return nil
}

// Update hook with contexte user
func {{$alias.UpSingular}}UpdateHook(ctx context.Context, exec boil.ContextExecutor, m *{{$alias.UpSingular}}) error {
	m.ModifiedUser = conn.GetUser(ctx)
	return nil
}
