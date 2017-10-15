package anvers

type AdminMenuEntry struct {
	Name        string
	Title       string
	CreateTitle string
	Link        string
	Entries     []*AdminMenuEntry
}

func newAdminMenu() *AdminMenuEntry {
	return &AdminMenuEntry{
		Name:  "main",
		Title: "Main Menu",
		Entries: []*AdminMenuEntry{
			{Name: "content_types", Title: "Content Types", Entries: []*AdminMenuEntry{}},
		},
	}
}

func (a *AdminMenuEntry) GetByName(name string) *AdminMenuEntry {
	if a.Name == name {
		return a
	}

	if a.Entries == nil || len(a.Entries) == 0 {
		return nil
	}

	for _, entry := range a.Entries {
		e := entry.GetByName(name)
		if e != nil {
			return e
		}
	}

	return nil
}

func (a *AdminMenuEntry) GetEntries() []*AdminMenuEntry {
	return a.Entries
}

func (a *AdminMenuEntry) AddEntry(entry *AdminMenuEntry) {
	a.Entries = append(a.Entries, entry)
}
