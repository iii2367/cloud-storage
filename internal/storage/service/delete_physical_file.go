package service

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Service) deletePhysicalFile(
	id uuid.UUID,
) error {

	path := filepath.Join(
		"storage/files",
		id.String(),
	)
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}
