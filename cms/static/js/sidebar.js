export function initSidebar() {
    const currentPath = window.location.pathname;
    const sidebarItems = document.querySelectorAll('.sidebar-item');
    
    sidebarItems.forEach(item => {
      if (item.getAttribute('href') === currentPath) {
        item.classList.add('active');
      }
    });
  }