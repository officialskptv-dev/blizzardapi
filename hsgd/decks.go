package hsgd

// Deck structure
type Deck struct {
	Version int    `json:"version"`
	Format  string `json:"format"`
	Hero    struct {
		ID            int    `json:"id"`
		Collectible   int    `json:"collectible"`
		Slug          string `json:"slug"`
		ClassID       int    `json:"classId"`
		MultiClassIds []int  `json:"multiClassIds"`
		CardTypeID    int    `json:"cardTypeId"`
		CardSetID     int    `json:"cardSetId"`
		RarityID      int    `json:"rarityId"`
		ArtistName    string `json:"artistName"`
		Health        int    `json:"health"`
		ManaCost      int    `json:"manaCost"`
		Name          string `json:"name"`
		Text          string `json:"text"`
		Image         string `json:"image"`
		ImageGold     string `json:"imageGold"`
		FlavorText    string `json:"flavorText"`
		CropImage     string `json:"cropImage"`
		ChildIds      []int  `json:"childIds"`
	} `json:"hero"`
	HeroPower struct {
		ID            int    `json:"id"`
		Collectible   int    `json:"collectible"`
		Slug          string `json:"slug"`
		ClassID       int    `json:"classId"`
		MultiClassIds []int  `json:"multiClassIds"`
		CardTypeID    int    `json:"cardTypeId"`
		CardSetID     int    `json:"cardSetId"`
		RarityID      int    `json:"rarityId"`
		ArtistName    string `json:"artistName"`
		ManaCost      int    `json:"manaCost"`
		Name          string `json:"name"`
		Text          string `json:"text"`
		Image         string `json:"image"`
		ImageGold     string `json:"imageGold"`
		FlavorText    string `json:"flavorText"`
		CropImage     string `json:"cropImage"`
	} `json:"heroPower"`
	Class struct {
		ID   int    `json:"id"`
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"class"`
	Cards []struct {
		ID            int    `json:"id"`
		Collectible   int    `json:"collectible"`
		Slug          string `json:"slug"`
		ClassID       int    `json:"classId"`
		MultiClassIds []int  `json:"multiClassIds"`
		CardTypeID    int    `json:"cardTypeId"`
		CardSetID     int    `json:"cardSetId"`
		RarityID      int    `json:"rarityId"`
		ArtistName    string `json:"artistName"`
		Health        int    `json:"health,omitempty"`
		Attack        int    `json:"attack,omitempty"`
		ManaCost      int    `json:"manaCost"`
		Name          string `json:"name"`
		Text          string `json:"text"`
		Image         string `json:"image"`
		ImageGold     string `json:"imageGold"`
		FlavorText    string `json:"flavorText"`
		CropImage     string `json:"cropImage"`
		KeywordIds    []int  `json:"keywordIds,omitempty"`
		Armor         int    `json:"armor,omitempty"`
		ChildIds      []int  `json:"childIds,omitempty"`
		MinionTypeID  int    `json:"minionTypeId,omitempty"`
		Durability    int    `json:"durability,omitempty"`
	} `json:"cards"`
	CardCount int `json:"cardCount"`
}
