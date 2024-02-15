package app

func (s *server) routes() {
	s.e.POST("/character", s.handleCharacterCreate())
}
