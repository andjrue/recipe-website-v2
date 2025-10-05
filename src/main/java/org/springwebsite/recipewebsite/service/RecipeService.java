package org.springwebsite.recipewebsite.service;



import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;
import org.springwebsite.recipewebsite.repository.RecipeRepository;

import java.util.List;
import java.util.NoSuchElementException;

@Service
@Transactional(readOnly = true)
public class RecipeService {

    private final RecipeRepository recipeRepository;

    public RecipeService(RecipeRepository recipeRepository) {
        this.recipeRepository = recipeRepository;
    }

    public List<Recipe> findAll() {
        return this.recipeRepository.findAll();
    }

    public Recipe findById(String id) {
        return this.recipeRepository.findById(id)
                .orElseThrow(() ->
                        new NoSuchElementException(
                                String.format("Recipe with id: [%s] not found", id)
                        ));
    }

    public List<Recipe> findByUserId(User user) {
        List<Recipe> recipes = this.recipeRepository.findByUserId(user);

        if (recipes.isEmpty()) {
            throw new NoSuchElementException(String.format("Recipe with id: [%s] not found", user.getId()));
        }

        return recipes;
    }

    public List<Recipe> findByTitle(String title) {
        List<Recipe> recipes = this.recipeRepository.findByTitle(title);
        if (recipes.isEmpty()) {
            throw new NoSuchElementException(String.format("Recipe with title: [%s] not found", title));
        }

        return recipes;
    }

   public List<Recipe> findTenMostRecentRecipes() {
        return this.recipeRepository.findTop10ByOrderByDatePostedDesc();
   }



}
