package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random int64 between min and max
// If an optional *rand.Rand instance is provided, it is used to generate the random integer.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomInt(min, max int64, r ...*rand.Rand) int64 {
	var randInstance *rand.Rand
	if len(r) > 0 {
		randInstance = r[0]
	} else {
		randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return min + randInstance.Int63n(max-min+1)
}

// RandomTwoOrThreeDigitInt generates a random uint8 between 10 and 999
// If an optional *rand.Rand instance is provided, it is used to generate the random integer.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomTwoOrThreeDigitInt(r ...*rand.Rand) uint8 {
	var randInstance *rand.Rand
	if len(r) > 0 {
		randInstance = r[0]
	} else {
		randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return uint8(randInstance.Intn(990) + 10)
}

// RandomString generates a random string of length n
// If an optional *rand.Rand instance is provided, it is used to generate the random string.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomString(n int, r ...*rand.Rand) string {
	var randInstance *rand.Rand
	if len(r) > 0 {
		randInstance = r[0]
	} else {
		randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[randInstance.Intn(len(charset))]
	}

	return string(b)
}

// RandomUsername generates a random username
// If an optional *rand.Rand instance is provided, it is used to generate the random username.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomUsername(r ...*rand.Rand) string {
	return RandomString(8, r...)
}

// RandomEmail generates a random email
// If an optional *rand.Rand instance is provided, it is used to generate the random email.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomEmail(r ...*rand.Rand) string {
	return RandomString(8, r...) + "@" + RandomString(4, r...) + ".com"
}

// RandomPasswordHash generates a random password hash
// If an optional *rand.Rand instance is provided, it is used to generate the random password hash.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomPassword(r ...*rand.Rand) string {
	return RandomString(8, r...)
}

// func RandomCountryCode(r ...*rand.Rand) string {
// 	// List of ISO 3166-1 alpha-2 country codes
// 	countryCodes := []string{
// 		"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ",
// 		"BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ",
// 		"CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ",
// 		"DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ",
// 		"LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ",
// 		"NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ",
// 		"TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW",
// 	}

// 	// Use the provided rand.Rand instance if available, otherwise create a new one.
// 	var randInstance *rand.Rand
// 	if len(r) > 0 {
// 		randInstance = r[0]
// 	} else {
// 		randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
// 	}

// 	// Generate a random index within the bounds of the countryCodes slice.
// 	randomIndex := randInstance.Intn(len(countryCodes))

// 	// Return the country code at the random index.
// 	return countryCodes[randomIndex]
// }
// RandomCountryCode generates a random country code
// If an optional *rand.Rand instance is provided, it is used to generate the random country code.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomCountryCode(r ...*rand.Rand) string {
	// return RandomString(2, r...)
	return "US" // <- i know ... but it's just a test
}

// RandomDate generates a random date within the last 100 years with time part set to 00:00:00 UTC
// If an optional *rand.Rand instance is provided, it is used to generate the random date.
// Otherwise, a new *rand.Rand instance is created using the current time as a seed.
func RandomDate(r ...*rand.Rand) time.Time {
	var randInstance *rand.Rand
	if len(r) > 0 {
		randInstance = r[0]
	} else {
		randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	startDate := time.Now().AddDate(-100, 0, 0)
	delta := time.Now().Unix() - startDate.Unix()
	sec := randInstance.Int63n(delta)

	return startDate.Add(time.Duration(sec) * time.Second).UTC().Truncate(24 * time.Hour)
}
