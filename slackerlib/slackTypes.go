package slackerlib

import (
	"fmt"
)

type AuthResponse struct {
	Bots []struct {
		Deleted bool `json:"deleted"`
		Icons   struct {
			Image48 string `json:"image_48"`
		} `json:"icons"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"bots"`
	CacheVersion string `json:"cache_version"`
	Channels     []struct {
		Created    float64 `json:"created"`
		Creator    string  `json:"creator"`
		ID         string  `json:"id"`
		IsArchived bool    `json:"is_archived"`
		IsChannel  bool    `json:"is_channel"`
		IsGeneral  bool    `json:"is_general"`
		IsMember   bool    `json:"is_member"`
		LastRead   string  `json:"last_read"`
		Latest     struct {
			BotID    string `json:"bot_id"`
			Subtype  string `json:"subtype"`
			Text     string `json:"text"`
			Ts       string `json:"ts"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"latest"`
		Members []string `json:"members"`
		Name    string   `json:"name"`
		Purpose struct {
			Creator string  `json:"creator"`
			LastSet float64 `json:"last_set"`
			Value   string  `json:"value"`
		} `json:"purpose"`
		Topic struct {
			Creator string  `json:"creator"`
			LastSet float64 `json:"last_set"`
			Value   string  `json:"value"`
		} `json:"topic"`
		UnreadCount float64 `json:"unread_count"`
	} `json:"channels"`
	Groups []interface{} `json:"groups"`
	Ims    []struct {
		Created  float64 `json:"created"`
		ID       string  `json:"id"`
		IsIm     bool    `json:"is_im"`
		IsOpen   bool    `json:"is_open"`
		LastRead string  `json:"last_read"`
		Latest   struct {
			Text string `json:"text"`
			Ts   string `json:"ts"`
			Type string `json:"type"`
			User string `json:"user"`
		} `json:"latest"`
		UnreadCount float64 `json:"unread_count"`
		User        string  `json:"user"`
	} `json:"ims"`
	LatestEventTs string `json:"latest_event_ts"`
	Ok            bool   `json:"ok"`
	Self          struct {
		Created        float64 `json:"created"`
		ID             string  `json:"id"`
		ManualPresence string  `json:"manual_presence"`
		Name           string  `json:"name"`
		Prefs          struct {
			AllChannelsLoud                 bool    `json:"all_channels_loud"`
			ArrowHistory                    bool    `json:"arrow_history"`
			AtChannelSuppressedChannels     string  `json:"at_channel_suppressed_channels"`
			AutoplayChatSounds              bool    `json:"autoplay_chat_sounds"`
			Collapsible                     bool    `json:"collapsible"`
			CollapsibleByClick              bool    `json:"collapsible_by_click"`
			ColorNamesInList                bool    `json:"color_names_in_list"`
			CommaKeyPrefs                   bool    `json:"comma_key_prefs"`
			ConvertEmoticons                bool    `json:"convert_emoticons"`
			DisplayRealNamesOverride        float64 `json:"display_real_names_override"`
			DropboxEnabled                  bool    `json:"dropbox_enabled"`
			EmailAlerts                     string  `json:"email_alerts"`
			EmailAlertsSleepUntil           float64 `json:"email_alerts_sleep_until"`
			EmailMisc                       bool    `json:"email_misc"`
			EmailWeekly                     bool    `json:"email_weekly"`
			EmojiMode                       string  `json:"emoji_mode"`
			EnterIsSpecialInTbt             bool    `json:"enter_is_special_in_tbt"`
			ExpandInlineImgs                bool    `json:"expand_inline_imgs"`
			ExpandInternalInlineImgs        bool    `json:"expand_internal_inline_imgs"`
			ExpandNonMediaAttachments       bool    `json:"expand_non_media_attachments"`
			ExpandSnippets                  bool    `json:"expand_snippets"`
			FKeySearch                      bool    `json:"f_key_search"`
			FullTextExtracts                bool    `json:"full_text_extracts"`
			FuzzyMatching                   bool    `json:"fuzzy_matching"`
			GraphicEmoticons                bool    `json:"graphic_emoticons"`
			GrowlsEnabled                   bool    `json:"growls_enabled"`
			HasCreatedChannel               bool    `json:"has_created_channel"`
			HasInvited                      bool    `json:"has_invited"`
			HasUploaded                     bool    `json:"has_uploaded"`
			HighlightWords                  string  `json:"highlight_words"`
			KKeyOmnibox                     bool    `json:"k_key_omnibox"`
			LastSnippetType                 string  `json:"last_snippet_type"`
			LoudChannels                    string  `json:"loud_channels"`
			LoudChannelsSet                 string  `json:"loud_channels_set"`
			LsDisabled                      bool    `json:"ls_disabled"`
			MacSpeakSpeed                   float64 `json:"mac_speak_speed"`
			MacSpeakVoice                   string  `json:"mac_speak_voice"`
			MacSsbBounce                    string  `json:"mac_ssb_bounce"`
			MacSsbBullet                    bool    `json:"mac_ssb_bullet"`
			MarkMsgsReadImmediately         bool    `json:"mark_msgs_read_immediately"`
			MessagesTheme                   string  `json:"messages_theme"`
			MuteSounds                      bool    `json:"mute_sounds"`
			MutedChannels                   string  `json:"muted_channels"`
			NeverChannels                   string  `json:"never_channels"`
			NewMsgSnd                       string  `json:"new_msg_snd"`
			NoCreatedOverlays               bool    `json:"no_created_overlays"`
			NoJoinedOverlays                bool    `json:"no_joined_overlays"`
			NoMacssb1Banner                 bool    `json:"no_macssb1_banner"`
			NoTextInNotifications           bool    `json:"no_text_in_notifications"`
			ObeyInlineImgLimit              bool    `json:"obey_inline_img_limit"`
			PagekeysHandled                 bool    `json:"pagekeys_handled"`
			PostsFormattingGuide            bool    `json:"posts_formatting_guide"`
			PrivacyPolicySeen               bool    `json:"privacy_policy_seen"`
			PromptedForEmailDisabling       bool    `json:"prompted_for_email_disabling"`
			PushAtChannelSuppressedChannels string  `json:"push_at_channel_suppressed_channels"`
			PushDmAlert                     bool    `json:"push_dm_alert"`
			PushEverything                  bool    `json:"push_everything"`
			PushIdleWait                    float64 `json:"push_idle_wait"`
			PushLoudChannels                string  `json:"push_loud_channels"`
			PushLoudChannelsSet             string  `json:"push_loud_channels_set"`
			PushMentionAlert                bool    `json:"push_mention_alert"`
			PushMentionChannels             string  `json:"push_mention_channels"`
			PushSound                       string  `json:"push_sound"`
			RequireAt                       bool    `json:"require_at"`
			SearchExcludeBots               bool    `json:"search_exclude_bots"`
			SearchExcludeChannels           string  `json:"search_exclude_channels"`
			SearchOnlyMyChannels            bool    `json:"search_only_my_channels"`
			SearchSort                      string  `json:"search_sort"`
			SeenChannelMenuTipCard          bool    `json:"seen_channel_menu_tip_card"`
			SeenChannelsTipCard             bool    `json:"seen_channels_tip_card"`
			SeenDomainInviteReminder        bool    `json:"seen_domain_invite_reminder"`
			SeenFlexpaneTipCard             bool    `json:"seen_flexpane_tip_card"`
			SeenMemberInviteReminder        bool    `json:"seen_member_invite_reminder"`
			SeenMessageInputTipCard         bool    `json:"seen_message_input_tip_card"`
			SeenSearchInputTipCard          bool    `json:"seen_search_input_tip_card"`
			SeenSsbPrompt                   bool    `json:"seen_ssb_prompt"`
			SeenTeamMenuTipCard             bool    `json:"seen_team_menu_tip_card"`
			SeenUserMenuTipCard             bool    `json:"seen_user_menu_tip_card"`
			SeenWelcome2                    bool    `json:"seen_welcome_2"`
			ShowMemberPresence              bool    `json:"show_member_presence"`
			ShowTyping                      bool    `json:"show_typing"`
			SidebarBehavior                 string  `json:"sidebar_behavior"`
			SidebarTheme                    string  `json:"sidebar_theme"`
			SidebarThemeCustomValues        string  `json:"sidebar_theme_custom_values"`
			SnippetEditorWrapLongLines      bool    `json:"snippet_editor_wrap_long_lines"`
			SpeakGrowls                     bool    `json:"speak_growls"`
			SsEmojis                        bool    `json:"ss_emojis"`
			StartScrollAtOldest             bool    `json:"start_scroll_at_oldest"`
			TabUiReturnSelects              bool    `json:"tab_ui_return_selects"`
			Time24                          bool    `json:"time24"`
			Tz                              string  `json:"tz"`
			UserColors                      string  `json:"user_colors"`
			WebappSpellcheck                bool    `json:"webapp_spellcheck"`
			WelcomeMessageHidden            bool    `json:"welcome_message_hidden"`
			WinSsbBullet                    bool    `json:"win_ssb_bullet"`
		} `json:"prefs"`
	} `json:"self"`
	Team struct {
		Domain      string `json:"domain"`
		EmailDomain string `json:"email_domain"`
		Icon        struct {
			Image102     string `json:"image_102"`
			Image132     string `json:"image_132"`
			Image34      string `json:"image_34"`
			Image44      string `json:"image_44"`
			Image68      string `json:"image_68"`
			Image88      string `json:"image_88"`
			ImageDefault bool   `json:"image_default"`
		} `json:"icon"`
		ID                string  `json:"id"`
		MsgEditWindowMins float64 `json:"msg_edit_window_mins"`
		Name              string  `json:"name"`
		OverStorageLimit  bool    `json:"over_storage_limit"`
		Prefs             struct {
			AllowMessageDeletion   bool     `json:"allow_message_deletion"`
			DefaultChannels        []string `json:"default_channels"`
			DisplayRealNames       bool     `json:"display_real_names"`
			DmRetentionDuration    float64  `json:"dm_retention_duration"`
			DmRetentionType        float64  `json:"dm_retention_type"`
			GatewayAllowIrcPlain   float64  `json:"gateway_allow_irc_plain"`
			GatewayAllowIrcSsl     float64  `json:"gateway_allow_irc_ssl"`
			GatewayAllowXmppSsl    float64  `json:"gateway_allow_xmpp_ssl"`
			GroupRetentionDuration float64  `json:"group_retention_duration"`
			GroupRetentionType     float64  `json:"group_retention_type"`
			HideReferers           bool     `json:"hide_referers"`
			MsgEditWindowMins      float64  `json:"msg_edit_window_mins"`
			RequireAtForMention    float64  `json:"require_at_for_mention"`
			RetentionDuration      float64  `json:"retention_duration"`
			RetentionType          float64  `json:"retention_type"`
			WhoCanArchiveChannels  string   `json:"who_can_archive_channels"`
			WhoCanAtChannel        string   `json:"who_can_at_channel"`
			WhoCanAtEveryone       string   `json:"who_can_at_everyone"`
			WhoCanCreateChannels   string   `json:"who_can_create_channels"`
			WhoCanCreateGroups     string   `json:"who_can_create_groups"`
			WhoCanKickChannels     string   `json:"who_can_kick_channels"`
			WhoCanKickGroups       string   `json:"who_can_kick_groups"`
			WhoCanPostGeneral      string   `json:"who_can_post_general"`
		} `json:"prefs"`
	} `json:"team"`
	URL   string `json:"url"`
	Users []struct {
		Color             string      `json:"color"`
		Deleted           bool        `json:"deleted"`
		HasFiles          bool        `json:"has_files"`
		ID                string      `json:"id"`
		IsAdmin           bool        `json:"is_admin"`
		IsBot             bool        `json:"is_bot"`
		IsOwner           bool        `json:"is_owner"`
		IsPrimaryOwner    bool        `json:"is_primary_owner"`
		IsRestricted      bool        `json:"is_restricted"`
		IsUltraRestricted bool        `json:"is_ultra_restricted"`
		Name              string      `json:"name"`
		Phone             interface{} `json:"phone"`
		Presence          string      `json:"presence"`
		Profile           struct {
			Email              string `json:"email"`
			Image192           string `json:"image_192"`
			Image24            string `json:"image_24"`
			Image32            string `json:"image_32"`
			Image48            string `json:"image_48"`
			Image72            string `json:"image_72"`
			RealName           string `json:"real_name"`
			RealNameNormalized string `json:"real_name_normalized"`
		} `json:"profile"`
		RealName string      `json:"real_name"`
		Skype    interface{} `json:"skype"`
		Status   interface{} `json:"status"`
		Tz       string      `json:"tz"`
		TzLabel  string      `json:"tz_label"`
		TzOffset float64     `json:"tz_offset"`
	} `json:"users"`
}

type Event struct {
	ID      int32    `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
	Ts      string `json:"ts,omitempty"`
	Bot	  *Bot
}

func (meta *AuthResponse) GetUserName(id string) string{
	for _,user := range meta.Users{
		if user.ID == id{
			return user.Name
		}
	}
	return ``
}

func (event *Event) Reply(s string){
	replyText:=fmt.Sprintf(`%s: %s`, event.Bot.Meta.GetUserName(event.User), s)
	response := Event{
      Type:    event.Type,
      Channel: event.Channel,
      Text:    replyText,
      }
   event.Bot.WriteThread.Chan <- response
}

func (event *Event) Respond(s string){
	response := Event{
      Type:    event.Type,
      Channel: event.Channel,
      Text:    s,
      }
   event.Bot.WriteThread.Chan <- response
}
