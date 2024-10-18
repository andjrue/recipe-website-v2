package recipes

import "github.com/andjrue/recipe-website-v2/internal/structs"


type Recipe struct {
    Title string 
    Description string 
    Ingredients string 
    TimeToMake string 
    Steps string 
}

type Server structs.Server

func NewRecipe(title, descrip, ingre, ttm, steps string) *Recipe {
    return &Recipe{
        Title: title,
        Description: descrip,
        Ingredients: ingre,
        TimeToMake: ttm,
        Steps: steps,
    }
}

func (s *Server) HandleAddRecipe() error {
    return nil
}

func (s *Server) HandleGetRecipes() error {
    return nil
}

