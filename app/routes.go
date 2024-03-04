package app

// instantiates routing to the server
func (s *server) routes() {
	// GET
	s.e.GET("/", s.handleHomePage())
	s.e.GET("/home", s.handleWelcome())
	s.e.GET("/prompt", s.handlePrompt())
	s.e.GET("/sheets", s.handleSheetsView())
	s.e.GET("/notes", s.handleSheetView())
	s.e.GET("/upload", s.handleFileUploadView())

	// POST
	s.e.POST("/character", s.handleCharacterCreate())
	s.e.POST("/upload", s.handleFileUpload())
}
