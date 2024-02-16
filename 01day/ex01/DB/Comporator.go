package DB

import (
	"fmt"
	"os"
)

type checkIngredients struct {
	oldIngr		Ingredients
	freshIngr	Ingredients
	CakeName	string
	comp		int8
}

type check struct {
	oldCake		Cake
	freshCake	Cake
	comp		int8
}

func CompareIngr(old *Ingredients, fresh *Ingredients, CakeName string) {
	if old.Count != fresh.Count {
		fmt.Fprintf(os.Stdout, "CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", old.Name, CakeName, fresh.Count, old.Count)
	}
	if old.Unit != fresh.Unit {
		if old.Unit == "" {
		fmt.Fprintf(os.Stdout, "ADDED unit for ingredient \"%s\" for cake \"%s\" - \"%s\"\n", old.Name, CakeName, fresh.Unit)
		} else if fresh.Unit == "" {
			fmt.Fprintf(os.Stdout, "REMOVED unit for ingredient \"%s\" for cake \"%s\" - \"%s\"\n", old.Name, CakeName, old.Unit)
		} else {
			fmt.Fprintf(os.Stdout, "CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", old.Name, CakeName, fresh.Unit, old.Unit)
		}
	}
}

func CompareCakes(oldCake *Cake, freshCake *Cake) {
	if oldCake.Time != freshCake.Time {
		fmt.Fprintf(os.Stdout, "CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, freshCake.Time, oldCake.Time)
	}

	HM := make(map[string]*checkIngredients)
	for _, val := range oldCake.Ingredients {
		if _, ok := HM[val.Name]; !ok {
			HM[val.Name] = new(checkIngredients)
		}
		v := HM[val.Name]
		v.comp -= 1
		v.oldIngr = val
		v.CakeName = oldCake.Name
	}

	for _, val := range freshCake.Ingredients {
		if _, ok := HM[val.Name]; !ok {
			HM[val.Name] = new(checkIngredients)
		}
		v := HM[val.Name]
		v.comp += 1
		v.freshIngr = val
		v.CakeName = freshCake.Name
	}

	for k, v := range HM {
		if v.comp == -1 {
			fmt.Fprintf(os.Stdout, "REMOVED ingredient \"%s\" for cake \"%s\"\n", k, v.CakeName)
		} else if v.comp == 1 {
			fmt.Fprintf(os.Stdout, "ADDED ingredient \"%s\" for cake \"%s\"\n", k, v.CakeName)
		} else {
			defer CompareIngr(&v.oldIngr, &v.freshIngr, v.CakeName)
		}
	}
}

func CompareBook(old *CookBook, fresh *CookBook) {
	HM := make(map[string]*check)

	for _, val := range old.Cakes {
		if _, ok := HM[val.Name]; !ok {
			HM[val.Name] = new(check)
		}
		v := HM[val.Name]
		v.comp -= 1
		v.oldCake = val
	}

	for _, val := range fresh.Cakes {
		if _, ok := HM[val.Name]; !ok {
			HM[val.Name] = new(check)
		}
		v := HM[val.Name];
		v.comp += 1
		v.freshCake = val
	}

	for k, v := range HM {
		if v.comp == -1 {
			fmt.Fprintf(os.Stdout, "REMOVED cake \"%s\"\n", k)
		} else if v.comp == 1 {
			fmt.Fprintf(os.Stdout, "ADDED cake \"%s\"\n", k)
		} else {
			defer CompareCakes(&v.oldCake, &v.freshCake)
		}
	}
}
