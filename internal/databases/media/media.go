// stub to get it to build :3
// TODO: implement anything at all lol
package media

import (
	"errors"
	"github.com/beeper/babbleserv/internal/types"
)

type MediaDatabase struct {}

func (d *MediaDatabase) GetMedia(... any) (*types.Media, error) {
	fakeMedia := types.NewMedia("", "", "", "")
	return fakeMedia, errors.New("not implemented")
}

func (d *MediaDatabase) SetMedia(... any) (error) {
	return errors.New("not implemented")
}

func (d *MediaDatabase) CreateMedia(... any) (error) {
	return errors.New("not implemented")
}

func (d *MediaDatabase) GenerateMediaID(... any) (string) {
	return "totally generated media id"
}
