import {
  CMS
} from "./formHandlers.js";

import {
  formatDate,
  createModalWindow
} from "./utils.js";

document.addEventListener('DOMContentLoaded', () => {

  // Load Blog Posts
  const cms = new CMS();

  const loadBlogPosts = async () => {
    // Load blog posts from the CMS
    const result = await cms.GetBlogPosts();
    const posts = result.data.posts;

    const postsContainer = document.querySelector('.posts-grid');
    // postsContainer.innerHTML = '';

    posts.forEach(post => {
      const postID = post.id;
      const postAuthor = DOMPurify.sanitize(post.author_name);
      const postTitle = DOMPurify.sanitize(post.title);
      const postBody = DOMPurify.sanitize(post.body);
      const postExcerpt =
        postBody.length > 50 ? postBody.substring(0, 50) + '...' : postBody
      const postPublishedAt = DOMPurify.sanitize(post.published_at);
      const postStatus = post.is_draft ? 'Draft' : 'Published';
      const postViews = post.views;

      const postCard = document.createElement('div');

      const publishedDate = formatDate(postPublishedAt);

      postCard.classList.add('post-card');
      postCard.innerHTML = `
        <div class="post-status ${postStatus.toLowerCase()}">${postStatus}</div>
        <h2 class="post-title">${postTitle}</h2>
        <hr/>
        <div class="post-excerpt">${postExcerpt}</div>
        <div class="post-meta">
            <span class="post-date">${postStatus == "Draft" ? "Last edited" : "Published"} ${publishedDate}</span>
            <span class="post-author">By ${postAuthor}</span>
            <span class="post-views">${postStatus == "Draft" ? "📝 Draft" : " 👁️ " + postViews + " views"}</span>
        </div>
        <div class="post-actions">
          <button class="action-button edit">Edit</button>
          <button class="action-button delete">Delete</button>
        </div>
      `;
      postCard.dataset.postId = postID;
      // Handle the click event for the post card
      postCard.addEventListener('click', (event) => {
        // Ensure event target is not the edit or delete buttons
        if (event.target.classList.contains('action-button')) return;
        window.location.href = `/app/posts/${postID}`;
      });
      postsContainer.appendChild(postCard);
    });

    addHandlers();
    animatePostCards();
  };

  loadBlogPosts();

  function animatePostCards() {
    // Get all post card
    const post_cards = document.querySelectorAll('.posts-grid .post-card');

    // Loop through post cards and assign a custom --index property
    post_cards.forEach((post_card, index) => {
      post_card.style.setProperty('--index', index); // Set the delay index
    });

  }

  function addHandlers() {
    // Handle New Post button click
    const newPostButton = document.querySelector('.new-post-button');
    if (newPostButton) {
      newPostButton.addEventListener('click', () => {
        console.log('New post button clicked');
        // Here you would typically redirect to the post editor
        window.location.href = '/app/posts/new';
      });
    }

    // Handle Edit buttons
    document.querySelectorAll('.action-button.edit').forEach(button => {
      button.addEventListener('click', (e) => {
        const postCard = e.target.closest('.post-card');
        const postTitle = postCard.querySelector('.post-title').textContent;
        const postID = postCard.dataset.postId;
        // Redirect to the post editor with the post ID
        window.location.href = `/app/posts/${postID}/edit`;
      });
    });

    // Handle Delete buttons
    document.querySelectorAll('.action-button.delete').forEach(button => {
      button.addEventListener('click', (e) => {
        const postCard = e.target.closest('.post-card');
        const postID = postCard.dataset.postId;

        const postTitle = postCard.querySelector('.post-title').textContent;
        createModalWindow(`Are you sure you want to delete the post with title "${postTitle}"?`, async () => {
          // Delete the post
          const result = await cms.DeleteBlogPost(postID);
          if (result && result.status === 'success') {
            // Reload the page
            window.location.reload();

          } else {
            alert('Could not delete post');
          }
        }, () => {
          // Do nothing
        });
      });
    });
  }

});