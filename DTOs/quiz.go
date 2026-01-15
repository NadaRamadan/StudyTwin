package DTOs

type QuizRequest struct {
	Topic        string `json:"topic" binding:"required"`
	NumQuestions int    `json:"num_questions" binding:"required"`
}

type QuizQuestion struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
}

type QuizResponse struct {
	Questions []QuizQuestion `json:"questions"`
}
