package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"botmg.com/go-server/config"
	"botmg.com/go-server/routes"
	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Option struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Option   Option `json:"option"`
}

type Exam struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Date      time.Time  `json:"date"`
	Marks     int8       `json:"marks"`
	Time      time.Time  `json:"time"`
	Questions []Question `json:"questions"`
}

var question = Question{
	Id:       "iw7eriase698y23h9834n",
	Question: "sdjbkbsdkfkjasbd,asdnbc,mansdvf,jahsvdlfj,hvWE,MNCV,BXMZNCVZ,SJDHVFAKJNMSDV BV,ZDJFLZKJSDBV,DKDJGF;ASKFALKSNNDF;LASOIFIRF;LSDNVV;AKSHRO;IFFHA;LSKDN;ORIHGALKGHG;ALKDFGHG;OIDIHFG;LKANDF;OVILH;RLKGNG;ALDKFFNVOIRG;ALKERNVO98I4TT0U340TOJ;LKDSNFF0FP49UTPJA;SDLVKJ[03QP94U[PAKSDN;VAOEW4UF[09WEUAKLNDSV;VO4FKRNF;LASKDHF[OE4FH;LASKDNV;OARIH4HFFOPAKNS;RVKLAS[HF'KWENF;ALSKNVDOIHUGAEROHGKJRKDFJHLKDJHFGGORIOQ[PWEOR[PEWOIR[QPWOEIR[QPOWIE[QPWOEITU[QWPEORITU[QWOPEIRRTUQ]PWEOIRKVAL;SKDJF;ALSDKNXCV.,ZXCMNVZNC.,MVZ.,XCV;JFHD;LVKASND,VMN;ORIHF'PWIRIJ;LVKZDF;LKVZ;LKDHF;LKHAER'PIGH'PAKSNDF;VLKNZ;FDMNVV;OIERHO;SNDFVBZKSDJB;OSIHG;OSHRG;HRE;GLK;SDFKJB;SIODH;SOR;GGKNGDKFBN;OSDHT;OGITNR;VNZ;DKF;OHD;ORIGH;DFNV;VLNZODFHIGOIERHGOAR;LKN;DFOKHG'AOHRGOIHAORIGH4PTPOJRFKAD;LKLFDKHGLAKDHG;LAHDF",
	Option: Option{
		A: "ksdjhf;oWIHEF;KJSDBFKJSDH;FHGW;EKLJFB.KSDJBCERGF;KJSBD.KJVbsdk>jVBLKSJDBVLKJAGS;KJFGA;SKDJFG;KSDJGF",
		B: "ksdjhf;oWIHEF;KJSDBFKJSDH;FHGW;EKLJFB.KSDJBCERGF;KJSBD.KJVbsdk>jVBLKSJDBVLKJAGS;KJFGA;SKDJFG;KSDJGF",
		C: "ksdjhf;oWIHEF;KJSDBFKJSDH;FHGW;EKLJFB.KSDJBCERGF;KJSBD.KJVbsdk>jVBLKSJDBVLKJAGS;KJFGA;SKDJFG;KSDJGF",
		D: "ksdjhf;oWIHEF;KJSDBFKJSDH;FHGW;EKLJFB.KSDJBCERGF;KJSBD.KJVbsdk>jVBLKSJDBVLKJAGS;KJFGA;SKDJFG;KSDJGF",
	},
	Answer: "A",
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	var router = config.Router
	config.ConnectDB()

	setupRoute(router)

	runtime.GOMAXPROCS(8)
	router.GET("/", cache.CachePage(config.CacheStore, time.Minute*3, func(ctx *gin.Context) {
		var list []Question
		for range 50 {
			list = append(list, question)
		}

		ctx.JSON(http.StatusOK, Exam{
			Id:        27763627,
			Name:      "history of india",
			Date:      time.Now(),
			Marks:     50,
			Time:      time.Now(),
			Questions: list,
		})
	}))

	fmt.Println("Web server is started.")
	router.Run(":8080")
}

func setupRoute(router *gin.Engine) {
	routes.UserRoute(router)
}
