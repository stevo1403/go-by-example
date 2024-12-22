import { handleSubmit, CMS, Config } from './formHandlers.js';

handleSubmit('passwordForm', async (data) => {
    console.log('Password update:', data);
    let errorElement = document.getElementById('password-update-error');
    let alertElement = document.getElementById('password-update-alert');

    let _data = {};
    // check if password and confirm password are same
    if (data.password !== data.confirm_password) {
        errorElement.textContent = '❌ Passwords do not match';
        errorElement.style.display = 'block';
        errorElement.classList.add('fade-in');
        return;
    }
    // check password length requirement
    if (data.password.length < Config.MIN_PASSWORD_LENGTH) {
        errorElement.textContent = `❌ Password is too short (min ${Config.MIN_PASSWORD_LENGTH} characters)`;
        errorElement.style.display = 'block';
        errorElement.classList.add('fade-in');
        return;
    }
    if (data.password.length > Config.MAX_PASSWORD_LENGTH) {
        errorElement.textContent = `❌ Password is too long (max ${Config.MAX_PASSWORD_LENGTH} characters)`;
        errorElement.style.display = 'block';
        errorElement.classList.add('fade-in');
        return;
    }

    const cms = new CMS();
    let result = await cms.updatePassword(data);

    if (result.error) {
        alertElement.style.display = 'none';
        errorElement.textContent = '❌ Password update error: ' + result.error;
        errorElement.style.display = 'block';
        errorElement.classList.add('fade-in');
    } else {
        errorElement.style.display = 'none';
        alertElement.textContent = '🎉 Password updated successfully 🔒';
        alertElement.style.display = 'block';
        alertElement.classList.add('fade-in');
        
        // clear form fields
        let form = document.getElementById('passwordForm');
        form.reset();
    }
});