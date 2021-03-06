package request

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"tynmarket/coffeehub-go/test"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func handle(handlerFun func(r *gin.Engine, w *httptest.ResponseRecorder)) {
	test.Setup()
	test.SetUpCoffees()

	r := test.SetupRouter()
	w := httptest.NewRecorder()

	handlerFun(r, w)

	test.TearDown()
}

func handleSetup(handlerFun func(r *gin.Engine, w *httptest.ResponseRecorder)) {
	test.Setup()

	r := test.SetupRouter()
	w := httptest.NewRecorder()

	handlerFun(r, w)

	test.TearDown()
}

func TestCoffees(t *testing.T) {
	handle(func(r *gin.Engine, w *httptest.ResponseRecorder) {
		req, _ := http.NewRequest("GET", "/api/coffees", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		exp := `[{
			"area":"フランコ・ロペス",
			"arrival_date":"5月26日",
			"arrival_month":5,
			"country":"コロンビア",
			"new":true,
			"roast":"city",
			"roast_text":"シティ",
			"shop":"name",
			"taste":"口に含んだ時のやわらかな食感とやさしいオレンジのような印象はこの地域の特徴です。心地よい軽めの濃縮感、飲みこんだ後には長い甘みの余韻が続きます。全てが高いレベルで調和しているコーヒーです。",
			"url":"url/SHOP/CO-CY001.html"
		}]`
		test.AssertJSONEq(t, exp, w.Body.String())
	})
}

func TestCoffeesRoastCity(t *testing.T) {
	handle(func(r *gin.Engine, w *httptest.ResponseRecorder) {
		req, _ := http.NewRequest("GET", "/api/coffees/roast/city", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		exp := `[{
			"area":"フランコ・ロペス",
			"arrival_date":"5月26日",
			"arrival_month":5,
			"country":"コロンビア",
			"new":true,
			"roast":"city",
			"roast_text":"シティ",
			"shop":"name",
			"taste":"口に含んだ時のやわらかな食感とやさしいオレンジのような印象はこの地域の特徴です。心地よい軽めの濃縮感、飲みこんだ後には長い甘みの余韻が続きます。全てが高いレベルで調和しているコーヒーです。",
			"url":"url/SHOP/CO-CY001.html"
		}]`
		test.AssertJSONEq(t, exp, w.Body.String())
	})
}

func TestCoffeesRoastHigh(t *testing.T) {
	handle(func(r *gin.Engine, w *httptest.ResponseRecorder) {
		req, _ := http.NewRequest("GET", "/api/coffees/roast/high", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		exp := `[]`
		test.AssertJSONEq(t, exp, w.Body.String())
	})
}

func TestCoffeesCreate(t *testing.T) {
	handle(func(r *gin.Engine, w *httptest.ResponseRecorder) {
		form := url.Values{}
		form.Add("path", "/aaa/bbb")
		form.Add("countory", "日本")
		form.Add("area", "沖縄")
		form.Add("roast", "4")
		form.Add("taste", "スパイシー")
		req, _ := http.NewRequest("POST", "/api/coffees", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}
