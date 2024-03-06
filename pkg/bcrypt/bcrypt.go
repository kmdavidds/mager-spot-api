package bcrypt

import "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
}

type crypt struct {
	cost int
}

func Init() Interface {
	return &crypt{
		cost: 10,
	}
}

func (c *crypt) GenerateFromPassword(password string) (string, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), c.cost)
	if err != nil {
		return "", err
	}

	return string(passwordByte), nil
}

func (c *crypt) CompareHashAndPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
