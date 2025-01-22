# Database Test Case

## Test Case 1: Add a New User to the Database
**Description:** Verify that a new user can be successfully added to the database.  

**Input Data:**
- Full Name: `John Doe`
- Phone Number: `1234567890`
- Gender: `Male`

**Steps to Execute:**
1. Call the function `AddUserToDatabase("John Doe", "1234567890", "Male")`.
2. Check the function response.
3. Query the database to verify that the data has been inserted.

**Expected Result:**
- The function returns an `InsertedID`.
- The database query confirms the new user with the correct details.

---
## Test Case 2: Find user by phone number
**Description:** Verify that a user can be searched by their phone number.

**Input Data:**
- Phone number: `1234567890`

**Steps to Execute:**
1. Call the function `AddUserToDatabase("John Doe", "1234567890", "Male")` to create a user for test.
2. Call the function `FoundUserByPhoneNumber("1234567890")`.
3. Check the function response.

**Expected Result:**
- The functiuon returns the correct user information:
```
Full Name: John Doe, Phone Number: 1234567890, Gender: Male
```
---
## Test Case 3: Get all users
**Description:** Veryfy that the function will get all users when the database contains user records.

**Input Data:**

- Prepare data before test
    - `John Doe`: `1234567890`, `Male`
    - `Jane Smith`: `0987654321`, `Female`
    - `Bob Johnson`: `9876543210`, `Male`

**Steps to Execute:**
1. Add users to the database before test by call function `AddUserToDatabase`
2. Call the function `GetAllUsers()`
3. Capture the output from stdout.
4. Check the function response.
**Expected Result:**
- The function return the complete user list in the format:
```
Full Name: John Doe, Phone Number: 1234567890, Gender: Male
Full Name: Jane Smith, Phone Number: 0987654321, Gender: Female
Full Name: Bob Johnson, Phone Number: 9876543210, Gender: Male
```
---
## Test Case 4: Retrieve All Users (No Data)
**Description:** Verify that the function handles the scenario when no users exist in the database.

**Steps to Excute:**
1. Ensure the database don't have any data
2. Call the function `GetAllUsers()`
3. Capture the output from stdout.
4. Check the function response.

**Expected Result:**
```
No users found
```




<!-- 
## Test Case 4: Retrieve All Users (No Data)
**Description:** Verify that the function handles the scenario when no users exist in the database.

**Steps to Execute:**
1. Ensure the database contains no users.
2. Call the function `GetAllUsers()`.
3. Check the function response.

**Expected Result:**
- The function returns the message: `No users found`.

---
 -->
