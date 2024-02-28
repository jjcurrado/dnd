package app

func (s *server) routes() {
	s.e.GET("/", s.handleLogging(s.handleHomePage()))
	s.e.GET("/home", s.handleLogging(s.handleWelcome()))
	s.e.GET("/prompt", s.handleLogging(s.handlePrompt()))
	s.e.GET("/sheets", s.handleLogging(s.handleSheetsView()))
	s.e.POST("/character", s.handleLogging(s.handleCharacterCreate()))
}
