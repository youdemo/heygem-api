package tts

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"heygem-api/internal/boot"
	"heygem-api/internal/consts"
	"heygem-api/internal/dao"
	"heygem-api/internal/model/entity"
	"heygem-api/internal/model/input/ttsin"
	"heygem-api/internal/pkg/tts"
	"heygem-api/internal/service"
	util "heygem-api/utility"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type sTts struct{}

func init() {
	service.RegisterTts(&sTts{})
}

func (s *sTts) Train(ctx context.Context, inp *ttsin.TrainInp) (res *ttsin.TrainOut, err error) {

	return
}

func (s *sTts) Invoke(ctx context.Context, inp *ttsin.InvokeInp) (res *ttsin.InvokeOut, err error) {
	var voice *entity.Voice
	err = dao.Voice.Ctx(ctx).WherePri(inp.VoiceId).Scan(&voice)
	if err != nil || voice == nil {
		return nil, gerror.Wrap(err, "获取语音失败")
	}
	audio, err := tts.Invoke(ctx, util.JoinPaths(boot.VideoPath, voice.AudioPath), inp.Text, voice.AudioText, inp.ResponseType)
	downloadName := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + consts.Wav
	r := g.RequestFromCtx(ctx).Response
	r.Header().Set("Content-Type", "application/force-download")
	r.Header().Set("Accept-Ranges", "bytes")
	r.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(downloadName)))
	r.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	g.RequestFromCtx(ctx).Response.ServeContent(downloadName, time.Now(), bytes.NewReader(audio))
	return
}
