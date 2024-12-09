# E-commerce Store API

This is an e-commerce store API that allows users to add items to their cart, checkout, and receive a discount on every third order. Additionally, there are admin endpoints for generating discount codes and viewing reports on user activity.

## Discount Logic

- **10% Discount**: Every 3rd order placed by a user is eligible for a 10% discount.
- **Coupon Code**: The coupon code for the discount is automatically generated for every 3rd order and can be applied to the order.
- **Usage**: The coupon can only be used once per order. After the discount is applied, the coupon is removed, and the user will need to wait for their next 3rd order to receive another coupon.

## API Endpoints

### 1. **Add Item to Cart**

- **URL**: `/v1/cart/add`
- **Method**: `POST`
- **Description**: Adds an item to the user's cart.

#### Request Parameters:
- `user_id` (Query Parameter): The ID of the user placing the order.
- Request Body (JSON):
    ```json
    {
      "id": 101,
      "name": "Laptop",
      "price": 1000.0
    }
    ```

    Where:
    - `id`: Unique identifier for the item.
    - `name`: Name of the item.
    - `price`: Price of the item.

#### Example Curl Request:
```bash
curl -X POST "http://localhost:8080/v1/cart/add?user_id=1" \
     -H "Content-Type: application/json" \
     -d '{"id": 101, "name": "Laptop", "price": 1000.0}'
```

#### Response:
- Status: `200 OK`
- Response Body (JSON):
    ```json
    {
      "items": [
        {
          "id": 101,
          "name": "Laptop",
          "price": 1000.0
        }
      ]
    }
    ```

### 2. **Checkout**

- **URL**: `/v1/cart/checkout`
- **Method**: `GET`
- **Description**: Performs the checkout for the user and applies any applicable discount.

#### Request Parameters:
- `user_id` (Query Parameter): The ID of the user checking out.

#### Example Curl Request:
```bash
curl -X GET "http://localhost:8080/v1/cart/checkout?user_id=1"
```

#### Response:
- Status: `200 OK`
- Response Body (JSON):
    ```json
    {
      "user_id": 1,
      "total_amount": 1800.0,
      "coupon_code": "DISCOUNT10",
      "order_number": 3,
      "amount_after_discount": 1620.0
    }
    ```

#### Response Fields:
- `user_id`: The ID of the user who placed the order.
- `total_amount`: The total amount of the cart before applying the discount.
- `coupon_code`: The generated discount coupon code (if eligible). For every 3rd order, a coupon is applied (e.g., `DISCOUNT10`).
- `order_number`: The current order number of the user.
- `amount_after_discount`: The total amount after applying the 10% discount (only if the user is eligible for the discount).

### 3. **Admin API: Generate Discount Code**

- **URL**: `/v1/admin/discount/generate`
- **Method**: `POST`
- **Description**: Generates a discount code for every 3rd order for a user.

#### Request Parameters:
- `user_id` (Query Parameter): The ID of the user for whom to generate a discount code.

#### Example Curl Request:
```bash
curl -X POST "http://localhost:8080/v1/admin/discount/generate?user_id=1" \
     -H "Authorization: root"
```

#### Response:
- Status: `200 OK`
- Response Body (JSON):
    ```json
    {
      "user_id": 1,
      "discount_code": "DISCOUNT10"
    }
    ```

#### Response Fields:
- `user_id`: The ID of the user for whom the discount code was generated.
- `discount_code`: The generated discount code (e.g., `DISCOUNT10`).

### 4. **Admin API: Get Admin Report**

- **URL**: `/v1/admin/report`
- **Method**: `GET`
- **Description**: Provides an admin report containing the count of items purchased, total purchase amount, discount codes used, and total discount amount for each user.

#### Example Curl Request:
```bash
curl -X GET "http://localhost:8080/v1/admin/report" \
     -H "Authorization: root"
```

#### Response:
- Status: `200 OK`
- Response Body (JSON):
    ```json
    {
      "1": {
        "total_items": 5,
        "total_amount": 2000.0,
        "discount_codes": ["DISCOUNT10"],
        "total_discount": 200.0
      },
      "2": {
        "total_items": 3,
        "total_amount": 1500.0,
        "discount_codes": ["DISCOUNT10"],
        "total_discount": 150.0
      }
    }
    ```

#### Response Fields:
- For each user (by `user_id`):
  - `total_items`: The total number of items purchased by the user.
  - `total_amount`: The total amount spent by the user.
  - `discount_codes`: A list of discount codes used by the user.
  - `total_discount`: The total discount amount for the user.

## Running the API Locally

### Prerequisites
- Go 1.18 or later installed
- `curl` for testing API calls
- A running Go server

### Steps:
1. Clone the repository to your local machine.
2. Navigate to the project directory and run the server:
   ```bash
   go run main.go
   ```
3. The API will be available at `http://localhost:8080`.

---

This updated README includes the **Admin API** endpoints, as well as instructions for generating discount codes and accessing admin reports, including the `Authorization` requirement (`root` password).
