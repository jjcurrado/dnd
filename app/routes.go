package app

func (s *server) routes() {
	s.e.GET("/", s.handleLogging(s.handleHomePage()))
	s.e.POST("/character", s.handleLogging(s.handleCharacterCreate()))
}
