# Диаграмма классов

```mermaid
classDiagram
class App{
    Run() error
}

class Server{
    New(logger *slog.Logger, handler *handler.Handler) *Server
    Run() error
}
```