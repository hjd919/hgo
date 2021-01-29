package qq

type GeocodeRegeo struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Result    struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Address            string `json:"address"`
		FormattedAddresses struct {
			Recommend string `json:"recommend"`
			Rough     string `json:"rough"`
		} `json:"formatted_addresses"`
		AddressComponent struct {
			Nation       string `json:"nation"`
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_component"`
		AdInfo struct {
			NationCode string `json:"nation_code"`
			Adcode     string `json:"adcode"`
			CityCode   string `json:"city_code"`
			Name       string `json:"name"`
			Location   struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
		} `json:"ad_info"`
		AddressReference struct {
			BusinessArea struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"business_area"`
			FamousArea struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"famous_area"`
			Crossroad struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"crossroad"`
			Town struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"town"`
			StreetNumber struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"street_number"`
			Street struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
				DirDesc  string  `json:"_dir_desc"`
			} `json:"street"`
			LandmarkL2 struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"landmark_l2"`
		} `json:"address_reference"`
		PoiCount int `json:"poi_count"`
		Pois     []struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Address  string `json:"address"`
			Category string `json:"category"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			AdInfo struct {
				Adcode   string `json:"adcode"`
				Province string `json:"province"`
				City     string `json:"city"`
				District string `json:"district"`
			} `json:"ad_info"`
			Distance int    `json:"_distance"`
			DirDesc  string `json:"_dir_desc"`
		} `json:"pois"`
	} `json:"result"`
}
