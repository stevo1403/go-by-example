import {
  CMS,
  BlogPostEditor
} from './formHandlers.js';

import {
  showNotification
} from './utils.js'

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

// Initialize the CMS
const cms = new CMS();

// Extract the postID from the URL
const match = window.location.pathname.match(/\/posts\/(\d+)\/edit/);
if (!match) {
  // Redirect to the posts page
  window.location.href = '/app/posts';
}
// Check if postID is a number using Regex
const postID = match[1];
if (!/^\d+$/.test(postID)) {
  // Redirect to the posts page
  window.location.href = '/app/posts';
}

let blogPostEditor = null;

  async function updateBlogPostUI(result) {
  if (result && result.status === 'success') {
    const blogPost = result.data.post;

    // Extract the blog post data from the response
    const blogPostTitle = blogPost.title;
    const blogPostBody = blogPost.body;
    const blogPostPublishedAT = blogPost.published_at;
    const blogPostViews = blogPost.views;
    const blogPostStatus = blogPost.is_draft ? 'Draft' : 'Published';
    let blogPostFeaturedImage = null;

    // const blogPostComments = blogPost.comments;
    // const blogPostLikes = blogPost.likes;

    // Fetch the blog post images
    const postID = blogPost.id;
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

    // Populate the form with the blog post data
    blogPostEditor = new BlogPostEditor(blogPostTitle, blogPostBody, blogPostPublishedAT, blogPostViews, 0, 0);
    blogPostEditor.setPublishedStatus(blogPostStatus);
    blogPostEditor.setPublishedAt(blogPostPublishedAT);
    blogPostEditor.setImageUrl(blogPostFeaturedImage);
    blogPostEditor.setTags(blogPost.tags);
    blogPostEditor.show();

  } else {
    console.log('Error fetching post');
  }
}

// Fetch the blog post data
const result = await cms.GetBlogPost(postID);

await updateBlogPostUI(result);

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
async function savePost(status) {
  const postData = {
    title: document.getElementById('post-title').value,
    excerpt: document.getElementById('post-excerpt').value,
    content: tinymce.get('post-content').getContent(),
    tags: blogPostEditor.getTags(),
    status: status
  };
  console.log('Saving post:', postData);
  const revisedPostData = {
    title: postData.title,
    body: postData.content,
    is_draft: postData.status === 'draft',
    tags: postData.tags
  };
  
  await uploadImage(postID);
  const result = await cms.UpdateBlogPost(postID, revisedPostData);
  if (result && result.status === 'success') {
    // Notify the user that the post was updated
    showNotification('Post updated successfully.', 'success');
    // Update the UI with the new post data
    await updateBlogPostUI(result);
  } else {
    // Notify the user that an error occurred
    showNotification('Failed to update post.', 'error');
  }
}
// Handle Save Draft
const draftButton = document.querySelector('.action-button.draft');
draftButton.addEventListener('click', async () => {
  await savePost('draft');
});

// Handle Update
const publishButton = document.querySelector('.action-button.publish');
publishButton.addEventListener('click', async () => {
  await savePost('published');
});

// Handle Image Upload Preview
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