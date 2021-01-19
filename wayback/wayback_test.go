package wayback

import "testing"

func TestPayloadFromBytes(t *testing.T) {
	originalURL := "http://www.donaldjtrump.com/images/site/banner.jpg"
	closest, err := FindClosestURL(originalURL, false)
	if err != nil {
		t.Errorf("expecting ok")
	}
	if closest != "http://web.archive.org/web/20151030190157/http://www.donaldjtrump.com/images/site/banner.jpg" {
		t.Errorf("want != got, got: %s", closest)
	}
}

func TestPayloadFromBytesRaw(t *testing.T) {
	originalURL := "http://www.donaldjtrump.com/images/site/banner.jpg"
	closest, err := FindClosestURL(originalURL, true)
	if err != nil {
		t.Errorf("expecting ok")
	}
	if closest != "http://web.archive.org/web/20151030190157if_/http://www.donaldjtrump.com/images/site/banner.jpg" {
		t.Errorf("want != got, got: %s", closest)
	}
}
