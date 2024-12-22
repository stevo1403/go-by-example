function LogOut(){
    let logOutBtn = document.getElementById('nav-logout-btn');
    logOutBtn.addEventListener('click', function(){
        localStorage.removeItem('authToken');
        localStorage.removeItem('user');
        window.location.href = 'login';
    });
}

document.addEventListener('DOMContentLoaded', function() {
    LogOut();
});