package intl

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitIntl() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	loc := i18n.NewLocalizer(bundle, language.English.String())

	// Step 3: Define messages
	messages := &i18n.Message{
		ID:          "Emails",
		Description: "The number of unread emails a user has",
		One:         "{{.Name}} has {{.Count}} email.",
		Other:       "{{.Name}} has {{.Count}} emails.",
	}

	// Step 3: Localize Messages
	messagesCount := 2

	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,

		TemplateData: map[string]interface{}{
			"Name":  "Theo",
			"Count": messagesCount,
		},

		PluralCount: messagesCount,
	})

	fmt.Println(translation)

	return nil
}
