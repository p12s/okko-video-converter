package processing

import (
	"fmt"
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
	"github.com/p12s/okko-video-converter/api/utils/cli/ffmpeg"
)

func ProcessEvent(event common.Event, repos *repository.Repository) {
	switch event.Type {
	case common.EVENT_VIDEO_CONVERT:
		videoConvertProcess(event, repos)
	default:
		defaultProcess(event, repos)
	}
}

func videoConvertProcess(event common.Event, repos *repository.Repository) {
	err := ffmpeg.ConvertVideo(event.Value.Path, event.Value.TargetFormat)
	if err != nil {
		err = repos.UpdateStatus(event.Value.UserCode, err.Error(), common.ERROR)
		if err != nil {
			fmt.Println("file processing error:", err.Error())
		}
	}
	err = repos.UpdateStatus(event.Value.UserCode, "", common.FINISHED)
	if err != nil {
		fmt.Println("file processing error:", err.Error())
	} else {
		fmt.Println("file processing end OK")
	}
}

func defaultProcess(event common.Event, repos *repository.Repository) {
	err := repos.UpdateStatus(event.Value.UserCode, fmt.Sprintf("undefined processing event case: %s", event.Type), common.ERROR)
	if err != nil {
		fmt.Println("undefined processing event:", err.Error())
	}
}
