# zabira-code-challenge


For this particular use case, I was asked to create a sorting solution that is extensible so that different teams can easily add their own sorters without affecting existing code. 

Using an OOP approach, as demonstrated in my implementation, can make the code more modular and easier to extend. By creating an interface for sorters and implementing that interface in different sorter classes/objects, the team can add new sorting algorithms without changing the existing code. This approach promotes code reuse and separates concerns, which can lead to a more maintainable and extensible codebase.


### Running ``` go run main.go ```
<img width="1680" alt="Screenshot 2023-04-12 at 15 51 31" src="https://user-images.githubusercontent.com/43158886/231497566-ccabbe87-75fe-44e2-b89d-39160dc5c91e.png">

### Test ``` go test ./... ```
<img width="1680" alt="Screenshot 2023-04-12 at 15 57 24" src="https://user-images.githubusercontent.com/43158886/231497965-ed9cebe9-5be8-4799-8904-c45bc5f3fdce.png">

## Bench Test ``` go test -bench=.
<img width="1680" alt="Screenshot 2023-04-12 at 15 55 05" src="https://user-images.githubusercontent.com/43158886/231497378-93f8e1fc-ae8a-4830-ac07-6e8b1148ffe3.png">

