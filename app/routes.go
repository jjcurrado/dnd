package app

// instantiates routing to the server
func (s *server) routes() {
	// GET
	s.e.GET("/", s.handleLogging(s.handleHomePage()))
	s.e.GET("/home", s.handleLogging(s.handleWelcome()))
	s.e.GET("/prompt", s.handleLogging(s.handlePrompt()))
	s.e.GET("/sheets", s.handleLogging(s.handleSheetsView()))
	s.e.GET("/notes", s.handleLogging(s.handleSheetView()))

	// POST
	s.e.POST("/character", s.handleLogging(s.handleCharacterCreate()))
}
