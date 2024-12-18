package user

import (
	"context"
	"database/sql"
	"encoding/base64"
	"strings"

	"mallchat-go/app/internal/common"
	"mallchat-go/app/internal/common/errors"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

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

func (l *UploadAvatarLogic) UploadAvatar(req *types.UploadAvatarReq) (resp *types.UploadAvatarResp, err error) {
	// 获取当前用户ID
	userId, ok := common.GetUserIDFromContext(l.ctx)
	if !ok {
		return nil, errors.NewAuthError("未登录")
	}

	// 验证文件内容
	if req.File == "" {
		return nil, errors.NewValidationError("请选择要上传的文件")
	}

	// 从base64中提取文件内容和类型
	parts := strings.Split(req.File, ";base64,")
	if len(parts) != 2 {
		return nil, errors.NewValidationError("无效的文件格式")
	}

	// 获取文件类型和内容
	fileType := strings.TrimPrefix(parts[0], "data:")
	fileContent := parts[1]

	// 验证文件类型
	if !strings.HasPrefix(fileType, "image/") {
		return nil, errors.NewValidationError("仅支持图片文件")
	}

	// 获取文件扩展名
	ext := ""
	switch fileType {
	case "image/jpeg", "image/jpg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		return nil, errors.NewValidationError("不支持的图片格式，仅支持jpg/png/gif")
	}

	// 解码base64内容
	fileBytes, err := base64.StdEncoding.DecodeString(fileContent)
	if err != nil {
		logx.Errorf("解码文件内容失败: %v", err)
		return nil, errors.NewValidationError("无效的文件内容")
	}

	// 验证文件大小（限制为2MB）
	if len(fileBytes) > 2*1024*1024 {
		return nil, errors.NewValidationError("文件大小不能超过2MB")
	}

	// 生成文件名
	fileName := "avatar" + ext

	// 上传到OSS
	url, err := l.svcCtx.OSS.UploadFile(fileBytes, fileName)
	if err != nil {
		logx.Errorf("上传文件失败: %v", err)
		return nil, errors.New(errors.UnknownError, "上传文件失败")
	}

	// 获取用户当前信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(userId))
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	// 如果用户有旧头像，删除它
	if user != nil && user.Avatar.Valid && user.Avatar.String != "" && !strings.Contains(user.Avatar.String, "default") {
		err = l.svcCtx.OSS.DeleteFile(user.Avatar.String)
		if err != nil {
			logx.Errorf("删除旧头像失败: %v", err)
		}
	}

	// 更新用户头像
	user = &model.User{
		Id:     uint64(userId),
		Avatar: sql.NullString{String: url, Valid: true},
	}
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		// 如果更新失败，删除已上传的新头像
		_ = l.svcCtx.OSS.DeleteFile(url)
		return nil, err
	}

	return &types.UploadAvatarResp{
		Url: url,
	}, nil
}
