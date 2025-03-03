package fs

import (
	"context"
	"errors"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/kennygrant/sanitize"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/spf13/afero"
)

type fileRepo struct {
	fs      afero.Fs
	urlBase *url.URL
}

func NewFile(fs afero.Fs, urlBase string) (gateway.File, error) {
	var b *url.URL
	var err error
	b, err = url.Parse(urlBase)
	if err != nil {
		return nil, errors.New("invalid base URL")
	}

	return &fileRepo{
		fs:      fs,
		urlBase: b,
	}, nil
}

// asset

func (f *fileRepo) ReadAsset(ctx context.Context, filename string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(assetDir, sanitize.Path(filename)))
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (*url.URL, error) {
	filename := sanitize.Path(id.New().String() + path.Ext(file.Path))
	if err := f.upload(ctx, filepath.Join(assetDir, filename), file.Content); err != nil {
		return nil, err
	}
	return getAssetFileURL(f.urlBase, filename), nil
}

func (f *fileRepo) RemoveAsset(ctx context.Context, u *url.URL) error {
	if u == nil {
		return nil
	}
	p := sanitize.Path(u.Path)
	if p == "" || f.urlBase == nil || u.Scheme != f.urlBase.Scheme || u.Host != f.urlBase.Host || path.Dir(p) != f.urlBase.Path {
		return gateway.ErrInvalidFile
	}
	return f.delete(ctx, filepath.Join(assetDir, path.Base(p)))
}

// plugin

func (f *fileRepo) ReadPluginFile(ctx context.Context, pid id.PluginID, filename string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(pluginDir, pid.String(), sanitize.Path(filename)))
}

func (f *fileRepo) UploadPluginFile(ctx context.Context, pid id.PluginID, file *file.File) error {
	return f.upload(ctx, filepath.Join(pluginDir, pid.String(), sanitize.Path(file.Path)), file.Content)
}

func (f *fileRepo) RemovePlugin(ctx context.Context, pid id.PluginID) error {
	return f.delete(ctx, filepath.Join(pluginDir, pid.String()))
}

// built scene

func (f *fileRepo) ReadBuiltSceneFile(ctx context.Context, name string) (io.ReadCloser, error) {
	return f.read(ctx, filepath.Join(publishedDir, sanitize.Path(name+".json")))
}

func (f *fileRepo) UploadBuiltScene(ctx context.Context, reader io.Reader, name string) error {
	return f.upload(ctx, filepath.Join(publishedDir, sanitize.Path(name+".json")), reader)
}

func (f *fileRepo) MoveBuiltScene(ctx context.Context, oldName, name string) error {
	return f.move(
		ctx,
		filepath.Join(publishedDir, sanitize.Path(oldName+".json")),
		filepath.Join(publishedDir, sanitize.Path(name+".json")),
	)
}

func (f *fileRepo) RemoveBuiltScene(ctx context.Context, name string) error {
	return f.delete(ctx, filepath.Join(publishedDir, sanitize.Path(name+".json")))
}

// helpers

func (f *fileRepo) read(ctx context.Context, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, rerror.ErrNotFound
	}

	file, err := f.fs.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, rerror.ErrNotFound
		}
		return nil, rerror.ErrInternalBy(err)
	}
	return file, nil
}

func (f *fileRepo) upload(ctx context.Context, filename string, content io.Reader) error {
	if filename == "" {
		return gateway.ErrFailedToUploadFile
	}

	if fnd := path.Dir(filename); fnd != "" {
		if err := f.fs.MkdirAll(fnd, 0755); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	dest, err := f.fs.Create(filename)
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	defer func() {
		_ = dest.Close()
	}()

	if _, err := io.Copy(dest, content); err != nil {
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) move(ctx context.Context, from, dest string) error {
	if from == "" || dest == "" || from == dest {
		return gateway.ErrInvalidFile
	}

	if destd := path.Dir(dest); destd != "" {
		if err := f.fs.MkdirAll(destd, 0755); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	if err := f.fs.Rename(from, dest); err != nil {
		if os.IsNotExist(err) {
			return rerror.ErrNotFound
		}
		return rerror.ErrInternalBy(err)
	}

	return nil
}

func (f *fileRepo) delete(ctx context.Context, filename string) error {
	if filename == "" {
		return gateway.ErrFailedToUploadFile
	}

	if err := f.fs.RemoveAll(filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func getAssetFileURL(base *url.URL, filename string) *url.URL {
	if base == nil {
		return nil
	}

	// https://github.com/golang/go/issues/38351
	b := *base
	b.Path = path.Join(b.Path, filename)
	return &b
}
