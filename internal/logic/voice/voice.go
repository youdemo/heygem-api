package voice

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"heygem-api/internal/boot"
	"heygem-api/internal/consts"
	"heygem-api/internal/dao"
	"heygem-api/internal/model/do"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/voicein"
	"heygem-api/internal/pkg/tts"
	"heygem-api/internal/service"
	util "heygem-api/utility"
)

type sVoice struct{}

func init() {
	service.RegisterVoice(&sVoice{})
}

func (s *sVoice) Page(ctx context.Context, inp *voicein.PageInp) (res *voicein.PageOut, err error) {
	col := dao.Voice.Columns()
	mod := dao.Voice.Ctx(ctx)
	if inp.Name != "" {
		mod = mod.WhereLike(col.Name, "%"+inp.Name+"%")
	}
	res = &voicein.PageOut{}
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

func (s *sVoice) Find(ctx context.Context, id int64) (res *entity.Voice, err error) {
	err = dao.Voice.Ctx(ctx).WherePri(id).Scan(&res)
	return
}

func (s *sVoice) Save(ctx context.Context, inp *voicein.SaveInp) (res *voicein.SaveOut, err error) {
	dayDir := gtime.Now().Format(consts.DayDir)

	filename, err := inp.File.Save(util.JoinPaths(boot.VideoPath, consts.Voice, dayDir), true)
	inp.AudioText, err = tts.Train(ctx, util.JoinPaths(boot.VideoPath, consts.Voice, dayDir, filename))
	if err != nil {
		err = gerror.Wrap(err, "解析语音失败")
		return
	}
	audioPath := "/" + consts.Voice + "/" + dayDir + "/" + filename
	data := do.Voice{
		Name:      inp.Name,
		AudioPath: audioPath,
		AudioText: inp.AudioText,
	}
	if inp.Name == "" {
		data.Name = gfile.Name(inp.File.Filename)
	}
	// 添加语音
	id, err := dao.Voice.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return
	}
	res = &voicein.SaveOut{Id: id, AudioPath: audioPath}
	return
}
