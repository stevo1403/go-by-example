document.addEventListener('DOMContentLoaded', () => {
    // Handle Edit buttons
    document.querySelectorAll('.action-button.edit').forEach(button => {
      button.addEventListener('click', (e) => {
        const commentCard = e.target.closest('.comment-card');
        const commentContent = commentCard.querySelector('.comment-content').textContent;
        console.log('Edit comment:', commentContent);
        // Here you would typically open an edit modal or redirect to an edit page
      });
    });
  
    // Handle Delete buttons
    document.querySelectorAll('.action-button.delete').forEach(button => {
      button.addEventListener('click', (e) => {
        const commentCard = e.target.closest('.comment-card');
        const commentContent = commentCard.querySelector('.comment-content').textContent;
        if (confirm('Are you sure you want to delete this comment?')) {
          console.log('Delete comment:', commentContent);
          // Here you would typically send a delete request to your backend
        }
      });
    });
  });