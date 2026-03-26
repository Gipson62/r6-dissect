package dissect

import (
	"github.com/rs/zerolog/log"
)

func (r *Reader) recordScoreBasedDestruction(username string, destructionType string, timeInSeconds float64) {
	operator := r.getOperatorByUsername(username)

	if destructionType == "camera" {
		r.CameraDestructions = append(r.CameraDestructions, CameraDestruction{
			DestroyedBy:   username,
			TimeInSeconds: timeInSeconds,
		})
		log.Debug().
			Str("username", username).
			Float64("time", timeInSeconds).
			Msg("camera_destruction_detected")
	} else {
		r.UtilityEvents = append(r.UtilityEvents, UtilityEvent{
			Operator:      operator,
			Username:      username,
			UtilityType:   ClassifyGadgetByOperator(operator),
			Action:        "destroyed",
			TimeInSeconds: timeInSeconds,
		})
		log.Debug().
			Str("username", username).
			Str("operator", operator).
			Float64("time", timeInSeconds).
			Msg("utility_destruction_detected")
	}
}

func (r *Reader) getOperatorByUsername(username string) string {
	for _, p := range r.Header.Players {
		if p.Username == username {
			return p.Operator.String()
		}
	}
	return "Unknown"
}

func ClassifyGadgetByOperator(operator string) string {
	gadgetMap := map[string]string{
		"Valkyrie":    "Hereford Cam",
		"Goyo":        "Volcan Shield",
		"Maestro":     "Evil Eye",
		"AceA":        "Armor Crate",
		"Thunderbird": "Perimeter Alarm",
		"Tachanka":    "Shurikens",
		"Nomad":       "Airjab Sting",
		"Gridlock":    "Stinger Grenades",
		"Capitao":     "Incendiary Bolts",
		"Twitch":      "Shock Drone",
		"IQ":          "Electronics Detector",
		"Jackal":      "Eyenox Network",
		"Drones":      "RC Drone",
		"Thermite":    "Exothermic Charge",
		"Maverick":    "Torch",
		"Blackbeard":  "Shield",
		"Grim":        "Stinger Bolts",
		"SolidSnake":  "Cardboard Box",
	}
	if g, ok := gadgetMap[operator]; ok {
		return g
	}
	return "Gadget"
}
