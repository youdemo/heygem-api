package video

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"heygem-api/internal/boot"
	"heygem-api/internal/consts"
	"heygem-api/internal/dao"
	"heygem-api/internal/model/do"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/videoin"
	"heygem-api/internal/pkg/f2f"
	"heygem-api/internal/pkg/tts"
	"heygem-api/internal/service"
	util "heygem-api/utility"
	"slices"
	"strconv"
	"strings"
)

type sVideo struct{}

func init() {
	service.RegisterVideo(&sVideo{})
}

func (s *sVideo) Page(ctx context.Context, inp *videoin.PageInp) (res *videoin.PageOut, err error) {
	col := dao.Video.Columns()
	mod := dao.Video.Ctx(ctx)
	if inp.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+inp.Name+"%")
	}
	res = &videoin.PageOut{}
	res.Total, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据行失败，请稍后重试！")
		return
	}
	err = mod.Page(inp.Page, inp.PageSize).OrderDesc(col.CreatedAt).Scan(&res.List)

	if err != nil {
		err = gerror.Wrap(err, "获取数据失败，请稍后重试！")
		return
	}
	if res.List != nil {
		statusVideos, err := s.SelectByStatus(ctx, &videoin.SelectByStatusInp{Status: "waiting"})
		if err != nil {
			return nil, err
		}
		for _, video := range res.List {
			if video.Status == "waiting" {
				video.Progress = fmt.Sprintf("%d/%d", slices.IndexFunc(statusVideos, func(item *entity.Video) bool {
					return item.Id == video.Id
				})+1, len(statusVideos))
			}
		}

	}

	return
}

func (s *sVideo) SelectByStatus(ctx context.Context, inp *videoin.SelectByStatusInp) (res []*entity.Video, err error) {

	err = dao.Video.Ctx(ctx).Where(do.Video{Status: inp.Status}).Scan(&res)

	if err != nil {
		err = gerror.Wrap(err, "获取数据失败，请稍后重试！")
		return
	}
	return
}

func (s *sVideo) Count(ctx context.Context, name string) (total int, err error) {
	col := dao.Video.Columns()
	mod := dao.Video.Ctx(ctx)
	if name != "" {
		mod = mod.WhereLike(col.Name, "%"+name+"%")
	}
	total, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据行失败，请稍后重试！")
		return
	}
	return
}

func (s *sVideo) Find(ctx context.Context, id int64) (res *entity.Video, err error) {
	err = dao.Video.Ctx(ctx).WherePri(id).Scan(&res)
	return
}

func (s *sVideo) FindByStatus(ctx context.Context, status string) (res *entity.Video, err error) {
	err = dao.Video.Ctx(ctx).WherePri(do.Video{Status: status}).Scan(&res)
	return
}

func (s *sVideo) Save(ctx context.Context, inp *videoin.SaveInp) (id int64, err error) {
	video, _ := s.Find(ctx, inp.Id)
	if video != nil {
		_, err = dao.Video.Ctx(ctx).OmitEmptyData().Data(inp).WherePri(video.Id).Update()
		if err != nil {
			return
		}
		return video.Id, nil
	}
	inp.Status = "draft"
	id, err = dao.Video.Ctx(ctx).OmitEmptyData().Data(inp).InsertAndGetId()
	return
}

func (s *sVideo) Del(ctx context.Context, id int64) (err error) {
	item, err := s.Find(ctx, id)
	if err != nil || item == nil {
		return
	}
	// 删除文件
	_ = gfile.RemoveFile(util.JoinPaths(boot.VideoPath, item.AudioPath))
	_ = gfile.RemoveFile(util.JoinPaths(boot.VideoPath, item.FilePath))
	_, err = dao.Video.Ctx(ctx).WherePri(id).Delete()
	return
}

