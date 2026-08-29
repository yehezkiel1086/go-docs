package port

import (
	"context"
	"iter"
	"mime/multipart"

	"google.golang.org/genai"
)

type PromptService interface {
	TextToTextPrompt(ctx context.Context, prompt string) (string, error)
	StreamTextToTextPrompt(ctx context.Context, prompt string) (iter.Seq2[*genai.GenerateContentResponse, error])
	ImageToTextPrompt(ctx context.Context, prompt string, file *multipart.FileHeader) (string, error)
}
