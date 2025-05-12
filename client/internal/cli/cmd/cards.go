package cmd

import (
	"fmt"
	"time"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/cardsdto"
	"github.com/spf13/cobra"
)

var (
	cardsMethod    string
	cardID         int
	cardNumber     string
	cardCardholder string
	cardExpiry     string
	cardCVV        string
)

var cardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "Work with cards",
	Run: func(cmd *cobra.Command, args []string) {
		switch cardsMethod {
		case "post":
			fmt.Println(postCard())
		case "put":
			fmt.Println(putCard())
		case "get":
			fmt.Println(getCards())
		case "delete":
			fmt.Println(deleteCard())
		default:
			fmt.Println("Method is not allowed")
			fmt.Println("Use with post, put, get or delete")
		}
	},
}

func getCards() string {
	cards, err := cardsService.Fetch()
	if err != nil {
		return err.Error()
	}
	var ret string
	ret += "Cards:\n"
	for i := range cards {
		ret += cardToString(cards[i])
	}
	return ret
}

func postCard() string {
	timeLayout := "2006-01-02"
	var ret string
	request := cardsdto.PostRequest{}
	request.Number = cardNumber
	request.CardholderName = cardCardholder
	expiry, err := time.Parse(timeLayout, cardExpiry)
	if err != nil {
		return err.Error()
	}
	request.Exipre = expiry
	request.Metadata = accountMetadata
	card, err := cardsService.Create(request)
	if err != nil {
		return err.Error()
	}
	ret += "Create successfull\n"
	ret += cardToString(*card)
	return ret
}

func putCard() string {
	var ret string

	timeLayout := "2006-01-02"

	request := cardsdto.PutRequest{}
	request.Number = cardNumber
	request.CardholderName = cardCardholder
	expiry, err := time.Parse(timeLayout, cardExpiry)
	if err != nil {
		return err.Error()
	}
	request.Exipre = expiry
	request.Metadata = accountMetadata
	err = cardsService.Update(request)
	if err != nil {
		return err.Error()
	}
	ret += "Updated successful\n"
	return ret
}

func deleteCard() string {
	var ret string
	request := cardsdto.DeleteRequest{
		ID: cardID,
	}
	fmt.Println(request)
	err := cardsService.Delete(request)
	if err != nil {
		return err.Error()
	}
	ret += "Deleted successfull"
	return ret
}

func cardToString(card domain.CardData) string {
	ret := fmt.Sprintf("[%d] number: %s, cardholder: %s, expire: %s, cvv: %s, created_at: %s, updated_at: %s\n",
		card.ID,
		card.Number,
		card.CardholderName,
		card.Expire,
		card.CVV,
		card.CreatedAt,
		card.UpdatedAt,
	)
	return ret
}

func init() {
	rootCmd.AddCommand(cardsCmd)

	cardsCmd.Flags().StringVarP(&cardsMethod, "method", "m", "get", "method")
	cardsCmd.Flags().IntVarP(&cardID, "id", "i", 0, "id")
	cardsCmd.Flags().StringVarP(&cardNumber, "number", "n", "", "number")
	cardsCmd.Flags().StringVar(&cardCardholder, "cardholder", "", "cardholder")
	cardsCmd.Flags().StringVarP(&cardExpiry, "expiry", "e", "", "expiry")
	cardsCmd.Flags().StringVarP(&cardCVV, "cvv", "c", "", "cvv")
}
