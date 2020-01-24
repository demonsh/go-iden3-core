package core

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-core/testgen"
	"github.com/stretchr/testify/assert"
)

func TestClaimEthId(t *testing.T) {
	// If generateTest is true, the checked values will be used to generate a test vector
	generateTest := true
	// Init test
	err := testgen.InitTest("claimEthId", generateTest)
	if err != nil {
		fmt.Println("error initializing test data:", err)
		return
	}
	// Add input data to the test vector
	if generateTest {
		testgen.SetTestValue("idEthAddr", "0xe0fbce58cfaa72812103f003adce3f284fe5fc7c")
		testgen.SetTestValue("idFactoryEthAddr", "0x66D0c2F85F1B717168cbB508AfD1c46e07227130")
	}
	ethId := common.HexToAddress(testgen.GetTestValue("idEthAddr").(string))
	identityFactoryAddr := common.HexToAddress(testgen.GetTestValue("idFactoryEthAddr").(string))

	c0 := NewClaimEthId(ethId, identityFactoryAddr)

	c1 := NewClaimEthIdFromEntry(c0.Entry())
	c2, err := NewClaimFromEntry(c0.Entry())
	assert.Nil(t, err)
	assert.Equal(t, c0, c1)
	assert.Equal(t, c0, c2)

	assert.Equal(t, c0.Address, ethId)
	assert.Equal(t, c0.IdentityFactory, identityFactoryAddr)
	assert.Equal(t, c0.Address, c1.Address)
	assert.Equal(t, c0.IdentityFactory, c1.IdentityFactory)

	assert.Equal(t, c0.Entry().Bytes(), c1.Entry().Bytes())
	assert.Equal(t, c0.Entry().Bytes(), c2.Entry().Bytes())

	e := c0.Entry()
	checkClaim(e, t)
	dataTestOutput(&e.Data)
	// Stop test (write new test vector if needed)
	err = testgen.StopTest()
	if err != nil {
		fmt.Println("Error stopping test:", err)
	}
}
