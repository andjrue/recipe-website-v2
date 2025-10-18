package org.springwebsite.recipewebsite.service;



import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;
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

    public List<Recipe> findByUserId(String userId) {

        User user = new User();
        user.setId(userId);

        return this.recipeRepository.findByUserIdAndIsDeletedFalse(user);
    }

    public List<Recipe> findByTitle(String title) {
        return this.recipeRepository.findByTitleContainingIgnoreCaseAndIsDeletedFalse(title);
    }

   public List<Recipe> findTenMostRecentRecipes() {
        return this.recipeRepository.findTop10ByOrderByDatePostedDesc();
   }

   @Transactional
    public Recipe save(Recipe recipe, String userId) {
        if (recipe == null || !StringUtils.hasText(recipe.getTitle())) {
            throw new IllegalArgumentException("Recipe title is required");
        }

        User creator = new User();
        creator.setId("1");

//        if (!StringUtils.hasText(creator.getId())) {
//            throw new IllegalArgumentException("Recipe creator is required");
//        }

        recipe.setUserId(creator);
        recipe.setDatePosted(java.time.LocalDateTime.now().toString());

        return this.recipeRepository.save(recipe);
   }

   @Transactional
    public Recipe update(Recipe recipe, String userId) {
        if (recipe == null || !StringUtils.hasText(recipe.getId())) {
            throw new IllegalArgumentException("Recipe is required");
        }

        Recipe existingRecipe = this.findById(recipe.getId());

        User editor = new User();
        editor.setId(userId);

        if (!StringUtils.hasText(editor.getId())) {
            throw new IllegalArgumentException("Recipe editor is required");
        }

        if (!existingRecipe.getUserId().getId().equals(editor.getId())) {
            throw new SecurityException("User must match the original creator");
        }

        existingRecipe.setLastEditedAt(java.time.LocalDateTime.now().toString());
        existingRecipe.setDescription(recipe.getDescription());
        existingRecipe.setTitle(recipe.getTitle());
        existingRecipe.setTimeToCook(recipe.getTimeToCook());

        return this.recipeRepository.save(existingRecipe);
   }

    @Transactional
    public void deleteById(String id, String userId) {
        // Not actually deleting, only soft deleting
        if (id == null) {
            throw new IllegalArgumentException("Recipe id is required");
        }

        User editor = new User();
        editor.setId(userId);

        if (!StringUtils.hasText(editor.getId())) {
            throw new IllegalArgumentException("Recipe editor is required");
        }

        Recipe existingRecipe = this.findById(id);

        if (!editor.getId().equals(existingRecipe.getUserId().getId())) {
            throw new SecurityException("User must match the original editor");
        }

        existingRecipe.setLastEditedAt(java.time.LocalDateTime.now().toString());
        existingRecipe.setDeleted(true);

        this.recipeRepository.save(existingRecipe);
    }

}
