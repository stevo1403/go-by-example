import {
  handleSubmit,
  CMS
} from './formHandlers.js';
import {
  showNotification
} from './utils.js';

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

  // Handle Tags
  tagInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      const tag = tagInput.value.trim();
      addTag(tag);
      tagInput.value = '';
    }
  });

  const imageInput = document.querySelector('.image-upload #featured-image');
  const imageUrlInput = document.querySelector('.image-upload #featured-image-url');
  const uploadPlaceholder = document.querySelector('.image-upload .upload-placeholder');

  async function uploadImage(postID) {
    // Upload the image to the server
    const file = imageInput.files[0];
    if (file) {
      // Upload the image to the server
      const formData = new FormData();
      formData.append('image', file);
      formData.append('image_type', 'preview');

      try {
        const cms = new CMS();
        const result = await cms.UploadBlogPostImage(postID, formData);

        if (result && result.status === 'success') {
          const imageUrl = result.data.image.url;
          imageUrlInput.value = imageUrl;
          showNotification('Image uploaded successfully.', 'success');
        } else {
          showNotification('An error occurred while uploading the image.', 'error');
        }
      } catch (error) {
        console.error('Error uploading image:', error);
        showNotification('An error occurred while uploading the image.', 'error');
      }
    }

  }
  // Handle Save Draft
  const draftButton = document.querySelector('.action-button.draft');
  draftButton.addEventListener('click', async () => {
    const postData = {
      title: document.getElementById('post-title').value,
      excerpt: document.getElementById('post-excerpt').value,
      content: tinymce.get('post-content').getContent(),
      tags: Array.from(tags),
      status: 'draft'
    };

    const cms = new CMS();

    // Request body here
    const revisedPostData = {
      title: postData.title,
      body: postData.content,
      author_id: cms.getUserID(),
      status: postData.status
    }


    console.log(revisedPostData);
    const result = await cms.createNewPost(revisedPostData);

    if (result && result.data) {
      const post = result.data.post;
      const postID = post.id;
      uploadImage(postID);

      // Redirect to the page for the post
      window.location.href = `/app/posts/${postID}`
    } else {
      alert("Something went wrong!!")
    }
  });

  // Handle Publish
  const publishButton = document.querySelector('.action-button.publish');
  publishButton.addEventListener('click', async () => {
    const postData = {
      title: document.getElementById('post-title').value,
      excerpt: document.getElementById('post-excerpt').value,
      content: tinymce.get('post-content').getContent(),
      tags: Array.from(tags),
      status: 'published'
    };
    const cms = new CMS();

    // Request body here
    const revisedPostData = {
      title: postData.title,
      body: postData.content,
      author_id: cms.getUserID(),
      status: postData.status,
      tags: postData.tags
    }

    console.log(revisedPostData);
    const result = await cms.createNewPost(revisedPostData);

    if (result && result.data) {
      const post = result.data.post;
      const postID = post.id;
      uploadImage(postID);
      // Redirect to the page for the post
      window.location.href = `/app/posts/${postID}`
    } else {
      alert("Something went wrong!!")
    }

  });

  imageInput.addEventListener('change', async (e) => {

    const file = e.target.files[0];
    if (file) {
      // Display the image preview
      const reader = new FileReader();
      reader.onload = (e) => {
        uploadPlaceholder.innerHTML = `
            <img src="${e.target.result}" alt="Preview" style="max-width: 100%; max-height: 200px; border-radius: 0.5rem;">
          `;
      };
      // Load the image data
      reader.readAsDataURL(file);
    }
  });
});