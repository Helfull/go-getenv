package env

import (
	"os"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestGetEnv(t *testing.T) {
	gunit.Run(new(LoadDotEnvFixture), t)
	gunit.Run(new(GetEnvStrFixture), t)
	gunit.Run(new(GetEnvBoolFixture), t)
}

type LoadDotEnvFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.
}

func (this *LoadDotEnvFixture) Setup() {
	os.Setenv("TEST_ENV_VAR", "true")
	LoadDotEnv("env.testing")
}

func (this *LoadDotEnvFixture) TeardownStuff() {
	os.Unsetenv("TEST_ENV_VAR")
	ClearDotEnv()
}

func (this *LoadDotEnvFixture) TestOsEnvironmentVariablesShouldTakePriority() {
	result, err := GetEnvStr("TEST_ENV_VAR", "false")
	this.So(result, should.Equal, "true")
	this.So(err, should.BeNil)
}

func (this *LoadDotEnvFixture) TestItShouldUseTheLoadedVariablesIfNoEnvironmentVariableIsSet() {
	result, err := GetEnvStr("ENV_VAR_TESTING", "true")
	this.So(result, should.Equal, "false")
	this.So(err, should.BeNil)
}

func (this *LoadDotEnvFixture) TestItShouldReturnTheFallbackIfTheVariableIsNotFoundInTheEnvironmentOrTheLoadedVariables() {
	result, err := GetEnvStr("ENV_VAR_TESTING_2", "false")
	this.So(result, should.Equal, "false")
	this.So(err, should.Equal, errEnvVarEmpty)
}

type GetEnvStrFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.
}

func (this *GetEnvStrFixture) Setup() {
	os.Setenv("ENV_TEST", "real value")
}

func (this *GetEnvStrFixture) TeardownStuff() {
	os.Unsetenv("ENV_TEST")
	ClearDotEnv()
}

func (this *GetEnvStrFixture) TestItShouldReturnTheValueOfTheEnvVariable() {
	result, err := GetEnvStr("ENV_TEST", "this.fallback")
	this.So(result, should.Equal, "real value")
	this.So(err, should.BeNil)
}

func (this *GetEnvStrFixture) TestItReturnAnErrorAndTheFallbackValueIfTheEnvVarDoesNotExist() {
	result, err := GetEnvStr("MISSING_123ENV123_VAR", "fallback")
	this.So(result, should.Equal, "fallback")
	this.So(err, should.Equal, errEnvVarEmpty)
}

type GetEnvBoolFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.
}

func (this *GetEnvBoolFixture) Setup() {
	os.Setenv("ENV_TEST", "true")
	os.Setenv("ENV_FAIL_BOOL", "ME NOT YAYA")
}

func (this *GetEnvBoolFixture) TeardownStuff() {
	os.Unsetenv("ENV_TEST")
	os.Unsetenv("ENV_FAIL_BOOL")
	ClearDotEnv()
}

func (this *GetEnvBoolFixture) TestItShouldReturnTheValueOfTheEnvVariable() {
	result, err := GetEnvBool("ENV_TEST", false)
	this.So(result, should.Equal, true)
	this.So(err, should.BeNil)
}

func (this *GetEnvBoolFixture) TestItReturnAnErrorAndTheFallbackValueIfTheEnvVarDoesNotExist() {
	result, err := GetEnvBool("MISSING_123ENV123_VAR", true)
	this.So(result, should.Equal, true)
	this.So(err, should.Equal, errEnvVarEmpty)
}

func (this *GetEnvBoolFixture) TestItShouldReturnAnErrorAndTheFallbackIfTheValueCouldNotBeParsedToAnBoolean() {
	result, err := GetEnvBool("ENV_FAIL_BOOL", false)
	this.So(result, should.Equal, false)
	this.So(err, should.NotBeNil)
}
