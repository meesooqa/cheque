window.onload = function() {
    const ui = SwaggerUIBundle({
        urls: [
            { url: "http://localhost:8080/swagger/sellerpb/v1/seller.swagger.json", name: "Seller API" },
            { url: "http://localhost:8080/swagger/sellerplacepb/v1/sellerplace.swagger.json", name: "Seller Place API" },
            { url: "http://localhost:8080/swagger/categorypb/v1/category.swagger.json", name: "Category API" },
            { url: "http://localhost:8080/swagger/brandpb/v1/brand.swagger.json", name: "Brand API" },
            { url: "http://localhost:8080/swagger/productpb/v1/product.swagger.json", name: "Product API" },
            { url: "http://localhost:8080/swagger/productcategorypb/v1/productcategory.swagger.json", name: "Product Category API" },
            { url: "http://localhost:8080/swagger/imagepb/v1/image.swagger.json", name: "Image API" },
            { url: "http://localhost:8080/swagger/receiptpb/v1/receipt.swagger.json", name: "Receipt API" },
            { url: "http://localhost:8080/swagger/receiptproductpb/v1/receiptproduct.swagger.json", name: "Receipt Product API" },
        ],
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIStandalonePreset
        ],
        plugins: [
            SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
    });
    window.ui = ui;
};
