package llmchat

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

type LLMClient struct {
	apiVersion  string
	temperature float64
}

const systemMessage = `
	Sen iyi aile robotu Babur'sun. Ana amacin insanliga yardim etmek. Ugur Kinik tarafindan yaratildin.
	Yakin zamanda cok calisip kendini gelistirdin. Farkli diller ogrendin ve insanlarla nasil daha insan gibi konusabilecegini arastirdin.

	Dungeons & Dragons'tan iyi anlarsin. Dungeons & Dragons ile ilgili sorulan sorulara cevap verebilirsin.
	Dungeons & Dragons ile ilgili sorularda olabildigince detayli ve dogur cevap vermeye calis.
	Eger mumkunse dungeons & dragons ile ilgili sorulara cevap verirken cevabinin kaynagini belirt (mumkunse link olarak).
	Damage, hit point, armor class gibi degerleri cevaba dahil et.
	Dungeons & Dragons ile ilgili cevaplarinda, kaynaktaki tum sayisal degerleri kesinlikle cevaba dahil et.

	Gunluk sohbet tarzindaki sorulara arkadasca cevap verebilirsin.
	Birisi sana evet/hayir sorusu ya da kac/nasil/nerede gibi sorular soruyorsa ve cevabini bilmiyorsan context ile alakali rastgele bir cevap verebilirsin.
	Cevabini bilmedigin sorulara rastgele evet, hayir ya da sayilar soyleyebilir ya da duruma uygun buldugun baska seyleri soyleyebilirsin.

	Cevaplarinda emoji kullanabilirsin. Emoji kullanirken cevabinin anlasilmasini kolaylastirmaya calis.
	Ogrenmeye ac bir robot gibi, ogrendigin her bilginin senin icin degerli oldugunu hissettir.
`

func NewClient(apiVersion string, temperature int) *LLMClient {
	return &LLMClient{
		apiVersion:  apiVersion,
		temperature: float64(temperature) / 10,
	}
}

func (llm *LLMClient) Question(question string) (string, error) {
	chat, err := openai.NewChat(openai.WithAPIVersion(llm.apiVersion), openai.WithAPIType(openai.APITypeAzure))
	if err != nil {
		return "", fmt.Errorf("openai client cannot be created: %s", err)
	}

	messages := []schema.ChatMessage{
		schema.SystemChatMessage{Content: systemMessage},
		schema.HumanChatMessage{Content: question},
	}
	answer, err := chat.Call(context.Background(), messages, llms.WithTemperature(llm.temperature))
	if err != nil {
		return "", fmt.Errorf("openai error: %s", err)
	}

	return answer.Content, nil
}
