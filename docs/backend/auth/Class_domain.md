# Диаграмма классов

```mermaid
classDiagram

class User { 
    <<Entity>>
    ID int

    Login string
    Password string
}

class Updates{
    
}

class Repository {
    Create(context.Context, *User) error
    Update(context.Context, Updates) (*User, error)
 }

class IUserUsecase {
    <<Interface>>

    +CreateUser(context.Context, Input) (Output, error)
    +VerifyUser(context.Context, Input) error
 }


class UserDTO{
    Login string
    Password string
}

class CreatedUserDTO{
    ID int
    Login string
    Password string
}

class Usecase { 
    
    +CreateUser(context.Context, UserDTO) (CreatedUserDTO, error)
    +VerifyUser(context.Context, UserDTO) error

}

Repository --> Updates : use

Usecase --> User : use

Usecase --> Repository : depents

Usecase ..|> IUserUsecase : implements

Usecase --> UserDTO : use
Usecase --> CreatedUserDTO : use
