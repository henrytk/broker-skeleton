package broker_test

import (
	"strings"

	. "github.com/henrytk/broker-skeleton/broker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		err               error
		config            Config
		validConfigSource = `
			{
				"basic_auth_username":"admin",
				"basic_auth_password":"1234"
			}
		`
	)

	BeforeEach(func() {
		config, err = NewConfig(strings.NewReader(validConfigSource))
		Expect(err).NotTo(HaveOccurred())
	})

	It("requires a basic auth username", func() {
		config.BasicAuthUsername = ""

		err = config.Validate()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Config error: basic auth username required"))
	})

	It("requires a basic auth password", func() {
		config.BasicAuthPassword = ""

		err = config.Validate()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Config error: basic auth password required"))
	})
})