package learning

import (
	"DEV_INVENTORY/internal/logics"
	"fmt"
)

func Learn_Choices(choice string) {
     logics.Guide_Choice(choice)
	 for _, choices := range choice {
           if choice == "1"  {
             fmt.Println(choices)
		   }
	 }
}