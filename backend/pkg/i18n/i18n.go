package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"

	"github.com/rs/zerolog/log"
)

// Translator provides SMS template formatting and language utilities.
type Translator struct {
	configDir string
	templates map[string]map[string]string // language -> templateKey -> template
	mu        sync.RWMutex
}

// NewTranslator creates a new Translator that reads SMS templates from the
// given configuration directory.
func NewTranslator(configDir string) *Translator {
	return &Translator{
		configDir: configDir,
		templates: make(map[string]map[string]string),
	}
}

// LoadSMSTemplates loads SMS templates for the specified language from a JSON
// file at {configDir}/sms_templates/{language}.json.
func (t *Translator) LoadSMSTemplates(language string) error {
	path := filepath.Join(t.configDir, "sms_templates", language+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("i18n load templates for %s: %w", language, err)
	}

	var templates map[string]string
	if err := json.Unmarshal(data, &templates); err != nil {
		return fmt.Errorf("i18n parse templates for %s: %w", language, err)
	}

	t.mu.Lock()
	t.templates[language] = templates
	t.mu.Unlock()

	log.Info().
		Str("language", language).
		Int("templates", len(templates)).
		Msg("SMS templates loaded")

	return nil
}

// FormatSMS formats an SMS message using the named template in the specified
// language, substituting the provided variables.
//
// Variables in templates use the {{key}} syntax:
//
//	FormatSMS("otp_verification", "en", map[string]string{"code": "1234", "minutes": "5"})
func (t *Translator) FormatSMS(templateKey string, language string, vars map[string]string) (string, error) {
	t.mu.RLock()
	langTemplates, ok := t.templates[language]
	t.mu.RUnlock()

	if !ok {
		// Try to load templates on demand.
		if err := t.LoadSMSTemplates(language); err != nil {
			// Fall back to English.
			t.mu.RLock()
			langTemplates, ok = t.templates["en"]
			t.mu.RUnlock()
			if !ok {
				return "", fmt.Errorf("i18n: no templates loaded for language %q or fallback", language)
			}
		} else {
			t.mu.RLock()
			langTemplates = t.templates[language]
			t.mu.RUnlock()
		}
	}

	tmpl, ok := langTemplates[templateKey]
	if !ok {
		return "", fmt.Errorf("i18n: template %q not found for language %q", templateKey, language)
	}

	// Substitute variables.
	result := tmpl
	for k, v := range vars {
		result = strings.ReplaceAll(result, "{{"+k+"}}", v)
	}

	return result, nil
}

// supportedLanguages lists the language codes supported by the platform.
// V1: English + Hindi + Marathi only. Gujarati planned for V2.
var supportedLanguages = []string{
	"en", // English
	"hi", // Hindi
	"mr", // Marathi
}

// SupportedLanguages returns the list of language codes supported by the
// platform.
func SupportedLanguages() []string {
	langs := make([]string, len(supportedLanguages))
	copy(langs, supportedLanguages)
	return langs
}

// scriptRanges maps Unicode script ranges to language codes for simple script-based
// language detection.
var scriptRanges = []struct {
	RangeStart rune
	RangeEnd   rune
	Language   string
}{
	{0x0900, 0x097F, "hi"}, // Devanagari -> Hindi
	{0x0D00, 0x0D7F, "ml"}, // Malayalam
	{0x0B80, 0x0BFF, "ta"}, // Tamil
	{0x0C80, 0x0CFF, "kn"}, // Kannada
	{0x0C00, 0x0C7F, "te"}, // Telugu
	{0x0980, 0x09FF, "bn"}, // Bengali
	{0x0A80, 0x0AFF, "gu"}, // Gujarati
	{0x0A00, 0x0A7F, "pa"}, // Gurmukhi -> Punjabi
}

// DetectLanguage performs simple script-based language detection. It examines
// the Unicode script of each character and returns the most likely language
// code. Returns "en" if no Indic script is detected.
func DetectLanguage(text string) string {
	if text == "" {
		return "en"
	}

	scriptCounts := make(map[string]int)

	for _, r := range text {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) {
			continue
		}

		for _, sr := range scriptRanges {
			if r >= sr.RangeStart && r <= sr.RangeEnd {
				scriptCounts[sr.Language]++
				break
			}
		}
	}

	// Find the script with the most characters.
	maxCount := 0
	detected := "en"
	for lang, count := range scriptCounts {
		if count > maxCount {
			maxCount = count
			detected = lang
		}
	}

	return detected
}

// transliterationMap provides basic transliterations between Devanagari and Latin
// scripts. This is a simplified mapping for common characters.
var devanagariToLatin = map[rune]string{
	'\u0905': "a", '\u0906': "aa", '\u0907': "i", '\u0908': "ee", '\u0909': "u", '\u090A': "oo",
	'\u090F': "e", '\u0910': "ai", '\u0913': "o", '\u0914': "au",
	'\u0915': "ka", '\u0916': "kha", '\u0917': "ga", '\u0918': "gha",
	'\u091A': "cha", '\u091B': "chha", '\u091C': "ja", '\u091D': "jha",
	'\u091F': "ta", '\u0920': "tha", '\u0921': "da", '\u0922': "dha",
	'\u0924': "ta", '\u0925': "tha", '\u0926': "da", '\u0927': "dha", '\u0928': "na",
	'\u092A': "pa", '\u092B': "pha", '\u092C': "ba", '\u092D': "bha", '\u092E': "ma",
	'\u092F': "ya", '\u0930': "ra", '\u0932': "la", '\u0935': "va",
	'\u0936': "sha", '\u0937': "sha", '\u0938': "sa", '\u0939': "ha",
	'\u0902': "n", '\u0903': "h",
	'\u093E': "a", '\u093F': "i", '\u0940': "ee", '\u0941': "u", '\u0942': "oo",
	'\u0947': "e", '\u0948': "ai", '\u094B': "o", '\u094C': "au",
	'\u094D': "",
}

// Transliterate performs basic transliteration between scripts. Currently
// supports Devanagari to Latin conversion. For production-quality transliteration,
// consider using a proper library or API.
func Transliterate(text string, fromScript, toScript string) string {
	if fromScript == "devanagari" && toScript == "latin" {
		var result strings.Builder
		for _, r := range text {
			if latin, ok := devanagariToLatin[r]; ok {
				result.WriteString(latin)
			} else {
				result.WriteRune(r)
			}
		}
		return result.String()
	}

	// Unsupported script pair; return original text.
	log.Warn().
		Str("from", fromScript).
		Str("to", toScript).
		Msg("unsupported transliteration pair")

	return text
}
