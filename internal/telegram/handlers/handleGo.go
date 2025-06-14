package handlers

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	answer "github.com/cOsm1cs/World-Cards-master/internal/telegram/service"
	"github.com/cOsm1cs/World-Cards-master/internal/logger"	
	tele "gopkg.in/telebot.v4"
)

type WordPair struct {
	Original  string
	Translate string
}

type QuizState struct {
	WordPairs    []WordPair // пара слов
	CurrentIndex int        // номер текущего слова
}

var(
	// Хранилище (словарь), где ключ - ID пользователя, значение - указатель на его состояние викторины
	userStates = make(map[int64]*QuizState)
	
	// Хранилище, где ключ - ID пользователя, значение - список его пары слов
	userWordPairs = make(map[int64][]WordPair)
	
	stateMutex sync.Mutex
)

// HandleGo - инициализирует викторину.
// 1. Получить ID пользователя и заблокировать мьютекс
// 2. Проверить наличие пользовательских слов и, если их нет, сообщить пользователю.
// 3. Скопировать пользовательские слова, перемешать их и создать новое состояние викторины
// 4. Сохранить состояние и отправить первое слово
func HandleGo(c tele.Context) error {
	userID := c.Sender().ID
	stateMutex.Lock()
	defer stateMutex.Unlock()

	pairs := getUserWordPairs(userID)
	if len(pairs) == 0 {
		return logger.SendWithLogError(c, answer.HandleGoNotAddWorks)
	}

	shuffledPairs, err := shuffleWordPairs(pairs)
	if err != nil {
		return err
	}

	state, err := createQuizState(shuffledPairs)
	if err != nil {
		return err
	}

	userStates[userID] = state

	if len(shuffledPairs) > 0 {
		return logger.SendWithLogError(c, shuffledPairs[0].Original)
	}

	return logger.SendWithLogError(c, "Больше нет слов")
}

func getUserWordPairs(userID int64) []WordPair {
	return userWordPairs[userID]
}

func shuffleWordPairs(pairs []WordPair) ([]WordPair, error) {
	if len(pairs) == 0 {
		return nil, nil
	}

	shuffled := make([]WordPair, len(pairs))
	copy(shuffled, pairs)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i] 
	})

	return shuffled, nil
}

func createQuizState(pairs []WordPair) (*QuizState, error) {
	return &QuizState{
		WordPairs: pairs,
		CurrentIndex: 0,
	}, nil
}

// HandleText обрабатывает ответы пользователя в викторине по переводу слов.
// Проверяет правильность ответа, отправляет обратную связь и переходит к следующему слову.
// При завершении викторины очищает состояние пользователя и предлагает начать заново.
func HandleText(c tele.Context) error {
	userID := c.Sender().ID
	stateMutex.Lock() // блокируемся, если несколько человек пользуются ботом
	defer stateMutex.Unlock()

	state, err := handleTextCheckingQuiz(c, userID)
	if err != nil {
		return err
	}

	// state.CurrentIndex - номер текущего слова
	// len(state.WordPairs) - общее количество слов
	// Если все слова закончились, предлагаем начать заново
	if state.CurrentIndex >= len(state.WordPairs) {
		delete(userStates, userID)
		return logger.SendWithLogError(c, answer.HandleGoEndQuiz)
	}

	// Берет текущую пару слов (например, "дом" и "house"), которую пользователь должен перевести.
	// Это слово, которое бот отправил ранее
	currentPair := state.WordPairs[state.CurrentIndex]
	
	// Берем текст пользователя. Убираем лишние пробелы и приводим к нижнему регистру
	// Это нужно, чтобы не путаться с большими и маленькими буквами
	userAnswer := strings.ToLower(strings.TrimSpace(c.Message().Text))
	
	// Аналогично нормализует правильный ответ (перевод слова), чтобы сравнение было честным
	currentAnswer := strings.ToLower(strings.TrimSpace(currentPair.Translate))

	if userAnswer == currentAnswer { // Сравнивает ответ пользователя с правильным ответом
		_ = logger.SendWithLogError(c, "Правильно")
	} else {
		_ = logger.SendWithLogError(c, "Неправильно!\nПравильный ответ: " + currentPair.Translate)
	}

	state.CurrentIndex++ // Увеличиваем номер слова на 1, чтобы перейти к следующему слову
	// Проверяем, остались ли еще слова в викторине
	if state.CurrentIndex < len(state.WordPairs) {
		// Если остались, отправляем следующее слово
		return logger.SendWithLogError(c, state.WordPairs[state.CurrentIndex].Original)
	}

	// Если не остались
	// delete(userStates, userID) // То удаляем данные викторины этого пользователя, чтобы не занимать место
	return logger.SendWithLogError(c, answer.HandleGoRepeatQuiz)
}

func handleTextCheckingQuiz(c tele.Context, userID int64) (*QuizState, error) {
	state, exists := userStates[userID] // проверяем, есть ли викторина
	if !exists {                        // если викторины нет выводим сообщение и завершаем
		return nil, logger.SendWithLogError(c, answer.HandleGoStartQuiz)
	}

	return state, nil
}
