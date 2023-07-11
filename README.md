# content-analyzer-helper
It will read a PDF file, uses it as dataset and be able to answer your questions based on the file content.

## How to use it

### Set up API Key to use OpenAi
- https://platform.openai.com/account/api-keys
 
 After creating your Api Key you can set it as an environment variable, so the app can access it. Just run the following command, changing `<YOUR-API-KEY>` to the actual key that you've created:

    export AUTOMATE_CODE_API_KEY=<YOUR-API-KEY>

 once the Api key is set you need to execute the app.

 ### Running the APP
In your terminal go to the project root folder and run the following command:

    go build

after building the app you a file named `content-analyzes-helper` will be created, you can execute it by running this command:

    go run content-analyzes-helper

### Interaction
The app will ask you to provide the location and name of the PDF that will be used as data set:

    Please, type the PDF name, with whole path, that will be used as dataset...

you should type the whole file path and name, such as `~/Desktop/test.pdf`

After this, the app will prompt you

    Ask your question...
    
Then, you can ask anything based on the PDF's content and wait the response.    