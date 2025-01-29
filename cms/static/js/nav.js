export function initNav() {
    const logoutButton = document.querySelector('.nav-logout');
    if (logoutButton) {
      logoutButton.addEventListener('click', (e) => {
        e.preventDefault();
        // Handle logout logic here
        console.log('Logout clicked');
      });
    }
  }