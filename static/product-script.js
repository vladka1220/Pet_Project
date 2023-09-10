document.getElementById("product-form").addEventListener("submit", function (event) {
    event.preventDefault();
    const productName = document.getElementById("product-name").value;  
    const lotepProduto = document.getElementById("lote-produto").value;
    const productDescription = document.getElementById("product-description").value;
    const productQuantity = document.getElementById("product-quantity").value;
    const maquina = document.getElementById("maquina").value;
    
    // действия при отправке формы
    console.log("Название товара:", productName);
    console.log("Описание товара:", productDescription);
    console.log("Количество:", productQuantity);
    console.log("Lote:", lotepProduto);
    console.log("Maquina:", maquina);

    
    // отправить данные на сервер, сохранить их в базе данных и т. д.
});


