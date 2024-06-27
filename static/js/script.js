document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('loginForm');
    const errorMessage = document.getElementById('errorMessage');
    const togglePassword = document.getElementById('togglePassword');
    const passwordInput = document.getElementById('site_password');

    togglePassword.addEventListener('click', function() {
        const type = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
        passwordInput.setAttribute('type', type);
        this.textContent = type === 'password' ? '👁️' : '👁️‍🗨️';
    });

    loginForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const formData = new FormData(loginForm);

        fetch('/login', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                showError(data.error);
            } else if (data.loginUrl) {
                // 登录成功，重定向到后端提供的URL
                window.location.href = data.loginUrl;
            } else {
                // 如果没有错误也没有loginUrl，显示一个通用成功消息
                showSuccess('登录成功！');
            }
        })
        .catch(error => {
            showError('登录时发生错误，请稍后再试。');
            console.error('Error:', error);
        });
    });

    function showError(message) {
        errorMessage.textContent = message;
        errorMessage.style.display = 'block';
        errorMessage.className = 'error-message';

        // 添加错误状态到表单
        loginForm.classList.add('error');

        // 3秒后移除错误状态
        setTimeout(() => {
            loginForm.classList.remove('error');
            errorMessage.style.display = 'none';
        }, 3000);
    }

    function showSuccess(message) {
        errorMessage.textContent = message;
        errorMessage.style.display = 'block';
        errorMessage.className = 'success-message';

        // 3秒后隐藏成功消息
        setTimeout(() => {
            errorMessage.style.display = 'none';
        }, 3000);
    }
});