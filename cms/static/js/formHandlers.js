import {
  formatDate,
  createModalWindow,
} from './utils.js';

function handleSubmit(formId, onSubmit) {
  document.getElementById(formId).addEventListener('submit', function (e) {
    e.preventDefault();

    const formData = {};
    const form = e.target;

    for (const element of form.elements) {
      if (element.name) {
        if (element.type === 'checkbox') {
          formData[element.name] = element.checked;
        } else {
          formData[element.name] = element.value;
        }
      }
    }

    onSubmit(formData);
  });
}

class CMS {
  constructor() {
    this.authToken = localStorage.getItem('authToken');
    let _user = localStorage.getItem('user');
    this.userObj = _user ? JSON.parse(_user) : null;

    this.UserID = this.userObj ? this.userObj.id : null;
    this.FirstName = this.userObj ? this.userObj.first_name : null;
    this.LastName = this.userObj ? this.userObj.last_name : null;
    this.Email = this.userObj ? this.userObj.email : null;
    this.Phone = this.userObj ? this.userObj.phone : null;

  }
  async login(data) {
    const response = await fetch('/api/v1/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;
  }
  async signup(data) {
    const response = await fetch('/api/v1/auth/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;
  }
  async updateProfile(data) {
    const response = await fetch(`/api/v1/users/${this.UserID}/profile`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    console.log('Profile update result:', result);
    return result;
  }
  async updatePassword(data) {
    const response = await fetch(`/api/v1/users/${this.UserID}/password`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;
  }

  async createNewPost(data) {
    const response = await fetch(`/api/v1/posts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;

  }

  async GetBlogPost(postID) {
    const response = await fetch(`/api/v1/posts/${postID}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  };

  async DeleteBlogPost(postID) {
    const response = await fetch(`/api/v1/posts/${postID}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }

  async GetBlogPosts() {
    const response = await fetch(`/api/v1/posts`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }

  async UpdateBlogPost(postID, data) {
    const response = await fetch(`/api/v1/posts/${postID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;
  }

  async GetBlogPostComments(postID) {
    const response = await fetch(`/api/v1/comments?post_id=${postID}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }
  async CreateBlogPostComment(data) {
    const response = await fetch(`/api/v1/comments`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + this.authToken,
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    return result;
  }
  async UploadBlogPostImage(postID, data) {
    // Ensure both postID and data are provided or throw an exception
    if (!postID || !data) {
      throw new Error('Both postID and data are required');
    }

    const response = await fetch(`/api/v1/posts/${postID}/images`, {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + this.authToken,
      },
      body: data
    });
    const result = await response.json();
    return result;
  }
  async DeleteBlogPostImage(postID, imageID) {
    const response = await fetch(`/api/v1/posts/${postID}/images/${imageID}`, {
      method: 'DELETE',
      headers: {
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }
  async GetBlogPostImages(postID) {
    const response = await fetch(`/api/v1/posts/${postID}/images`, {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }
  async GetBlogPostImage(postID, imageID) {
    const response = await fetch(`/api/v1/posts/${postID}/images/${imageID}`, {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer ' + this.authToken,
      }
    });
    const result = await response.json();
    return result;
  }
  getUserID() {
    return this.UserID;
  }
  getAuthToken() {
    return this.authToken;
  }

}

class Config {
  static MIN_PASSWORD_LENGTH = 7;
  static MAX_PASSWORD_LENGTH = 20;
  static API_BASE_URL = '/api/v1';
  static TOKEN_STORAGE_KEY = 'authToken';
  static USER_STORAGE_KEY = 'user';
}

class BlogPost {
  constructor(title, body, published_at, views, comments, likes) {

    this.title = title;
    this.body = body;
    this.published_at = published_at;
    this.views = views;
    this.comments = comments || 0;
    this.likes = likes;
    this.tags = [];
    this.commentsData = [];
    this.status = '';
    this.image = null;

    this.postPage = document.querySelector('main.main .post');

    this.pagePostTitle = this.postPage.querySelector('.post .post-content .post-title');
    this.pagePostStatus = this.postPage.querySelector('.post .post-header .post-meta .post-status');
    this.pagePostPublishedAt = this.postPage.querySelector('.post .post-header .post-meta .post-date');
    this.pagePostTags = this.postPage.querySelector('.post .post-content .post-tags');
    this.pagePostBody = this.postPage.querySelector('.post .post-content .post-body');
    this.pagePostViews = this.postPage.querySelector('.post .post-content .post-stats .stat-value.views');
    this.pagePostComments = this.postPage.querySelector('.post .post-content .post-stats .stat-value.comments');
    this.pagePostLikes = this.postPage.querySelector('.post .post-content .post-stats .stat-value.likes');
    this.pagePostFeaturedImage = this.postPage.querySelector('.post .post-content .post-featured-image img');
  }

  setPublishedAt(published_at) {
    this.published_at = published_at;
    // Set the published date
    this.pagePostPublishedAt.textContent = formatDate(published_at);
    return this;
  }
  setPublishedStatus(status) {
    this.status = status;

    // Set the status text
    this.pagePostStatus.textContent = status;
    return this;
  }
  setTitle(title) {
    this.title = title;
    // Set the window and document title
    window.title = title;
    document.title = title;
    // Set the title
    this.pagePostTitle.textContent = title;
    return this;
  }
  setBody(body) {
    this.body = body;
    // Set the body
    this.pagePostBody.innerHTML = body;
    return
  }
  setTags(tags) {
    this.tags = tags;

    // Clear the tags container
    this.pagePostTags.innerHTML = '';

    // Add the tags to the tags container
    for (let tag of tags) {
      const tagElement = document.createElement('span');
      tagElement.className = 'tag';
      tagElement.textContent = tag;
      this.pagePostTags.appendChild(tagElement);
    }

    return this;
  }
  setViews(views) {
    this.views = views;

    // Set the views count
    this.pagePostViews.textContent = views;
    return this;
  }
  setComments(comments) {
    this.comments = comments;

    // Set the comments count
    this.pagePostComments.textContent = comments;

    return this;
  }
  setLikes(likes) {
    this.likes = likes;

    // Set the likes count
    this.pagePostLikes.textContent = likes;
    return this;
  }
  setImageUrl(image) {
    if (image) {
      this.image = image;
      this.pagePostFeaturedImage.src = `/${image}`;
    }
    return this;
  }
  updateCommentsCount() {
    let commentTitle = document.getElementById('commentsTitle');
    let content = `Comments (${this.comments})`;
    commentTitle.textContent = content;

    this.pagePostComments.textContent = this.comments;
    return this;
  }
  incrementCommentsCount() {
    this.comments += 1;
    this.updateCommentsCount();
    return this;
  }
  show() {

    this.setPublishedAt(this.published_at);
    this.setTitle(this.title);
    this.setBody(this.body);
    this.setTags(this.tags);
    this.setViews(this.views);
    this.setComments(this.comments);
    this.setLikes(this.likes);
    this.setPublishedStatus(this.status);

    // Show the post
    this.postPage.classList.remove('hide');
    this.postPage.classList.add('show');
  }
}

class BlogPostEditor {
  constructor(title, body, published_at, views, comments, likes) {

    this.title = title;
    this.body = body;
    this.published_at = published_at;
    this.views = views;
    this.comments = comments;
    this.likes = likes;
    this.tags = [];
    this.status = '';

    this.postPage = document.querySelector('main.main .editor');
    this.loading = document.querySelector('main.main .loading');

    this.pagePostTitle = this.postPage.querySelector('.editor-container #post-title');
    this.pagePostStatus = this.postPage.querySelector('.editor-header .post-meta .post-status');
    this.pagePostPublishedAt = this.postPage.querySelector('.editor-header .post-meta .post-date');
    this.pagePostTagInput = this.postPage.querySelector('.editor-container .tags-input #tag-input');
    this.pagePostTags = this.postPage.querySelector('.editor-container .tags-input .tags-container');
    this.pagePostBody = this.postPage.querySelector('.editor-container  #post-content');
    this.pagePostViews = this.postPage.querySelector('.editor-container .post-stats .stat-value.views');
    this.pagePostComments = this.postPage.querySelector('.editor-container .post-stats .stat-value.comments');
    this.pagePostLikes = this.postPage.querySelector('.editor-container .post-stats .stat-value.likes');
    this.pagePostFeaturedImage = this.postPage.querySelector('.editor-container .image-upload .upload-placeholder img');

    this.configureEditor();
  }

  setPublishedAt(published_at) {
    this.published_at = published_at;
    // Set the published date
    if (this.status.toLowerCase() === 'draft') {
      this.pagePostPublishedAt.textContent = 'Last edited: ' + formatDate(published_at);
    } else {
      this.pagePostPublishedAt.textContent = 'Last published: ' + formatDate(published_at);
    }
    return this;
  }
  setPublishedStatus(status) {
    this.status = status;

    // Set the status text
    this.pagePostStatus.textContent = status.charAt(0).toUpperCase() + status.slice(1).toLowerCase();
    // Remove the existing status classes
    this.pagePostStatus.classList.remove('draft', 'published');
    this.pagePostStatus.classList.add(status.toLowerCase());
    return this;
  }
  setTitle(title) {
    this.title = title;
    // Set the window and document title
    window.title = title;
    document.title = title;
    // Set the title
    this.pagePostTitle.value = title;
    return this;
  }
  setBody(body) {
    this.body = body;
    // Set the body
    this.pagePostBody.innerHTML = body;
    return
  }
  addTag(tag) {
    this.tags.push(tag);

    // Create a new element for housing the tag
    const tagElement = document.createElement('span');
    tagElement.className = 'tag';
    tagElement.textContent = tag;

    // Create a new element for the tag remove button
    const tagRemoveElement = document.createElement('span');
    tagRemoveElement.className = 'tag-remove';
    tagRemoveElement.textContent = '×';

    // Add a click event listener to the tag remove button
    tagRemoveElement.addEventListener('click', () => {
      this.tags.delete(tag);
      tagElement.remove();
    });
    tagElement.appendChild(tagRemoveElement);

    // Add the tag to the tags container
    this.pagePostTags.appendChild(tagElement);

    return this;
  }
  getTags() {
    return this.tags;
  }
  setTags(tags) {
    this.tags = tags;

    // Clear the tags container
    this.pagePostTags.innerHTML = '';

    const addTag = (tag) => {
      // Create a new element for housing the tag
      const tagElement = document.createElement('span');
      tagElement.className = 'tag';
      tagElement.textContent = tag;

      // Create a new element for the tag remove button
      const tagRemoveElement = document.createElement('span');
      tagRemoveElement.className = 'tag-remove';
      tagRemoveElement.textContent = '×';

      // Add a click event listener to the tag remove button
      tagRemoveElement.addEventListener('click', () => {
        this.tags = this.tags.filter((t) => t !== tag);
        tagElement.remove();
      });
      tagElement.appendChild(tagRemoveElement);

      return tagElement;
    }
    // Add the tags to the tags container
    for (let tag of tags) {
      const tagElement = addTag(tag);
      this.pagePostTags.appendChild(tagElement);
    }

    return this;
  }
  setViews(views) {
    this.views = views;

    // Set the views count
    this.pagePostViews.textContent = views;
    return this;
  }
  setComments(comments) {
    this.comments = comments;

    // Set the comments count
    this.pagePostComments.textContent = comments;

    return this;
  }
  setLikes(likes) {
    this.likes = likes;

    // Set the likes count
    this.pagePostLikes.textContent = likes;
    return this;
  }
  setImageUrl(image) {
    if (image) {
      this.image = image;
      this.pagePostFeaturedImage.src = `/${image}`;
    }
    return this;
  }
  configureEditor() {
    // Input box to add tags
    this.pagePostTagInput.addEventListener('keydown', (e) => {
      if (e.key === 'Enter') {
        e.preventDefault();
        const tag = this.pagePostTagInput.value.trim();
        this.addTag(tag);
        this.pagePostTagInput.value = '';
      }
    });
  }
  show() {
    this.setPublishedAt(this.published_at);
    this.setTitle(this.title);
    this.setBody(this.body);
    this.setTags(this.tags);
    // this.setViews(this.views);
    // this.setComments(this.comments);
    // this.setLikes(this.likes);
    console.log(this.status);
    this.setPublishedStatus(this.status);

    if (this.loading) {
      // Hide the loading spinner
      this.loading.classList.remove('show');
      this.loading.classList.add('hide');
    }

    if (this.postPage) {
      // Show the post
      this.postPage.classList.remove('hide');
      this.postPage.classList.add('show');
    }
  }
}

export {
  handleSubmit,
  CMS,
  Config,
  BlogPost,
  BlogPostEditor
};