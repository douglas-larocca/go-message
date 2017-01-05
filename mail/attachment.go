package mail

import (
	"mime"
	"net/textproto"

	"github.com/emersion/go-message/internal"
)

// An AttachmentHeader represents an attachment's header.
type AttachmentHeader struct {
	textproto.MIMEHeader
}

// NewAttachmentHeader creates a new AttachmentHeader.
func NewAttachmentHeader() AttachmentHeader {
	h := AttachmentHeader{make(textproto.MIMEHeader)}
	h.Set("Content-Disposition", "attachment")
	h.Set("Content-Transfer-Encoding", "base64")
	return h
}

// Filename parses the attachment's filename.
func (h AttachmentHeader) Filename() (string, error) {
	_, params, err := mime.ParseMediaType(h.Get("Content-Disposition"))

	filename, ok := params["filename"]
	if !ok {
		// Using "name" in Content-Type is discouraged
		_, params, err = mime.ParseMediaType(h.Get("Content-Type"))
		filename = params["name"]
	}
	if err != nil {
		return filename, err
	}

	return internal.DecodeHeader(filename)
}

// SetFilename formats the attachment's filename.
func (h AttachmentHeader) SetFilename(filename string) {
	filename = internal.EncodeHeader(filename)
	dispParams := map[string]string{"filename": filename}
	h.Set("Content-Disposition", mime.FormatMediaType("attachment", dispParams))
}