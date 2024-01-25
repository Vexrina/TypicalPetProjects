package shorterurl_test

import (
	"testing"
	"typicalypetprojects/pkg/shorterurl"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		inputString    string
		expectedString string
	}{
		{"123", "321"},
		{"111", "111"},
		{"helloworld", "dlrowolleh"},
		{"", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.inputString, func(t *testing.T) {
			reversedString := shorterurl.Reverse(tc.inputString)
			if reversedString != tc.expectedString {
				t.Errorf("Expected string: %s, got %s", tc.expectedString, reversedString)
			}
		})
	}
}

func TestMakeHash(t *testing.T) {
	testCases := []struct {
		inputId      string
		expectedHash string
	}{
		{"123", "GxbC"},
		{"ABC", "Hx2"},
		{"SomethingNew", "QwkI"},
		{"5fe31cacb0fa449c8af7ff6a27a62fa4", "30ha"},
	}
	for _, tc := range testCases {
		t.Run(tc.inputId, func(t *testing.T) {
			hash := shorterurl.MakeHash(tc.inputId)
			if hash != tc.expectedHash {
				t.Errorf("Expected hash: %s, got %s", tc.expectedHash, hash)
			}
		})
	}
}

/*
The hash is based on the UUID.
The UUID is generated using the library.
To verify exactly that they were unique:
len(set[uuids]) = 3.
The third test case checks that the function
understands that there is no need to shorten the URL.
The fourth test case checks that the function will
not work with an empty input.
*/
func TestCreateTiniUrl(t *testing.T) {
	uuids := make(map[string]bool)
	shortUrls := make(map[string]bool)

	testCases := []struct {
		inputUrl           string
		expectedSiteName   string
		expectedPathToPage string
	}{
		{"https://stackoverflow.com/questions/49826038/how-to-add-variable-to-string-variable-in-golang", "stackoverflow.com", "questions/49826038/how-to-add-variable-to-string-variable-in-golang"},
		{"stackoverflow.com/questions/49826038/how-to-add-variable-to-string-variable-in-golang", "stackoverflow.com", "questions/49826038/how-to-add-variable-to-string-variable-in-golang"},
		{"stackoverflow.com", "", ""},
		{"", "", ""},
		{"example.org/smth", "example.org", "smth"},
	}
	for _, tc := range testCases {
		t.Run(tc.inputUrl, func (t *testing.T)  {
			id, siteName, pathToPage, result := shorterurl.СreateTinyUrl(tc.inputUrl)
			if len(id) > 0 && !uuids[id] && !shortUrls[result]{
				uuids[id] = true
				shortUrls[result] = true
			}
			if tc.expectedPathToPage!=pathToPage || tc.expectedSiteName!=siteName {
				t.Errorf("Expected siteName: %s, got %s\nExpected pathToPage: %s, got %s", tc.expectedSiteName, siteName, tc.expectedPathToPage, pathToPage)
			}
		})
	}
	if len(uuids)!=3 || len(shortUrls)!=3{
		t.Errorf("Expected unique urls and uuids, got len(sets):\nlen(UUIDS) = %d\nlen(URLs) = %d\n", len(uuids), len(shortUrls))
	}
}
