package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestForm_Valid tests the validity of the form
func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Form is not valid")
	}
}

// TestForm_Required tests the field requirement in forms
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid but required fields are missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form shows InValid but all fields are present")
	}

}

// TestForm_Has tests if a field is present inside a form or not
func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	x := form.Get("whatever")
	if x != "" {
		t.Error("Got a field value but it is not present in the form")
	}

	postData := url.Values{}
	postData.Add("test", "checking")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	x = form.Get("test")
	if x == "" {
		t.Error("Got a empty value but field is present in the form")
	}
}

// TestForm_MinLength tests the minimum length of a field inside a form
func TestForm_MinLength(t *testing.T) {
	length := 8
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	x := form.Get("whatever")
	if len(x) == length {
		t.Error("Field length is present but field doesn't exist")
	}

	postData := url.Values{}
	postData.Add("test", "checking")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	x = form.Get("test")
	if len(x) != length {
		t.Error("Got length different from expected length")
	}
}

// TestForm_IsEmail Check if the field value is EMAIL or not
func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.IsEmail("whatever")
	if form.Valid() {
		t.Error("Not an email but still passed")
	}

	postData := url.Values{}
	postData.Add("email", "rr@rr.com")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Email exists but still failed")
	}
}
