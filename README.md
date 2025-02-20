# Billing Engine API

This is a billing engine for a loan management system that handles loan creation, payment tracking, delinquency checks, and automatic EMI status updates through a daily cron job.

## How to Run the App

### Prerequisites
- Go 1.18+
- Git

### Steps
1. **Clone the Repository**
```bash
  git clone https://github.com/sharkx018/billing-engine
  cd billing-engine
```

2. **Install Dependencies**
```bash
  go mod tidy
```

3. **Run the Server**
```bash
  go run main.go
```

4. **Server will start on port 3000**
```bash
  Cron job started for marking the missed payments
  Billing Server Started at port :3000
```

---

## API Endpoints

### 1. **User Sign-Up**
- **Endpoint:** `POST /sign-up`
- **Description:** Registers a new user and returns a JWT token.
- **Request:**
```json
{
    "mobile":"8318183466",
    "password":"123"
}
```
- 
- **Response:**
```json
{
  "data": {
    "message": "User registered successfully",
    "token": "<jwt_token>",
    "userId": 1
  },
  "success": true
}
```

### 2. **User Sign-In**
- **Endpoint:** `POST /sign-in`
- **Description:** Authenticates the user and returns a JWT token.
- **Request:**
```json
{
    "mobile":"8318183466",
    "password":"123"
}
```
-
- **Response:**
```json
{
  "data": {
    "message": "User logged in successfully",
    "token": "<jwt_token>",
    "userId": 1
  },
  "success": true
}
```


### 3. **Create Loan**
- **Endpoint:** `POST /create-loan`
- **Authorization:** Bearer Token (JWT) or Headers `x-bypass: true` and `x-user-id: <user_id>`
- **Description:** Creates a new loan for the authenticated user.
- **Response:**
```json
{
    "data": {
        "loan_info": {
            "loan_id": 1,
            "user_id": 1,
            "principal": 30000,
            "interest": 0.1,
            "total_amount": 33000,
            "weekly_payment": 660,
            "outstanding": 33000,
            "missed_payments": 0,
            "next_payment_date": "2025-02-20",
            "pending_payments": 50,
            "emi_schedule": [
                {
                    "week_number": 1,
                    "due_date": "2025-02-20",
                    "amount": 660,
                    "status": "pending"
                },
                {
                    "week_number": 2,
                    "due_date": "2025-02-27",
                    "amount": 660,
                    "status": "pending"
                }
               
            ]
        },
        "message": "Loan created successfully"
    },
    "success": true
}
```

### 4. **Make Payment**
- **Endpoint:** `POST /make-payment`
- **Authorization:** Bearer Token (JWT) or Headers `x-bypass: true` and `x-user-id: <user_id>`
- **Description:** Allows the user to make a payment for a specific EMI.
- **Request Body:**
```json
{
    "loan_id": 1,
    "emi_number": 1
}
```
- **Response:**
```json
{
  "data": {
    "loan_info": {
      "loan_id": 1,
      "user_id": 1,
      "principal": 30000,
      "interest": 0.1,
      "total_amount": 33000,
      "weekly_payment": 660,
      "outstanding": 31020,
      "missed_payments": 0,
      "next_payment_date": "2025-03-13",
      "pending_payments": 47,
      "emi_schedule": [
        {
          "week_number": 1,
          "due_date": "2025-02-20",
          "amount": 660,
          "status": "paid"
        },
        {
          "week_number": 2,
          "due_date": "2025-02-27",
          "amount": 660,
          "status": "paid"
        }
      ]
    },
    "message": "Payment processed successfully"
  },
  "success": true
}
```

### 5. **Get Outstanding Loans**
- **Endpoint:** `GET /get-outstanding`
- **Authorization:** Bearer Token (JWT) or Headers `x-bypass: true` and `x-user-id: <user_id>`
- **Description:** Fetches all loans for the authenticated user along with outstanding amounts.
- Response:
```json
{
    "data": {
        "message": "Loan info is fetched for the user",
        "total_outstanding": 31020,
        "active_loans": [
            {
                "loan_id": 1,
                "user_id": 1,
                "principal": 30000,
                "interest": 0.1,
                "total_amount": 33000,
                "weekly_payment": 660,
                "outstanding": 31020,
                "missed_payments": 0,
                "next_payment_date": "2025-03-13",
                "pending_payments": 47,
                "emi_schedule": [
                    {
                        "week_number": 1,
                        "due_date": "2025-02-20",
                        "amount": 660,
                        "status": "paid"
                    },
                    {
                        "week_number": 2,
                        "due_date": "2025-02-27",
                        "amount": 660,
                        "status": "paid"
                    }
                ]
            }
        ]
        
    },
    "success": true
}
```

### 6. **Check Delinquency Status**
- **Endpoint:** `GET /is-delinquent`
- **Authorization:** Bearer Token (JWT) or Headers `x-bypass: true` and `x-user-id: <user_id>`
- **Description:** Checks if the user is delinquent (missed 2 or more payments).
- **Response:** 
```json
{
    "data": {
        "message": "Loan info is fetched for the user",
        "is_delinquent": false
    },
    "success": true
}
```


---

## Assumptions
- **In-Memory Storage:** The app uses in-memory data structures (`map`) for storing users and loans. The data will reset when the server restarts.
- **EMI Status:** The status of each EMI can be `pending`, `paid`, or `missed`.
- **Daily Cron Job:** Automatically marks overdue EMIs as `missed` every 24 hours.
- **Authentication Bypass:** By using `x-bypass` and `x-user-id` headers, you can bypass token authentication.

---

## Edge Cases Handled
- **Duplicate Payments:** Prevents paying an EMI that is already marked as `paid`.
- **Invalid Loan or EMI Number:** Returns appropriate error messages for invalid loan IDs or EMI numbers.
- **User Mismatch:** Ensures that users can only access their own loans and payments.
- **Missed Payments:** Automatically increments the `MissedPayments` counter if an EMI is overdue.
- **Sequential Payment Enforcement:** Ensures all previous weekly EMIs are paid before paying the current one.

---

## Future Enhancements
- **Database Integration:** Replace in-memory storage with a persistent database like PostgreSQL.
- **Notification System:** Send alerts for upcoming due dates or missed payments.

---

