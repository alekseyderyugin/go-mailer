package infrastructure

import (
	"fmt"
	"github.com/jordan-wright/email"
	"go-mailer/letters/domain"
	"sync"
)

type Worker struct {
	consecutiveErrors int
	mutex             *sync.Mutex
	waitGroup         *sync.WaitGroup
	repository        *LetterRepository
}

func NewWorker(mutex *sync.Mutex, wg *sync.WaitGroup, repository *LetterRepository) *Worker {
	return &Worker{
		consecutiveErrors: 0,
		mutex:             mutex,
		waitGroup:         wg,
		repository:        repository,
	}
}

func (w *Worker) resetErrors() {
	w.consecutiveErrors = 0
}

func (w *Worker) Run(task *Task) {
	for {
		letter := task.Next()

		if letter == nil {
			break
		}

		go w.sendLetter(letter, w.repository)
	}

	w.waitGroup.Wait()
}

func (w *Worker) sendLetter(letter *domain.Letter, repository *LetterRepository) {
	defer w.waitGroup.Done()

	w.mutex.Lock()
	if w.consecutiveErrors >= 5 {
		w.mutex.Unlock()
		panic("Too many consecutive errors")
	}
	w.mutex.Unlock()

	letter.Status = domain.Processing

	err := repository.Save(letter)

	if err != nil {
		w.onConsecutiveError(err)
	}

	err = send(letter)
	if err != nil {
		letter.Status = domain.SendFailed
		w.onConsecutiveError(err)
	} else {
		letter.Status = domain.Sent
	}

	err = repository.Save(letter)
	if err != nil {
		w.onConsecutiveError(err)
	} else {
		w.resetErrors()
	}
}

func (w *Worker) onConsecutiveError(err error) {
	fmt.Println(err)
	w.mutex.Lock()
	w.consecutiveErrors++
	w.mutex.Unlock()
}

func send(letter *domain.Letter) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", letter.FromName, letter.GetFrom())
	e.To = letter.GetTo()
	e.HTML = []byte(letter.GetHtmlMessage())
	err := e.Send("0.0.0.0:1029", nil)
	if err != nil {
		return err
	}

	return nil
}
