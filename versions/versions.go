package main

import (
	"log"

	"github.com/neil-stoker/goav/avcodec"
	"github.com/neil-stoker/goav/avdevice"
	"github.com/neil-stoker/goav/avfilter"
	"github.com/neil-stoker/goav/avformat"
	"github.com/neil-stoker/goav/avutil"
	"github.com/neil-stoker/goav/swresample"
	"github.com/neil-stoker/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
