document.addEventListener('DOMContentLoaded', () => {
    // Handle form submissions
    const forms = [
      'generalSettingsForm',
      'contentSettingsForm',
      'emailSettingsForm',
      'apiSettingsForm'
    ];
  
    forms.forEach(formId => {
      const form = document.getElementById(formId);
      if (form) {
        form.addEventListener('submit', (e) => {
          e.preventDefault();
          const formData = new FormData(form);
          const data = Object.fromEntries(formData);
          console.log(`${formId} submitted:`, data);
          // Here you would typically send the data to your backend
        });
      }
    });
  
    // Handle copy button
    const copyButtons = document.querySelectorAll('.copy-button');
    copyButtons.forEach(button => {
      button.addEventListener('click', () => {
        const textToCopy = button.dataset.clipboard;
        navigator.clipboard.writeText(textToCopy).then(() => {
          const originalText = button.textContent;
          button.textContent = '✅ Copied!';
          setTimeout(() => {
            button.textContent = originalText;
          }, 2000);
        });
      });
    });
  
    // Handle API key regeneration
    const regenerateButton = document.querySelector('.settings-button.danger');
    if (regenerateButton) {
      regenerateButton.addEventListener('click', () => {
        if (confirm('Are you sure you want to regenerate your API key? This will invalidate the current key.')) {
          console.log('Regenerating API key...');
          // Here you would typically send a request to regenerate the API key
        }
      });
    }
  });