package update

type Format struct {
	Bold          *FormatField `json:"bold,omitempty"`
	Italic        *FormatField `json:"italic,omitempty"`
	Underline     *FormatField `json:"underline,omitempty"`
	Strikethrough *FormatField `json:"strikethrough,omitempty"`
	Link          *FormatField `json:"link,omitempty"`
	Mention       *FormatField `json:"mention,omitempty"`
	InlineCode    *FormatField `json:"inline_code,omitempty"`
	Pre           *FormatField `json:"pre,omitempty"`
	OrderedList   *FormatField `json:"ordered_list,omitempty"`
	UnorderedList *FormatField `json:"unordered_list,omitempty"`
	Quote         *FormatField `json:"quote,omitempty"`
}

type FormatField struct {
	Offset int    `json:"offset,omitempty"`
	Length int    `json:"length,omitempty"`
	URL    string `json:"url,omitempty"`
}
