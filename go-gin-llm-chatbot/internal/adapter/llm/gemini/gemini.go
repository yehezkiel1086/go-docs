package gemini

import (
	"context"
	"iter"
	"mime/multipart"

	"github.com/yehezkiel1086/go-gin-llm-chatbot/internal/adapter/config"
	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
	conf *config.LLM
}

func InitGemini(ctx context.Context, conf *config.LLM) (*Gemini, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  conf.APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &Gemini{
		conf: conf,
		client: client,
	}, nil
}

func (g *Gemini) GenerateTextToText(ctx context.Context, prompt string) (string, error) {
	result, err := g.client.Models.GenerateContent(
		ctx,
		g.conf.AIModel,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

func (g *Gemini) StreamTextToText(ctx context.Context, prompt string) (iter.Seq2[*genai.GenerateContentResponse, error]) {
	return g.client.Models.GenerateContentStream(
		ctx,
		g.conf.AIModel,
		genai.Text(prompt),
		nil,
	)
}

func (g *Gemini) GenerateImageToText(ctx context.Context, prompt string, file *multipart.FileHeader) (string, error) {
	uploadedFile, err := g.client.Files.UploadFromPath(ctx, "./assets/" + file.Filename, nil)
	if err != nil {
		return "", err
	}

  parts := []*genai.Part{
		genai.NewPartFromText(prompt),
		genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType),
  }

  contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
  }

  result, err := g.client.Models.GenerateContent(
		ctx,
		g.conf.AIModel,
		contents,
		nil,
  )
	if err != nil {
		return "", err
	}

  return result.Text(), nil
}
