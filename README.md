# Task_Managment_Zocket


## Details of the API
1=> Signin API => User will enter the email address and password and if password corresponding to that email matches then user will be login successfully and will generate the JWT token


2=> Signup API => New user has to enter the FirstName,LastName,Email and Password and user will be stored in the database

3=> createtask API => User who want to create the new task first has to do login and it willl veerify using JWT token and user has to give the detail of the task such as title,description,assignedby,assignedto,Status and Priority rest all other field will be autogenraed by mongodb

4=> getalltask API => User can  see all the task which has to be completed 

5=> assignedtask API => User can see all the task which are asigned to the particular user and user will see the task according to the Priority such as High Priotity will be shown first then Medium Prioeity then Low Priority

6=> deletetask API => If any user has competed the task user will give the his email address and the title of the task which has been completed so that the task will be deleted from the list of all the task which has to do

## How AI tools ChatGPT helped
To debug the code really fast and give the featues of the Golang to implement the API