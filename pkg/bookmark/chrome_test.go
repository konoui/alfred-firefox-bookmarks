package bookmark

import (
	"testing"
)

func TestChromeBookmarks(t *testing.T) {
	tests := []struct {
		description  string
		bookmarkPath string
		want         Bookmarks
		expectErr    bool
	}{
		{
			description:  "correct bookmark file",
			bookmarkPath: "test-chrome-bookmarks.json",
			want: Bookmarks{
				&Bookmark{
					Folder: "",
					Title:  "Google",
					Domain: "www.google.com",
					URI:    "https://www.google.com/",
				},
				&Bookmark{
					Folder: "/1-hierarchy-a",
					Title:  "GitHub",
					Domain: "github.com",
					URI:    "https://github.com/",
				},
				&Bookmark{
					Folder: "/1-hierarchy-a/2-hierarchy-a/3-hierarchy-a",
					Title:  "Stack Overflow",
					Domain: "stackoverflow.com",
					URI:    "https://stackoverflow.com/",
				},
				&Bookmark{
					Folder: "/1-hierarchy-a/2-hierarchy-a/3-hierarchy-a",
					Title:  "Amazon Web Services",
					Domain: "aws.amazon.com",
					URI:    "https://aws.amazon.com/?nc1=h_ls",
				},
				&Bookmark{
					Folder: "/1-hierarchy-b",
					Title:  "Yahoo",
					Domain: "www.yahoo.com",
					URI:    "https://www.yahoo.com/",
				},
				&Bookmark{
					Folder: "/1-hierarchy-b/2-hierarchy-a",
					Title:  "Facebook",
					Domain: "www.facebook.com",
					URI:    "https://www.facebook.com/",
				},
				&Bookmark{
					Folder: "/1-hierarchy-b/2-hierarchy-a",
					Title:  "Twitter",
					Domain: "twitter.com",
					URI:    "https://twitter.com/login",
				},
				&Bookmark{
					Folder: "/1-hierarchy-b/2-hierarchy-b",
					Title:  "Amazon.com",
					Domain: "www.amazon.com",
					URI:    "https://www.amazon.com/",
				},
			},
			expectErr: false,
		},
		{
			description:  "invalid bookmark file",
			bookmarkPath: "test",
			expectErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			b := chromeBookmark{
				bookmarkPath: tt.bookmarkPath,
			}

			bookmarks, err := b.Bookmarks()
			if tt.expectErr && err == nil {
				t.Errorf("expect error happens, but got response")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error got: %+v", err.Error())
			}

			diff := DiffBookmark(bookmarks, tt.want)
			if !tt.expectErr && diff != "" {
				t.Errorf("unexpected response: (+want -got)\n%+v", diff)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		description  string
		bookmarkPath string
		expectErr    bool
	}{
		{
			description:  "correct bookmark file",
			bookmarkPath: "test-chrome-bookmarks.json",
			expectErr:    false,
		},
		{
			description:  "invalid bookmark file",
			bookmarkPath: "test",
			expectErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			b := chromeBookmark{
				bookmarkPath: tt.bookmarkPath,
			}

			err := b.unmarshal()
			if tt.expectErr && err == nil {
				t.Errorf("expect error happens, but got response")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error got: %+v", err.Error())
			}
		})
	}
}