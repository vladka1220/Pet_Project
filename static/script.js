document.getElementById("login-form").addEventListener("submit", function (event) {
    event.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    
    // Проверяем учетные данные
    if ((username === "admin" && password === "1") || (username === "vlad" && password === "123")) {
        // Перенаправляем на страницу продукции
        window.location.href = "product.html";
    } else {
        alert("Ошибка аутентификации");
    }
});
