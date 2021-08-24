package htgotts

import (
	"htgo-tts/handlers"
	"htgo-tts/voices"

	"testing"
)

func TestSpeech_Speak(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.English}
	speech.Speak("Test")
}

func TestSpeech_Speak_MPlayer_Handler(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.English, Handler: &handlers.MPlayer{}}
	speech.Speak("Test")
}

func TestSpeech_Speak_voice_UkEnglish(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.EnglishUK}
	speech.Speak("Lancaster")
}

func TestSpeech_Speak_voice_Japanese(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.Japanese}
	speech.Speak("Test")
}
