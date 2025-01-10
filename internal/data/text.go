package data

type ParseMode int

const (
	None ParseMode = iota
	MarkdownV2
)

func (mode ParseMode) String() string {
	switch mode {
	case None:
		return ""
	case MarkdownV2:
		return "MarkdownV2"
	default:
		return "unknown"
	}
}

//TODO: Написать про новые команды

const TimeLayout = "15:04"
const DateLayout = "02.01.2006"

const URL = "https://api.telegram.org/bot"

const StartDescription = "Запускает бота и выводит служебную информацию"
const SetLoginDescription = "Позволяет изменить имя своего профиля"
const GetDiffDescription = "Позволяет узнать количество использованного трафика с времени последнего такого запроса"
const GetLoginDescription = "Позволяет узнать свой текущий логин"
const GetStatDescription = "Позволяет узнать статистику за всё время пользования"
const HelpDescription = "Печатает стартовое сообщение со справкой о командах бота"

const StartText =
"Привет\\! Я бот, который сообщает количество использованного трафика\\.\n\n" +
"С помощью команды `/set\\_login <login>` укажи имя своего профиля\\.\n\n" +
"С помощью команды `/get\\_diff` ты можешь узнать, сколько ты использовал" + 
"трафика с последнего использования\\."

const SuccessGetLoginText = 
"Текущий логин: %s"

const SuccessChangeLoginText = 
"Установленный профиль: %s"

const FailureChangeLoginText =
"Профиля %s не существует\\."

const SuccessGetDiffText = "С прошлого запроса (%s) твой расход трафика составил:\n" +
"Download: %s\nUpload: %s"

const FailureGetText = "Статистика не найдена. Возможно, ты не указал логин"

const SuccessGetStatText = "За всё время использования твой расход трафика составил:\n" +
"Download: %s\nUpload: %s"