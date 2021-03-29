package db

const (
	racesList = "list"
	sportsList = "list"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
			ORDER BY advertised_start_time 
		`,
	}
}

func getSportsQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM sports
			ORDER BY advertised_start_time 
		`,
	}
}
