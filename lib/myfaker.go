package lib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"syreclabs.com/go/faker"
)

func GenerateFake(t string) string {
	switch t {
	case "cardExpiry":
		return fmt.Sprintf("%d/%d", rand.Int31n(13), time.Now().Year()+rand.Intn(3))
	case "cardToken":
		return faker.Numerify("#########")
	case "creditCard":
		return strings.ReplaceAll(faker.Business().CreditCardNumber(), "-", "")
	case "cvv":
		return faker.Numerify("###")
	case "email":
		return faker.Internet().Email()
	case "ipv4":
		return faker.Internet().IpV4Address()
	case "ipv6":
		return faker.Internet().IpV6Address()
	case "ip":
		return fakeIP()
	case "name":
		return faker.Name().Name()
	case "title":
		return faker.Name().Title()
	case "password":
		return faker.Internet().Password(8, 14)
	case "phone":
		return faker.PhoneNumber().CellPhone()
	case "url":
		return faker.Internet().Url()
	case "username":
		return faker.Internet().UserName()
	default:
		log.Warnf("Faker %s not found. Using random string instead\n", t)
		return faker.RandomString(10)
	}
}

func fakeEmail() string {
	if rand.Intn(10) < 5 {
		return faker.Internet().Email()
	}
	return faker.Internet().FreeEmail()
}

func fakeIP() string {
	if rand.Intn(10) < 5 {
		return faker.Internet().IpV6Address()
	}
	return faker.Internet().IpV4Address()
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
