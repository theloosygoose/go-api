package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
    "gusti": {
        AuthToken: "123ABC",
        Username: "theloosygoose",
    },
    "dad": {
        AuthToken: "456DEF",
        Username: "dad",
    },
    "dewi": {
        AuthToken: "789GHI",
        Username: "dewi",
    },
}

var mockCoinDetils = map[string]CoinDetails{
    "gusti": {
        Coins:  100,
        Username: "theloosygoose",
    },
    "dad": {
        Coins: 200,
        Username: "dad",
    },
    "dewi": {
        Coins: 300,
        Username: "dewi",
    },
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
    time.Sleep(time.Second * 2)

    var clientData = LoginDetails {}
    clientData, ok := mockLoginDetails[username]
    if !ok {
        return nil
    }

    return &clientData
}

func (d *mockDB) GetUsercoins(username string) *CoinDetails {
    time.Sleep(time.Second * 2)

    var clientData  = CoinDetails{}
    clientData, ok := mockCoinDetils[username]
    if !ok {
        return nil
    }
    
    return &clientData
}

func (d *mockDB) SetupDatabase() error {
    return nil
}
