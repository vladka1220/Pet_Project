document.getElementById("login-form").addEventListener("submit", function (event) {
    event.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    
    // Проверяем учетные данные, вместо **** можем добавить свои данные пока нет подключенной БД
    if ((username === "****" && password === "*****") || (username === "*****" && password === "*****")) {
        // Перенаправляем на страницу продукции
        window.location.href = "product.html";
    } else {
        alert("Ошибка аутентификации");
    }
});
