package org.springwebsite.recipewebsite.service;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springwebsite.recipewebsite.entity.Comment;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;
import org.springwebsite.recipewebsite.repository.CommentRepository;

import java.util.List;

@Service
@Transactional(readOnly = true)
public class CommentService {

    private final CommentRepository commentRepository;

    public CommentService(CommentRepository commentRepository) {
        this.commentRepository = commentRepository;
    }

    public Comment save(Comment comment) {
        return commentRepository.save(comment);
    }

    public long countByUser(String userId) {

        User user = new User();
        user.setId(userId);

        return commentRepository.countByUser(user);
    }

    public List<Comment> findByUser(String userId) {

        User user = new User();
        user.setId(userId);

        return commentRepository.findByUser(user);

    }

    public List<Comment> findByRecipeId(String recipeId) {
        if (recipeId == null) {
            throw new IllegalArgumentException("recipeId cannot be null");
        }

        Recipe recipe = new Recipe();
        recipe.setId(recipeId);

        return commentRepository.findByRecipe(recipe.getId());
    }

    public List<Comment> findMostRecent() {
        return commentRepository.findTop10ByOrderByDatePostedDesc();
    }
}
