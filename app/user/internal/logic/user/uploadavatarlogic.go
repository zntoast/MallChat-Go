package user

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"
	"path"
	"strconv"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAvatarLogic) UploadAvatar(file *multipart.FileHeader) (resp *types.UploadAvatarResp, err error) {
	// 检查文件大小
	if file.Size > l.svcCtx.Config.Upload.MaxSize*1024*1024 {
		return nil, fmt.Errorf("文件大小不能超过%dMB", l.svcCtx.Config.Upload.MaxSize)
	}

	// 检查文件类型
	ext := path.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return nil, fmt.Errorf("只支持jpg/jpeg/png格式图片")
	}

	// 保存文件
	relativePath, err := utils.SaveUploadedFile(file, l.svcCtx.Config.Upload.SaveDir)
	if err != nil {
		return nil, err
	}

	// 更新用户头像
	userIdStr := l.ctx.Value("X-User-ID").(string)
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	avatarUrl := l.svcCtx.Config.Upload.BaseUrl + "/" + relativePath
	user.Avatar = sql.NullString{String: avatarUrl, Valid: true}
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.UploadAvatarResp{
		Url: avatarUrl,
	}, nil
}
