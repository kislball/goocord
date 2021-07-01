package types

import "fmt"

func EndpointGateway(zlib bool, version int, encoding string) (st string) {
	st = fmt.Sprintf("wss://gateway.discord.gg?v=%d&encoding=%s", version, encoding)
	if zlib == true {
		st += "&compress=zlib-stream"
	}
	return
}
