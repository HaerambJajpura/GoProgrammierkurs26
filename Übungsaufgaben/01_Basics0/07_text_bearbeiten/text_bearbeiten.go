package text_bearbeiten

// Text mit strings bearbeiten
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import "strings"

func ChangeCase(text string) (string, string) {
    // TODO: Erzeuge Groß- und Kleinschreibung.
    upper := strings.ToUpper(text)
    lower := strings.ToLower(text)
    return upper, lower
}
