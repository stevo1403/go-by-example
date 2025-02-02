function formatDate(dateString) {
    // Format a date string as "Month Day, Year"

    // Step 1: Parse the date string into a Date object
    const date = new Date(dateString);

    // Step 2: Use Date methods to get the month, day, and year
    const month = date.toLocaleString('default', {
        month: 'long'
    }); // Full month name (e.g., "January")
    const day = date.getDate(); // Day of the month (e.g., 25)
    const year = date.getFullYear(); // Full year (e.g., 2024)

    // Step 3: Format the date as "Month Day, Year"
    const formattedDate = `${month} ${day}, ${year}`;

    return formattedDate
}

function createModalWindow(question, onConfirm, onCancel) {
    // Remove existing modal windows
    const existingModals = document.querySelectorAll('.modal');
    for (let modal of existingModals) {
        modal.remove();
    }
    const modal = document.createElement('div');
    modal.className = 'modal';
    modal.innerHTML = `
    <div class="modal-content">
        <p>${question}</p>
        <hr>
        <div class="modal-actions">
            <button class="action-button confirm">Yes</button>
            <button class="action-button cancel">No</button>
        </div>
    </div>`;

    modal.classList.add('show');
    document.body.appendChild(modal);

    const confirmButton = modal.querySelector('.modal-content .modal-actions .action-button.confirm');
    const cancelButton = modal.querySelector('.modal-content .modal-actions .action-button.cancel');

    confirmButton.addEventListener('click', () => {
        onConfirm();
        modal.remove();
    });

    cancelButton.addEventListener('click', () => {
        onCancel();
        modal.remove();
    });
}

/**
 * Displays a notification to the user.
 * @param {string} message - The message to display.
 * @param {string} type - The type of notification ('success' or 'error').
 * @param {number} duration - The duration (in milliseconds) the notification should be visible.
 */
function showNotification(message, type = 'success', duration = 3000) {
    // Create the notification element
    const notification = document.createElement('div');
    notification.classList.add('notification', type);
  
    // Add emoji based on the type
    const emoji = type === 'success' ? '✅' : '❌';
    notification.innerHTML = `<span class="emoji">${emoji}</span>${message}`;
  
    // Append the notification to the container
    const container = document.getElementById('notification-container');
    container.appendChild(notification);
  
    // Trigger the show animation
    setTimeout(() => notification.classList.add('show'), 10);
  
    // Remove the notification after the specified duration
    setTimeout(() => {
      notification.classList.remove('show');
      setTimeout(() => notification.remove(), 500); // Wait for the fade-out transition
    }, duration);
  }

export {
    formatDate,
    createModalWindow,
    showNotification
}

