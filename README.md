# WIP: Cheque

## Getting Started and Using

1.
```bash
docker compose build
```

2.
```bash
docker compose up
```

3.
```bash
make db_scheme
```

4. Add `var/data/extract.json`.

5.
```bash
make import
```

6. Check. Available common params: `page_size`, `page`, `sort_by`, `sort_order`.
- http://localhost:8080/api/v1/receipts
- http://localhost:8080/api/v1/sellers
- http://localhost:8080/api/v1/sellerplaces
- http://localhost:8080/api/v1/categories
- http://localhost:8080/api/v1/brands
- http://localhost:8080/api/v1/products
- http://localhost:8080/api/v1/productscategories
- http://localhost:8080/api/v1/images
- http://localhost:8080/api/v1/receiptproducts

7.
```bash
docker compose down --remove-orphans
```

## Swagger
- http://localhost:8081/

## Dev: Frontend
1. `npx create-refine-app@latest cheque_frontend`
   - Vite
   - REST API
   - Ant Design
   - Example Yes
   - Auth None
   - Npm
   - Email No
2. package.json > projectId, search > projectId: `uuidgen`