package pokeapi

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	IsDefault bool `json:"is_default"`
	Order int `json:"order"`
	Weight int `json:"weight"`
	Abilities []struct {
		IsHidden bool `json:"is_hidden"`
		Slot int `json:"slot"`
		Ability struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"forms"`
	GameIndicies []struct {
		GameIndex int `json:"game_index"`
		Version []struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"version"`
	} `json:"game_indicies"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup struct {
				Name string `json:"name"`
				Url string `json:"url"`		
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault string `json:"back_default"`
		BackFemale *string `json:"back_female"`
		BackShiny string `json:"back_shiny"` 
		BackShinyFemale *string `json:"back_shiny_female"`
		FrontDefault string `json:"front_default"`
		FrontFemale *string `json:"front_female"`
		FrontShiny string `json:"front_shiny"`   
		FrontShinyFemale *string `json:"front_shiny_female"`  
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			Url string `json:"url"`			
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			Url string `json:"url"`	
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			Url string `json:"url"`	
		} `json:"generation"`
		Abilities []struct {
			IsHidden bool `json:"is_hidden"`
			Slot int `json:"slot"`
			Ability *struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}