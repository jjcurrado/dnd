package app

func (s *server) routes() {
	s.e.GET("/", s.handleHomePage())
	s.e.POST("/character", s.handleCharacterCreate())
}
