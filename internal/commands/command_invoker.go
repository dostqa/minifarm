package commands

var (
	DefaultInvoker Invoker
)

// Запускает выполнение команд
// Попавших к нему в очередь на выполнение
type Invoker struct {
	queue []Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.queue = append(i.queue, cmd)
}

func (i *Invoker) ExecuteCommmands() {
	for _, cmd := range i.queue {
		cmd.Execute()
	}
	i.queue = i.queue[:0]
}
