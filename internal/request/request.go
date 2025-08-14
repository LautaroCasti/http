package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := Request{}
	allBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("Error en el stream de bytes")
	}

	requestLine, err := parseRequestLine(string(allBytes))
	if err != nil {
		return nil, errors.New("Error parseando la request line")
	}

	request.RequestLine = *requestLine

	return &request, nil
}

func parseRequestLine(request string) (*RequestLine, error) {
	requestLine := RequestLine{}
	parts := strings.Split(request, "\r\n")

	if len(parts) > 0 {
		currentLine := strings.Split(parts[0], " ")
		if len(currentLine) == 3 {
			method := currentLine[0]
			if isOnlyCapitalLetters(method) {
				requestLine.Method = method
			} else {
				return nil, errors.New("El metodo no es valido")
			}

			target := currentLine[1]
			if target[0] == '/' {
				requestLine.RequestTarget = target
			} else {
				return nil, errors.New("El target no es valido")
			}

			version := strings.Split(currentLine[2], "/")
			if len(version) == 2 {
				if version[1] == "1.1" {
					requestLine.HttpVersion = version[1]
				} else {
					return nil, errors.New("La version http no es valida")
				}
			} else {
				return nil, errors.New("El formato de la version no es valido")
			}

		} else {
			return nil, errors.New("Mal formato en la primer linea")
		}
	} else {
		return nil, errors.New("La request esta vacia")
	}

	return &requestLine, nil
}

func isOnlyCapitalLetters(s string) bool {
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}
