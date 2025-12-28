package mapper

import (
	"testing"

	"github.com/imkarthi24/sf-backend/internal/entities"
	requestModel "github.com/imkarthi24/sf-backend/internal/model/request"
	"github.com/stretchr/testify/require"
)

var ChitgroupID = 25
var CustomerID = 12
var homeBranch_CBE = "CBE"

var discountAmount100 = "100"  // decimal.NewFromInt(100),
var payableAmount9800 = "9800" // decimal.NewFromInt(9800),
var balanceAmount4000 = "4000" // decimal.NewFromInt(5000),
var actualAmount5000 = "5000"  // decimal.NewFromInt(5000),
var paymentMethodCASH = "CASH"

var dueDate = "2022-11-21" // time.Date(2022, time.November, 21, 0, 0, 0, 0, time.UTC),

func Test_User(t *testing.T) {

	mapper := ProvideMapper()

	reqModelInput := requestModel.User{
		ID:                  0,
		IsActive:            false,
		FirstName:           "K",
		LastName:            "L",
		Extension:           "9",
		PhoneNumber:         "",
		Email:               "",
		Password:            "",
		Role:                "",
		IsLoginDisabled:     false,
		IsLoggedIn:          false,
		LastLoginTime:       "",
		LoginFailureCounter: 0,
	}
	actualResult, err := mapper.User(reqModelInput)

	expectedResult := &entities.User{
		Model:               &entities.Model{},
		FirstName:           "K",
		LastName:            "L",
		Extension:           "9",
		PhoneNumber:         "",
		Email:               "",
		Password:            "",
		Role:                "",
		IsLoginDisabled:     false,
		IsLoggedIn:          false,
		LastLoginTime:       nil,
		LoginFailureCounter: 0,
		ResetPasswordString: nil,
	}

	require.Nil(t, err)
	require.Equal(t, expectedResult, actualResult)

}

func Test_User_Negative(t *testing.T) {

	mapper := ProvideMapper()

	reqModelInput := requestModel.User{
		ID:                  0,
		IsActive:            false,
		FirstName:           "K",
		LastName:            "L",
		Extension:           "9",
		PhoneNumber:         "",
		Email:               "useremail@gmail.com",
		Password:            "password",
		Role:                "",
		IsLoginDisabled:     false,
		IsLoggedIn:          false,
		LastLoginTime:       "",
		LoginFailureCounter: 0,
	}
	actualResult, err := mapper.User(reqModelInput)

	expectedResult := &entities.User{
		Model:               &entities.Model{},
		FirstName:           "K",
		LastName:            "L",
		Extension:           "9",
		PhoneNumber:         "",
		Email:               "",
		Password:            "",
		Role:                "",
		IsLoginDisabled:     false,
		IsLoggedIn:          false,
		LastLoginTime:       nil,
		LoginFailureCounter: 0,
		ResetPasswordString: nil,
	}

	require.Nil(t, err)
	require.NotEqual(t, expectedResult, actualResult)

}
