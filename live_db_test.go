package talon_test

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/bbuck/talon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file contains tests that require an actual connection to a Neo4j instance.
// Communication on authentication/user/pass/host/port details should be done
// via the env variables listed below: (possible values are listed as a list)
//   TALON_PERFORM_LIVE_TEST = [1,0]
//   TALON_LIVE_TEST_AUTHENTICATED = [1,0]
//   TALON_LIVE_TEST_USER = <string>
//   TALON_LIVE_TEST_PASS = <string>
//   TALON_LIVE_TEST_HOST = <string> -- default: localhost
//   TALON_LIVE_TEST_PORT = <uint16>

const (
	defaultHost        = "localhost"
	defaultUser        = "neo4j"
	defaultPort uint16 = 7687
)

var (
	performLiveTest       bool
	liveTestAuthenticated bool
	liveTestUser          string
	liveTestPassword      string
	liveTestHost          string
	liveTestPort          uint16
)

var _ = Describe("LiveDb", func() {
	loadLiveTestEnvVariables()

	if !performLiveTest {
		fmt.Println("Skipping live test. Set TALON_PERFORM_LIVE_TEST to 1 to execute live tests.")

		return
	}

	co := ConnectOptions{
		User: liveTestUser,
		Host: liveTestHost,
		Port: liveTestPort,
	}

	if liveTestAuthenticated {
		co.Pass = liveTestPassword
	}

	Describe("Connecting", func() {
		var err error

		BeforeEach(func() {
			_, err = co.Connect()
		})

		It("doesn't fail", func() {
			Ω(err).Should(BeNil())
		})
	})

	Context("when connected, without a pool", func() {
		var (
			db  *DB
			err error
		)

		BeforeEach(func() {
			db, _ = co.Connect()
		})

		Describe("Cypher/CypherP", func() {
			Context("when making a query", func() {
				BeforeEach(func() {
					_, err = db.Cypher("MATCH (n) RETURN n").Query()
				})

				It("should not fail", func() {
					Ω(err).Should(BeNil())
				})
			})

			Context("single node", func() {
				It("allows creating, accessing and deleting", func() {
					By("creating a node")

					result, err := db.Cypher(`CREATE (:TalonSingleNodeTest {hello: "world"})`).Exec()

					Ω(err).Should(BeNil())
					Ω(result.Stats.LabelsAdded).Should(Equal(int64(1)))
					Ω(result.Stats.NodesCreated).Should(Equal(int64(1)))
					Ω(result.Stats.PropertiesSet).Should(Equal(int64(1)))

					By("accessing the node")

					rows, err := db.Cypher(`MATCH (n:TalonSingleNodeTest) RETURN n`).Query()
					defer rows.Close()

					Ω(err).Should(BeNil())

					debug("rows.Metadata()", rows.Metadata())
					a, b, c := rows.NextNeo()
					debug("a", a)
					debug("b", b)
					debug("c", c)

					a, b, c = rows.NextNeo()
					debug("a2", a)
					debug("b2", b)
					debug("c2", c)

					// TODO: Delete nodes
				})
			})
		})
	})
})

func debug(lbl string, v interface{}) {
	fmt.Printf("\n\nDEBUGGING %q ->\n\n%+v\n\n----------\n\n", lbl, v)
}

func loadLiveTestEnvVariables() {
	if val, ok := os.LookupEnv("TALON_PERFORM_LIVE_TEST"); ok {
		performLiveTest = val == "1"
	}

	if val, ok := os.LookupEnv("TALON_LIVE_TEST_AUTHENTICATED"); ok {
		liveTestAuthenticated = val == "1"
	}

	if val, ok := os.LookupEnv("TALON_LIVE_TEST_USER"); ok {
		liveTestUser = val
	} else {
		liveTestUser = defaultUser
	}

	if val, ok := os.LookupEnv("TALON_LIVE_TEST_PASS"); ok {
		liveTestPassword = val
	}

	if val, ok := os.LookupEnv("TALON_LIVE_TEST_HOST"); ok {
		liveTestHost = val
	} else {
		liveTestHost = defaultHost
	}

	if val, ok := os.LookupEnv("TALON_TEST_PORT"); ok {
		if ui, err := strconv.ParseUint(val, 10, 16); err != nil {
			fmt.Println("Failed to parse TALON_LIVE_TEST_PORT")
			liveTestPort = defaultPort
		} else {
			liveTestPort = uint16(ui)
		}
	} else {
		liveTestPort = defaultPort
	}
}
