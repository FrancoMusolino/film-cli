package menu

type Item struct {
	Name, Headers string
	Action        func() error
}

type Menu struct {
	Items map[string]Item
}

func InitMenu() *Menu {
	menu := &Menu{
		map[string]Item{
			"top-rated": {
				Name:    "Mejor calificadas",
				Headers: "Obten las películas mejores calificadas",
				Action: func() error {
					return nil
				},
			},
			"now-playing": {
				Name:    "Reproduciendo ahora",
				Headers: "Obten las películas que se están reproduciendo en las pantallas de los cines",
				Action: func() error {
					return nil
				},
			},
			"popular": {
				Name:    "Populares",
				Headers: "Obten las películas más populares del momento",
				Action: func() error {
					return nil
				},
			},
			"upcoming": {
				Name:    "Próximamente",
				Headers: "Obten las películas que estarán en cartelera en los próximos días",
				Action: func() error {
					return nil
				},
			},
		},
	}

	return menu
}
