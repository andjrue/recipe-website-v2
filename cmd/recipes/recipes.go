package recipes

import (
    "github.com/andjrue/recipe-website-v2/cmd/router.go"
)

type Recipe struct {
    Title string 
    Description string 
    Ingredients string 
    TimeToMake string 
    Steps string 
}

func NewRecipe(title, descrip, ingre, ttm, steps string) *Recipe {
    return &Recipe{
        Title: title,
        Description: descrip,
        Ingredients: ingre,
        TimeToMake: ttm,
        Steps: steps,
    }
}

func (s *main.Server) HandleAddRecipe() error {
    return nil
}

func (s *main.Server) HandleGetRecipes() error {
    return nil
}

