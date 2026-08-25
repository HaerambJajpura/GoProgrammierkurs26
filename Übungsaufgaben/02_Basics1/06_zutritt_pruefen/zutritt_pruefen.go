package zutritt_pruefen

// Zutritt prüfen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func CanEnter(age int, hasID bool) bool {
    // TODO: Kombiniere Alter und Ausweisstatus.
    if age>=18 && hasID==true {
        return true
    } else {
        return false
    }
    return false
}
