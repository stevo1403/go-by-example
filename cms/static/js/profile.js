import { handleSubmit, CMS } from './formHandlers.js';

handleSubmit('profileForm', async (data) => {
    console.log('Profile update:', data);
    let errorElement = document.getElementById('profile-update-error');
    let alertElement = document.getElementById('profile-update-alert');

    let _data = {};
    if (data.first_name) _data.first_name = data.first_name;
    if (data.second_name) _data.second_name = data.second_name;
    if (data.phone) _data.phone = data.phone;

    const cms = new CMS();
    let result = await cms.updateProfile(data);

    if (result.error) {
        alertElement.style.display = 'none';
        errorElement.textContent = '❌ Profile update error: ' + result.error;
        errorElement.style.display = 'block';
        errorElement.classList.add('fade-in');
    } else {
        updateLocalData(result.data);
        errorElement.style.display = 'none';
        alertElement.textContent = '🎉 Profile updated successfully 🎉';
        alertElement.style.display = 'block';
        alertElement.classList.add('fade-in');
    }
});

function checkAuthState(){
    // Check if local storage has a token
    let authToken = localStorage.getItem('authToken');
    let user = localStorage.getItem('user');
    
    if (!authToken || !user) {
        // redirect to login page
        console.log('No token or user found');
        // window.location.href = 'login';
    }else{
        let tokenHeader = authToken.split('.')[0];
        let tokenBody = authToken.split('.')[1];
        
        let parsedTokenBody = JSON.parse(atob(tokenBody));
        let tokenExpiration = parsedTokenBody.exp;
        // check if token is expired
        if (Date.now() >= tokenExpiration * 1000) {
            // redirect to login page
            console.log('Token expired');
            window.location.href = 'login';
        }
    }
}

function FillProfileForm(){
    let user = localStorage.getItem('user');
    let userObj = user ? JSON.parse(user) : null;
    
    let form = document.getElementById('profileForm');

    if (userObj) {
        let firstName = form.querySelector('input[name="first_name"]');
        let lastName = form.querySelector('input[name="last_name"]');
        let phone = form.querySelector('input[name="phone"]');
        
        firstName.value = userObj.first_name;
        lastName.value = userObj.last_name;
        phone.value = userObj.phone;
    }
}

function FillLabels(){
    let user = localStorage.getItem('user');
    let userObj = user ? JSON.parse(user) : null;
    
    let fullNameLabel = document.getElementById('nav-fullname-label');
    
    if (userObj) {
        fullNameLabel.textContent = `${userObj.first_name} ${userObj.last_name}`;
    }
}

function updateLocalData(data){
    let user = localStorage.getItem('user');
    let userObj = user ? JSON.parse(user) : null;
    
    if (userObj) {
        if (data.id) userObj.id = data.id;
        if (data.first_name) userObj.first_name = data.first_name;
        if (data.last_name) userObj.last_name = data.last_name;
        if (data.phone) userObj.phone = data.phone;
        if (data.email) userObj.email = data.email;
        
        localStorage.setItem('user', JSON.stringify(userObj));
    }
}
checkAuthState();

document.addEventListener('DOMContentLoaded', function() {
    FillProfileForm();
    FillLabels();
});