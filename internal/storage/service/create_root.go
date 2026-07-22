package service

import ("context")

func (s *Service) CreateRoot(
	ctx context.Context,
	userID uint,
) error {

	root, err := s.treeNodeRepo.FindRoot(ctx, userID)

	if err == nil && root != nil {
    	return nil
	}
	_, err = s.createNode(
		ctx,
		"root",
		"root folder",
		"folder",
		nil,
		userID,
	)

	return err
}
