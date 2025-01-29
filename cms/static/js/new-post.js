// Initialize TinyMCE
tinymce.init({
    selector: '#post-content',
    height: 500,
    menubar: true,
    plugins: [
      'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview',
      'anchor', 'searchreplace', 'visualblocks', 'code', 'fullscreen',
      'insertdatetime', 'media', 'table', 'help', 'wordcount'
    ],
    toolbar: 'undo redo | blocks | ' +
      'bold italic backcolor | alignleft aligncenter ' +
      'alignright alignjustify | bullist numlist outdent indent | ' +
      'removeformat | help',
    content_style: 'body { font-family: -apple-system, BlinkMacSystemFont, San Francisco, Segoe UI, Roboto, Helvetica Neue, sans-serif; font-size: 14px; }'
  });
  
  // Handle Tags
  document.addEventListener('DOMContentLoaded', () => {
    const tagInput = document.getElementById('tag-input');
    const tagsContainer = document.querySelector('.tags-container');
    const tags = new Set();
  
    function addTag(tag) {
      if (tag && !tags.has(tag)) {
        tags.add(tag);
        const tagElement = document.createElement('span');
        tagElement.className = 'tag';
        tagElement.innerHTML = `
          ${tag}
          <span class="tag-remove">×</span>
        `;
        
        tagElement.querySelector('.tag-remove').addEventListener('click', () => {
          tags.delete(tag);
          tagElement.remove();
        });
        
        tagsContainer.appendChild(tagElement);
      }
    }
  
    tagInput.addEventListener('keydown', (e) => {
      if (e.key === 'Enter') {
        e.preventDefault();
        const tag = tagInput.value.trim();
        addTag(tag);
        tagInput.value = '';
      }
    });
  
    // Handle Save Draft
    const draftButton = document.querySelector('.action-button.draft');
    draftButton.addEventListener('click', () => {
      const postData = {
        title: document.getElementById('post-title').value,
        excerpt: document.getElementById('post-excerpt').value,
        content: tinymce.get('post-content').getContent(),
        tags: Array.from(tags),
        status: 'draft'
      };
      console.log('Saving draft:', postData);
      // Here you would typically send the data to your backend
    });
  
    // Handle Publish
    const publishButton = document.querySelector('.action-button.publish');
    publishButton.addEventListener('click', () => {
      const postData = {
        title: document.getElementById('post-title').value,
        excerpt: document.getElementById('post-excerpt').value,
        content: tinymce.get('post-content').getContent(),
        tags: Array.from(tags),
        status: 'published'
      };
      console.log('Publishing post:', postData);
      // Here you would typically send the data to your backend
    });
  
    // Handle Image Upload Preview
    const imageInput = document.getElementById('featured-image');
    const uploadPlaceholder = document.querySelector('.upload-placeholder');
  
    imageInput.addEventListener('change', (e) => {
      const file = e.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          uploadPlaceholder.innerHTML = `
            <img src="${e.target.result}" alt="Preview" style="max-width: 100%; max-height: 200px; border-radius: 0.5rem;">
          `;
        };
        reader.readAsDataURL(file);
      }
    });
  });