package internal

import (
	"TrelloTribbles/mocks"
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

func TestTrelloTribbles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TrelloTribbles Suite")
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
		trello = NewTrelloClient(mockClient, "apiKey", "apiToken")
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Get Board Lists", func() {
		It("Returns a list of Lists", func() {
			//given
			mockClient.EXPECT().Do(gomock.Any()).Return(responseFromFile(200, "test_data/board.json"), nil)

			//when
			lists, err := trello.getBoardLists("boardId")

			//then
			Expect(err).To(BeNil())
			Expect(len(lists)).To(Equal(3))
		})
	})

	Context("Get Lists Cards", func() {
		It("Returns a list of Cards", func() {
			//given
			mockClient.EXPECT().Do(gomock.Any()).Return(responseFromFile(200, "test_data/lists.json"), nil)

			//when
			lists, err := trello.getListCards("listId")

			//then
			Expect(err).To(BeNil())
			Expect(len(lists)).To(Equal(3))
		})
	})

	Context("Sets Auth", func() {
		It("Adds the appropriate auth header to requests", func() {
			//given
			req, _ := http.NewRequest(http.MethodGet, "http://apiurl.com", nil)

			//when
			trello.setAuth(req)

			//then
			authString := req.Header.Get("Authorization")
			Expect(authString).To(Equal("OAuth oauth_consumer_key=\"apiKey\", oauth_token=\"apiToken\""))
		})
	})

	Context("SortCards", func() {
		It("Adds the appropriate auth header to requests", func() {
			//given
			boardSorter := &BoardSorter{}
			cards := []Card{
				{"2", "item2", 1, []string{"user1", "user2"}},
				{"3", "item3", 2, []string{"user1", "user2", "user3"}},
				{"1", "item1", 4, []string{"user1"}},
			}

			//when
			boardSorter.sortCards(cards)

			//then
			Expect(cards[0].Id).To(Equal("3"))
			Expect(cards[1].Id).To(Equal("2"))
			Expect(cards[2].Id).To(Equal("1"))
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
