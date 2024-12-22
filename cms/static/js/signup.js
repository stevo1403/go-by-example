import {
    handleSubmit,
    CMS
} from './formHandlers.js';

handleSubmit('signupForm', async (data) => {
    let errorElement = document.getElementById('error');
    const cms = new CMS();
    let result = await cms.signup(data);

    if (result.error) {
        if (result.error.toLowerCase().includes('first')) {
            errorElement.textContent = 'First name isn\'t valid: ' + result.error;
            // show the error
            errorElement.style.display = 'block';
            // add a fade in animation to errorElement
            errorElement.classList.add('fade-in');
        }else if (result.error.toLowerCase().includes('last')) {
            errorElement.textContent = 'Last name isn\'t valid: ' + result.error;
            // show the error
            errorElement.style.display = 'block';
            // add a fade in animation to errorElement
            errorElement.classList.add('fade-in');
        } else if (result.error.toLowerCase().includes('email')) {
            errorElement.textContent = 'Email isn\'t valid: ' + result.error;
            // show the error
            errorElement.style.display = 'block';
            // add a fade in animation to errorElement
            errorElement.classList.add('fade-in');
        } else if (result.error.toLowerCase().includes('password')) {
            errorElement.textContent = 'Password isn\'t valid: ' + result.error;
            // show the error
            errorElement.style.display = 'block';
            // add a fade in animation to errorElement
            errorElement.classList.add('fade-in');
        } else if (result.error.toLowerCase().includes('phone')) {
            errorElement.textContent = 'Phone isn\'t valid: ' + result.error;
            // show the error
            errorElement.style.display = 'block';
            // add a fade in animation to errorElement
            errorElement.classList.add('fade-in');
        } 
    } else {
        // hide any previous auth error
        errorElement.style.display = 'none';
        // redirect to login page
        window.location.href = 'login';
    }
});