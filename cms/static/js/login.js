import { handleSubmit, CMS } from './formHandlers.js';

handleSubmit('loginForm', async (data) => {
    let errorElement = document.getElementById('error');
    const cms = new CMS();
    let result = await cms.login(data);
    
    console.log('Login attempt:', data);
    console.log(result);

    if (
        result.error
    ) {
        errorElement.textContent = 'Authentication error: ' + result.error;
            // show the error
        errorElement.style.display = 'block';
        // add a fade in animation to errorElement
        errorElement.classList.add('fade-in');
    }else{
        // hide any previous auth error
        errorElement.style.display = 'none';
        let authToken = result.data?.token
        let user = result.data?.user
        
        // store token & user info in local storage
        localStorage.setItem('authToken', authToken);
        localStorage.setItem('user', JSON.stringify(user));

        window.location.href = 'home';
    }
});