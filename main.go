package main

import (
	"context"
	"fmt"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"src.agwa.name/tlshacks"
)

func main() {
	fsthttp.ServeFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)

		ch := tlshacks.UnmarshalClientHello(r.TLSInfo.ClientHello)
		if ch == nil {
			fmt.Fprintln(w, "Not a TLS connection")
			return
		}

		fmt.Fprintln(w, ch.Info.JA3String)
		fmt.Fprintln(w)
		fmt.Fprintln(w, ch.Info.JA3Fingerprint)
	})
}
