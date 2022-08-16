package workers

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"server/engine/utils"

	"server/engine/session"

	"github.com/vladimirok5959/golang-worker/worker"
)

func SessionCleaner(www_dir string) *worker.Worker {
	return worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		if www_dir, ok := (*o)[0].(string); ok {
			session_clean(ctx, www_dir)
		}
		select {
		case <-ctx.Done():
		case <-time.After(1 * time.Hour):
			return
		}
	}, &[]worker.Iface{
		www_dir,
	})
}

func session_clean(ctx context.Context, www_dir string) {
	files, err := ioutil.ReadDir(www_dir)
	if err == nil {
		for _, file := range files {
			select {
			case <-ctx.Done():
				return
			default:
				tmpdir := strings.Join([]string{www_dir, file.Name(), "tmp"}, string(os.PathSeparator))
				if utils.IsDirExists(tmpdir) {
					session.Clean(tmpdir)
				}
			}
		}
	}
}
