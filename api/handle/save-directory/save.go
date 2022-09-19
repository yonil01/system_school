package save_directory

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
)

func SaveImageinDirectorio(name string, b64 string, typ string) (error) {
	unbased, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic("Cannot decode b64")
		return errors.New("Error al cargar el archivo")
	}

	r := bytes.NewReader(unbased)
	if typ == "data:image/png" {
		im, err := png.Decode(r)
		if err != nil {
		panic("Bad png")
		return errors.New("Error al cargar el archivo")

		}
		f, err := os.OpenFile("images/"+name, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
			return errors.New("Error al cargar el archivo")
		}

		png.Encode(f, im)
		return nil
	} else {
		im, err := jpeg.Decode(r)
		if err != nil {
			panic("Bad png")
			return errors.New("Error al cargar el archivo")

		}
		f, err := os.OpenFile("images/"+name, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
			return errors.New("Error al cargar el archivo")
		}

		jpeg.Encode(f, im, nil)
		return nil
	}

	return nil
}
