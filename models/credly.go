package models

import "time"

type Credly struct {
	Data struct {
		ID                string      `json:"id"`
		ExpiresAtDate     string      `json:"expires_at_date"`
		IssuedAtDate      string      `json:"issued_at_date"`
		IssuedTo          string      `json:"issued_to"`
		Locale            string      `json:"locale"`
		Public            bool        `json:"public"`
		State             string      `json:"state"`
		TranslateMetadata bool        `json:"translate_metadata"`
		AcceptedAt        time.Time   `json:"accepted_at"`
		ExpiresAt         time.Time   `json:"expires_at"`
		IssuedAt          time.Time   `json:"issued_at"`
		LastUpdatedAt     time.Time   `json:"last_updated_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		EarnerPath        string      `json:"earner_path"`
		EarnerPhotoURL    interface{} `json:"earner_photo_url"`
		IsPrivateBadge    bool        `json:"is_private_badge"`
		UserIsEarner      bool        `json:"user_is_earner"`
		Issuer            struct {
			Summary  string `json:"summary"`
			Entities []struct {
				Label   string `json:"label"`
				Primary bool   `json:"primary"`
				Entity  struct {
					Type                           string `json:"type"`
					ID                             string `json:"id"`
					Name                           string `json:"name"`
					URL                            string `json:"url"`
					VanityURL                      string `json:"vanity_url"`
					InternationalizeBadgeTemplates bool   `json:"internationalize_badge_templates"`
					ShareToZiprecruiter            bool   `json:"share_to_ziprecruiter"`
					Verified                       bool   `json:"verified"`
				} `json:"entity"`
			} `json:"entities"`
		} `json:"issuer"`
		BadgeTemplate struct {
			ID                string      `json:"id"`
			Cost              interface{} `json:"cost"`
			Description       string      `json:"description"`
			GlobalActivityURL string      `json:"global_activity_url"`
			Level             string      `json:"level"`
			Name              string      `json:"name"`
			VanitySlug        string      `json:"vanity_slug"`
			TimeToEarn        interface{} `json:"time_to_earn"`
			TypeCategory      string      `json:"type_category"`
			ShowBadgeLmi      bool        `json:"show_badge_lmi"`
			ShowSkillTagLinks bool        `json:"show_skill_tag_links"`
			Translatable      bool        `json:"translatable"`
			Image             struct {
				ID  string `json:"id"`
				URL string `json:"url"`
			} `json:"image"`
			ImageURL string `json:"image_url"`
			URL      string `json:"url"`
			Issuer   struct {
				Summary  string `json:"summary"`
				Entities []struct {
					Label   string `json:"label"`
					Primary bool   `json:"primary"`
					Entity  struct {
						Type                           string `json:"type"`
						ID                             string `json:"id"`
						Name                           string `json:"name"`
						URL                            string `json:"url"`
						VanityURL                      string `json:"vanity_url"`
						InternationalizeBadgeTemplates bool   `json:"internationalize_badge_templates"`
						ShareToZiprecruiter            bool   `json:"share_to_ziprecruiter"`
						Verified                       bool   `json:"verified"`
					} `json:"entity"`
				} `json:"entities"`
			} `json:"issuer"`
			Alignments              []interface{} `json:"alignments"`
			BadgeTemplateActivities []struct {
				ID                      string      `json:"id"`
				ActivityType            string      `json:"activity_type"`
				RequiredBadgeTemplateID interface{} `json:"required_badge_template_id"`
				Title                   string      `json:"title"`
				URL                     interface{} `json:"url"`
			} `json:"badge_template_activities"`
			Endorsements []interface{} `json:"endorsements"`
			Skills       []struct {
				ID         string `json:"id"`
				Name       string `json:"name"`
				VanitySlug string `json:"vanity_slug"`
			} `json:"skills"`
		} `json:"badge_template"`
		Image struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"image"`
		ImageURL        string        `json:"image_url"`
		Evidence        []interface{} `json:"evidence"`
		Recommendations []interface{} `json:"recommendations"`
	} `json:"data"`
	Metadata struct {
	} `json:"metadata"`
}
