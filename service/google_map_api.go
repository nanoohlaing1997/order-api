package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func GetDistance(origin, destination []string) (float64, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey != "" {
		url := fmt.Sprintf(
			"https://maps.googleapis.com/maps/api/distancematrix/json?units=metric&origins=%s&destinations=%s&key=%s",
			origin[0]+","+origin[1],
			destination[0]+","+destination[1],
			apiKey,
		)
		res, err := http.Get(url)
		if err != nil {
			return 0, err
		}
		defer res.Body.Close()

		var data map[string]interface{}

		// testing
		data = map[string]interface{}{
			"rows": []interface{}{
				map[string]interface{}{
					"elements": []interface{}{
						map[string]interface{}{
							"distance": map[string]interface{}{
								"value": 1235325101,
							},
						},
					},
				},
			},
		}
		//

		// if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		// 	return 0, err
		// }

		rows, ok := data["rows"].([]interface{})
		if !ok || len(rows) == 0 {
			return 0, errors.New("Distance information not found")
		}

		elements := rows[0].(map[string]interface{})["elements"].([]interface{})
		if len(elements) == 0 {
			return 0, errors.New("Distance information not found")
		}

		distanceVal, ok := elements[0].(map[string]interface{})["distance"].(map[string]interface{})["value"]
		if !ok {
			return 0, errors.New("Distance information not found")
		}

		switch distance := distanceVal.(type) {
		case int:
			return float64(distance), nil
		case int16:
			return float64(distance), nil
		case int64:
			return float64(distance), nil
		case float32:
			return float64(distance), nil
		case float64:
			return float64(distance), nil
		}
	}
	return 0, errors.New("Google map API key is missing")
}
