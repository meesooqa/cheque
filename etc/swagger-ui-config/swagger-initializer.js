window.onload = function() {
    const ui = SwaggerUIBundle({
        urls: [
            { url: "http://localhost:8080/swagger/sellerpb.swagger.json", name: "Seller API" },
            { url: "http://localhost:8080/swagger/sellerplacepb.swagger.json", name: "Seller Place API" },
            { url: "http://localhost:8080/swagger/categorypb.swagger.json", name: "Category API" },
            { url: "http://localhost:8080/swagger/brandpb.swagger.json", name: "Brand API" },
            { url: "http://localhost:8080/swagger/productpb.swagger.json", name: "Product API" },
            { url: "http://localhost:8080/swagger/productcategorypb.swagger.json", name: "Product Category API" },
            { url: "http://localhost:8080/swagger/imagepb.swagger.json", name: "Image API" },
            { url: "http://localhost:8080/swagger/receiptpb.swagger.json", name: "Receipt API" },
            { url: "http://localhost:8080/swagger/receiptproductpb.swagger.json", name: "Receipt Product API" },
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
