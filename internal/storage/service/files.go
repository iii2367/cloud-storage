package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Service) saveFile(
	file *multipart.FileHeader,
	id uuid.UUID,
) error {

	root := "storage/files"

	err := os.MkdirAll(
		root,
		0755,
	)

	if err != nil {
		return err
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	path := filepath.Join(
		root,
		id.String(),
	)
	dst, err := os.Create(path)

	if err != nil {
		return err
	}

	defer dst.Close()
	_, err = io.Copy(
		dst,
		src,
	)

	return err
}
