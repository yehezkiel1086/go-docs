package service

import (
	"context"
	"iter"
	"mime/multipart"

	"github.com/yehezkiel1086/go-gin-llm-chatbot/internal/adapter/llm/gemini"
	"google.golang.org/genai"
)

type PromptService struct {
	llm *gemini.Gemini
}

func InitPromptService(llm *gemini.Gemini) *PromptService {
	return &PromptService{
		llm: llm,
	}
}

func (s *PromptService) TextToTextPrompt(ctx context.Context, prompt string) (string, error) {
	res, err := s.llm.GenerateTextToText(ctx, prompt)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *PromptService) StreamTextToTextPrompt(ctx context.Context, prompt string) (iter.Seq2[*genai.GenerateContentResponse, error]) {
	return s.llm.StreamTextToText(ctx, prompt)
}

func (s *PromptService) ImageToTextPrompt(ctx context.Context, prompt string, file *multipart.FileHeader) (string, error) {
	res, err := s.llm.GenerateImageToText(ctx, prompt, file)
	if err != nil {
		return "", err
	}

	return res, nil
}
