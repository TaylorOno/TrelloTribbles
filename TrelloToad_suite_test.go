package main

import (
	"TrelloToad/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrelloToad(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TrelloToad Suite")
}

var _ = Describe("trello", func() {
	var (
		mockCtrl   *gomock.Controller
		mockClient *mocks.MockClient
		trello     *Trello
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockClient = mocks.NewMockClient(mockCtrl)
		trello = NewTrelloClient(mockClient, "apiKey", "apiString")
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Get Board Lists", func() {
		It("Returns a list", func() {
			//given
			mockClient.EXPECT().Do(gomock.Any()).Return(responseFromFile(200, "test_data/board.json"))

			//when
			lists, err := trello.getBoardLists("boardId")

			//then
			Expect(err).To(BeNil())
			Expect(len(lists)).To(Equal(3))
		})
	})
})

func responseFromFile(i int, s string) *http.Response {
	jsonContentFile, err := os.Open(s)
	if err != nil {
		Fail(fmt.Sprintf("unable to open test file %s", s))
	}
	defer jsonContentFile.Close()
	byteValue, err := ioutil.ReadAll(jsonContentFile)
	if err != nil {
		Fail(fmt.Sprintf("unable to read test file %s", s))
	}
	contentBody := ioutil.NopCloser(strings.NewReader(string(byteValue)))
	return &http.Response{StatusCode: i, Body: contentBody}
}
