document.addEventListener('DOMContentLoaded', () => {
    // Handle Upload button
    const uploadButton = document.querySelector('.upload-button');
    if (uploadButton) {
      uploadButton.addEventListener('click', () => {
        console.log('Upload button clicked');
        // Here you would typically open a file upload dialog
      });
    }
  
    // Handle media type filter
    const typeFilter = document.getElementById('typeFilter');
    if (typeFilter) {
      typeFilter.addEventListener('change', (e) => {
        console.log('Filter by type:', e.target.value);
        // Here you would typically filter the media items
      });
    }
  
    // Handle date filter
    const dateFilter = document.getElementById('dateFilter');
    if (dateFilter) {
      dateFilter.addEventListener('change', (e) => {
        console.log('Filter by date:', e.target.value);
        // Here you would typically filter the media items
      });
    }
  
    // Handle overlay buttons
    document.querySelectorAll('.overlay-button').forEach(button => {
      button.addEventListener('click', (e) => {
        e.stopPropagation();
        const action = button.classList[1];
        const mediaCard = button.closest('.media-card');
        const fileName = mediaCard.querySelector('.media-name').textContent;
  
        switch (action) {
          case 'preview':
            console.log('Preview:', fileName);
            // Here you would typically open a preview modal
            break;
          case 'copy':
            console.log('Copy link:', fileName);
            // Here you would typically copy the media URL to clipboard
            break;
          case 'delete':
            if (confirm(`Are you sure you want to delete ${fileName}?`)) {
              console.log('Delete:', fileName);
              // Here you would typically send a delete request to your backend
            }
            break;
        }
      });
    });
  });