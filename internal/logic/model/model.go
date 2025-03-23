package model

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
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
	"heygem-api/internal/model/input/modelin"
	"heygem-api/internal/pkg/tts"
	"heygem-api/internal/service"
	util "heygem-api/utility"
	"strconv"
	"strings"
)

type sModel struct{}

func init() {
	service.RegisterModel(&sModel{})
}

func (s *sModel) Page(ctx context.Context, inp *modelin.PageInp) (res *modelin.PageOut, err error) {
	col := dao.Model.Columns()
	mod := dao.Model.Ctx(ctx)
	if inp.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+inp.Name+"%")
	}
	res = &modelin.PageOut{}
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
	return
}

func (s *sModel) Count(ctx context.Context, inp *modelin.CountInp) (total int, err error) {
	col := dao.Model.Columns()
	mod := dao.Model.Ctx(ctx)
	if inp.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+inp.Name+"%")
	}
	total, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据行失败，请稍后重试！")
		return
	}
	return
}

func (s *sModel) Find(ctx context.Context, id int64) (res *entity.Model, err error) {
	err = dao.Model.Ctx(ctx).WherePri(id).Scan(&res)
	return
}

// Add 添加模特
func (s *sModel) Add(ctx context.Context, inp *modelin.AddInp) (res *modelin.AddOut, err error) {
	// 分离视频/音频
	dayDir := gtime.Now().Format(consts.DayDir)
	filename, err := inp.File.Save(util.JoinPaths(boot.VideoPath, consts.Model, dayDir), true)
	if err != nil {
		return nil, err
	}
	modelPath := util.NewBuilder().Append(util.Separator).Append(util.JoinPaths(consts.Model, dayDir, filename)).String()
	inputFile := boot.VideoPath + modelPath

	outputName := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	// 输出音频文件路径
	outputAudio := outputName + consts.Wav
	audioPath := util.NewBuilder().Append(util.Separator).Append(util.JoinPaths(consts.Voice, dayDir, outputAudio)).String()
	// 输出视频文件路径（无音频）
	// 分离音频
	outputAudioDir := util.JoinPaths(boot.VideoPath, consts.Voice, dayDir)
	if !gfile.Exists(outputAudioDir) {
		_ = gfile.Mkdir(util.JoinPaths(boot.VideoPath, consts.Voice, dayDir))
	}
	err = ffmpeg.Input(inputFile).
		Output(util.JoinPaths(boot.VideoPath, consts.Voice, dayDir, outputAudio), ffmpeg.KwArgs{"q:a": 0, "map": "a"}).
		Run()
	if err != nil {
		return nil, gerror.Wrap(err, "分离音频失败")
	}
	g.Log().Info(ctx, "音频分离成功")
	// 调用api解析语音文本
	text, err := tts.Train(ctx, util.JoinPaths(boot.VideoPath, audioPath))
	if err != nil {
		return nil, gerror.Wrap(err, "解析语音失败")
	}
	err = dao.Model.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		voice := do.Voice{
			Name:      inp.Name,
			AudioPath: audioPath,
			AudioText: text,
		}
		// 添加语音
		voiceId, err := dao.Voice.Ctx(ctx).Data(voice).InsertAndGetId()
		if err != nil {
			return err
		}
		// 保存模特
		modelId, err := dao.Model.Ctx(ctx).Data(do.Model{
			Name:      inp.Name,
			VideoPath: modelPath,
			AudioPath: voice.AudioPath,
			VoiceId:   voiceId,
		}).InsertAndGetId()
		res = &modelin.AddOut{
			Id:   modelId,
			Path: modelPath,
		}
		return
	})
	return
}

// Del 删除模特
func (s *sModel) Del(ctx context.Context, id int64) (err error) {
	model, err := s.Find(ctx, id)
	if err != nil || model == nil {
		return
	}
	// 删除文件
	_ = gfile.RemoveFile(util.JoinPaths(boot.VideoPath, model.VideoPath))
	_ = gfile.RemoveFile(util.JoinPaths(boot.VideoPath, model.AudioPath))

	return dao.Model.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Model.Ctx(ctx).WherePri(id).Delete()
		if err != nil {
			return err
		}
		_, err = dao.Voice.Ctx(ctx).WherePri(model.VoiceId).Delete()
		return
	})
}
