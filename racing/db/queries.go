package db

const (
	racesList = "list"
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
				advertised_start_time,
				CASE 
					WHEN advertised_start_time > date() THEN 'OPEN'
					ELSE 'CLOSED'
				END AS status
			FROM races
			ORDER BY advertised_start_time 
		`,
	}
}
