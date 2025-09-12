package jfile

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/ComepassDeveloper/jd-utils/pkg/jstr"
)

func DetectFileType(fileHeader *multipart.FileHeader) (*string, *string) {
	// Check header first
	mimeType := fileHeader.Header.Get("Content-Type")
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	if strings.HasPrefix(mimeType, "image/") {
		return jstr.ToPtr("image"), &ext
	}
	if strings.HasPrefix(mimeType, "audio/") {
		return jstr.ToPtr("audio"), &ext
	}
	if strings.HasPrefix(mimeType, "video/") {
		return jstr.ToPtr("video"), &ext
	}

	// Fallback if header can not check
	switch ext {
	// Case image
	case ".jpg", ".jpeg", ".png":
		return jstr.ToPtr("image"), &ext
	// Case audio
	case ".mp3":
		return jstr.ToPtr("audio"), &ext
	// Case video
	case ".mp4", ".mov", ".avi":
		return jstr.ToPtr("video"), &ext
	}
	// Return "" for error if not debug mode
	return nil, nil
}
