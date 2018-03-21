package fasthttp_router

import (
	"testing"
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
)

func TestRouter(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "fasthttp-Router tests")
}
