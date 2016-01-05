package schema

type MessageFeed struct {
	Messages   []*Message   `json:"messages"`
	References []*Reference `json:"references"`
	Meta       *Meta        `json:"meta"`
}

type Message struct {
	Attachments    []*Attachment `json:"attachments"`
	Body           *MessageBody  `json:"body"`
	Id             int           `json:"id"`
	GroupId        int           `json:"group_id"`
	GroupCreatedId int           `json:"group_created_id"`
	CreatedAt      *Time         `json:"created_at"`
	ClientType     string        `json:"client_type"`
	ClientURL      string        `json:"client_url"`
	ContentExcerpt string        `json:"content_excerpt"`
	DirectMessage  bool          `json:"direct_message"`
	Language       string        `json:"language"`
	MessageType    string        `json:"message_type"`
	NetworkId      int           `json:"network_id"`
	Privacy        string        `json:"privacy"`
	RepliedToId    int           `json:"replied_to_id"`
	SenderId       int           `json:"sender_id"`
	SenderType     string        `json:"sender_type"`
	SystemMessage  bool          `json:"system_message"`
	ThreadId       int           `json:"thread_id"`
	URL            string        `json:"url"`
	WebURL         string        `json:"web_url"`
}

func (m *Message) IsThreadStarter() bool {
	if m.RepliedToId == 0 {
		return true
	}
	return false
}

type Attachment struct {
	Id                 int    `json:"id"`
	NetworkId          int    `json:"network_id"`
	URL                string `json:"url"`
	WebURL             string `json:"web_url"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	OriginalName       string `json:"original_name"`
	FullName           string `json:"full_name"`
	Description        string `json:"description"`
	ContentType        string `json:"content_type"`
	ContentClass       string `json:"content_class"`
	CreatedAt          *Time  `json:"created_at"`
	OwnerId            int    `json:"owner_id"`
	OwnerType          string `json:"owner_type"`
	Official           bool   `json:"official"`
	SmallIconURL       string `json:"small_icon_url"`
	LargeIconURL       string `json:"large_icon_url"`
	DownloadURL        string `json:"download_url"`
	ThumbnailURL       string `json:"thumbnail_url"`
	PreviewURL         string `json:"preview_url"`
	LargePreviewURL    string `json:"large_preview_url"`
	Size               int    `json:"size"`
	LastUploadedAt     *Time  `json:"last_uploaded_at"`
	LastUploadedById   int    `json:"last_uploaded_by_id"`
	LastUploadedByType string `json:"last_uploaded_by_type"`
}

type MessageBody struct {
	Parsed string `json:"parsed`
	Plain  string `json:"plain"`
	Rich   string `json:"rich"`
}

type Meta struct {
	CurrentUserId         int       `json:"current_user_id"`
	LastSeenMessageId     int       `json:"last_seen_message_id"`
	RequestedPollInterval int       `json:"requested_poll_interval"`
	Realtime              *Realtime `json:"realtime,omitempty"`
}

type Realtime struct {
	AuthenticationToken string `json:"authentication_token"`
	ChannelId           string `json:"channel_id"`
	URI                 string `json:"uri"`
}
