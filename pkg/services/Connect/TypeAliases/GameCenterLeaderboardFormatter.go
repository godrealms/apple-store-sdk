package TypeAliases

// GameCenterLeaderboardFormatter The values you can select to describe the format of a leaderboard.
type GameCenterLeaderboardFormatter string

const (
	GameCenterLeaderboardFormatterInteger                GameCenterLeaderboardFormatter = "INTEGER"
	GameCenterLeaderboardFormatterDecimalPoint1Place     GameCenterLeaderboardFormatter = "DECIMAL_POINT_1_PLACE"
	GameCenterLeaderboardFormatterDecimalPoint2Place     GameCenterLeaderboardFormatter = "DECIMAL_POINT_2_PLACE"
	GameCenterLeaderboardFormatterDecimalPoint3Place     GameCenterLeaderboardFormatter = "DECIMAL_POINT_3_PLACE"
	GameCenterLeaderboardFormatterElapsedTimeMillisecond GameCenterLeaderboardFormatter = "ELAPSED_TIME_MILLISECOND"
	GameCenterLeaderboardFormatterElapsedTimeCentisecond GameCenterLeaderboardFormatter = "ELAPSED_TIME_CENTISECOND"
	GameCenterLeaderboardFormatterElapsedTimeMinute      GameCenterLeaderboardFormatter = "ELAPSED_TIME_MINUTE"
	GameCenterLeaderboardFormatterElapsedTimeSecond      GameCenterLeaderboardFormatter = "ELAPSED_TIME_SECOND"
	GameCenterLeaderboardFormatterMoneyPoundDecimal      GameCenterLeaderboardFormatter = "MONEY_POUND_DECIMAL"
	GameCenterLeaderboardFormatterMoneyPound             GameCenterLeaderboardFormatter = "MONEY_POUND"
	GameCenterLeaderboardFormatterMoneyDollarDecimal     GameCenterLeaderboardFormatter = "MONEY_DOLLAR_DECIMAL"
	GameCenterLeaderboardFormatterMoneyDollar            GameCenterLeaderboardFormatter = "MONEY_DOLLAR"
	GameCenterLeaderboardFormatterMoneyEuroDecimal       GameCenterLeaderboardFormatter = "MONEY_EURO_DECIMAL"
	GameCenterLeaderboardFormatterMoneyEuro              GameCenterLeaderboardFormatter = "MONEY_EURO"
	GameCenterLeaderboardFormatterMoneyFrancDecimal      GameCenterLeaderboardFormatter = "MONEY_FRANC_DECIMAL"
	GameCenterLeaderboardFormatterMoneyFranc             GameCenterLeaderboardFormatter = "MONEY_FRANC"
	GameCenterLeaderboardFormatterMoneyKronerDecimal     GameCenterLeaderboardFormatter = "MONEY_KRONER_DECIMAL"
	GameCenterLeaderboardFormatterMoneyKroner            GameCenterLeaderboardFormatter = "MONEY_KRONER"
	GameCenterLeaderboardFormatterMoneyYen               GameCenterLeaderboardFormatter = "MONEY_YEN"
)

func (gc GameCenterLeaderboardFormatter) String() string {
	return string(gc)
}
