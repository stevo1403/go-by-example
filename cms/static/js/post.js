import {
    CMS,
    BlogPost
} from "./formHandlers.js";

import {
    createModalWindow,
    formatDate,
    showNotification,
} from "./utils.js"

async function loadBlogPost() {

    // Create a new CMS instance
    const cms = new CMS();

    // Extract the post ID from the URL
    const postID = window.location.pathname.split('/').pop();
    // Check if postID is a number using Regex
    if (!/^\d+$/.test(postID)) {
        // Redirect to the posts page
        window.location.href = '/app/posts';
    }

    let blog_post = null;
    
    // Fetch the post from the CMS
    const result = await cms.GetBlogPost(postID);
    if (result && result.status === 'success') {
        const blogPost = result.data.post;

        const blogPostTitle = blogPost.title;
        const blogPostBody = blogPost.body;
        const blogPostPublishedAT = blogPost.published_at;
        const blogPostViews = blogPost.views;
        const blogPostStatus = blogPost.is_draft ? 'Draft' : 'Published';
        let blogPostFeaturedImage = null;
        // const blogPostComments = blogPost.comments;
        // const blogPostLikes = blogPost.likes;

        // Fetch the blog post images
        const blogPostImages = await cms.GetBlogPostImages(postID);
        if (blogPostImages && blogPostImages.status === 'success') {
            const images = blogPostImages.data.images;
            // Extract the featured image
            for (let image of images) {
                if (image.image_type === 'preview') {
                    blogPostFeaturedImage = image.url;
                    break;
                }
            }
        }
        blog_post = new BlogPost(blogPostTitle, blogPostBody, blogPostPublishedAT, blogPostViews, 0, 0);

        blog_post.setPublishedStatus(blogPostStatus);
        blog_post.setImageUrl(blogPostFeaturedImage);
        blog_post.show();
    }

    handleCommentForm();
    handleCommentButtons();

    function handleCommentForm() {
        // Handle Comment Form
        const commentForm = document.getElementById('commentForm');
        // Check if the form has already been assigned a submit event
        if (commentForm.dataset.submitEvent) return;
        commentForm.dataset.submitEvent = true;
        commentForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const commentContent = document.getElementById('comment-content').value;
            console.log('New comment:', commentContent);
            console.log(postID);
            const postData = {
                "author_id": parseInt(cms.getUserID(), 10),
                "body": commentContent,
                "post_id": parseInt(postID, 10)
            }
            const result = await cms.CreateBlogPostComment(postData);
            if (result && result.status === 'success') {
                const commentsList = document.querySelector('#commentsContainer #commentsList');
                commentsList.appendChild(createCommentElement(result.data.comment));
                console.log('Comment created:', result.data.comment);
                document.getElementById('comment-content').value = '';
                blog_post.incrementCommentsCount();
            }
            else{
                showNotification('Could not create comment', 'error');
            }
        });
    }

    function handleDeleteComment(commentItem) {
        if (confirm('Are you sure you want to delete this comment?')) {
            console.log('Deleting comment');
            // Here you would typically send a delete request to your backend
            // and then remove the comment from the DOM
            commentItem.remove();
        }
    }

    function handleReplyComment(commentItem) {
        const authorName = commentItem.querySelector('.author-name').textContent;
        const commentContent = document.getElementById('comment-content');
        commentContent.value = `@${authorName} `;
        commentContent.focus();
    }

    function handleCommentButtons() {
        // Handle Reply Buttons
        document.querySelectorAll('.action-button.reply').forEach(button => {
            // Check if the button has already been assigned a click event
            if (button.dataset.clickEvent) return;
            button.dataset.clickEvent = true;
            button.addEventListener('click', (e) => {
                const commentItem = e.target.closest('.comment-item');
                handleReplyComment(commentItem);
            });
        });

        // Handle Delete Buttons
        document.querySelectorAll('.action-button.delete').forEach(button => {
            // Check if the button has already been assigned a click event
            if (button.dataset.clickEvent) return;
            button.dataset.clickEvent = true;
            button.addEventListener('click', async (e) => {
                const commentItem = e.target.closest('.comment-item');
                handleDeleteComment(commentItem);
            });
        });
    }

    function createCommentElement(comment) {
        const commentItem = document.createElement('div');
        commentItem.className = 'comment-item';
        commentItem.innerHTML = `
        <div class="comment-header">
            <div class="comment-author">
                <img src="/static/images/profile-image-32x32.svg" alt="User Avatar" class="comment-avatar" />
                <span class="author-name">${comment.author_name}</span>
            </div>
            <span class="comment-date">${formatDate(comment.published_at)}</span>

        </div>
        <div class="comment-content">${comment.body}</div>
        <div class="comment-actions">
            <button class="action-button reply">Reply</button>
            <button class="action-button delete">Delete</button>
        </div>`;
        return commentItem;
    }

    // Fetch the comments from the CMS
    const rComments = await cms.GetBlogPostComments(postID);
    if (rComments && rComments.status === 'success') {
        const comments = rComments.data.comments;

        const commentsList = document.querySelector('#commentsContainer #commentsList');
        commentsList.innerHTML = '';

        for (let comment of comments) {
            commentsList.appendChild(createCommentElement(comment));
            blog_post.incrementCommentsCount();
        }
    }
    handleCommentButtons();

}

async function main() {
    loadBlogPost();

    const deletePostButton = document.querySelector('.post .post-header .header-actions .action-button.delete');
    const editPostButton = document.querySelector('.post .post-header .header-actions .action-button.edit');

    deletePostButton.addEventListener('click', async () => {
        const cms = new CMS();

        const postID = window.location.pathname.split('/').pop();
        // Check if postID is a number using Regex
        if (!/^\d+$/.test(postID)) {
            // Show a nice alert
            alert('Could not delete post');
            return;
        }

        // Create a modal window confirming the delete action
        console.log('here');
        createModalWindow('Are you sure you want to delete this post?', async () => {
            // Delete the post
            const result = await cms.DeleteBlogPost(postID);
            if (result && result.status === 'success') {
                window.location.href = '/app/posts';
            } else {
                alert('Could not delete post');
            }
        }, () => {
            // Do nothing
        });

    });

    editPostButton.addEventListener('click', () => {
        const postID = window.location.pathname.split('/').pop();
        // Check if postID is a number using Regex
        if (!/^\d+$/.test(postID)) {
            // Show a nice alert
            alert('Could not edit post');
            return;
        }

        window.location.href = `/app/posts/${postID}/edit`
    });
}

// Run the main function when the DOM is loaded
document.addEventListener('DOMContentLoaded', main);