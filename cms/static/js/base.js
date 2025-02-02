function LogOut(){
    let logOutBtn = document.getElementById('nav-logout-btn');
    logOutBtn.addEventListener('click', function(){
        localStorage.removeItem('authToken');
        localStorage.removeItem('user');
        window.location.href = '/app/login';
    });
}

function checkAuth(){
    // Check if we are already on the login page
    if(window.location.pathname.startsWith('/app/login')){
        return;
    }
    // Retrieve the token from local storage
    let authToken = localStorage.getItem('authToken');
    // Check if token is present
    if(!authToken){
        window.location.href = '/app/login';
    }
    else{
        // Parse auth token
        let token_parts = authToken.split('.');
        let tokenPayload = token_parts[1];
        let tokenPayloadDecoded = atob(tokenPayload);
        let payload = JSON.parse(tokenPayloadDecoded);

        // Check if token is expired
        let current_time = new Date().getTime() / 1000;
        if(current_time > payload.exp){
            localStorage.removeItem('authToken');
            localStorage.removeItem('user');
            window.location.href = '/app/login';
        }
    }
}
document.addEventListener('DOMContentLoaded', function() {
    LogOut();
    checkAuth();
});