function handleSubmit(formId, onSubmit) {
  document.getElementById(formId).addEventListener('submit', function(e) {
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

class CMS{
  constructor(){
    this.authToken = localStorage.getItem('authToken');
    let _user = localStorage.getItem('user');
    this.userObj = _user ? JSON.parse(_user) : null;
    
    this.UserID = this.userObj ? this.userObj.id : null;
    this.FirstName = this.userObj ? this.userObj.first_name : null;
    this.LastName = this.userObj ? this.userObj.last_name : null;
    this.Email = this.userObj ? this.userObj.email : null;
    this.Phone = this.userObj ? this.userObj.phone : null;

  }
  async login(data){
      const response = await fetch('/api/v1/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });
        const result = await response.json();
        console.log('Login result:', result);
        return result;
  }
  async signup(data){
      const response = await fetch('/api/v1/auth/signup', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });
        const result = await response.json();
        console.log('Signup result:', result);
        return result;
  }
  async updateProfile(data){
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
  async updatePassword(data){
      const response = await fetch(`/api/v1/users/${this.UserID}/password`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + this.authToken
          },
          body: JSON.stringify(data)
        });
        const result = await response.json();
        console.log('Password update result:', result);
        return result;
  }

}
class Config {
  static MIN_PASSWORD_LENGTH = 7;
  static MAX_PASSWORD_LENGTH = 20;
  static API_BASE_URL = '/api/v1';
  static TOKEN_STORAGE_KEY = 'authToken';
  static USER_STORAGE_KEY = 'user';
}

export { handleSubmit, CMS, Config };