package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	e, err := NewEntry([]byte(`
		{
			"time": "2009-11-10T23:05:06.000000+00:00",
			"file": "main.go:123",
			"level": "notice",
			"host": "localhost",
			"keyi": 123,
			"keyf": 3.1415926535,
			"key": "val",
			"msg": "{\"inside\":2,\"inside2\":4}"
			}`))
	assert.NoError(t, err)
	if !assert.NotNil(t, e) {
		t.FailNow()
	}
	assert.Equal(t, 2009, e.Time.Year())
	assert.Equal(t, 6, e.Time.Second())
	assert.Equal(t, "main.go:123", e.File)
	assert.Equal(t, "notice", e.Level)
	i, ok := e.I("keyi")
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assert.Equal(t, 123, i)
	f, ok := e.F("keyf")
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assert.Equal(t, 3.1415926535, f)
	s, ok := e.S("key")
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assert.Equal(t, "val", s)
	assert.Equal(t, "{\"inside\":2,\"inside2\":4}", e.Msg)
}

func TestError(t *testing.T) {
	buf := `{"time":"2017-04-28T15:28:07.275599+02:00", "file":"requester_ss.go:269", "host":"app1", "app":"backend_api", "level":"info", "url":"http://upis_validator.service.sd/validacija", "listic_id":"132ABD96-883A-4781-A50E-CDF6029D326A", "request":"{\"listic_id\":\"132ABD96-883A-4781-A50E-CDF6029D326A\",\"listic\":{\"ulog_bez\":3.8,\"ulog_mt\":0.2,\"ulog_po_listicu\":0,\"tip\":0,\"tecaj\":162.3446,\"dobitak\":616.91,\"isplata\":555.6,\"porez\":61.31,\"tecajevi\":[592575030,592697949,592698007,592573743,592573706,592573821,592771883,592757192,592761762,592094372,592695621,592228284,592347751,592714900,592574718,592228182,592228104,592639122,592099161,592396633,592745295,592751139,592054510,592182036,592592143,592639065,592697544,592697561,592031487,592697111],\"fixevi\":null},\"listicTip\":\"sport\",\"tecajevi\":{\"592031487\":{\"tecaj_id\":592031487,\"ponuda_id\":10484743,\"dogadjaj_id\":10484743,\"tip\":\"1\",\"naziv\":\"B.Dortmund-1.FC K\u00f6ln\",\"broj\":\"625\",\"ponuda_db_id\":10484743,\"vrijeme\":\"29.04. 15:30\",\"tecaj\":1.25,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592054510\":{\"tecaj_id\":592054510,\"ponuda_id\":10408864,\"dogadjaj_id\":10408864,\"tip\":\"1\",\"naziv\":\"RB Salzburg-SV Ried\",\"broj\":\"322\",\"ponuda_db_id\":10408864,\"vrijeme\":\"29.04. 18:30\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592094372\":{\"tecaj_id\":592094372,\"ponuda_id\":10580557,\"dogadjaj_id\":10580557,\"tip\":\"1\",\"naziv\":\"BenficaL-GD Estoril Praia\",\"broj\":\"671\",\"ponuda_db_id\":10580557,\"vrijeme\":\"29.04. 19:15\",\"tecaj\":1.15,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592099161\":{\"tecaj_id\":592099161,\"ponuda_id\":10457768,\"dogadjaj_id\":10457768,\"tip\":\"1\",\"naziv\":\"Sheffield Utd-Chesterfield\",\"broj\":\"389\",\"ponuda_db_id\":10457768,\"vrijeme\":\"30.04. 13:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592182036\":{\"tecaj_id\":592182036,\"ponuda_id\":10448158,\"dogadjaj_id\":10448158,\"tip\":\"1\",\"naziv\":\"Pardub-1.SK Prost\u011bjov\",\"broj\":\"443\",\"ponuda_db_id\":10448158,\"vrijeme\":\"29.04. 10:15\",\"tecaj\":1.25,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592228104\":{\"tecaj_id\":592228104,\"ponuda_id\":10416699,\"dogadjaj_id\":10416699,\"tip\":\"1\",\"naziv\":\"HNK Rijeka-Inter Zapre\u0161i\u0107\",\"broj\":\"811\",\"ponuda_db_id\":10416699,\"vrijeme\":\"29.04. 19:00\",\"tecaj\":1.15,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592228182\":{\"tecaj_id\":592228182,\"ponuda_id\":10416703,\"dogadjaj_id\":10416703,\"tip\":\"1\",\"naziv\":\"Dinamo Zagreb-Cibalia\",\"broj\":\"815\",\"ponuda_db_id\":10416703,\"vrijeme\":\"29.04. 15:00\",\"tecaj\":1.1,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592228284\":{\"tecaj_id\":592228284,\"ponuda_id\":10421316,\"dogadjaj_id\":10421316,\"tip\":\"1\",\"naziv\":\"De Graafschap-Achilles\",\"broj\":\"618\",\"ponuda_db_id\":10421316,\"vrijeme\":\"28.04. 20:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592347751\":{\"tecaj_id\":592347751,\"ponuda_id\":10570435,\"dogadjaj_id\":10570435,\"tip\":\"1\",\"naziv\":\"B.Honv\u00e9d FC-Gyirm\u00f3t FC\",\"broj\":\"497\",\"ponuda_db_id\":10570435,\"vrijeme\":\"29.04. 18:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592396633\":{\"tecaj_id\":592396633,\"ponuda_id\":11601049,\"dogadjaj_id\":11601049,\"tip\":\"1\",\"naziv\":\"AhlyCairo-Entag El Harby\",\"broj\":\"770\",\"ponuda_db_id\":11601049,\"vrijeme\":\"30.04. 19:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592573706\":{\"tecaj_id\":592573706,\"ponuda_id\":10629183,\"dogadjaj_id\":10629183,\"tip\":\"2\",\"naziv\":\"Espanyol Bar.-FC Barcelona\",\"broj\":\"504\",\"ponuda_db_id\":10629183,\"vrijeme\":\"29.04. 20:45\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592573743\":{\"tecaj_id\":592573743,\"ponuda_id\":10629185,\"dogadjaj_id\":10629185,\"tip\":\"1\",\"naziv\":\"Real Madrid-CF Valencia\",\"broj\":\"506\",\"ponuda_db_id\":10629185,\"vrijeme\":\"29.04. 16:15\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592573821\":{\"tecaj_id\":592573821,\"ponuda_id\":10629186,\"dogadjaj_id\":10629186,\"tip\":\"1\",\"naziv\":\"Real Sociedad-Granada CF\",\"broj\":\"507\",\"ponuda_db_id\":10629186,\"vrijeme\":\"29.04. 13:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592574718\":{\"tecaj_id\":592574718,\"ponuda_id\":12041422,\"dogadjaj_id\":12041422,\"tip\":\"1\",\"naziv\":\"Kairat Almaty-FK Akzhayik\",\"broj\":\"657\",\"ponuda_db_id\":12041422,\"vrijeme\":\"29.04. 14:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592575030\":{\"tecaj_id\":592575030,\"ponuda_id\":12197388,\"dogadjaj_id\":12197388,\"tip\":\"1\",\"naziv\":\"ShakhtarD-FC Oleksandria\",\"broj\":\"824\",\"ponuda_db_id\":12197388,\"vrijeme\":\"30.04. 16:00\",\"tecaj\":1.15,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592592143\":{\"tecaj_id\":592592143,\"ponuda_id\":10351359,\"dogadjaj_id\":10351359,\"tip\":\"1\",\"naziv\":\"Monaco-Toulouse\",\"broj\":\"903\",\"ponuda_db_id\":10351359,\"vrijeme\":\"29.04. 17:00\",\"tecaj\":1.25,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592639065\":{\"tecaj_id\":592639065,\"ponuda_id\":10732352,\"dogadjaj_id\":10732352,\"tip\":\"1\",\"naziv\":\"Panathin.Atena-AE Larissa\",\"broj\":\"523\",\"ponuda_db_id\":10732352,\"vrijeme\":\"30.04. 18:00\",\"tecaj\":1.25,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592639122\":{\"tecaj_id\":592639122,\"ponuda_id\":10732358,\"dogadjaj_id\":10732358,\"tip\":\"1\",\"naziv\":\"PAOK Solun-AOK Kerkyra\",\"broj\":\"526\",\"ponuda_db_id\":10732358,\"vrijeme\":\"30.04. 18:00\",\"tecaj\":1.1,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592695621\":{\"tecaj_id\":592695621,\"ponuda_id\":10769457,\"dogadjaj_id\":10769457,\"tip\":\"1\",\"naziv\":\"Wolfsburg II-SV Eichede\",\"broj\":\"1819\",\"ponuda_db_id\":10769457,\"vrijeme\":\"29.04. 13:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592697111\":{\"tecaj_id\":592697111,\"ponuda_id\":11595564,\"dogadjaj_id\":11595564,\"tip\":\"1\",\"naziv\":\"Al Ahli Dubai-Al Sharjah\",\"broj\":\"1192\",\"ponuda_db_id\":11595564,\"vrijeme\":\"29.04. 16:00\",\"tecaj\":1.25,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592697544\":{\"tecaj_id\":592697544,\"ponuda_id\":11653999,\"dogadjaj_id\":11653999,\"tip\":\"1\",\"naziv\":\"LevadTal-P\u00e4rnu JK Vaprus\",\"broj\":\"779\",\"ponuda_db_id\":11653999,\"vrijeme\":\"28.04. 18:00\",\"tecaj\":1.03,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592697561\":{\"tecaj_id\":592697561,\"ponuda_id\":11654001,\"dogadjaj_id\":11654001,\"tip\":\"2\",\"naziv\":\"ViljanTul-Flora Tallinn\",\"broj\":\"780\",\"ponuda_db_id\":11654001,\"vrijeme\":\"28.04. 18:00\",\"tecaj\":1.15,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592697949\":{\"tecaj_id\":592697949,\"ponuda_id\":11983617,\"dogadjaj_id\":11983617,\"tip\":\"1\",\"naziv\":\"Buriram Utd-Sisaket FC\",\"broj\":\"848\",\"ponuda_db_id\":11983617,\"vrijeme\":\"29.04. 15:00\",\"tecaj\":1.1,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592698007\":{\"tecaj_id\":592698007,\"ponuda_id\":11983620,\"dogadjaj_id\":11983620,\"tip\":\"1\",\"naziv\":\"MuangTU-Suphanburi FC\",\"broj\":\"851\",\"ponuda_db_id\":11983620,\"vrijeme\":\"30.04. 14:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592714900\":{\"tecaj_id\":592714900,\"ponuda_id\":12375577,\"dogadjaj_id\":12375577,\"tip\":\"1\",\"naziv\":\"Keflav\u00edk-Vidir Gardur\",\"broj\":\"254\",\"ponuda_db_id\":12375577,\"vrijeme\":\"28.04. 21:00\",\"tecaj\":1.2,\"dogadjaj_grupa\":5,\"sport_razrada_id\":22},\"592745295\":{\"tecaj_id\":592745295,\"ponuda_id\":10448153,\"dogadjaj_id\":10448153,\"tip\":\"1\",\"naziv\":\"Olomouc-Ban\u0...", "msg":"validate_service"}`

	e, err := NewEntry([]byte(buf))
	assert.Nil(t, err)
	assert.Equal(t, "backend_api", e.App)
	//	testu.PP(e)
}
