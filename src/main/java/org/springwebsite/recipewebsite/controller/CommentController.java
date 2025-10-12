package org.springwebsite.recipewebsite.controller;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springwebsite.recipewebsite.entity.Comment;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.service.CommentService;

import java.util.List;

@RestController
@RequestMapping("api/v1/comments")
public class CommentController {

    private final CommentService commentService;

    public CommentController(CommentService commentService) {
        this.commentService = commentService;
    }

    @GetMapping("/{recipeId}")
    public ResponseEntity<List<Comment>> findAllByRecipeId(@PathVariable("recipeId") String recipeId) {
        List<Comment> comments = this.commentService.findByRecipeId(recipeId);
        return ResponseEntity.ok(comments);
    }

}
