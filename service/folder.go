package service

import (
	"context"
	"github.com/wq1019/cloud_disk/model"
)

type folderService struct {
	model.FolderStore
}

func LoadFolder(ctx context.Context, id, userId int64, isLoadRelated bool) (folder *model.Folder, err error) {
	return FromContext(ctx).LoadFolder(id, userId, isLoadRelated)
}

func CreateFolder(ctx context.Context, folder *model.Folder) (err error) {
	return FromContext(ctx).CreateFolder(folder)
}

func ExistFolder(ctx context.Context, userId int64, folderName string) (isExist bool) {
	return FromContext(ctx).ExistFolder(userId, folderName)
}

func NewFolderService(ds model.FolderStore) model.FolderService {
	return &folderService{ds}
}
