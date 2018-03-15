package presenters_test

import (
	. "github.com/alphagov/paas-metric-exporter/presenters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MyStruct struct {
    Foo string
    Bar string
}

var _ = Describe("PathPresenter", func() {
	Describe("#Present", func() {
        It("should present the data according to the template", func() {
            presenter := PathPresenter{Template: "{{.Foo}}-{{.Bar}}"}
            data := MyStruct{Foo: "foo", Bar: "bar"}
            output, err := presenter.Present(data)

            Expect(err).NotTo(HaveOccurred())
            Expect(output).To(Equal("foo-bar"))
        })

        It("should fail to present the data due to lack of dot", func() {
            presenter := PathPresenter{Template: "{{Foo}}"}
            data := MyStruct{Foo: "foo", Bar: "bar"}
            _, err := presenter.Present(data)

            Expect(err).To(HaveOccurred())
        })

        It("should fail to present the data due to unknown property in template", func() {
            presenter := PathPresenter{Template: "{{.Missing}}"}
            data := MyStruct{Foo: "foo", Bar: "bar"}
            _, err := presenter.Present(data)

            Expect(err).To(HaveOccurred())
        })
	})
})
