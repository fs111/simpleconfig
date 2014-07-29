package simpleconfig

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_NewConfig(t *testing.T) {

	c, err := NewConfig("/dev/null")
	if err != nil {
		t.Error("creating new config struct failed")
	}
	if c == nil {
		t.Error("creating new config struct failed")
	}
}

func Test_NewConfigParsing(t *testing.T) {

	tmpFile := MakeConfigFile("\n#answer=42\nfoo=bar\n\n")
	defer os.Remove(tmpFile.Name())

	c, err := NewConfig(tmpFile.Name())
	if err != nil {
		t.Error("creating new config struct failed")
	}
	if c == nil {
		t.Error("creating new config struct failed")
	}

	answer, e := c.GetInt("answer")
	if e == nil {
		t.Error("parsing of comments does not work")
	}
	if answer != -1 {
		t.Error("got an unexpexted value")
	}

	foo, _ := c.GetString("foo")
	if foo != "bar" {
		t.Error("the value of foo should be bar")
	}
}

func Test_GetStringEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	_, err := c.GetString("test")
	if err == nil {
		t.Error("an error for an non existing key should have been returned")
	}
}

func Test_GetStringDefaultEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	defaultValue := "default"
	val := c.GetStringDefault("test", defaultValue)
	if val != defaultValue {
		t.Error("the default value should have been returned")
	}
}

func Test_GetStringWithExistingKeys(t *testing.T) {
	tmpFile := MakeConfigFile("foo=bar\nbla=1")
	defer os.Remove(tmpFile.Name())

	c, _ := NewConfig(tmpFile.Name())
	val, _ := c.GetString("foo")
	if val != "bar" {
		t.Error("the key 'foo' does not have the value 'bar'")
	}
	val, _ = c.GetString("bla")
	if val != "1" {
		t.Error("the key 'bla' does not have the value '1'")
	}
}

func Test_GetIntEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	_, err := c.GetInt("test")
	if err == nil {
		t.Error("an error for an non existing key should have been returned")
	}
}

func Test_GetIntDefaultEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	defaultValue := 42
	val := c.GetIntDefault("test", defaultValue)
	if val != defaultValue {
		t.Error("the default value should have been returned")
	}
}

func Test_GetInt(t *testing.T) {

	tmpFile := MakeConfigFile("answer=42")
	defer os.Remove(tmpFile.Name())

	c, _ := NewConfig(tmpFile.Name())
	val, _ := c.GetInt("answer")
	if val != 42 {
		t.Error("the answer is 42")
	}
}

func Test_GetBoolEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	val, err := c.GetBool("somekey")
	if err == nil {
		t.Error("the boolean setting should not exist")
	}
	if val {
		t.Error("the value must be false")
	}
}


func Test_GetStringTrim(t *testing.T) {
	tmpFile := MakeConfigFile("val = 42abc")
	defer os.Remove(tmpFile.Name())
	c, _ := NewConfig(tmpFile.Name())
    val, _ := c.GetString("val")
	if val != "42abc" {
		t.Error("the value of foo should be bar")
    }
}


func Test_GetBool(t *testing.T) {
	tmpFile := MakeConfigFile("bool1=1\nbool2=0\nbool3=true\nbool4=false")
	defer os.Remove(tmpFile.Name())

	c, _ := NewConfig(tmpFile.Name())

	val, err := c.GetBool("bool1")
	if err != nil {
		t.Error("no error should be returned")
	}
	if !val {
		t.Error("the value must be true")
	}
	val, err = c.GetBool("bool3")
	if err != nil {
		t.Error("no error should be returned")
	}
	if !val {
		t.Error("the value must be true")
	}
	val, err = c.GetBool("bool2")
	if err != nil {
		t.Error("no error should be returned")
	}
	if val {
		t.Error("the value must be false")
	}
	val, err = c.GetBool("bool4")
	if err != nil {
		t.Error("no error should be returned")
	}
	if val {
		t.Error("the value must be false")
	}
}

func Test_GetBoolDefaultEmptyFile(t *testing.T) {
	c, _ := NewConfig("/dev/null")
	val := c.GetBoolDefault("somekey", true)
	if !val {
		t.Error("the value must be true")
	}
	val = c.GetBoolDefault("somekey", false)
	if val {
		t.Error("the value must be false")
	}
}

func MakeConfigFile(content string) *os.File {
	tmpFile, _ := ioutil.TempFile("", "simpleconfig")
	tmpFile.WriteString(content)
	tmpFile.Close()
	return tmpFile
}
