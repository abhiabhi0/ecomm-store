Here's a sample README document explaining the discount logic and how to call the APIs using `curl`:

---

# E-commerce Store API

This is an e-commerce store API that allows users to add items to their cart, checkout, and receive a discount on every third order.

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
- `coupon_code`: The generated discount coupon code (if eligible). For every 3rd order, a coupon is applied (e.g., `DISCOUNT10-3`).
- `order_number`: The current order number of the user.
- `amount_after_discount`: The total amount after applying the 10% discount (only if the user is eligible for the discount).

### Example Flow

1. **Add Items to Cart**: 
    - User adds multiple items to their cart.
    - Example: First add a Laptop, then a Phone, and a Tablet.
    
2. **Checkout**: 
    - After adding items to the cart, the user proceeds to checkout.
    - If this is the 3rd order, a coupon will be applied automatically and a 10% discount will be calculated.
    - The response will include the order number, the total amount before the discount, the applied coupon, and the amount after the discount.

### Example Scenario:
- **User 1**:
  1. Adds Laptop (Price: 1000.0)
  2. Adds Phone (Price: 500.0)
  3. Adds Tablet (Price: 300.0)
  4. Checkout: Discount applied on 3rd order, total amount after discount is 1620.0 (10% off applied).

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

This README should provide a clear understanding of how the e-commerce API works, how to call the APIs with `curl`, and how the discount logic is implemented.