func (s *sVideo) MakeByF2F(ctx context.Context, id int64) (err error) {
	dayDir := gtime.Now().Format(consts.DayDir)
	video, err := s.Find(ctx, id)
	if err != nil || video == nil {
		return
	}
	// 根据modelId获取model信息
	model, err := service.Model().Find(ctx, video.ModelId)
	if err != nil {
		return
	}
	if video.AudioPath == "" {
		// 调用TTS生成语音
		var voice *entity.Voice
		voice, err = service.Voice().Find(ctx, video.VoiceId)
		if err != nil {
			return gerror.Wrap(err, "生成语音失败")
		}
		audio, err := tts.Invoke(ctx, util.JoinPaths(boot.VideoPath, voice.AudioPath), video.TextContent, voice.AudioText, consts.RespStream)
		downloadName := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + consts.Wav
		err = gfile.PutBytes(util.JoinPaths(boot.VideoPath, consts.Voice, dayDir, downloadName), audio)
		if err != nil {
			return gerror.Wrap(err, "生成语音失败")
		}
		video.AudioPath = util.NewBuilder().Append(util.Separator).Append(util.JoinPaths(consts.Voice, dayDir, downloadName)).String()
	}
	// 调用视频生成接口生成视频
	code := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	param, res, err := f2f.Submit(ctx, code, video.AudioPath, model.VideoPath)
	if err != nil {
		_, _ = dao.Video.Ctx(ctx).Data(do.Video{
			Status:  consts.StatusFailed,
			Message: err.Error(),
		}).WherePri(id).Update()
		return
	}
	g.Log().Info(ctx, "res: ", res.MustToJsonString())
	if res.Get("code").Int() == 10000 {
		_, _ = dao.Video.Ctx(ctx).Data(do.Video{
			Status:    consts.StatusPending,
			Message:   res.MustToJsonString(),
			AudioPath: video.AudioPath,
			Param:     gjson.New(param).MustToJsonString(),
			Code:      code,
		}).WherePri(id).Update()
	} else {
		_, _ = dao.Video.Ctx(ctx).Data(do.Video{
			Status:    consts.StatusFailed,
			Message:   res.Get("message").String(),
			AudioPath: video.AudioPath,
			Param:     gjson.New(param).MustToJsonString(),
			Code:      code,
		}).WherePri(id).Update()
	}
	return
}

// LoopPending 处理未合成的视频
func (s *sVideo) LoopPending(ctx context.Context) {
	var pendingVideo *entity.Video
	_ = dao.Video.Ctx(ctx).Where(do.Video{Status: consts.StatusPending}).Scan(&pendingVideo)
	if pendingVideo == nil {
		s.synthesisNext(ctx)
		return
	}
	res, err := f2f.Query(ctx, pendingVideo.Code)
	if err != nil {
		return
	}
	var updateData do.Video
	if slices.Contains(consts.F2FFailed, res.Code) {
		updateData = do.Video{Status: consts.StatusFailed, Message: res.Msg}
	} else if consts.F2FSuccess == res.Code {
		if res.Data.Status == 1 {
			updateData = do.Video{Status: consts.StatusPending, Message: res.Data.Msg, Progress: res.Data.Progress}
		} else if res.Data.Status == 2 {
			// 移动视频到result目录
			dayDir := gtime.Now().Format(consts.DayDir)
			filePath := util.NewBuilder().Append(boot.VideoPath).Append(util.JoinPaths(consts.Voice, dayDir, res.Data.Result)).String()
			if !gfile.Exists(filePath) {
				_ = gfile.Mkdir(filePath)
			}
			err = gfile.Move(util.JoinPaths(boot.VideoPath, consts.Temp, res.Data.Result), util.JoinPaths(boot.VideoPath, filePath))
			if err != nil {
				g.Log().Error(ctx, "移动视频失败", err.Error())
				return
			}
			// 获取视频时长
			info, err := ffmpeg.Probe(util.JoinPaths(boot.VideoPath, filePath))
			if err != nil {
				g.Log().Error(ctx, "获取视频信息失败", err)
				return
			}
			duration := gjson.New(info).Get("streams.0.duration").Float64()
			updateData = do.Video{Status: consts.StatusSuccess, Message: res.Data.Msg, Progress: res.Data.Progress, FilePath: filePath, Duration: duration}
		} else if res.Data.Status == 3 {
			updateData = do.Video{Status: consts.StatusFailed, Message: res.Data.Msg}
		}
	}
	_, _ = dao.Video.Ctx(ctx).Data(updateData).WherePri(pendingVideo.Id).Update()
}

// 合成下一个视频
func (s *sVideo) synthesisNext(ctx context.Context) {
	var waitingVideo *entity.Video
	_ = dao.Video.Ctx(ctx).Where(do.Video{Status: consts.StatusWaiting}).Scan(&waitingVideo)
	if waitingVideo != nil {
		_ = s.MakeByF2F(ctx, waitingVideo.Id)
	}

}
