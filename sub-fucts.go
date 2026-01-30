package main


func checkRep(sl []Room, s Room) bool {
	for _, r := range sl {
		if r.name == s.name {
			return false
		}
	}
	return true
}