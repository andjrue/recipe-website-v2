package recipes

type Recipe struct {
    Title string 
    Description string 
    Ingredients string 
    TimeToMake string 
    Steps string 
}

func newRecipe(title, descrip, ingre, ttm, steps string) *Recipe {
    return &Recipe{
        Title: title,
        Description: descrip,
        Ingredients: ingre,
        TimeToMake: ttm,
        Steps: steps,
    }
}


