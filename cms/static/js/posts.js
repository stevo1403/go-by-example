document.addEventListener('DOMContentLoaded', () => {
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
        console.log('Edit post:', postTitle);
        // Here you would typically redirect to the post editor with the post ID
      });
    });
  
    // Handle Delete buttons
    document.querySelectorAll('.action-button.delete').forEach(button => {
      button.addEventListener('click', (e) => {
        const postCard = e.target.closest('.post-card');
        const postTitle = postCard.querySelector('.post-title').textContent;
        if (confirm(`Are you sure you want to delete "${postTitle}"?`)) {
          console.log('Delete post:', postTitle);
          // Here you would typically send a delete request to your backend
        }
      });
    });
  });